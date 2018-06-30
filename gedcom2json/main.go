package main

import (
	"flag"
	"github.com/elliotchance/gedcom"
	"os"
	"log"
	"encoding/json"
	"strings"
)

var (
	optionGedcomFile       string
	optionPrettyJSON       bool
	optionPrettyTags       bool
	optionNoPointers       bool
	optionTagKeys          bool
	optionStringName       bool
	optionExcludeTags      string
	optionOnlyOfficialTags bool
	optionOnlyTags         string
)

func main() {
	flag.StringVar(&optionGedcomFile, "gedcom", "", "Input GEDCOM file.")
	flag.BoolVar(&optionPrettyJSON, "pretty-json", false, "Pretty print JSON.")
	flag.BoolVar(&optionPrettyTags, "pretty-tags", false,
		"Output tags with their descriptive name instead of their raw tag "+
			`value. For example, "BIRT" would be output as "Birth".`)
	flag.BoolVar(&optionNoPointers, "no-pointers", false,
		`Do not include Pointer values ("ptr" attribute) in the output JSON. `+
			`This is useful to activate when comparing GEDCOM files that have `+
			`had pointers generated from different sources.`)
	flag.BoolVar(&optionTagKeys, "tag-keys", false,
		`Use tags (pretty or raw) as object keys rather than arrays.`)
	flag.BoolVar(&optionStringName, "string-name", false,
		`Convert NAME tags to a string (instead of the object parts).`)
	flag.StringVar(&optionExcludeTags, "exclude-tags", "",
		`Comma-separated list of tags to ignore.`)
	flag.BoolVar(&optionOnlyOfficialTags, "only-official-tags", false,
		`Only include tags from the GEDCOM standard in the output.`)
	flag.StringVar(&optionOnlyTags, "only-tags", "",
		`Only include these tags in the output.`)
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
		PrettyTags:       optionPrettyTags,
		NoPointers:       optionNoPointers,
		TagKeys:          optionTagKeys,
		StringName:       optionStringName,
		ExcludeTags:      splitTags(optionExcludeTags),
		OnlyOfficialTags: optionOnlyOfficialTags,
		OnlyTags:         splitTags(optionOnlyTags),
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

func splitTags(s string) []gedcom.Tag {
	if s == "" {
		return []gedcom.Tag{}
	}

	tags := []gedcom.Tag{}
	for _, t := range strings.Split(s, ",") {
		tags = append(tags, gedcom.Tag(t))
	}

	return tags
}
