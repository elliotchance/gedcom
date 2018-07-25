package main

import (
	"flag"
	"fmt"
	"github.com/elliotchance/gedcom"
	"log"
	"os"
)

var (
	optionGedcomFile string
	optionOutputDir  string
)

func main() {
	flag.StringVar(&optionGedcomFile, "gedcom", "", "Input GEDCOM file.")
	flag.StringVar(&optionOutputDir, "output-dir", ".", "Output directory. It"+
		" will use the current directory if output-dir is not provided. "+
		"Output files will only be added or replaced. Existing files will not"+
		" be deleted.")
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
	createFile(pageIndividuals(), newIndividualListPage(document))

	for _, individual := range getIndividuals(document) {
		if individual.IsLiving() {
			continue
		}

		page := newIndividualPage(document, individual)
		createFile(pageIndividual(document, individual), page)
	}

	createFile(pagePlaces(), newPlaceListPage(document))

	for key, place := range getPlaces(document) {
		page := newPlacePage(document, key)
		createFile(pagePlace(document, place.prettyName), page)
	}

	createFile(pageFamilies(), newFamilyListPage(document))

	createFile(pageSources(), newSourceListPage(document))

	for _, source := range document.Sources() {
		page := newSourcePage(document, source)
		createFile(pageSource(source), page)
	}
}

func createFile(name string, contents fmt.Stringer) {
	path := fmt.Sprintf("%s/%s", optionOutputDir, name)
	log.Printf("Writing %s...", path)

	out, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	out.Write([]byte(contents.String()))
}
