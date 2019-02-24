package gedcom

import "sync"

// IndividualNodesCompareOptions provides more optional attributes for
// IndividualNodes.Compare.
//
// You should use NewIndividualNodesCompareOptions to start with sensible
// defaults.
type IndividualNodesCompareOptions struct {
	// SimilarityOptions controls the weights of the comparisons. See
	// SimilarityOptions for more information. Any matches below
	// MinimumSimilarity will not be used.
	//
	// Since this can take a long time to run with a lot of individuals there is
	// an optional notifier channel that can be listened to for progress
	// updates. Or pass nil to ignore this feature.
	SimilarityOptions SimilarityOptions

	// Notifier will be sent progress updates throughout the comparison if it is
	// not nil. If it is nil then this feature is ignored.
	//
	// You can control how precise this is with NotifierStep.
	//
	// You may close this Notifier to abort the comparison early.
	Notifier chan Progress

	// NotifierStep is the number of comparisons that must happen before the
	// Notifier be notified. The default is zero so all comparisons will cause a
	// notify. You should set this to a higher amount of reduce the frequency of
	// chatter.
	NotifierStep int64

	// Jobs controls the parallelism of the processing. The default value of 0
	// works the same as a value of 1 (no concurrency). Any number higher than 1
	// will create extra go routines to increase the CPU utilization of the
	// processing.
	//
	// It is important to note that the parallelism is still bound by
	// GOMAXPROCS.
	Jobs int

	// These lengths start with the respective sizes of the IndividualNodes
	// slices. In the simplest case all left will need to be compared with all
	// right individuals resulting in leftLen * rightLen number of comparisons.
	//
	// However, shortcuts can be found through pointers and other unique
	// identifiers that can severely reduce the number of comparisons needed.
	// For this to be calculated properly we have to maintain the size of the
	// left and right through each stage accordingly. See adjustTotal.
	//
	// SentA and sentB describe individuals that should be excluded from the
	// brute force matrix. Related to leftLen and rightLen.
	leftLen, rightLen int64
	sentA, sentB      *sync.Map
	totalMutex        *sync.Mutex
}

// NewIndividualNodesCompareOptions creates sensible defaults for
// IndividualNodesCompareOptions. In the majority of cases you will not need to
// change any further options.
//
// Important: A single IndividualNodesCompareOptions must be created for each
// comparison.
func NewIndividualNodesCompareOptions() *IndividualNodesCompareOptions {
	return &IndividualNodesCompareOptions{
		SimilarityOptions: NewSimilarityOptions(),
		totalMutex:        &sync.Mutex{},
		sentA:             &sync.Map{},
		sentB:             &sync.Map{},
	}
}

func (o *IndividualNodesCompareOptions) notify(m Progress) {
	defer func() {
		// Catch "panic: send on closed channel". This means Notifier was closed
		// prematurely to abort the comparisons.
		recover()
	}()

	if o.Notifier != nil {
		o.Notifier <- m
	}
}

func (o *IndividualNodesCompareOptions) notifierStep() int64 {
	return maxInt64(o.NotifierStep, 1)
}

func (o *IndividualNodesCompareOptions) ConcurrentJobs() int {
	return maxInt(o.Jobs, 1)
}

func (o *IndividualNodesCompareOptions) adjustTotal(totals chan int64) {
	// See getTotals(). We need to notify that there will be now less jobs.
	o.totalMutex.Lock()
	totals <- -int64(o.leftLen + o.rightLen - 2)
	o.leftLen--
	o.rightLen--
	o.totalMutex.Unlock()
}
