package main

import (
	"flag"
	"github.com/elliotchance/gedcom"
	"os"
	"log"
	"encoding/json"
)

var (
	optionGedcomFile string
	optionPrettyJSON bool
	optionPrettyTags bool
)

func main() {
	flag.StringVar(&optionGedcomFile, "gedcom", "", "Input GEDCOM file.")
	flag.BoolVar(&optionPrettyJSON, "pretty-json", false, "Pretty print JSON.")
	flag.BoolVar(&optionPrettyTags, "pretty-tags", false,
		"Use descriptive names instead of raw tags.")
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

	options := gedcom.TransformOptions{
		PrettyTags: optionPrettyTags,
	}

	var bytes []byte
	transformedDocument := gedcom.Transform(document, options)

	if optionPrettyJSON {
		bytes, err = json.MarshalIndent(transformedDocument, "", "  ")
	} else {
		bytes, err = json.Marshal(transformedDocument)
	}
	if err != nil {
		log.Fatal(err)
	}

	os.Stdout.Write(bytes)
	os.Stdout.Write([]byte{'\n'})
}
