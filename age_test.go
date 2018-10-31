package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"testing"
	"time"
)

func TestNewUnknownAge(t *testing.T) {
	NewUnknownAge := tf.Function(t, gedcom.NewUnknownAge)

	NewUnknownAge().Returns(gedcom.Age{
		Age:        0,
		IsEstimate: false,
		IsKnown:    false,
		Constraint: gedcom.AgeConstraintUnknown,
	})
}

func TestNewAge(t *testing.T) {
	NewAge := tf.Function(t, gedcom.NewAge)

	NewAge(0, false, gedcom.AgeConstraintUnknown).Returns(gedcom.Age{
		Age:        0,
		IsEstimate: false,
		IsKnown:    true,
		Constraint: gedcom.AgeConstraintUnknown,
	})

	NewAge(100*time.Hour, true, gedcom.AgeConstraintLiving).Returns(gedcom.Age{
		Age:        360000000000000,
		IsEstimate: true,
		IsKnown:    true,
		Constraint: gedcom.AgeConstraintLiving,
	})

	NewAge(200*time.Hour, false, gedcom.AgeConstraintAfterDeath).Returns(gedcom.Age{
		Age:        720000000000000,
		IsEstimate: false,
		IsKnown:    true,
		Constraint: gedcom.AgeConstraintAfterDeath,
	})
}

func TestNewAgeWithYears(t *testing.T) {
	NewAgeWithYears := tf.Function(t, gedcom.NewAgeWithYears)

	NewAgeWithYears(0, false, gedcom.AgeConstraintBeforeBirth).Returns(gedcom.Age{
		Age:        0,
		IsEstimate: false,
		IsKnown:    true,
		Constraint: gedcom.AgeConstraintBeforeBirth,
	})

	NewAgeWithYears(2.5, true, gedcom.AgeConstraintLiving).Returns(gedcom.Age{
		Age:        78894000000000000,
		IsEstimate: true,
		IsKnown:    true,
		Constraint: gedcom.AgeConstraintLiving,
	})
}

func TestAge_IsAfter(t *testing.T) {
	IsAfter := tf.Function(t, gedcom.Age.IsAfter)

	IsAfter(gedcom.NewAgeWithYears(23, false, gedcom.AgeConstraintLiving),
		gedcom.NewAgeWithYears(45, false, gedcom.AgeConstraintLiving)).False()

	IsAfter(gedcom.NewAgeWithYears(63, false, gedcom.AgeConstraintLiving),
		gedcom.NewAgeWithYears(45, false, gedcom.AgeConstraintLiving)).True()

	IsAfter(gedcom.NewAgeWithYears(23.1, false, gedcom.AgeConstraintLiving),
		gedcom.NewAgeWithYears(23, false, gedcom.AgeConstraintLiving)).True()

	IsAfter(gedcom.NewAgeWithYears(30, false, gedcom.AgeConstraintLiving),
		gedcom.NewAgeWithYears(30, false, gedcom.AgeConstraintLiving)).False()
}
