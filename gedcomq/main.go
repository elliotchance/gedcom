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
	"flag"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/q"
	"github.com/elliotchance/gedcom/util"
	"log"
	"os"
)

var (
	optionGedcomFiles util.CLIStringSlice
	optionFormat      string
)

func main() {
	parseCLIFlags()

	parser := q.NewParser()
	engine, err := parser.ParseString(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	docs := []*gedcom.Document{}

	for _, gedcomFile := range optionGedcomFiles {
		doc, err := gedcom.NewDocumentFromGEDCOMFile(gedcomFile)
		if err != nil {
			log.Fatal(err)
		}

		docs = append(docs, doc)
	}

	if len(docs) == 0 {
		log.Fatal("you must provide at least one gedcom file")
	}

	result, err := engine.Evaluate(docs)
	if err != nil {
		log.Fatal(err)
	}

	output(result)
}

func output(result interface{}) {
	var formatter q.Formatter

	switch optionFormat {
	case "json":
		formatter = &q.JSONFormatter{os.Stdout}
	case "pretty-json":
		formatter = &q.PrettyJSONFormatter{os.Stdout}
	case "csv":
		formatter = &q.CSVFormatter{os.Stdout}
	case "gedcom":
		formatter = &q.GEDCOMFormatter{os.Stdout}
	default:
		log.Panicf("unsupported format: %s", optionFormat)
	}

	formatter.Write(result)
}

func parseCLIFlags() {
	flag.Var(&optionGedcomFiles, "gedcom", util.CLIDescription(`
		Path to the GEDCOM file. You may specify more than one document by
		providing -gedcom with an argument multiple times. You must provide at
		least one document.`))
	flag.StringVar(&optionFormat, "format", "json", util.CLIDescription(`
		Output format, can be one of the following: "json", "pretty-json",
		"gedcom" or "csv".`))

	flag.Parse()
}
