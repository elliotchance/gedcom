package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
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
	return html.Sprintf(`
		<tr>
			<td>%s</td>
		</tr>`, newSourceLink(c.source))
}
