// Package gedcomdiff is a command line tool for comparing GEDCOM files and
// producing a HTML report.
package main

import (
	"flag"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/util"
	"log"
	"os"
)

var (
	optionLeftGedcomFile    string
	optionRightGedcomFile   string
	optionOutputFile        string
	optionSubset            bool
	optionGoogleAnalyticsID string
)

var filterFlags = &util.FilterFlags{}

func main() {
	parseCLIFlags()

	leftGedcom, err := gedcom.NewDocumentFromGEDCOMFile(optionLeftGedcomFile)
	if err != nil {
		log.Fatal(err)
	}

	rightGedcom, err := gedcom.NewDocumentFromGEDCOMFile(optionRightGedcomFile)
	if err != nil {
		log.Fatal(err)
	}

	// Run compare.
	leftIndividuals := leftGedcom.Individuals()
	rightIndividuals := rightGedcom.Individuals()

	log.Printf("Writing %s...", optionOutputFile)

	out, err := os.Create(optionOutputFile)
	if err != nil {
		log.Fatal(err)
	}

	options := gedcom.NewSimilarityOptions()
	comparisons := leftIndividuals.Compare(rightIndividuals, options)

	page := newDiffPage(comparisons, options, filterFlags, optionGoogleAnalyticsID)
	out.Write([]byte(page.String()))
}

func parseCLIFlags() {
	// Input files. Must be provided.
	flag.StringVar(&optionLeftGedcomFile, "left-gedcom", "", "Left GEDCOM file.")
	flag.StringVar(&optionRightGedcomFile, "right-gedcom", "", "Right GEDCOM file.")
	flag.StringVar(&optionOutputFile, "output", "", "Output file.")
	flag.BoolVar(&optionSubset, "subset", false, "When -subset is enabled the "+
		"right side will be considered a smaller part of the larger left "+
		"side. This means that individuals that entirely exist on the left "+
		"side will not be included.")
	flag.StringVar(&optionGoogleAnalyticsID, "google-analytics-id", "",
		"The Google Analytics ID, like 'UA-78454410-2'.")

	filterFlags.SetupCLI()

	flag.Parse()
}
