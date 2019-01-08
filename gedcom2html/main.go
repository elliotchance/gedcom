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
// Using the Checksum File
//
// The "-checksum" generates a file called "checksum.csv". This file contains
// file names and their SHA-1 checksum like:
//
//   amos-adams.html,b0538fb8186a50c4079c902fec2b4ba0af843061
//   massachusetts-united-states.html,79db811c089e8ab5653d34551e6540cb2ea2c947
//
// The lines are ordered by the file name so the output is ideal for comparison.
//
// Here is an example of using the previous and current checksum file to
// generate sync commands:
//
//   join -a 1 -a 2 -t, -o 0.1,1.2,2.2 /old/checksum.csv /new/checksum.csv | \
//      awk -F, '$2 == $3 { next } { print $3 == "" \
//          ? "rm /some/folder/" $1 \
//          : "cp" " " $1 " /some/folder/" $1 }'
//
// Will produce commands like:
//
//   cp abos-adams.html /some/folder/abos-adams.html
//   rm /some/folder/massachusetts-united-states.html
//
package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

var (
	optionGedcomFile        string
	optionOutputDir         string
	optionGoogleAnalyticsID string
	optionChecksum          bool

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
	flag.BoolVar(&optionChecksum, "checksum", false,
		"Output a checksum file, helpful for syncing large trees.")

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

	// Create the pages.
	if !optionNoIndividuals {
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
	}

	if !optionNoPlaces {
		page := newPlaceListPage(document, optionGoogleAnalyticsID)
		createFile(pagePlaces(), page)

		// Sort the places so that the generated page names will be more
		// deterministic.
		places := getPlaces(document)
		placeKeys := []string{}

		for key := range places {
			placeKeys = append(placeKeys, key)
		}

		sort.Strings(placeKeys)

		for _, key := range placeKeys {
			place := places[key]
			page := newPlacePage(document, key, optionGoogleAnalyticsID)
			createFile(pagePlace(document, place.prettyName), page)
		}
	}

	if !optionNoFamilies {
		createFile(pageFamilies(), newFamilyListPage(document, optionGoogleAnalyticsID))
	}

	if !optionNoSurnames {
		createFile(pageSurnames(), newSurnameListPage(document, optionGoogleAnalyticsID))
	}

	if !optionNoSources {
		createFile(pageSources(), newSourceListPage(document, optionGoogleAnalyticsID))

		for _, source := range document.Sources() {
			page := newSourcePage(document, source, optionGoogleAnalyticsID)
			createFile(pageSource(source), page)
		}
	}

	if !optionNoStatistics {
		createFile(pageStatistics(), newStatisticsPage(document, optionGoogleAnalyticsID))
	}

	// Calculate checksum
	if optionChecksum {
		lines := []string{}
		fileInfos, err := ioutil.ReadDir(optionOutputDir)
		if err != nil {
			log.Fatal(err)
		}

		for _, fileInfo := range fileInfos {
			checksum := fileSha1(fileInfo.Name())
			line := fmt.Sprintf("%s,%s", fileInfo.Name(), checksum)
			lines = append(lines, line)
		}

		sort.Strings(lines)

		createFile("checksum.csv", html.NewText(strings.Join(lines, "\n")))
	}
}

func fileSha1(path string) string {
	f, err := os.Open(optionOutputDir + "/" + path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

func createFile(name string, contents fmt.Stringer) {
	path := fmt.Sprintf("%s/%s", optionOutputDir, name)
	log.Printf("Writing %s...", path)

	out, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	// ghost:ignore
	out.Write([]byte(contents.String()))

	out.Close()
}
