package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
	"sort"
)

// PlaceListPage lists all places.
type PlaceListPage struct {
	document          *gedcom.Document
	googleAnalyticsID string
	options           *PublishShowOptions
	indexLetters      []rune
}

func NewPlaceListPage(document *gedcom.Document, googleAnalyticsID string, options *PublishShowOptions, indexLetters []rune) *PlaceListPage {
	return &PlaceListPage{
		document:          document,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
		indexLetters:      indexLetters,
	}
}

func (c *PlaceListPage) WriteHTMLTo(w io.Writer) (int64, error) {
	table := []core.Component{}

	places := GetPlaces(c.document)

	// Get all countries
	countries := gedcom.NewStringSet()
	for _, place := range places {
		countries.Add(place.country)
	}

	sortedPlaces := []*place{}
	for _, placeName := range places {
		sortedPlaces = append(sortedPlaces, placeName)
	}

	sort.SliceStable(sortedPlaces, func(i, j int) bool {
		if sortedPlaces[i].country != sortedPlaces[j].country {
			return sortedPlaces[i].country < sortedPlaces[j].country
		}

		return sortedPlaces[i].PrettyName < sortedPlaces[j].PrettyName
	})

	lastCountry := ""

	for _, place := range sortedPlaces {
		if lastCountry != place.country {
			anchor := core.NewAnchor(place.country)
			heading := core.NewHeading(2, "", core.NewText(place.country))
			components := core.NewComponents(anchor, heading)
			cell := core.NewTableCell(components)
			countryTitle := core.NewTableRow(cell)
			table = append(table, countryTitle)
		}

		placeEntry := NewPlaceInList(c.document, place)
		table = append(table, placeEntry)

		lastCountry = place.country
	}

	// Render
	pills := []core.Component{}
	for _, country := range countries.Strings() {
		pills = append(pills, core.NewNavLink(country, "#"+country, false))
	}

	return core.NewPage("Places", core.NewComponents(
		NewPublishHeader(c.document, "", selectedPlacesTab, c.options,
			c.indexLetters),
		core.NewNavPillsRow(pills),
		core.NewSpace(),
		core.NewRow(
			core.NewColumn(core.EntireRow, core.NewTable("", table...)),
		),
	), c.googleAnalyticsID).WriteHTMLTo(w)
}
