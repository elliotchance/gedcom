package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

// IndividualNameAndSex shows the name parts and sex of an individual in the
// "Name & Sex" section of the individuals page.
type IndividualNameAndSex struct {
	individual *gedcom.IndividualNode
}

func NewIndividualNameAndSex(individual *gedcom.IndividualNode) *IndividualNameAndSex {
	return &IndividualNameAndSex{
		individual: individual,
	}
}

func (c *IndividualNameAndSex) WriteTo(w io.Writer) (int64, error) {
	primaryName := c.individual.Names()[0]
	title := primaryName.Title()
	prefix := primaryName.Prefix()
	name := primaryName.GivenName()
	surnamePrefix := primaryName.SurnamePrefix()
	surname := primaryName.Surname()
	suffix := primaryName.Suffix()

	titleRow := NewKeyedTableRow("Title", NewText(title), title != "")
	prefixRow := NewKeyedTableRow("Prefix", NewText(prefix), prefix != "")
	givenNameRow := NewKeyedTableRow("Given Name", NewText(name), name != "")
	surnamePrefixRow := NewKeyedTableRow("Surname Prefix", NewText(surnamePrefix), surnamePrefix != "")
	surnameRow := NewKeyedTableRow("Surname", NewText(surname), surname != "")
	suffixRow := NewKeyedTableRow("Suffix", NewText(suffix), suffix != "")
	sexRow := NewKeyedTableRow("Sex", NewSexBadge(c.individual.Sex()), true)

	s := NewComponents(
		titleRow,
		prefixRow,
		givenNameRow,
		surnamePrefixRow,
		surnameRow,
		suffixRow,
		sexRow,
	)

	return NewCard("Name & Sex", noBadgeCount, NewTable("", s)).WriteTo(w)
}
