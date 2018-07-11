package main

import (
	"flag"
	"github.com/elliotchance/gedcom"
	"os"
	"log"
	"fmt"
	"sort"
)

var (
	optionGedcomFile       string
	optionNoSources        bool
	optionOnlyOfficialTags bool
	optionSplitDir         string
	optionSingleName       bool
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

		printLine(fmt.Sprintf("  Spouses:"))
		for _, spouse := range individual.Spouses(document) {
			for _, name := range spouse.Names() {
				printLine(fmt.Sprintf("    Name: %s", name.String()))

				if optionSingleName {
					break
				}
			}
		}
	}
}

func outputFileName(individual *gedcom.IndividualNode) string {
	names := individual.Names()
	if len(names) == 0 {
		return ""
	}

	// Include the birth/death information to make the name more unique.
	birth := ""
	if node := individual.FirstNodeWithTag(gedcom.TagBirth); node != nil {
		if node2 := node.FirstNodeWithTag(gedcom.TagDate); node2 != nil {
			birth = node2.Value()
		}
	}

	death := ""
	if node := individual.FirstNodeWithTag(gedcom.TagDeath); node != nil {
		if node2 := node.FirstNodeWithTag(gedcom.TagDate); node2 != nil {
			death = node2.Value()
		}
	}

	// TODO: Need to sanitise the name so it is safe for a file name.
	return fmt.Sprintf("%s/%s (%s - %s).txt", optionSplitDir, names[0].String(), birth, death)
}

func printNodes(parent gedcom.Node, tag gedcom.Tag) {
	for _, node := range parent.NodesWithTag(tag) {
		printLine(fmt.Sprintf("  %s:", tag.String()))
		for _, n := range node.Nodes() {
			if n.Tag() == gedcom.TagSource && optionNoSources {
				continue
			}

			if !n.Tag().IsOfficial() && optionOnlyOfficialTags {
				continue
			}

			printLine(fmt.Sprintf("    %s: %s", n.Tag().String(), n.Value()))
		}
	}
}

func printLine(line string) {
	fmt.Fprintf(outFile, "%s\n", line)
}
