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
	"strings"
)

var (
	optionLeftGedcomFile            string
	optionRightGedcomFile           string
	optionOutputFile                string
	optionSubset                    bool
	optionGoogleAnalyticsID         string
	optionProgress                  bool
	optionJobs                      int
	optionMinimumSimilarity         float64
	optionMinimumWeightedSimilarity float64
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

	similarityOptions := gedcom.NewSimilarityOptions()
	similarityOptions.MinimumWeightedSimilarity = optionMinimumWeightedSimilarity
	similarityOptions.MinimumSimilarity = optionMinimumSimilarity

	compareOptions := &gedcom.IndividualNodesCompareOptions{
		SimilarityOptions: similarityOptions,
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

	pageString := page.String()
	out.Write([]byte(pageString))
}

func parseCLIFlags() {
	// Input files. Must be provided.
	flag.StringVar(&optionLeftGedcomFile, "left-gedcom", "",
		"Required. Left GEDCOM file.")

	flag.StringVar(&optionRightGedcomFile, "right-gedcom", "",
		"Required. Right GEDCOM file.")

	flag.StringVar(&optionOutputFile, "output", "", "Output file.")

	flag.BoolVar(&optionSubset, "subset", false, CLIDescription(`When -subset is
		enabled the right side will be considered a smaller part of the larger
		left side. This means that individuals that entirely exist on the left
		side will not be included.`))

	flag.StringVar(&optionGoogleAnalyticsID, "google-analytics-id", "",
		"The Google Analytics ID, like 'UA-78454410-2'.")

	flag.BoolVar(&optionProgress, "progress", false, "Show progress bar.")

	flag.IntVar(&optionJobs, "jobs", 1, CLIDescription(`Number of jobs to run in
		parallel. If you are comparing large trees this will make the process
		faster but will consume more CPU.`))

	flag.Float64Var(&optionMinimumWeightedSimilarity,
		"minimum-weighted-similarity", gedcom.DefaultMinimumSimilarity,
		CLIDescription(`The weighted minimum similarity is the threshold for
			whether two individuals should be the seen as the same person when
			the surrounding immediate family is taken into consideration.

			This value must be between 0 and 1 and is the primary way to adjust
			the sensitivity of matches. It is best to also set
			"-minimum-similarity" to the same value.

			A higher value means you will get less matches but they will be of
			higher quality. If you are comparing trees that do not share many of
			the same individuals you should consider raising this to prevent
			false-positives.`))

	flag.Float64Var(&optionMinimumSimilarity,
		"minimum-similarity", gedcom.DefaultMinimumSimilarity,
		CLIDescription(`The minimum similarity is the threshold for matching
			individuals as the same person. This is used to compare only the
			individual (not surrounding family) like spouses and children.

			This value must be between 0 and 1 and should be set to the same
			value as "minimum-weighted-similarity" if you are unsure.`))

	filterFlags.SetupCLI()

	flag.Parse()
}

func CLIDescription(s string) (r string) {
	lines := strings.Split(s, "\n")

	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			r += "\n\n"
		} else {
			r += strings.Replace(line, "\t", "", -1) + " "
		}
	}

	return WrapToMargin(r, 80)
}

func WrapToMargin(s string, width int) (r string) {
	lines := strings.Split(s, "\n")

	for _, line := range lines {
		words := strings.Split(line, " ")
		newLine := ""

		for _, word := range words {
			if len(newLine)+len(word)+1 > width {
				r += strings.TrimSpace(newLine) + "\n"
				newLine = word
			} else {
				newLine += " " + word
			}
		}

		r += strings.TrimSpace(newLine) + "\n"
	}

	// Remove last new line
	r = r[:len(r)-1]

	return
}
