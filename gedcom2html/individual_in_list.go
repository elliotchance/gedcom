package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

// individualInList is a single row in the table of individuals on the list
// page.
type individualInList struct {
	individual *gedcom.IndividualNode
	document   *gedcom.Document
}

func newIndividualInList(document *gedcom.Document, individual *gedcom.IndividualNode) *individualInList {
	return &individualInList{
		individual: individual,
		document:   document,
	}
}

func (c *individualInList) String() string {
	birth := gedcom.First(c.individual.Births())
	birthDate := gedcom.String(gedcom.First(gedcom.Dates(birth)))
	birthPlace := gedcom.String(gedcom.First(gedcom.Places(birth)))

	death := gedcom.First(c.individual.Deaths())
	deathDate := gedcom.String(gedcom.First(gedcom.Dates(death)))
	deathPlace := gedcom.String(gedcom.First(gedcom.Places(death)))

	birthPlace = prettyPlaceName(birthPlace)
	deathPlace = prettyPlaceName(deathPlace)

	if birthDate == "" {
		birthDate = "-"
	}

	if deathDate == "" {
		deathDate = "-"
	}

	return html.Sprintf(`
		<tr>
			<td nowrap="nowrap">%s</td>
			<td>%s<br/>%s</td>
			<td>%s<br/>%s</td>
		</tr>`,
		newIndividualLink(c.document, c.individual),
		birthDate, newPlaceLink(c.document, birthPlace),
		deathDate, newPlaceLink(c.document, deathPlace))
}
