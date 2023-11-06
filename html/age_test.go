package html_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html"
)

func TestAge_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "Age")

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

	c(html.NewAge(unknown, unknown)).Returns(``)
	c(html.NewAge(age1, unknown)).Returns(`after 43y 2m`)
	c(html.NewAge(unknown, age7)).Returns(`until ~ 12y 1m`)
	c(html.NewAge(age7, age3)).Returns(`from ~ 12y 1m to 45y 2m`)
	c(html.NewAge(age2, age3)).Returns(`44y 2m`)
	c(html.NewAge(age5, age6)).Returns(`~ 10y`)
	c(html.NewAge(age0, age0)).Returns(`0y`)

	c(html.NewAge(age1, age1)).Returns(`43y 2m`)
	c(html.NewAge(age2, age2)).Returns(`44y 2m`)
	c(html.NewAge(age3, age3)).Returns(`45y 2m`)
	c(html.NewAge(age4, age4)).Returns(``)
	c(html.NewAge(age5, age5)).Returns(`~ 10y`)
	c(html.NewAge(age6, age6)).Returns(`~ 11y`)
	c(html.NewAge(age7, age7)).Returns(`~ 12y 1m`)
	c(html.NewAge(age8, age8)).Returns(``)
}
