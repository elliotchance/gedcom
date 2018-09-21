package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
)

type placePage struct {
	document          *gedcom.Document
	placeKey          string
	googleAnalyticsID string
}

func newPlacePage(document *gedcom.Document, placeKey string, googleAnalyticsID string) *placePage {
	return &placePage{
		document:          document,
		placeKey:          placeKey,
		googleAnalyticsID: googleAnalyticsID,
	}
}

func (c *placePage) String() string {
	place := getPlaces(c.document)[c.placeKey]

	table := []fmt.Stringer{
		html.NewTableHead("Date", "Event", "Individual"),
	}

	for _, node := range place.nodes {
		placeEvent := newPlaceEvent(c.document, node)
		table = append(table, placeEvent)
	}

	return html.NewPage(
		place.prettyName,
		html.NewComponents(
			newHeader(c.document, place.prettyName, selectedExtraTab),
			html.NewBigTitle(place.prettyName),
			html.NewSpace(),
			html.NewRow(
				html.NewColumn(html.EntireRow, html.NewTable("", table...)),
			),
		),
		c.googleAnalyticsID,
	).String()
}
