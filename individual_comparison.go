package gedcom

import "strings"

// IndividualComparison is the result of two compared individuals.
type IndividualComparison struct {
	// Left or Right may be nil, but never both.
	Left, Right *IndividualNode

	// Similarity will only contain a usable value if Left and Right are not
	// nil. Otherwise, Similarity may contain any unexpected value.
	Similarity *SurroundingSimilarity

	// This is an internal flag when a comparison is known to be equal so that
	// it doesn't need to be tested again.
	certainMatch bool
}

// IndividualComparisons is a slice of IndividualComparison instances.
type IndividualComparisons []*IndividualComparison

func NewIndividualComparison(Left, Right *IndividualNode, Similarity *SurroundingSimilarity) *IndividualComparison {
	return &IndividualComparison{
		Left:       Left,
		Right:      Right,
		Similarity: Similarity,
	}
}

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
