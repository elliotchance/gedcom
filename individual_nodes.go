package gedcom

import (
	"sort"
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
const DefaultMinimumSimilarity = 0.735

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
func (nodes IndividualNodes) Similarity(other IndividualNodes, options *SimilarityOptions) float64 {
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
	Similarity SurroundingSimilarity
}

// CompareProgress contains information about the progress of a Comparison.
type CompareProgress struct {
	Done, Total int64
}

// IndividualNodesCompareOptions provides more optional attributes for
// IndividualNodes.Compare.
type IndividualNodesCompareOptions struct {
	// Options controls the weights of the comparisons. See SimilarityOptions
	// for more information. Any matches below MinimumSimilarity will not be
	// used.
	//
	// Since this can take a long time to run with a lot of individuals there is an
	// optional notifier channel that can be listened to for progress updates. Or
	// pass nil to ignore this feature.
	//
	SimilarityOptions *SimilarityOptions

	// Notifier will be sent progress updates throughout the comparison if it is
	// not nil. If it is nil then this feature is ignored.
	//
	// You can control how precise this is with NotifierStep.
	Notifier chan CompareProgress

	// NotifierStep is the number of comparisons that must happen before the
	// Notifier is used. The default is zero so all comparisons will cause a
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

func (o *IndividualNodesCompareOptions) end() {
	if o.Notifier != nil {
		close(o.Notifier)
	}
}

func (o *IndividualNodesCompareOptions) notify(m CompareProgress) {
	if o.Notifier != nil {
		o.Notifier <- m
	}
}

func (o *IndividualNodesCompareOptions) notifierStep() int64 {
	return maxInt64(o.NotifierStep, 1)
}

func (o *IndividualNodesCompareOptions) jobs() int {
	return maxInt(o.Jobs, 1)
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
func (nodes IndividualNodes) Compare(other IndividualNodes, options *IndividualNodesCompareOptions) []IndividualComparison {
	// Calculate all the similarities of the matrix.
	similarities := []IndividualComparison{}

	// ghost:ignore
	total := int64(len(nodes)) * int64(len(other))

	jobs := make(chan IndividualComparison, 100)
	results := make(chan IndividualComparison, 100)

	// Start workers.
	numberOfJobs := options.jobs()
	for w := 1; w <= numberOfJobs; w++ {
		go nodes.compareWorker(options.SimilarityOptions, jobs, results)
	}

	// Send jobs.
	go func() {
		for _, a := range nodes {
			for _, b := range other {
				jobs <- IndividualComparison{
					Left:  a,
					Right: b,
				}
			}
		}

		close(jobs)
	}()

	// Collect results.
	for done := int64(1); done <= total; done++ {
		similarities = append(similarities, <-results)

		if done%options.notifierStep() == 0 {
			options.notify(CompareProgress{
				Done:  done,
				Total: total,
			})
		}
	}

	// Make sure we notify that all comparisons have completed.
	options.notify(CompareProgress{
		Done:  total,
		Total: total,
	})

	// Sort by similarity.
	sort.SliceStable(similarities, func(i, j int) bool {
		similarityOptions := options.SimilarityOptions

		return similarities[i].Similarity.WeightedSimilarity(similarityOptions) >
			similarities[j].Similarity.WeightedSimilarity(similarityOptions)
	})

	// Find the winners.
	found := map[*IndividualNode]bool{}
	winners := []IndividualComparison{}
	for _, s := range similarities {
		// Once we have gone below the acceptable similarity we can bail out.
		if s.Similarity.WeightedSimilarity(options.SimilarityOptions) < options.SimilarityOptions.MinimumWeightedSimilarity {
			break
		}

		// We can only proceed with a match if both sides are unmatched.
		if found[s.Left] == true || found[s.Right] == true {
			continue
		}

		winners = append(winners, s)
		found[s.Left] = true
		found[s.Right] = true
	}

	// All of the remaining need to be added.
	for _, left := range nodes {
		if !found[left] {
			winners = append(winners, IndividualComparison{
				Left: left,
			})
		}
	}

	for _, right := range other {
		if !found[right] {
			winners = append(winners, IndividualComparison{
				Right: right,
			})
		}
	}

	options.end()

	return winners
}

func (nodes IndividualNodes) compareWorker(similarityOptions *SimilarityOptions, jobs <-chan IndividualComparison, results chan<- IndividualComparison) {
	for j := range jobs {
		j.Similarity = j.Left.SurroundingSimilarity(j.Right, similarityOptions)
		results <- j
	}
}
