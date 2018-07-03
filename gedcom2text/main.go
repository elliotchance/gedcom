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
		// The first name of the individual is important

		if optionSplitDir != "" {
			names := individual.Names()
			if len(names) == 0 {
				// TODO: Should probably print out an error message here.
				continue
			}

			// TODO: Need to sanitise the name so it is safe for a file name.
			outFile, err = os.Create(optionSplitDir + "/" + names[0].String() + ".txt")
			if err != nil {
				log.Fatal(err)
			}
		}

		printLine("---")
		printLine("Individual:")

		for _, name := range individual.Names() {
			printLine(fmt.Sprintf("  Name: %s", name.String()))
		}
		printLine(fmt.Sprintf("  Sex: %s", individual.Sex()))

		printNodes(individual, gedcom.Birth)
		printNodes(individual, gedcom.Death)
	}
}

func printNodes(parent gedcom.Node, tag gedcom.Tag) {
	for _, node := range parent.NodesWithTag(tag) {
		printLine(fmt.Sprintf("  %s:", tag.String()))
		for _, n := range node.Nodes() {
			if n.Tag() == gedcom.Source && optionNoSources {
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
