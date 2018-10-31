package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"testing"
)

func TestAgeConstraint_String(t *testing.T) {
	String := tf.Function(t, gedcom.AgeConstraint.String)

	String(gedcom.AgeConstraintUnknown).Returns("Unknown")
	String(gedcom.AgeConstraintBeforeBirth).Returns("Before Birth")
	String(gedcom.AgeConstraintLiving).Returns("Living")
	String(gedcom.AgeConstraintAfterDeath).Returns("After Death")
	String(gedcom.AgeConstraint(-1)).Returns("Unknown")
}
