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
	optionGedcomFile string
	optionFormat     string
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
	default:
		log.Panicf("unsupported format: %s", optionFormat)
	}

	formatter.Write(result)
}

func parseCLIFlags() {
	flag.StringVar(&optionGedcomFile, "gedcom", "", util.CLIDescription(`
		Path to the GEDCOM file (required).`))
	flag.StringVar(&optionFormat, "format", "json", util.CLIDescription(`
		Output format, can be one of the following: "json", "pretty-json" or
		"csv".`))

	flag.Parse()
}
