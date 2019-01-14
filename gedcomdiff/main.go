// Gedcomdiff is a tool for comparing GEDCOM files and producing a HTML report.
//
// Usage
//
//   gedcomdiff -left-gedcom file1.ged -right-gedcom file2.ged -output out.html
//
// For a complete list of options use:
//
//   gedcomdiff -help
//
package main

import (
	"flag"
	"fmt"
	"github.com/cheggaaa/pb"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
	"github.com/elliotchance/gedcom/util"
	"log"
	"os"
	"os/signal"
	"time"
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
	optionPreferPointerAbove        float64
)

var filterFlags = &util.FilterFlags{}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	parseCLIFlags()

	similarityOptions := gedcom.NewSimilarityOptions()
	similarityOptions.MinimumWeightedSimilarity = optionMinimumWeightedSimilarity
	similarityOptions.PreferPointerAbove = optionPreferPointerAbove
	similarityOptions.MinimumSimilarity = optionMinimumSimilarity

	compareOptions := gedcom.NewIndividualNodesCompareOptions()
	compareOptions.SimilarityOptions = similarityOptions
	compareOptions.Notifier = make(chan gedcom.Progress)
	compareOptions.NotifierStep = 100
	compareOptions.Jobs = optionJobs

	// Gracefully handle a ctrl+c.
	signaller := make(chan os.Signal, 1)
	signal.Notify(signaller, os.Interrupt)
	go func() {
		<-signaller

		defer func() {
			// The panic may occur if crtl+c happens after the comparisons after
			// the comparisons are finished but it's still rendering the output.
			//
			// We tried our best, exit with a failure code.
			recover()
			log.Fatal("aborted")
		}()

		// This is the correct way to abort the comparisons.
		close(compareOptions.Notifier)
	}()

	leftGedcom, err := gedcom.NewDocumentFromGEDCOMFile(optionLeftGedcomFile)
	check(err)

	rightGedcom, err := gedcom.NewDocumentFromGEDCOMFile(optionRightGedcomFile)
	check(err)

	// Run compare.
	leftIndividuals := leftGedcom.Individuals()
	rightIndividuals := rightGedcom.Individuals()

	out, err := os.Create(optionOutputFile)
	check(err)

	var comparisons gedcom.IndividualComparisons

	go func() {
		comparisons = leftIndividuals.Compare(rightIndividuals, compareOptions)
	}()

	if optionProgress {
		progressBar := pb.StartNew(0).Prefix("Comparing Documents")
		progressBar.SetRefreshRate(500 * time.Millisecond)
		progressBar.ShowElapsedTime = true
		progressBar.ShowTimeLeft = true

		for n := range compareOptions.Notifier {
			progressBar.SetTotal64(n.Total)
			progressBar.Set64(n.Done)
		}

		progressBar.Finish()
	} else {
		// Wait for notifier channel to be closed.
		for range compareOptions.Notifier {
		}
	}

	diffProgress := make(chan gedcom.Progress)

	page := html.NewDiffPage(comparisons, filterFlags, optionGoogleAnalyticsID,
		optionShow, optionSort, diffProgress, compareOptions, html.LivingVisibilityShow)

	go func() {
		_, err = page.WriteTo(out)
		if err != nil {
			log.Fatal(err)
		}

		close(diffProgress)
	}()

	if optionProgress {
		progressBar := pb.StartNew(0).Prefix("Comparing Individuals")
		progressBar.SetRefreshRate(500 * time.Millisecond)
		progressBar.ShowElapsedTime = true
		progressBar.ShowTimeLeft = true

		for p := range diffProgress {
			if p.Total != 0 {
				progressBar.SetTotal64(p.Total)
			}

			progressBar.Add64(p.Add)
		}

		progressBar.Finish()
	} else {
		for range diffProgress {
		}
	}
}

func parseCLIFlags() {
	// Input files. Must be provided.
	flag.StringVar(&optionLeftGedcomFile, "left-gedcom", "",
		"Required. Left GEDCOM file.")

	flag.StringVar(&optionRightGedcomFile, "right-gedcom", "",
		"Required. Right GEDCOM file.")

	flag.StringVar(&optionOutputFile, "output", "", "Output file.")

	flag.StringVar(&optionShow, "show", html.DiffPageShowAll, util.CLIDescription(`
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
		util.CLIDescription(`The minimum similarity is the threshold for
			matching individuals as the same person. This is used to compare
			only the individual (not surrounding family) like spouses and
			children.

			This value must be between 0 and 1 and should be set to the same
			value as "minimum-weighted-similarity" if you are unsure.`))

	flag.StringVar(&optionSort, "sort", html.DiffPageSortWrittenName,
		util.CLIDescription(`
			Controls how the individuals are sorted in the output:

			"written-name": Sort individuals by written their written name.

			"highest-similarity": Sort the individuals by their match
			similarity. Highest matches will appear first.`))

	flag.Float64Var(&optionPreferPointerAbove, "prefer-pointer-above",
		gedcom.DefaultMinimumSimilarity, util.CLIDescription(fmt.Sprintf(`
			Controls if two individuals should be considered a match by their
			pointer value.

			The default value is %f which means that the individuals will be
			considered a match if they share the same pointer and hit the same
			default minimum similarity.

			A value of 1.0 would have to be a perfect match to be considered
			equal on their pointer, this is the same as disabling the feature.

			A value of 0.0 would mean that it always trusts the pointer match,
			even if the individuals are nothing alike.

			This option makes sense when you are comparing documents that have
			come from the same base and retained the pointers between
			individuals of the existing data.
			`, gedcom.DefaultMinimumSimilarity)))

	filterFlags.SetupCLI()

	flag.Parse()

	validateOptions()
}

func validateOptions() {
	if optionLeftGedcomFile == "" {
		log.Fatalf(`-left-gedcom is required`)
	}

	if optionRightGedcomFile == "" {
		log.Fatalf(`-right-gedcom is required`)
	}

	if optionOutputFile == "" {
		log.Fatalf(`-output is required`)
	}

	optionShowValues := []string{
		html.DiffPageShowAll,
		html.DiffPageShowSubset,
		html.DiffPageShowOnlyMatches,
	}

	if !util.StringSliceContains(optionShowValues, optionShow) {
		log.Fatalf(`invalid "-show" value: %s`, optionShow)
	}

	optionSortValues := []string{
		html.DiffPageSortWrittenName,
		html.DiffPageSortHighestSimilarity,
	}

	if !util.StringSliceContains(optionSortValues, optionSort) {
		log.Fatalf(`invalid "-sort" value: %s`, optionSort)
	}
}
