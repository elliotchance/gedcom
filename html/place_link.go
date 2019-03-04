package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
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

func (c *PlaceLink) WriteHTMLTo(w io.Writer) (int64, error) {
	if c.place == "" {
		return writeNothing()
	}

	icon := core.NewOcticon("location", "")
	text := core.NewComponents(icon, core.NewText(c.place))

	return core.NewLink(text, PagePlace(c.document, c.place)).WriteHTMLTo(w)
}
