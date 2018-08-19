package main

import (
	"flag"
	"github.com/elliotchance/gedcom"
	"log"
	"os"
)

var (
	optionLeftGedcomFile  string
	optionRightGedcomFile string
	optionOutputFile      string
	optionNoPlaces        bool
	optionHideSame        bool
)

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
	comparisons := leftIndividuals.Compare(leftGedcom, rightGedcom, rightIndividuals, options)

	out.Write([]byte(newDiffPage(comparisons, options, leftGedcom, rightGedcom, !optionNoPlaces, optionHideSame).String()))
}

func parseCLIFlags() {
	// Input files. Must be provided.
	flag.StringVar(&optionLeftGedcomFile, "left-gedcom", "", "Left GEDCOM file.")
	flag.StringVar(&optionRightGedcomFile, "right-gedcom", "", "Right GEDCOM file.")
	flag.StringVar(&optionOutputFile, "output", "", "Output file.")
	flag.BoolVar(&optionNoPlaces, "no-places", false, "Do not include places.")
	flag.BoolVar(&optionHideSame, "hide-same", false, "Hide equal values.")

	flag.Parse()
}
