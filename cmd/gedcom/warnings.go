package main

import (
	"flag"
	"fmt"
	"github.com/elliotchance/gedcom"
	"os"
)

func runWarningsCommand() {
	err := flag.CommandLine.Parse(os.Args[2:])
	if err != nil {
		fatalln(err)
	}

	gedcomFile := flag.Arg(0)
	if gedcomFile == "" {
		fatalln("you must provide a gedcom file")
	}

	doc, err := gedcom.NewDocumentFromGEDCOMFile(gedcomFile)
	if err != nil {
		fatalln(err)
	}

	for _, warning := range doc.Warnings() {
		fmt.Println(warning)
	}
}
