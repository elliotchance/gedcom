package gedcom

import "sort"

// DefaultMinimumSimilarity is a sensible value to provide to the
// minimumSimilarity parameter of IndividualNodes.Similarity.
//
// It is quite possible that this value will change in the future if a more
// accurate figure is found or the algorithm is generally tuned with different
// weightings.
const DefaultMinimumSimilarity = 0.7

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
func (nodes IndividualNodes) Similarity(other IndividualNodes, minimumSimilarity float64) float64 {
	// We have to catch this because otherwise it would lead to a divide-by-zero
	// at the end.
	if len(nodes) == 0 && len(other) == 0 {
		return 1
	}

	// 0.5 is actually the same value that would be returned in all cases if the
	// logic were to fall through.
	if len(nodes) == 0 || len(other) == 0 {
		return 0.5
	}

	// Calculate all the similarities of the matrix.
	similarities := []*individualSimilarity{}

	for _, a := range nodes {
		for _, b := range other {
			similarities = append(similarities, &individualSimilarity{
				a:          a,
				b:          b,
				similarity: a.Similarity(b),
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
		if s.similarity < minimumSimilarity {
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

	wantedLength := len(nodes)
	if len(other) > wantedLength {
		wantedLength = len(other)
	}

	total += 0.5 * float64(wantedLength-len(winners))

	return total / float64(wantedLength)
}

// IndividualComparison is the result of two compared individuals.
type IndividualComparison struct {
	// Left or Right may be nil, but never both.
	Left, Right *IndividualNode

	// Similarity will only contain a usable value if Left and Right are not
	// nil. Otherwise, Similarity may contain any unexpected value.
	Similarity SurroundingSimilarity
}

// Compare returns the matching individuals from two lists.
//
// It is expected that all of the individuals in a single slice come from the
// same document. doc1 and doc2 are used as the Documents for the current and
// other nodes respectively. If both sets of individuals come from the same
// Document you must specify the same Document for both values.
//
// The length of the result slice will be no larger than the largest slice
// provided and no smaller than the smallest slice provided.
//
// Individuals can only be matched once on their respective side so you can
// guarantee that all the Left's are unique and belong to the current nodes.
// Likewise all Right's will be unique and only belong to the other set.
//
// The minimumSimilarity sets a threshold of WeightedSimilarity(). Any matches
// below minimumSimilarity will not be used.
func (nodes IndividualNodes) Compare(doc1, doc2 *Document, other IndividualNodes, minimumSimilarity float64) []IndividualComparison {
	comparisons := []IndividualComparison{}

	// Tracks individuals that already part of a match.
	matched := map[*IndividualNode]bool{}

	for _, left := range nodes {
		comparison := IndividualComparison{
			Left: left,
		}

		for _, right := range other {
			s := left.SurroundingSimilarity(doc1, doc2, right)
			weighted := s.WeightedSimilarity()
			if weighted >= minimumSimilarity &&
				weighted >= comparison.Similarity.WeightedSimilarity() {
				comparison.Right = right
				comparison.Similarity = s

				matched[right] = true
			}
		}

		comparisons = append(comparisons, comparison)
	}

	// All of the remaining right side need to be added.
	for _, right := range other {
		if !matched[right] {
			comparisons = append(comparisons, IndividualComparison{
				// Left and Similarity will be nil.
				Right: right,
			})
		}
	}

	return comparisons
}
