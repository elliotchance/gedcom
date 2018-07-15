package main

import (
	"encoding/json"
	"flag"
	"github.com/elliotchance/gedcom"
	"log"
	"os"
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
	optionSingleName       bool
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
	flag.BoolVar(&optionSingleName, "single-name", false,
		`When there are multiple names for an individual this will return the `+
			`first of the name nodes only.`)
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
		ExcludeTags:      splitTags(optionExcludeTags),
		NoPointers:       optionNoPointers,
		OnlyOfficialTags: optionOnlyOfficialTags,
		OnlyTags:         splitTags(optionOnlyTags),
		PrettyTags:       optionPrettyTags,
		SingleName:       optionStringName,
		StringName:       optionStringName,
		TagKeys:          optionTagKeys,
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
