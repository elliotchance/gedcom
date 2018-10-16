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

	text := newOcticon("location", "").String() + c.place

	return html.NewLink(text, pagePlace(c.document, c.place)).String()
}
