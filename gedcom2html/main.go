// Gedcom2html renders a GEDCOM file into HTML pages that can be shared and
// published easily.
//
// Usage
//
//   gedcom2html -gedcom file.ged
//
// You can view the full list of options using:
//
//   gedcom2html -help
//
// Example
//
// You can see an online example at http://dechauncy.family.
//
package main

import (
	"flag"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
	"github.com/elliotchance/gedcom/html/core"
	"github.com/elliotchance/gedcom/util"
	"log"
	"os"
)

var (
	optionGedcomFile        string
	optionOutputDir         string
	optionGoogleAnalyticsID string
	optionLivingVisibility  string
	optionJobs              int

	optionNoIndividuals bool
	optionNoPlaces      bool
	optionNoFamilies    bool
	optionNoSurnames    bool
	optionNoSources     bool
	optionNoStatistics  bool
)

func main() {
	flag.StringVar(&optionGedcomFile, "gedcom", "", "Input GEDCOM file.")
	// ghost:ignore
	flag.StringVar(&optionOutputDir, "output-dir", ".", "Output directory. It"+
		" will use the current directory if output-dir is not provided. "+
		"Output files will only be added or replaced. Existing files will not"+
		" be deleted.")
	flag.StringVar(&optionGoogleAnalyticsID, "google-analytics-id", "",
		"The Google Analytics ID, like 'UA-78454410-2'.")
	flag.StringVar(&optionLivingVisibility, "living",
		html.LivingVisibilityPlaceholder, util.CLIDescription(`
			Controls how information for living individuals are handled:

			"show": Show all living individuals and their information.

			"hide": Remove all living individuals as if they never existed.

			"placeholder": Show a "Hidden" placeholder that only that
			individuals are known but will not be displayed.`))
	flag.IntVar(&optionJobs, "jobs", 1,
		"Increasing this value will consume more resources but render the"+
			"website faster. An ideal value would be the number of CPUs "+
			"available, if you can spare it.")

	flag.BoolVar(&optionNoIndividuals, "no-individuals", false,
		"Exclude Individuals.")
	flag.BoolVar(&optionNoPlaces, "no-places", false,
		"Exclude Places.")
	flag.BoolVar(&optionNoFamilies, "no-families", false,
		"Exclude Families.")
	flag.BoolVar(&optionNoSurnames, "no-surnames", false,
		"Exclude Surnames.")
	flag.BoolVar(&optionNoSources, "no-sources", false,
		"Exclude Sources.")
	flag.BoolVar(&optionNoStatistics, "no-statistics", false,
		"Exclude Statistics.")

	flag.Parse()

	if optionGedcomFile == "" {
		log.Fatal("-gedcom is required")
	}

	file, err := os.Open(optionGedcomFile)
	if err != nil {
		log.Fatal(err)
	}

	decoder := gedcom.NewDecoder(file)
	document, err := decoder.Decode()
	if err != nil {
		log.Fatal(err)
	}

	options := &html.PublishShowOptions{
		ShowIndividuals:  !optionNoIndividuals,
		ShowPlaces:       !optionNoPlaces,
		ShowFamilies:     !optionNoFamilies,
		ShowSurnames:     !optionNoSurnames,
		ShowSources:      !optionNoSources,
		ShowStatistics:   !optionNoStatistics,
		LivingVisibility: html.NewLivingVisibility(optionLivingVisibility),
	}

	writer := core.NewDirectoryFileWriter(optionOutputDir)
	writer.WillWriteFile = func(file *core.File) {
		log.Printf("%s/%s\n", optionOutputDir, file.Name)
	}

	publisher := html.NewPublisher(document, options)
	err = publisher.Publish(writer, optionJobs)
	if err != nil {
		log.Fatal(err)
	}
}
