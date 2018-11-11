package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/elliotchance/gedcom"
	"log"
)

var (
	optionGedcomFile string
)

func main() {
	parseCLIFlags()

	parser := NewParser()
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
