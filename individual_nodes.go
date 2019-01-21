package gedcom

import (
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

// DefaultMinimumSimilarity is a sensible value to provide to the
// minimumSimilarity parameter of IndividualNodes.Similarity.
//
// It is quite possible that this value will change in the future if a more
// accurate figure is found or the algorithm is generally tuned with different
// weightings.
//
// The value was chosen by running comparison experiments with gedcomtune, a
// tool to try and find ideal values for constants like this.
const DefaultMinimumSimilarity = 0.733

// IndividualNodes is a collection of individuals.
type IndividualNodes []*IndividualNode

type individualSimilarity struct {
	a, b       *IndividualNode
	similarity float64
}

// Similarity calculates how similar two collections of individuals are. The
// individuals could be children of the same parents, spouses that share the
// same partner or some other similar group of individuals.
//
// Individuals that have a similarity below minimumSimilarity will not be
// considered a match. Making this a larger value will lead to less
// false-positives but also cause more true-negatives. You can use the
// DefaultMinimumSimilarity constant if you are unsure.
//
// As a rule of thumb you should not trust a return value of 0.5 of less as any
// kind of match. 0.7 and above is getting in the range of what is likely to be
// considered the same.
//
// Although there is no theoretical limit to the amount of individuals that can
// be compared, this function invokes a complex chain of comparisons that can be
// very CPU intensive. It is not a good idea to use it compare a large amount of
// individuals, like an entire file.
//
// If both slices have zero elements then 1.0 is returned. However, if one slice
// is not empty then 0.5 is returned. 0.5 is actually the same value that would
// be returned in all cases if the logic were to fall through.
//
// The algorithm works through several stages explained in more detail below but
// the basic principle is this:
//
// 1. Find an individual from both slices that have the highest similarity.
// There similarity is added to a total. These individuals are now excluded from
// the next pass.
//
// 2. Repeat the first step until one of the list runs out of individuals.
//
// 3. Increase the number of matched individuals to the length of the largest
// slice by using the score 0.5 for each individual.
//
// 4. Take the average. The total score (including the 0.5 for padding
// individuals) and divide it by the maximum slice size.
//
// The easiest way to explain how the process works is by working through an
// example. Consider the two following slices of individuals:
//
//   [a, b, c]
//   [d, e]
//
// There is 5 individuals in total, but the largest slice length is 3.
//
// The first step is to calculate the individual similarity matrix between the
// inputs. This is the expensive part because it takes O(nm) time where n and m
// are the slice lengths respectively. The results may look like:
//
//   a.Similarity(d) = 0.234
//   a.Similarity(e) = 0.123
//   b.Similarity(d) = 0.546
//   b.Similarity(e) = 0.678
//   c.Similarity(d) = 0.235
//   c.Similarity(e) = 0.456
//
// The results are sorted by their similarity score, highest first:
//
//   b.Similarity(e) = 0.678
//   b.Similarity(d) = 0.546
//   c.Similarity(e) = 0.456
//   c.Similarity(d) = 0.235
//   a.Similarity(d) = 0.234
//   a.Similarity(e) = 0.123
//
// In reality the scores should be closer to 1.0 for many of the individuals,
// but I will use a wide range of values for this example.
//
// The above list states that individual "b" is the most similar individual to
// "e". So we should consider these a match, and record the fact that "b" and
// "e" are matched.
//
// Moving on to the second iteration, "b" has already been matched. Even though
// "d" has not we still skip this iteration. The same is true for the third
// iteration, since "e" has already been matched.
//
// The forth iteration is the first time we see two new individuals, "c" and
// "d". For the second time we consider these two to be a match.
//
// The process continues and can be visualised like this:
//
//   b.Similarity(e) = 0.678   b = e
//   b.Similarity(d) = 0.546   b already matched
//   c.Similarity(e) = 0.456   e already matched
//   c.Similarity(d) = 0.235   c = d
//   a.Similarity(d) = 0.234   d already matched
//   a.Similarity(e) = 0.123   e already matched
//
// There should be the same amount of matches as the smallest slice size. As we
// can see the two matches is the same number as the smallest slice.
//
// Finally, we pad out the remaining individuals with a score of 0.5 so that the
// total matches equals the longest slice. Using a middle-score (and not 0.0) is
// important because 0.0 would represent that individuals were a complete
// non-match and throw out the final score quite heavily. We must treat missing
// individuals as neither a perfect or non-match because we simply don't have
// the information to make the call either way, hence 0.5.
//
// The returned score is now the average:
//
//   (0.678 + 0.234 + 0.5) / 3 = 0.4707
//
// There are some known caveats to the algorithm:
//
// When calculating the similarity of individuals where data is missing from
// both sides (such as both individuals having a missing birth date) then the
// result will be less than 1.0 even though the individuals are arguably the
// same.
//
// This is designed this way on purpose as to not so eagerly match individuals
// with incomplete information. Otherwise these would take a higher score for
// matches of individuals that have a slightly different birth date.
//
// The options.MaxYears allows the error margin on dates to be adjusted. See
// DefaultMaxYearsForSimilarity for more information.
func (nodes IndividualNodes) Similarity(other IndividualNodes, options SimilarityOptions) float64 {
	nodesLen := float64(len(nodes))
	otherLen := float64(len(other))

	// We have to catch this because otherwise it would lead to a divide-by-zero
	// at the end.
	if nodesLen == 0 && otherLen == 0 {
		return 1
	}

	// 0.5 is actually the same value that would be returned in all cases if the
	// logic were to fall through.
	if nodesLen == 0 || otherLen == 0 {
		return 0.5
	}

	// Calculate all the similarities of the matrix.
	similarities := []*individualSimilarity{}

	for _, a := range nodes {
		for _, b := range other {
			similarities = append(similarities, &individualSimilarity{
				a:          a,
				b:          b,
				similarity: a.Similarity(b, options),
			})
		}
	}

	// Sort by similarity.
	sort.SliceStable(similarities, func(i, j int) bool {
		return similarities[i].similarity > similarities[j].similarity
	})

	// Find the winners.
	found := map[*IndividualNode]bool{}
	winners := []*individualSimilarity{}
	for _, s := range similarities {
		// Once we have gone below the acceptable similarity we can bail out.
		if s.similarity < options.MinimumSimilarity {
			break
		}

		// We can only proceed with a match if both sides are unmatched.
		if found[s.a] == true || found[s.b] == true {
			continue
		}

		winners = append(winners, s)
		found[s.a] = true
		found[s.b] = true
	}

	// Tally up what we have and fill out the missing individuals.
	total := 0.0
	for _, s := range winners {
		total += s.similarity
	}

	if otherLen > nodesLen {
		nodesLen = otherLen
	}

	winnersLength := float64(len(winners))
	total += 0.5 * (nodesLen - winnersLength)

	return total / nodesLen
}

// IndividualComparison is the result of two compared individuals.
type IndividualComparison struct {
	// Left or Right may be nil, but never both.
	Left, Right *IndividualNode

	// Similarity will only contain a usable value if Left and Right are not
	// nil. Otherwise, Similarity may contain any unexpected value.
	Similarity *SurroundingSimilarity
}

// IndividualComparisons is a slice of IndividualComparison instances.
type IndividualComparisons []*IndividualComparison

// String returns each comparison string on its own like, like:
//
//   John Smith <-> John H Smith (0.833333)
//   Jane Doe <-> (none) (?)
//   (none) <-> Joe Bloggs (?)
//
func (comparisons IndividualComparisons) String() string {
	lines := []string{}

	for _, comparison := range comparisons {
		lines = append(lines, comparison.String())
	}

	return strings.Join(lines, "\n")
}

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
}

// NewIndividualNodesCompareOptions creates sensible defaults for
// IndividualNodesCompareOptions. In the majority of cases you will not need to
// change any further options.
func NewIndividualNodesCompareOptions() *IndividualNodesCompareOptions {
	return &IndividualNodesCompareOptions{
		SimilarityOptions: NewSimilarityOptions(),
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

func createJobs(totals chan int64, left, right IndividualNodes, options *IndividualNodesCompareOptions) chan *IndividualComparison {
	// Because the jobs are so small I've found that using a buffered channel
	// can make the processing up to 30% faster on my 4 cores. I'm not sure what
	// the best number for this should be, or if it could be dynamic.
	jobs := make(chan *IndividualComparison, 10)

	// Before we send of the matrix of comparisons we should attempt to find (if
	// any) the matching pointers.
	//
	// Even though 1.0 would have the effect of only matching perfect matches we
	// still want to cover this case as we might match some perfectly matched
	// individuals that do not need to go into the matrix below.
	//
	// Since the pointers have to exist in both sides and similarities are
	// calculated the same in both directions we only need to check one side of
	// the individuals.
	//
	// We should really only check the side with the least amount of items but
	// this is such a fast process that I won't complicate the code with this
	// right now.
	go func() {
		// Describes the pointers that we have found to match and have already
		// been emitted so there is no need to send them again.
		sentPointers := &sync.Map{}

		// This check is important because if the right side is empty we will
		// not be able to get the Document on the right side to to the
		// comparison.
		leftLen := int64(len(left))
		rightLen := int64(len(right))
		if len(right) > 0 {
			m := sync.Mutex{}
			ws := options.ConcurrentJobs()

			wg := sync.WaitGroup{}
			for w := 0; w < ws; w++ {
				wg.Add(1)
				go func(w int) {
					for leftI := w; leftI < len(left); leftI += ws {
						a := left[leftI]
						b := right.ByPointer(a.Pointer())
						if b != nil {
							ss := a.SurroundingSimilarity(b, options.SimilarityOptions, true)
							if ss.WeightedSimilarity() >= options.SimilarityOptions.PreferPointerAbove {
								// See getTotals(). We need to notify that there will
								// be now less jobs.
								//
								// We are removing one individual from both sides, but
								// the number of comparisons goes down by the sum of
								// each side?
								m.Lock()
								totals <- -int64(leftLen + rightLen - 2)
								leftLen--
								rightLen--
								m.Unlock()

								jobs <- &IndividualComparison{
									Left:       a,
									Right:      b,
									Similarity: ss,
								}

								// The true value here does not actually matter.
								sentPointers.Store(a.Pointer(), nil)
							}
						}
					}

					wg.Done()
				}(w)
			}

			wg.Wait()
		}

		close(totals)

		// Send the remaining matrix of individuals to be compared.
		for _, a := range left {
			if _, ok := sentPointers.Load(a.Pointer()); ok {
				continue
			}

			for _, b := range right {
				if _, ok := sentPointers.Load(b.Pointer()); ok {
					continue
				}

				jobs <- &IndividualComparison{
					Left:  a,
					Right: b,
				}
			}
		}

		close(jobs)
	}()

	return jobs
}

func (o *IndividualNodesCompareOptions) processJobs(jobs chan *IndividualComparison) chan *IndividualComparison {
	// See description in createJobs().
	results := make(chan *IndividualComparison, 10)

	go func() {
		wg := sync.WaitGroup{}

		for i := 0; i < 1; i++ {
			wg.Add(1)
			go func() {
				for j := range jobs {
					// The similarity may already be calculated from when it was
					// comparing on the pointer.
					if j.Similarity == nil {
						j.Similarity = j.Left.SurroundingSimilarity(j.Right, o.SimilarityOptions, false)
					}
					results <- j
				}

				wg.Done()
			}()
		}

		wg.Wait()
		close(results)
	}()

	return results
}

func (o *IndividualNodesCompareOptions) collectResults(results chan *IndividualComparison, totals chan int64) chan *IndividualComparison {
	// See description in createJobs().
	similarities := make(chan *IndividualComparison, 10)

	go func() {
		total := int64(0)
		done := int64(0)
		for {
			select {
			case t, ok := <-totals:
				if !ok {
					totals = nil
					continue
				}

				total += t

			case next, ok := <-results:
				if !ok {
					results = nil
					continue
				}

				similarities <- next

				if done%o.notifierStep() == 0 {
					o.notify(Progress{
						Done:  done,
						Total: total,
					})
				}

				done++

			default:
				time.Sleep(1 * time.Millisecond)
			}

			if totals == nil && results == nil {
				break
			}
		}

		// Make sure we notify that all comparisons have completed.
		o.notify(Progress{
			Done:  total,
			Total: total,
		})

		close(similarities)
	}()

	return similarities
}

func (o *IndividualNodesCompareOptions) calculateWinners(a, b IndividualNodes, similarityResults chan *IndividualComparison, options SimilarityOptions) chan *IndividualComparison {
	// See description in createJobs().
	winners := make(chan *IndividualComparison, 10)

	go func() {
		similarities := IndividualComparisons{}
		found := map[*IndividualNode]bool{}

		// We have to collect all items before they can be sorted.
		for similarity := range similarityResults {
			// If any of the matches were made through pointers we have to
			// remove them from the possible winners.
			if similarity.Left.Pointer() == similarity.Right.Pointer() &&
				similarity.Similarity.WeightedSimilarity() >= options.PreferPointerAbove {
				winners <- similarity
				found[similarity.Left] = true
				found[similarity.Right] = true
				continue
			}

			similarities = append(similarities, similarity)
		}

		// Sort by similarity.
		sort.SliceStable(similarities, func(i, j int) bool {
			return similarities[i].Similarity.WeightedSimilarity() >
				similarities[j].Similarity.WeightedSimilarity()
		})

		// Find the winners.
		for _, s := range similarities {
			// Once we have gone below the acceptable similarity we can bail out.
			minWS := o.SimilarityOptions.MinimumWeightedSimilarity
			if s.Similarity.WeightedSimilarity() < minWS {
				break
			}

			// We can only proceed with a match if both sides are unmatched.
			if found[s.Left] == true || found[s.Right] == true {
				continue
			}

			winners <- s
			found[s.Left] = true
			found[s.Right] = true
		}

		// All of the remaining need to be added.
		for _, left := range a {
			if !found[left] {
				winners <- &IndividualComparison{
					Left: left,
				}
			}
		}

		for _, right := range b {
			if !found[right] {
				winners <- &IndividualComparison{
					Right: right,
				}
			}
		}

		close(winners)
	}()

	return winners
}

func (o *IndividualNodesCompareOptions) getTotals(nodes, other IndividualNodes) chan int64 {
	// This channel must be buffered to allow the initial value below. It cannot
	// be a separate goroutine because in cases where the comparison is
	// extremely small the totals cannot could be closed before It writes the
	// initial total.
	//
	// Yes, it could be rewritten with an extra select{} but this works just as
	// well.
	totals := make(chan int64, 1)

	// The totals channel starts with the maximum possible comparisons. As we
	// find matches that allow us to skip blocks of comparisons new totals will
	// be calculated and pushed through.
	totals <- int64(len(nodes)) * int64(len(other))

	return totals
}

// Compare returns the matching individuals from two lists.
//
// The length of the result slice will be no larger than the largest slice
// provided and no smaller than the smallest slice provided.
//
// Individuals can only be matched once on their respective side so you can
// guarantee that all the Left's are unique and belong to the current nodes.
// Likewise all Right's will be unique and only belong to the other set.
//
// See IndividualNodesCompareOptions for more options.
func (nodes IndividualNodes) Compare(other IndividualNodes, options *IndividualNodesCompareOptions) IndividualComparisons {
	defer func() {
		if options.Notifier != nil {
			close(options.Notifier)
			options.Notifier = nil
		}
	}()

	totals := options.getTotals(nodes, other)
	jobs := createJobs(totals, nodes, other, options)
	results := options.processJobs(jobs)
	similarities := options.collectResults(results, totals)
	winners := options.calculateWinners(nodes, other, similarities, options.SimilarityOptions)

	realWinners := IndividualComparisons{}
	for winner := range winners {
		realWinners = append(realWinners, winner)
	}

	return realWinners
}

// Merge uses the expensive but more accurate Compare algorithm to determine the
// best way to merge two slices of individuals.
//
// Merge relies exclusively on the logic of Compare so there's no need to repeat
// those details here.
//
// Any individuals that do not match on either side will be appended to the end.
func (nodes IndividualNodes) Merge(other IndividualNodes, options *IndividualNodesCompareOptions) (IndividualNodes, error) {
	comparisons := nodes.Compare(other, options)
	merged := IndividualNodes{}

	for _, comparison := range comparisons {
		left := comparison.Left
		right := comparison.Right

		switch {
		case left != nil && right != nil:
			node, err := MergeNodes(left, right)
			if err != nil {
				return nil, err
			}

			merged = append(merged, node.(*IndividualNode))

		case left != nil:
			merged = append(merged, left)

		case right != nil:
			merged = append(merged, right)
		}
	}

	return merged, nil
}

// Nodes returns a slice containing the same individuals.
//
// Individuals that are manipulated will affect the original individuals.
func (nodes IndividualNodes) Nodes() (ns []Node) {
	for _, individual := range nodes {
		ns = append(ns, individual)
	}

	return
}

// GEDCOMString returns the GEDCOM for all individuals. See GEDCOMStringer for
// more details.
func (nodes IndividualNodes) GEDCOMString(indent int) string {
	return NewDocumentWithNodes(nodes.Nodes()).GEDCOMString(0)
}

func (nodes IndividualNodes) String() string {
	s := []string{}

	for _, individual := range nodes {
		s = append(s, individual.String())
	}

	return strings.Join(s, "\n")
}

func (nodes IndividualNodes) ByPointer(pointer string) *IndividualNode {
	for _, node := range nodes {
		if node.Pointer() == pointer {
			return node
		}
	}

	return nil
}

func (c IndividualComparison) stringOrDefault(s fmt.Stringer, def string) string {
	if IsNil(s) {
		return def
	}

	return s.String()
}

// String returns the comparison in a human-readable format, like one of:
//
//   John Smith <-> John H Smith (0.833333)
//   Jane Doe <-> (none) (?)
//   (none) <-> Joe Bloggs (?)
//
func (c IndividualComparison) String() string {
	left := c.stringOrDefault(c.Left, "(none)")
	right := c.stringOrDefault(c.Right, "(none)")
	similarity := c.stringOrDefault(c.Similarity, "?")

	return fmt.Sprintf("%s <-> %s (%s)", left, right, similarity)
}
