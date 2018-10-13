// Package gedcomdiff is a command line tool for comparing GEDCOM files and
// producing a HTML report.
package main

import (
	"flag"
	"github.com/cheggaaa/pb"
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
	optionProgress          bool
	optionJobs              int
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

	out, err := os.Create(optionOutputFile)
	if err != nil {
		log.Fatal(err)
	}

	var comparisons []gedcom.IndividualComparison
	compareOptions := &gedcom.IndividualNodesCompareOptions{
		SimilarityOptions: gedcom.NewSimilarityOptions(),
		Notifier:          make(chan gedcom.CompareProgress),
		NotifierStep:      100,
		Jobs:              optionJobs,
	}

	if optionProgress {
		go func() {
			comparisons = leftIndividuals.Compare(rightIndividuals, compareOptions)
		}()

		progressBar := pb.StartNew(0)

		for {
			n, ok := <-compareOptions.Notifier
			if !ok {
				break
			}

			progressBar.SetTotal(n.Total)
			progressBar.SetCurrent(n.Done)
		}

		progressBar.Finish()
	} else {
		comparisons = leftIndividuals.Compare(rightIndividuals, compareOptions)
	}

	page := newDiffPage(comparisons, compareOptions.SimilarityOptions,
		filterFlags, optionGoogleAnalyticsID)
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
	flag.BoolVar(&optionProgress, "progress", false, "Show progress bar.")
	flag.IntVar(&optionJobs, "jobs", 1, "Number of jobs to run in parallel.")

	filterFlags.SetupCLI()

	flag.Parse()
}
