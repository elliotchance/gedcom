// Package gedcom2text is a command line tool for converting GEDCOM to a
// human-readable text format.
package main

import (
	"flag"
	"fmt"
	"github.com/elliotchance/gedcom"
	"log"
	"os"
	"sort"
)

var (
	optionGedcomFile       string
	optionNoSources        bool
	optionOnlyOfficialTags bool
	optionSplitDir         string
	optionSingleName       bool
	optionNoPlaces         bool
	optionNoChangeTimes    bool
	optionNoEmptyDeaths    bool
)

var outFile *os.File

func main() {
	flag.StringVar(&optionGedcomFile, "gedcom", "", "Input GEDCOM file.")
	flag.BoolVar(&optionNoSources, "no-sources", false,
		"Do not include sources.")
	flag.BoolVar(&optionOnlyOfficialTags, "only-official-tags", false,
		"Only output official GEDCOM tags.")
	flag.StringVar(&optionSplitDir, "split-dir", "",
		"Split the individuals into separate files in this directory.")
	flag.BoolVar(&optionSingleName, "single-name", false,
		"Only output the primary name.")
	flag.BoolVar(&optionNoPlaces, "no-places", false,
		"Do not include places.")
	flag.BoolVar(&optionNoChangeTimes, "no-change-times", false,
		"Do not change timestamps.")
	flag.BoolVar(&optionNoEmptyDeaths, "no-empty-deaths", false,
		"Do not include Death node if there are no visible details.")
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

	// Sort individuals by name.
	individuals := document.Individuals()
	sort.SliceStable(individuals, func(i, j int) bool {
		return individuals[i].Names()[0].String() < individuals[j].Names()[0].String()
	})

	outFile = os.Stdout
	for _, individual := range individuals {
		if optionSplitDir != "" {
			outputFile := outputFileName(individual)
			if outputFile == "" {
				// TODO: Should probably print out an error message here.
				continue
			}

			outFile, err = os.Create(outputFile)
			if err != nil {
				log.Fatal(err)
			}
		}

		printLine("---")
		printLine("Individual:")

		for _, name := range individual.Names() {
			printLine(fmt.Sprintf("  Name: %s", name.String()))

			if optionSingleName {
				break
			}
		}
		printLine(fmt.Sprintf("  Sex: %s", individual.Sex()))

		printNodes(individual, gedcom.TagBirth)
		printNodes(individual, gedcom.TagDeath)

		printLine(fmt.Sprintf("  Families:"))
		for _, family := range individual.Families() {
			if !family.HasChild(individual) {
				continue
			}

			printLine("    -")
			if father := family.Husband(); father != nil {
				printLine(fmt.Sprintf("      Father: %s", father.Name()))
			}
			if mother := family.Wife(); mother != nil {
				printLine(fmt.Sprintf("      Mother: %s", mother.Name()))
			}
		}

		printLine(fmt.Sprintf("  Spouses:"))
		spouses := individual.Spouses()

		// Make sure the spouses are sorted as to not interfere with the
		// diffing.
		sort.SliceStable(spouses, func(i, j int) bool {
			return spouses[i].Names()[0].String() < spouses[j].Names()[0].String()
		})

		for _, spouse := range spouses {
			for _, name := range spouse.Names() {
				// Only use the primary name for the spouse here as it needs to
				// be the key. If they have enabled multiple names then the
				// extra names will appear on the individual page.
				printLine(fmt.Sprintf("    - Name: %s", name.String()))
				break
			}

			// Marriage events.
			withSpouse := individual.FamilyWithSpouse(spouse)
			m := gedcom.First(gedcom.NodesWithTagPath(withSpouse,
				gedcom.TagMarriage, gedcom.TagDate))
			if m != nil {
				printLine(fmt.Sprintf("      Marriage: %s", m.Value()))
			}

			// Children of the spouse.
			children := withSpouse.Children()
			sort.SliceStable(children, func(i, j int) bool {
				return children[i].Names()[0].String() < children[j].Names()[0].String()
			})
			printLine("      Children:")
			for _, child := range children {
				n := "Unknown"
				if n2 := child.Name(); n2 != nil {
					n = n2.String()
				}
				printLine(fmt.Sprintf("        - Name: %s", n))
			}
		}
	}
}

func tagShouldBeExcluded(tag gedcom.Tag) bool {
	if tag == gedcom.TagSource && optionNoSources {
		return true
	}

	if tag == gedcom.TagPlace && optionNoPlaces {
		return true
	}

	if tag == gedcom.TagChange && optionNoChangeTimes {
		return true
	}

	if !tag.IsOfficial() && optionOnlyOfficialTags {
		return true
	}

	return false
}

func printNodes(parent gedcom.Node, tag gedcom.Tag) {
	for _, node := range gedcom.NodesWithTag(parent, tag) {
		// Death is a special case because it's common to have a Death node with
		// no details to signify that the person is not living.
		//
		// This can lead to problems comparing files when one side has not
		// followed this pattern.
		//
		// We have to look forward and be sensitive to data that otherwise would
		// not have been shown to make sure we do not include empty Death tags.
		if tag == gedcom.TagDeath {
			foundChild := false

			for _, n := range node.Nodes() {
				if tagShouldBeExcluded(n.Tag()) {
					continue
				}

				foundChild = true
			}

			if !foundChild {
				continue
			}
		}

		printLine(fmt.Sprintf("  %s:", tag.String()))
		for _, n := range node.Nodes() {
			if tagShouldBeExcluded(n.Tag()) {
				continue
			}

			printLine(fmt.Sprintf("    %s: %s", n.Tag().String(), n.Value()))
		}
	}
}

func printLine(line string) {
	fmt.Fprintf(outFile, "%s\n", line)
}
