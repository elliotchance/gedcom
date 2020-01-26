package gedcom

import "fmt"

// SimilarityOptions is used by all of the functions that calculate the
// similarity or otherwise compare entities. This struct allows many things to
// be tweaked. However, not all of the values are used by all of the functions.
//
// Use NewSimilarityOptions() to choose sensible defaults that are best for most
// general cases.
type SimilarityOptions struct {
	// MinimumSimilarity is the threshold for matching individuals as the same
	// person. This is used to compare only the individual (not surrounding
	// family) like spouses and children. See DefaultMinimumSimilarity.
	MinimumSimilarity float64

	// MinimumWeightedSimilarity is the threshold for whether two individuals
	// should be the seen as the same person when the surrounding immediate
	// family is taken into consideration. See WeightedSimilarity().
	MinimumWeightedSimilarity float64

	// MaxYears is the maximum error margin (in years) that two dates can be
	// different before they are assume to not be the same. See
	// DefaultMaxYearsForSimilarity.
	MaxYears float64

	// All four of these must sum up to 1.0.
	IndividualWeight, ParentsWeight, SpousesWeight, ChildrenWeight float64

	// NameToDateRatio describes the ratio between the weight of the individuals
	// name to their combined estimated birth and death dates. A value of 0.0
	// would not take into account the individuals name at all, whereas 1.0
	// would not take into account any dates. A sensible default is 0.5.
	NameToDateRatio float64

	// JaroBoostThreshold and JaroPrefixSize are used by the JaroWinkler
	// function. They affect the properties of names are compared. The default
	// values for each of these can be found in the constants
	// DefaultJaroWinklerBoostThreshold and DefaultJaroWinklerPrefixSize. Their
	// values have been chosen with "gedcom tune".
	JaroBoostThreshold float64
	JaroPrefixSize     int

	// PreferPointerAbove controls if two individuals should be considered a
	// match by their pointer value.
	//
	// The default value is DefaultMinimumSimilarity which means that the
	// individuals will be considered a match if they share the same pointer and
	// hit the same default minimum similarity.
	//
	// A value of 1.0 would have to be a perfect match to be considered equal on
	// their pointer, this is the same as disabling the feature.
	//
	// A value of 0.0 would mean that it always trusts the pointer match, even
	// if the individuals are nothing alike.
	//
	// PreferPointerAbove makes sense when you are comparing documents that have
	// come from the same base and retained the pointers between individuals of
	// the existing data.
	PreferPointerAbove float64
}

// NewSimilarityOptions returns sensible defaults that are used around many of
// the similarity functions.
func NewSimilarityOptions() SimilarityOptions {
	return SimilarityOptions{
		MaxYears:                  DefaultMaxYearsForSimilarity,
		MinimumSimilarity:         DefaultMinimumSimilarity,
		MinimumWeightedSimilarity: DefaultMinimumSimilarity,

		// All four of these must sum up to 1.0.
		IndividualWeight: 0.8,
		ParentsWeight:    0.2 / 3,
		SpousesWeight:    0.2 / 3,
		ChildrenWeight:   0.2 - (0.4 / 3),

		NameToDateRatio:    0.5,
		JaroBoostThreshold: DefaultJaroWinklerBoostThreshold,
		JaroPrefixSize:     DefaultJaroWinklerPrefixSize,

		// Allow individuals to me matched using their pointer if they hit the
		// same default minimum threshold.
		PreferPointerAbove: DefaultMinimumSimilarity,
	}
}

// String renders the options as a comma-separated string.
func (options SimilarityOptions) String() string {
	s := fmt.Sprintf("%#v", options)
	sLen := len(s)

	return s[25 : sLen-1]
}

func (options SimilarityOptions) canSkipExtraProcessing(individualSimilarity float64) bool {
	actual := individualSimilarity * options.IndividualWeight
	threshold := options.MinimumWeightedSimilarity -
		options.ParentsWeight -
		options.SpousesWeight -
		options.ChildrenWeight

	return actual <= threshold
}
