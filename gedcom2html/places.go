package main

import (
	"github.com/elliotchance/gedcom"
	"sort"
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

		for key := range placesMap {
			// Make sure the events are sorted otherwise the pages will be
			// different.
			sort.Slice(placesMap[key].nodes, func(i, j int) bool {
				left := placesMap[key].nodes[i]
				right := placesMap[key].nodes[j]

				// Years.
				leftYears := gedcom.Years(left)
				rightYears := gedcom.Years(right)

				if leftYears != rightYears {
					return leftYears < rightYears
				}

				// Tag.
				leftTag := left.Tag().String()
				rightTag := right.Tag().String()

				if leftTag != rightTag {
					return leftTag < rightTag
				}

				// Individual name.
				leftIndividual := individualForNode(left)
				rightIndividual := individualForNode(right)

				if leftIndividual != nil && rightIndividual != nil {
					leftName := gedcom.String(leftIndividual.Name())
					rightName := gedcom.String(rightIndividual.Name())

					return leftName < rightName
				}

				// Value.
				valueLeft := gedcom.Value(left)
				valueRight := gedcom.Value(right)

				return valueLeft < valueRight
			})
		}
	}

	return placesMap
}
