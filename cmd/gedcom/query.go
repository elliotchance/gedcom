// "gedcom query" is a command line tool and query language for GEDCOM files
// heavily inspired by jq, in name and syntax.
//
// The basic syntax of the tool is:
//
//   gedcom query -gedcom file.ged '.Individuals | .Name'
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
	"os"
)

func runQueryCommand() {
	var gedcomFiles util.CLIStringSlice
	var format string

	flag.Var(&gedcomFiles, "gedcom", util.CLIDescription(`
		Path to the GEDCOM file. You may specify more than one document by
		providing -gedcom with an argument multiple times. You must provide at
		least one document.`))

	flag.StringVar(&format, "format", "json", util.CLIDescription(`
		Output format, can be one of the following: "json", "pretty-json",
		"gedcom" or "csv".`))

	err := flag.CommandLine.Parse(os.Args[2:])
	if err != nil {
		fatalln(err)
	}

	if gedcomFiles.String() == "" {
		fatalln("you must specify at least one -gedcom file")
	}

	parser := q.NewParser()
	engine, err := parser.ParseString(flag.Arg(0))
	if err != nil {
		fatalln(err)
	}

	var mutDocs []*gedcom.Document

	for _, gedcomFile := range gedcomFiles {
		doc, err := gedcom.NewDocumentFromGEDCOMFile(gedcomFile)
		if err != nil {
			fatalln(err)
		}

		mutDocs = append(mutDocs, doc)
	}

	if len(mutDocs) == 0 {
		fatalln("you must provide at least one gedcom file")
	}

	result, err := engine.Evaluate(mutDocs)
	if err != nil {
		fatalln(err)
	}

	err = output(result, format)
	if err != nil {
		fatalln(err)
	}
}

func output(result interface{}, format string) error {
	var formatter q.Formatter

	switch format {
	case "json":
		formatter = &q.JSONFormatter{os.Stdout}
	case "pretty-json":
		formatter = &q.PrettyJSONFormatter{os.Stdout}
	case "csv":
		formatter = &q.CSVFormatter{os.Stdout}
	case "gedcom":
		formatter = &q.GEDCOMFormatter{os.Stdout}
	case "html":
		formatter = &q.HTMLFormatter{os.Stdout}
	default:
		fatalln("unsupported format: %s", format)
	}

	return formatter.Write(result)
}
