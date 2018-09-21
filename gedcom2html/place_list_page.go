package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
	"github.com/elliotchance/gedcom/util"
	"sort"
)

// placeListPage lists all places.
type placeListPage struct {
	document          *gedcom.Document
	googleAnalyticsID string
}

func newPlaceListPage(document *gedcom.Document, googleAnalyticsID string) *placeListPage {
	return &placeListPage{
		document:          document,
		googleAnalyticsID: googleAnalyticsID,
	}
}

func (c *placeListPage) String() string {
	table := []fmt.Stringer{}

	places := getPlaces(c.document)

	// Get all countries
	countries := []string{}
	for _, place := range places {
		if util.StringSliceContains(countries, place.country) {
			continue
		}

		countries = append(countries, place.country)
	}

	sortedPlaces := []*place{}
	for _, placeName := range places {
		sortedPlaces = append(sortedPlaces, placeName)
	}

	sort.SliceStable(sortedPlaces, func(i, j int) bool {
		if sortedPlaces[i].country != sortedPlaces[j].country {
			return sortedPlaces[i].country < sortedPlaces[j].country
		}

		return sortedPlaces[i].prettyName < sortedPlaces[j].prettyName
	})

	lastCountry := ""

	for _, place := range sortedPlaces {
		if lastCountry != place.country {
			anchor := html.NewAnchor(place.country)
			heading := html.NewHeading(2, "", place.country)
			components := html.NewComponents(anchor, heading)
			cell := html.NewTableCell("", components)
			countryTitle := html.NewTableRow(cell)
			table = append(table, countryTitle)
		}

		placeEntry := newPlaceInList(c.document, place)
		table = append(table, placeEntry)

		lastCountry = place.country
	}

	// Sort countries
	sort.Strings(countries)

	// Render
	pills := []fmt.Stringer{}
	for _, country := range countries {
		pills = append(pills, newNavLink(country, "#"+country, false))
	}

	return html.NewPage("Places", html.NewComponents(
		newHeader(c.document, "", selectedPlacesTab),
		newNavPillsRow(pills),
		html.NewSpace(),
		html.NewRow(
			html.NewColumn(html.EntireRow, html.NewTable("", table...)),
		),
	), c.googleAnalyticsID).String()
}
