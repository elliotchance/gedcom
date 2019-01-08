package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type PlaceLink struct {
	document *gedcom.Document
	place    string
}

func NewPlaceLink(document *gedcom.Document, place string) *PlaceLink {
	return &PlaceLink{
		document: document,
		place:    place,
	}
}

func (c *PlaceLink) WriteTo(w io.Writer) (int64, error) {
	if c.place == "" {
		return writeNothing()
	}

	text := NewComponents(NewOcticon("location", ""), NewText(c.place))

	return NewLink(text, PagePlace(c.document, c.place)).WriteTo(w)
}
