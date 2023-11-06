package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestSimilarityOptions_String(t *testing.T) {
	String := tf.NamedFunction(t, "SimilarityOptions_String", gedcom.SimilarityOptions.String)
	options := gedcom.NewSimilarityOptions()

	String(options).Returns("MinimumSimilarity:0.733, " +
		"MinimumWeightedSimilarity:0.733, " +
		"MaxYears:3, " +
		"IndividualWeight:0.8, " +
		"ParentsWeight:0.06666666666666667, " +
		"SpousesWeight:0.06666666666666667, " +
		"ChildrenWeight:0.06666666666666667, " +
		"NameToDateRatio:0.5, " +
		"JaroBoostThreshold:0, " +
		"JaroPrefixSize:8, " +
		"PreferPointerAbove:0.733")
}

func TestNewSimilarityOptions(t *testing.T) {
	options := gedcom.NewSimilarityOptions()

	t.Run("Weights", func(t *testing.T) {
		shouldBeOne := options.IndividualWeight + options.ParentsWeight +
			options.SpousesWeight + options.ChildrenWeight

		assert.Equal(t, 1.0, shouldBeOne)
	})

	t.Run("PreferPointerAbove", func(t *testing.T) {
		assert.Equal(t, gedcom.DefaultMinimumSimilarity, options.PreferPointerAbove)
	})
}
