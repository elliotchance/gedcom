package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
)

const symbolLetter = '#'

func pageIndividuals(firstLetter rune) string {
	if firstLetter == symbolLetter {
		return "individuals-symbol.html"
	}

	return fmt.Sprintf("individuals-%c.html", firstLetter)
}

func pageIndividual(document *gedcom.Document, individual *gedcom.IndividualNode) string {
	individuals := getIndividuals(document)

	for key, value := range individuals {
		if value.Is(individual) {
			return fmt.Sprintf("%s.html", key)
		}
	}

	return "#"
}

func pagePlaces() string {
	return "places.html"
}

func pagePlace(document *gedcom.Document, place string) string {
	places := getPlaces(document)

	for key, value := range places {
		if value.prettyName == place {
			return fmt.Sprintf("%s.html", key)
		}
	}

	return "#"
}

func pageFamilies() string {
	return "families.html"
}

func pageSources() string {
	return "sources.html"
}

func pageSource(source *gedcom.SourceNode) string {
	return fmt.Sprintf("%s.html", source.Pointer())
}

func pageStatistics() string {
	return "statistics.html"
}

func pageSurnames() string {
	return "surnames.html"
}
