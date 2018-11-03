package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"testing"
)

func TestAge_String(t *testing.T) {
	String := tf.Function(t, (*Age).String)

	unknown := gedcom.NewUnknownAge()
	age0 := gedcom.NewAgeWithYears(0, true, gedcom.AgeConstraintUnknown)
	age1 := gedcom.NewAgeWithYears(43.2, false, gedcom.AgeConstraintUnknown)
	age2 := gedcom.NewAgeWithYears(44.2, false, gedcom.AgeConstraintLiving)
	age3 := gedcom.NewAgeWithYears(45.2, false, gedcom.AgeConstraintBeforeBirth)
	age4 := gedcom.NewAgeWithYears(46.2, false, gedcom.AgeConstraintAfterDeath)
	age5 := gedcom.NewAgeWithYears(10, true, gedcom.AgeConstraintUnknown)
	age6 := gedcom.NewAgeWithYears(11, true, gedcom.AgeConstraintLiving)
	age7 := gedcom.NewAgeWithYears(12.113, true, gedcom.AgeConstraintBeforeBirth)
	age8 := gedcom.NewAgeWithYears(13.5, true, gedcom.AgeConstraintAfterDeath)

	String(NewAge(unknown, unknown)).Returns(``)
	String(NewAge(age1, unknown)).Returns(`after 43y 2m`)
	String(NewAge(unknown, age7)).Returns(`until ~ 12y 1m`)
	String(NewAge(age7, age3)).Returns(`from ~ 12y 1m to 45y 2m`)
	String(NewAge(age2, age3)).Returns(`44y 2m`)
	String(NewAge(age5, age6)).Returns(`~ 10y`)
	String(NewAge(age0, age0)).Returns(`0y`)

	String(NewAge(age1, age1)).Returns(`43y 2m`)
	String(NewAge(age2, age2)).Returns(`44y 2m`)
	String(NewAge(age3, age3)).Returns(`45y 2m`)
	String(NewAge(age4, age4)).Returns(``)
	String(NewAge(age5, age5)).Returns(`~ 10y`)
	String(NewAge(age6, age6)).Returns(`~ 11y`)
	String(NewAge(age7, age7)).Returns(`~ 12y 1m`)
	String(NewAge(age8, age8)).Returns(``)
}
