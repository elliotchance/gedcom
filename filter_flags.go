package gedcom

import (
	"flag"

	"github.com/elliotchance/gedcom/v39/util"
)

type FilterFlags struct {
	// Specific exclusions.
	NoEvents         bool
	NoResidences     bool
	NoPlaces         bool
	NoSources        bool
	NoMaps           bool
	NoChanges        bool
	NoObjects        bool
	NoLabels         bool
	NoCensuses       bool
	NoEmptyDeaths    bool
	NoDuplicateNames bool

	// Only vitals (name, birth, baptism, death and burial).
	OnlyVitals bool

	// Only official tags.
	OnlyOfficial bool

	// When comparing, hide lines that are equal on both sides.
	HideEqual bool

	// Condense NAME nodes to a simple string.
	NameFormat string
}

func (ff *FilterFlags) SetupCLI() {
	flag.BoolVar(&ff.NoPlaces, "no-places", false, "Exclude places.")
	flag.BoolVar(&ff.NoEvents, "no-events", false, "Exclude events.")
	flag.BoolVar(&ff.NoResidences, "no-residences", false, "Exclude residence events.")
	flag.BoolVar(&ff.NoSources, "no-sources", false, "Exclude sources.")
	flag.BoolVar(&ff.NoMaps, "no-maps", false, "Exclude maps (locations).")
	flag.BoolVar(&ff.NoChanges, "no-changes", false, "Exclude change timestamps.")
	flag.BoolVar(&ff.NoObjects, "no-objects", false, "Exclude objects.")
	flag.BoolVar(&ff.NoLabels, "no-labels", false, "Exclude labels.")
	flag.BoolVar(&ff.NoCensuses, "no-censuses", false, "Exclude censuses.")
	flag.BoolVar(&ff.NoDuplicateNames, "no-duplicate-names", false, "Exclude names that are duplicates.")

	flag.BoolVar(&ff.NoEmptyDeaths, "no-empty-deaths", false, util.CLIDescription(
		`Remove death nodes (DEAT) that do not have children. This is caused by
		applications signalling that the individual is not living but can lead
		to unwanted discrepancies in the comparison.`))

	flag.BoolVar(&ff.OnlyVitals, "only-vitals", false, util.CLIDescription(`
		Remove all data except for vital information. The vital nodes are (or
		multiples in the same individual of): Name, birth, baptism, death and
		burial. Within these only the date and place is retained.`))

	flag.BoolVar(&ff.OnlyOfficial, "only-official", false,
		"Only include official GEDCOM tags.")

	flag.BoolVar(&ff.HideEqual, "hide-equal", false, "Hide equal values.")

	flag.StringVar(&ff.NameFormat, "name-format", "written", util.CLIDescription(`
		The NAME node can be represented a single string, or name parts such as
		Given name, Surname, Title, etc. When enabled, this option flattens name
		parts into a single string with the given format:

		"written": Default. Flatten names to their written names, like
		"John Smith".

		"gedcom": Flatten names to their GEDCOM name, like "John /Smith/".

		"index": Flatten names to their index name, like "Smith, John".

		"unmodified": Do not make any modifications to the name or name parts.

		You can also provide a custom format (see NameFormat) by not using one
		of the presets above.`))
}

func (ff *FilterFlags) FilterFunctions() []FilterFunction {
	m := map[*bool]Tag{
		&ff.NoEvents:     TagEvent,
		&ff.NoResidences: TagResidence,
		&ff.NoPlaces:     TagPlace,
		&ff.NoSources:    TagSource,
		&ff.NoMaps:       TagMap,
		&ff.NoChanges:    TagChange,
		&ff.NoObjects:    TagObject,
		&ff.NoLabels:     TagLabel,
		&ff.NoCensuses:   TagCensus,
	}

	blacklistTags := []Tag{TagFamilyChild, TagFamilySpouse}
	for k, v := range m {
		if *k {
			blacklistTags = append(blacklistTags, v)
		}
	}

	filters := []FilterFunction{
		BlacklistTagFilter(blacklistTags...),
	}

	if ff.OnlyOfficial {
		filters = append(filters, OfficialTagFilter())
	}

	// This must be before NameFormat because NameFormat may simplify/destroy
	// NAME nodes.
	if ff.NoDuplicateNames {
		filters = append(filters, RemoveDuplicateNamesFilter())
	}

	if ff.NameFormat != "unmodified" {
		format, _ := NewNameFormatByName(ff.NameFormat)
		filters = append(filters, SimpleNameFilter(format))
	}

	if ff.OnlyVitals {
		filters = append(filters, OnlyVitalsTagFilter())
	}

	if ff.NoEmptyDeaths {
		filters = append(filters, RemoveEmptyDeathTagFilter())
	}

	return filters
}

func (ff *FilterFlags) Filter(node Node, document *Document) Node {
	if IsNil(node) {
		return nil
	}

	for _, filter := range ff.FilterFunctions() {
		node = Filter(node, document, filter)
	}

	return node
}
