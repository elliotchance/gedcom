package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/tf"
)

func TestAgeConstraint_String(t *testing.T) {
	String := tf.Function(t, gedcom.AgeConstraint.String)

	String(gedcom.AgeConstraintUnknown).Returns("Unknown")
	String(gedcom.AgeConstraintBeforeBirth).Returns("Before Birth")
	String(gedcom.AgeConstraintLiving).Returns("Living")
	String(gedcom.AgeConstraintAfterDeath).Returns("After Death")
	String(gedcom.AgeConstraint(-1)).Returns("Unknown")
}
