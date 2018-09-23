package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

type placeLink struct {
	document *gedcom.Document
	place    string
}

func newPlaceLink(document *gedcom.Document, place string) *placeLink {
	return &placeLink{
		document: document,
		place:    place,
	}
}

func (c *placeLink) String() string {
	if c.place == "" {
		return ""
	}

	return html.Sprintf(`<a href="%s">%s%s</a>`,
		pagePlace(c.document, c.place), newOcticon("location", ""), c.place)
}
