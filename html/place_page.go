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
	indexLetters      []rune
	placesMap         map[string]*place
}

func NewPlacePage(document *gedcom.Document, placeKey string, googleAnalyticsID string, options *PublishShowOptions, indexLetters []rune, placesMap map[string]*place) *PlacePage {
	return &PlacePage{
		document:          document,
		placeKey:          placeKey,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
		indexLetters:      indexLetters,
		placesMap:         placesMap,
	}
}

func (c *PlacePage) WriteHTMLTo(w io.Writer) (int64, error) {
	place := c.placesMap[c.placeKey]

	table := []core.Component{
		core.NewTableHead("Date", "Event", "Individual"),
	}

	for _, node := range place.nodes {
		placeEvent := NewPlaceEvent(c.document, node,
			c.options.LivingVisibility, c.placesMap)
		table = append(table, placeEvent)
	}

	return core.NewPage(
		place.PrettyName,
		core.NewComponents(
			NewPublishHeader(c.document, place.PrettyName, selectedExtraTab,
				c.options, c.indexLetters, c.placesMap),
			core.NewBigTitle(1, core.NewText(place.PrettyName)),
			core.NewSpace(),
			core.NewRow(
				core.NewColumn(core.EntireRow, core.NewTable("", table...)),
			),
		),
		c.googleAnalyticsID,
	).WriteHTMLTo(w)
}
