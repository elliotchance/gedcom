package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

// individualNameAndSex shows the name parts and sex of an individual in the
// "Name & Sex" section of the individuals page.
type individualNameAndSex struct {
	individual *gedcom.IndividualNode
}

func newIndividualNameAndSex(individual *gedcom.IndividualNode) *individualNameAndSex {
	return &individualNameAndSex{
		individual: individual,
	}
}

func (c *individualNameAndSex) String() string {
	primaryName := c.individual.Names()[0]
	title := primaryName.Title()
	prefix := primaryName.Prefix()
	name := primaryName.GivenName()
	surnamePrefix := primaryName.SurnamePrefix()
	surname := primaryName.Surname()
	suffix := primaryName.Suffix()

	titleRow := newKeyedTableRow("Title", title, title != "")
	prefixRow := newKeyedTableRow("Prefix", prefix, prefix != "")
	givenNameRow := newKeyedTableRow("Given Name", name, name != "")
	surnamePrefixRow := newKeyedTableRow("Surname Prefix", surnamePrefix, surnamePrefix != "")
	surnameRow := newKeyedTableRow("Surname", surname, surname != "")
	suffixRow := newKeyedTableRow("Suffix", suffix, suffix != "")
	sexRow := newKeyedTableRow("Sex", newSexBadge(c.individual.Sex()).String(), true)

	s := html.NewComponents(
		titleRow,
		prefixRow,
		givenNameRow,
		surnamePrefixRow,
		surnameRow,
		suffixRow,
		sexRow,
	)

	return newCard("Name & Sex", noBadgeCount, html.NewTable("", s)).String()
}
