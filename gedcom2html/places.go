package main

import (
	"github.com/elliotchance/gedcom"
	"strings"
)

var placesMap map[string]*place

type place struct {
	prettyName string
	country    string
	nodes      []gedcom.Node
}

func prettyPlaceName(s string) string {
	s = strings.Replace(s, ",,", ",", -1)
	s = strings.Replace(s, ",,", ",", -1)
	s = strings.Replace(s, ",", ", ", -1)
	s = strings.Trim(s, ", ")

	return strings.TrimSpace(s)
}

func getPlaces(document *gedcom.Document) map[string]*place {
	if placesMap == nil {
		placesMap = map[string]*place{}

		// Get all of the unique place names.
		for placeTag, node := range document.Places() {
			prettyName := prettyPlaceName(placeTag.Value())

			if prettyName == "" {
				prettyName = "(none)"
			}

			key := alnumOrDashRegexp.
				ReplaceAllString(strings.ToLower(prettyName), "-")

			if _, ok := placesMap[key]; !ok {
				country := placeTag.Country()
				if country == "" {
					country = "(unknown)"
				}

				placesMap[key] = &place{
					prettyName: prettyName,
					country:    country,
					nodes:      []gedcom.Node{},
				}
			}

			placesMap[key].nodes = append(placesMap[key].nodes, node)
		}
	}

	return placesMap
}
