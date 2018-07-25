package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
)

type familyInList struct {
	document *gedcom.Document
	family   *gedcom.FamilyNode
}

func newFamilyInList(document *gedcom.Document, family *gedcom.FamilyNode) *familyInList {
	return &familyInList{
		document:   document,
		family: family,
	}
}

func (c *familyInList) String() string {
	date := "-"
	if n := c.family.FirstNodeWithTagPath(gedcom.TagMarriage, gedcom.TagDate); n != nil {
		date = n.Value()
	}

	return fmt.Sprintf(fmt.Sprintf(`
		<tr>
			<td>%s</td>
			<td nowrap="nowrap" class="text-center">%s</td>
			<td>%s</td>
		</tr>`,
		newIndividualLink(c.document, c.family.Husband(c.document)),
		date,
		newIndividualLink(c.document, c.family.Wife(c.document))))
}
