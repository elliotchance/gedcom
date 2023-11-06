package html

import (
	"io"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
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

func (c *IndividualNameAndSex) WriteHTMLTo(w io.Writer) (int64, error) {
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
	sexRow := core.NewKeyedTableRow("Sex", sexBadge, true)

	s := core.NewComponents(
		titleRow,
		prefixRow,
		givenNameRow,
		surnamePrefixRow,
		surnameRow,
		suffixRow,
		sexRow,
	)

	return core.NewCard(core.NewText("Name & Sex"), core.CardNoBadgeCount,
		core.NewTable("", s)).WriteHTMLTo(w)
}

func keyedRow(title, value string) *core.KeyedTableRow {
	text := core.NewText(value)
	visible := value != ""

	return core.NewKeyedTableRow(title, text, visible)
}
