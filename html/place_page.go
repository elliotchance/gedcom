package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
)

type PlacePage struct {
	document          *gedcom.Document
	placeKey          string
	googleAnalyticsID string
	options           *PublishShowOptions
}

func NewPlacePage(document *gedcom.Document, placeKey string, googleAnalyticsID string, options *PublishShowOptions) *PlacePage {
	return &PlacePage{
		document:          document,
		placeKey:          placeKey,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
	}
}

func (c *PlacePage) WriteHTMLTo(w io.Writer) (int64, error) {
	place := GetPlaces(c.document)[c.placeKey]

	table := []core.Component{
		core.NewTableHead("Date", "Event", "Individual"),
	}

	for _, node := range place.nodes {
		placeEvent := NewPlaceEvent(c.document, node, c.options.LivingVisibility)
		table = append(table, placeEvent)
	}

	return core.NewPage(
		place.PrettyName,
		core.NewComponents(
			NewPublishHeader(c.document, place.PrettyName, selectedExtraTab, c.options),
			core.NewBigTitle(1, core.NewText(place.PrettyName)),
			core.NewSpace(),
			core.NewRow(
				core.NewColumn(core.EntireRow, core.NewTable("", table...)),
			),
		),
		c.googleAnalyticsID,
	).WriteHTMLTo(w)
}
