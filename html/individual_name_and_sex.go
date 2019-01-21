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

	titleRow := keyedRow("Title", title)
	prefixRow := keyedRow("Prefix", prefix)
	givenNameRow := keyedRow("Given Name", name)
	surnamePrefixRow := keyedRow("Surname Prefix", surnamePrefix)
	surnameRow := keyedRow("Surname", surname)
	suffixRow := keyedRow("Suffix", suffix)

	sexBadge := NewSexBadge(c.individual.Sex())
	sexRow := NewKeyedTableRow("Sex", sexBadge, true)

	s := NewComponents(
		titleRow,
		prefixRow,
		givenNameRow,
		surnamePrefixRow,
		surnameRow,
		suffixRow,
		sexRow,
	)

	return NewCard(NewText("Name & Sex"), noBadgeCount, NewTable("", s)).WriteTo(w)
}

func keyedRow(title, value string) *KeyedTableRow {
	text := NewText(value)
	visible := value != ""

	return NewKeyedTableRow(title, text, visible)
}
