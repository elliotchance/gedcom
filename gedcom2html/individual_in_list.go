package main

import (
	"fmt"
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
	birthDate, birthPlace := html.GetBirth(c.individual)
	deathDate, deathPlace := html.GetDeath(c.individual)

	birthPlace = prettyPlaceName(birthPlace)
	deathPlace = prettyPlaceName(deathPlace)

	if birthDate == "" {
		birthDate = "-"
	}

	if deathDate == "" {
		deathDate = "-"
	}

	return fmt.Sprintf(fmt.Sprintf(`
		<tr>
			<td nowrap="nowrap">%s</td>
			<td>%s<br/>%s</td>
			<td>%s<br/>%s</td>
		</tr>`,
		newIndividualLink(c.document, c.individual),
		birthDate, newPlaceLink(c.document, birthPlace),
		deathDate, newPlaceLink(c.document, deathPlace)))
}
