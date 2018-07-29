package main

import (
	"github.com/elliotchance/gedcom"
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

	s := newComponents(
		newKeyedTableRow("Title", primaryName.Title(), primaryName.Title() != ""),
		newKeyedTableRow("Prefix", primaryName.Prefix(), primaryName.Prefix() != ""),
		newKeyedTableRow("Given Name", primaryName.GivenName(), primaryName.GivenName() != ""),
		newKeyedTableRow("Surname Prefix", primaryName.SurnamePrefix(), primaryName.SurnamePrefix() != ""),
		newKeyedTableRow("Surname", primaryName.Surname(), primaryName.Surname() != ""),
		newKeyedTableRow("Suffix", primaryName.Suffix(), primaryName.Suffix() != ""),
		newKeyedTableRow("Sex", newSexBadge(c.individual.Sex()).String(), true),
	)

	return newCard("Name & Sex", noBadgeCount, newTable("", s)).String()
}
