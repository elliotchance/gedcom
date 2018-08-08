package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimilarityOptions_String(t *testing.T) {
	String := tf.Function(t, gedcom.SimilarityOptions.String)
	options := *gedcom.NewSimilarityOptions()

	String(options).Returns("MinimumSimilarity:0.735, " +
		"MinimumWeightedSimilarity:0.735, " +
		"MaxYears:3, " +
		"IndividualWeight:0.8, " +
		"ParentsWeight:0.06666666666666667, " +
		"SpousesWeight:0.06666666666666667, " +
		"ChildrenWeight:0.06666666666666667, " +
		"NameToDateRatio:0.5, " +
		"JaroBoostThreshold:0, " +
		"JaroPrefixSize:8")
}

func TestNewSimilarityOptions(t *testing.T) {
	options := gedcom.NewSimilarityOptions()

	shouldBeOne := options.IndividualWeight + options.ParentsWeight +
		options.SpousesWeight + options.ChildrenWeight

	assert.Equal(t, 1.0, shouldBeOne)
}
