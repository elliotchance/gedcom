package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

type PlacePage struct {
	document          *gedcom.Document
	placeKey          string
	googleAnalyticsID string
	options           PublishShowOptions
}

func NewPlacePage(document *gedcom.Document, placeKey string, googleAnalyticsID string, options PublishShowOptions) *PlacePage {
	return &PlacePage{
		document:          document,
		placeKey:          placeKey,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
	}
}

func (c *PlacePage) WriteTo(w io.Writer) (int64, error) {
	place := GetPlaces(c.document)[c.placeKey]

	table := []Component{
		NewTableHead("Date", "Event", "Individual"),
	}

	for _, node := range place.nodes {
		placeEvent := NewPlaceEvent(c.document, node)
		table = append(table, placeEvent)
	}

	return NewPage(
		place.PrettyName,
		NewComponents(
			NewPublishHeader(c.document, place.PrettyName, selectedExtraTab, c.options),
			NewBigTitle(1, NewText(place.PrettyName)),
			NewSpace(),
			NewRow(
				NewColumn(EntireRow, NewTable("", table...)),
			),
		),
		c.googleAnalyticsID,
	).WriteTo(w)
}
