package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/stretchr/testify/assert"
)

func assertAge(t *testing.T, age gedcom.Age, years float64, isEstimate bool, constraint gedcom.AgeConstraint) {
	assert.True(t, age.IsKnown, "is known")
	assert.Equal(t, isEstimate, age.IsEstimate, "is estimate")
	assert.Equal(t, constraint.String(), age.Constraint.String(), "constraint")

	expected := float64(age.Age) / 3.15576e16
	assert.InDelta(t, years, expected, 0.1)
}

func assertNodeEqual(t *testing.T, expected, actual gedcom.Node, msgAndArgs ...interface{}) {
	if gedcom.IsNil(expected) || gedcom.IsNil(actual) {
		assert.True(t, gedcom.IsNil(expected))
		assert.True(t, gedcom.IsNil(actual))
	} else {
		assert.True(t, expected.Equals(actual), msgAndArgs...)
	}
}

func assertDocumentEqual(t *testing.T, expected, actual *gedcom.Document, msgAndArgs ...interface{}) {
	assert.Equal(t, expected.String(), actual.String(), msgAndArgs...)

	if !assert.Equal(t, len(expected.Nodes()), len(actual.Nodes()), msgAndArgs...) {
		return
	}

	for i, n := range expected.Nodes() {
		assertNodeEqual(t, n, actual.Nodes()[i])
	}

	assert.Equal(t, expected.MaxLivingAge, actual.MaxLivingAge, msgAndArgs...)
	assert.Equal(t, expected.HasBOM, actual.HasBOM, msgAndArgs...)
}

func assertError(t *testing.T, expected, actual error) {
	if expected == nil {
		assert.NoError(t, actual)
	} else {
		assert.EqualError(t, expected, actual.Error())
	}
}
