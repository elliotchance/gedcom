// Gedcomdiff is a tool for comparing GEDCOM files and producing a HTML report.
//
// Usage
//
//   gedcomdiff -left-gedcom file1.ged -right-gedcom file2.ged
//
// For a complete list of options use:
//
//   gedcomdiff -help
//
package main

import (
	"flag"
	"github.com/cheggaaa/pb"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/util"
	"log"
	"os"
)

// These are used for optionShow. If you update these options you will also
// need to adjust validateOptions.
const (
	optionShowAll         = "all" // default
	optionShowOnlyMatches = "only-matches"
	optionShowSubset      = "subset"
)

// These are used for optionSort. If you update these options you will also
// need to adjust validateOptions.
const (
	optionSortWrittenName       = "written-name" // default
	optionSortHighestSimilarity = "highest-similarity"
)

var (
	optionLeftGedcomFile            string
	optionRightGedcomFile           string
	optionOutputFile                string
	optionShow                      string // see optionShow constants.
	optionGoogleAnalyticsID         string
	optionProgress                  bool
	optionJobs                      int
	optionMinimumSimilarity         float64
	optionMinimumWeightedSimilarity float64
	optionSort                      string // see optionSort constants.
)

var filterFlags = &util.FilterFlags{}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	parseCLIFlags()

	leftGedcom, err := gedcom.NewDocumentFromGEDCOMFile(optionLeftGedcomFile)
	check(err)

	rightGedcom, err := gedcom.NewDocumentFromGEDCOMFile(optionRightGedcomFile)
	check(err)

	// Run compare.
	leftIndividuals := leftGedcom.Individuals()
	rightIndividuals := rightGedcom.Individuals()

	out, err := os.Create(optionOutputFile)
	check(err)

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

	flag.StringVar(&optionShow, "show", optionShowAll, util.CLIDescription(`
		The "-show" option controls which individuals are shown in the output:

		"all": Default. Show all individuals from both files.

		"only-matches": Only show individuals that match in both files. You can
		control the threshold with the "-minimum-weighted-similarity" and
		"-minimum-similarity" options. This is useful when comparing trees that
		are unlikely to have many matches.

		"subset": The right side will be considered a smaller part of the larger
		left side. This means that individuals that entirely exist on the left
		side will not be shown. This is useful when comparing a smaller part of
		a tree with a larger tree.`))

	flag.StringVar(&optionGoogleAnalyticsID, "google-analytics-id", "",
		"The Google Analytics ID, like 'UA-78454410-2'.")

	flag.BoolVar(&optionProgress, "progress", false, "Show progress bar.")

	flag.IntVar(&optionJobs, "jobs", 1, util.CLIDescription(`Number of jobs to run in
		parallel. If you are comparing large trees this will make the process
		faster but will consume more CPU.`))

	flag.Float64Var(&optionMinimumWeightedSimilarity,
		"minimum-weighted-similarity", gedcom.DefaultMinimumSimilarity,
		util.CLIDescription(`The weighted minimum similarity is the threshold
			for whether two individuals should be the seen as the same person
			when the surrounding immediate family is taken into consideration.

			This value must be between 0 and 1 and is the primary way to adjust
			the sensitivity of matches. It is best to also set
			"-minimum-similarity" to the same value.

			A higher value means you will get less matches but they will be of
			higher quality. If you are comparing trees that do not share many of
			the same individuals you should consider raising this to prevent
			false-positives.`))

	flag.Float64Var(&optionMinimumSimilarity,
		"minimum-similarity", gedcom.DefaultMinimumSimilarity,
		util.CLIDescription(`The minimum similarity is the threshold for matching
			individuals as the same person. This is used to compare only the
			individual (not surrounding family) like spouses and children.

			This value must be between 0 and 1 and should be set to the same
			value as "minimum-weighted-similarity" if you are unsure.`))

	flag.StringVar(&optionSort, "sort", optionSort, util.CLIDescription(`
			Controls how the individuals are sorted in the output:

			"written-name": Default. Sort individuals by written their written
			name.

			"highest-similarity": Sort the individuals by their match
			similarity. Highest matches will appear first.`))

	filterFlags.SetupCLI()

	flag.Parse()

	validateOptions()
}

func validateOptions() {
	optionShowValues := []string{
		optionShowAll,
		optionShowSubset,
		optionShowOnlyMatches,
	}

	if !util.StringSliceContains(optionShowValues, optionShow) {
		log.Fatalf(`invalid "-show" value: %s`, optionShow)
	}

	optionSortValues := []string{
		optionSortWrittenName,
		optionSortHighestSimilarity,
	}

	if !util.StringSliceContains(optionSortValues, optionSort) {
		log.Fatalf(`invalid "-sort" value: %s`, optionSort)
	}
}
