package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
)

type sourceInList struct {
	document *gedcom.Document
	source   *gedcom.SourceNode
}

func newSourceInList(document *gedcom.Document, source *gedcom.SourceNode) *sourceInList {
	return &sourceInList{
		document: document,
		source:   source,
	}
}

func (c *sourceInList) String() string {
	return fmt.Sprintf(fmt.Sprintf(`
		<tr>
			<td>%s</td>
		</tr>`, newSourceLink(c.source)))
}
