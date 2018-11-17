// Gedcomq is a command line tool and query language for GEDCOM files heavily
// inspired by jq, in name and syntax.
//
// The basic syntax of the tool is:
//
//   gedcomq -gedcom file.ged '.Individuals | .Name'
//
// You can find the full language documentation in the q package:
//
// https://godoc.org/github.com/elliotchance/gedcom/q
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/q"
	"log"
)

var (
	optionGedcomFile string
)

func main() {
	parseCLIFlags()

	parser := q.NewParser()
	engine, err := parser.ParseString(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	doc, err := gedcom.NewDocumentFromGEDCOMFile(optionGedcomFile)
	if err != nil {
		log.Fatal(err)
	}

	result, err := engine.Evaluate(doc)
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}

func parseCLIFlags() {
	flag.StringVar(&optionGedcomFile, "gedcom", "", "GEDCOM file.")

	flag.Parse()
}
