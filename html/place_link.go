package html

import (
	"io"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html/core"
)

type PlaceLink struct {
	document  *gedcom.Document
	place     string
	placesMap map[string]*place
}

func NewPlaceLink(document *gedcom.Document, place string, placesMap map[string]*place) *PlaceLink {
	return &PlaceLink{
		document:  document,
		place:     place,
		placesMap: placesMap,
	}
}

func (c *PlaceLink) WriteHTMLTo(w io.Writer) (int64, error) {
	if c.place == "" {
		return writeNothing()
	}

	icon := core.NewOcticon("location", "")
	text := core.NewComponents(icon, core.NewText(c.place))

	return core.NewLink(text, PagePlace(c.place, c.placesMap)).WriteHTMLTo(w)
}
