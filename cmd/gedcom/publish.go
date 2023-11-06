// "gedcom publish" renders a GEDCOM file into HTML pages that can be shared and
// published easily.
//
// Usage
//
//   gedcom publish -gedcom file.ged
//
// You can view the full list of options using:
//
//   gedcom publish -help
//
package main

import (
	"flag"
	"log"
	"os"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html"
	"github.com/elliotchance/gedcom/v39/html/core"
	"github.com/elliotchance/gedcom/v39/util"
)

func runPublishCommand() {
	var optionGedcomFile string
	var optionOutputDir string
	var optionGoogleAnalyticsID string
	var optionLivingVisibility string
	var optionJobs int

	var optionNoIndividuals bool
	var optionNoPlaces bool
	var optionNoFamilies bool
	var optionNoSurnames bool
	var optionNoSources bool
	var optionNoStatistics bool

	flag.StringVar(&optionGedcomFile, "gedcom", "", "Input GEDCOM file.")

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

	err := flag.CommandLine.Parse(os.Args[2:])
	if err != nil {
		fatalln(err)
	}

	if optionGedcomFile == "" {
		fatalln("-gedcom is required")
	}

	file, err := os.Open(optionGedcomFile)
	if err != nil {
		fatalln(err)
	}

	decoder := gedcom.NewDecoder(file)
	document, err := decoder.Decode()
	if err != nil {
		fatalln(err)
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
		fatalln(err)
	}
}
