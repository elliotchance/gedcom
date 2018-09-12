// Package gedcom2html is a command line tool for rendering a GEDCOM file into
// HTML pages that shared and published easily.
package main

import (
	"flag"
	"fmt"
	"github.com/elliotchance/gedcom"
	"log"
	"os"
)

var (
	optionGedcomFile        string
	optionOutputDir         string
	optionGoogleAnalyticsID string
)

func main() {
	flag.StringVar(&optionGedcomFile, "gedcom", "", "Input GEDCOM file.")
	flag.StringVar(&optionOutputDir, "output-dir", ".", "Output directory. It"+
		" will use the current directory if output-dir is not provided. "+
		"Output files will only be added or replaced. Existing files will not"+
		" be deleted.")
	flag.StringVar(&optionGoogleAnalyticsID, "google-analytics-id", "",
		"The Google Analytics ID, like 'UA-78454410-2'.")
	flag.Parse()

	file, err := os.Open(optionGedcomFile)
	if err != nil {
		log.Fatal(err)
	}

	decoder := gedcom.NewDecoder(file)
	document, err := decoder.Decode()
	if err != nil {
		log.Fatal(err)
	}

	// Create the pages.
	for _, letter := range getIndexLetters(document) {
		createFile(pageIndividuals(letter),
			newIndividualListPage(document, letter, optionGoogleAnalyticsID))
	}

	for _, individual := range getIndividuals(document) {
		if individual.IsLiving() {
			continue
		}

		page := newIndividualPage(document, individual, optionGoogleAnalyticsID)
		createFile(pageIndividual(document, individual), page)
	}

	createFile(pagePlaces(), newPlaceListPage(document, optionGoogleAnalyticsID))

	for key, place := range getPlaces(document) {
		page := newPlacePage(document, key, optionGoogleAnalyticsID)
		createFile(pagePlace(document, place.prettyName), page)
	}

	createFile(pageFamilies(), newFamilyListPage(document, optionGoogleAnalyticsID))

	createFile(pageSources(), newSourceListPage(document, optionGoogleAnalyticsID))

	for _, source := range document.Sources() {
		page := newSourcePage(document, source, optionGoogleAnalyticsID)
		createFile(pageSource(source), page)
	}

	createFile(pageStatistics(), newStatisticsPage(document, optionGoogleAnalyticsID))
}

func createFile(name string, contents fmt.Stringer) {
	path := fmt.Sprintf("%s/%s", optionOutputDir, name)
	log.Printf("Writing %s...", path)

	out, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	out.Write([]byte(contents.String()))

	out.Close()
}
