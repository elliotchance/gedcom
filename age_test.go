package gedcom_test

import (
	"testing"
	"time"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
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

func TestAge_String(t *testing.T) {
	String := tf.Function(t, gedcom.Age.String)

	String(gedcom.NewUnknownAge()).Returns("unknown")

	String(gedcom.NewAgeWithYears(20, false, gedcom.AgeConstraintLiving)).Returns("20y")
	String(gedcom.NewAgeWithYears(20.5, false, gedcom.AgeConstraintLiving)).Returns("20y 6m")
	String(gedcom.NewAgeWithYears(20.01, false, gedcom.AgeConstraintLiving)).Returns("20y")

	String(gedcom.NewAgeWithYears(21, true, gedcom.AgeConstraintLiving)).Returns("~ 21y")
	String(gedcom.NewAgeWithYears(22.5, true, gedcom.AgeConstraintLiving)).Returns("~ 22y 6m")
	String(gedcom.NewAgeWithYears(23.01, true, gedcom.AgeConstraintLiving)).Returns("~ 23y")

	String(gedcom.NewAgeWithYears(22.5, true, gedcom.AgeConstraintUnknown)).Returns("~ 22y 6m")
	String(gedcom.NewAgeWithYears(22.5, true, gedcom.AgeConstraintBeforeBirth)).Returns("~ 22y 6m")
	String(gedcom.NewAgeWithYears(22.5, true, gedcom.AgeConstraintAfterDeath)).Returns("~ 22y 6m")

	String(gedcom.NewAgeWithYears(0.0, false, gedcom.AgeConstraintLiving)).Returns("0y")
	String(gedcom.NewAgeWithYears(0.0, true, gedcom.AgeConstraintLiving)).Returns("0y")
}

func TestAge_Years(t *testing.T) {
	Years := tf.Function(t, gedcom.Age.Years)

	Years(gedcom.NewUnknownAge()).Returns(0)

	Years(gedcom.NewAgeWithYears(23.5, false, gedcom.AgeConstraintUnknown)).Returns(23.5)
	Years(gedcom.NewAgeWithYears(23.5, false, gedcom.AgeConstraintLiving)).Returns(23.5)
	Years(gedcom.NewAgeWithYears(23.5, false, gedcom.AgeConstraintBeforeBirth)).Returns(23.5)
	Years(gedcom.NewAgeWithYears(23.5, false, gedcom.AgeConstraintAfterDeath)).Returns(23.5)

	Years(gedcom.NewAgeWithYears(24.5, true, gedcom.AgeConstraintUnknown)).Returns(24.5)
	Years(gedcom.NewAgeWithYears(24.5, true, gedcom.AgeConstraintLiving)).Returns(24.5)
	Years(gedcom.NewAgeWithYears(24.5, true, gedcom.AgeConstraintBeforeBirth)).Returns(24.5)
	Years(gedcom.NewAgeWithYears(24.5, true, gedcom.AgeConstraintAfterDeath)).Returns(24.5)
}
