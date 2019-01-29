package html

import (
	"github.com/elliotchance/gedcom"
	"io"
	"sort"
)

// PlaceListPage lists all places.
type PlaceListPage struct {
	document          *gedcom.Document
	googleAnalyticsID string
	options           PublishShowOptions
}

func NewPlaceListPage(document *gedcom.Document, googleAnalyticsID string, options PublishShowOptions) *PlaceListPage {
	return &PlaceListPage{
		document:          document,
		googleAnalyticsID: googleAnalyticsID,
		options:           options,
	}
}

func (c *PlaceListPage) WriteTo(w io.Writer) (int64, error) {
	table := []Component{}

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
			anchor := NewAnchor(place.country)
			heading := NewHeading(2, "", NewText(place.country))
			components := NewComponents(anchor, heading)
			cell := NewTableCell(components)
			countryTitle := NewTableRow(cell)
			table = append(table, countryTitle)
		}

		placeEntry := NewPlaceInList(c.document, place)
		table = append(table, placeEntry)

		lastCountry = place.country
	}

	// Render
	pills := []Component{}
	for _, country := range countries.Strings() {
		pills = append(pills, NewNavLink(country, "#"+country, false))
	}

	return NewPage("Places", NewComponents(
		NewPublishHeader(c.document, "", selectedPlacesTab, c.options),
		NewNavPillsRow(pills),
		NewSpace(),
		NewRow(
			NewColumn(EntireRow, NewTable("", table...)),
		),
	), c.googleAnalyticsID).WriteTo(w)
}
