package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
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
	return fmt.Sprintf(`
		<a href="%s">%s</a>`,
		pagePlace(c.document, c.place), c.place)
}
