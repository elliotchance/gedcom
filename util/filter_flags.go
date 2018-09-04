package util

import (
	"flag"
	"github.com/elliotchance/gedcom"
)

type FilterFlags struct {
	// Specific exclusions.
	NoEvents     bool
	NoResidences bool
	NoPlaces     bool
	NoSources    bool
	NoMaps       bool
	NoChanges    bool
	NoObjects    bool
	NoLabels     bool

	// Only official tags.
	OnlyOfficial bool

	// When comparing, hide lines that are equal on both sides.
	HideEqual bool

	// Condense NAME nodes to a simple string.
	SimpleNames bool
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

	flag.BoolVar(&ff.OnlyOfficial, "only-official", false, "Only include official GEDCOM tags.")

	flag.BoolVar(&ff.HideEqual, "hide-equal", false, "Hide equal values.")

	flag.BoolVar(&ff.SimpleNames, "simple-names", false, "Simplify names.")
}

func (ff *FilterFlags) FilterFunctions() []gedcom.FilterFunction {
	m := map[*bool]gedcom.Tag{
		&ff.NoEvents:     gedcom.TagEvent,
		&ff.NoResidences: gedcom.TagResidence,
		&ff.NoPlaces:     gedcom.TagPlace,
		&ff.NoSources:    gedcom.TagSource,
		&ff.NoMaps:       gedcom.TagMap,
		&ff.NoChanges:    gedcom.TagChange,
		&ff.NoObjects:    gedcom.TagObject,
		&ff.NoLabels:     gedcom.TagLabel,
	}

	blacklistTags := []gedcom.Tag{gedcom.TagFamilyChild, gedcom.TagFamilySpouse}
	for k, v := range m {
		if *k {
			blacklistTags = append(blacklistTags, v)
		}
	}

	filters := []gedcom.FilterFunction{
		gedcom.BlacklistTagFilter(blacklistTags...),
	}

	if ff.OnlyOfficial {
		filters = append(filters, gedcom.OfficialTagFilter())
	}

	if ff.SimpleNames {
		filters = append(filters, gedcom.SimpleNameFilter())
	}

	return filters
}

func (ff *FilterFlags) Filter(node gedcom.Node) gedcom.Node {
	if gedcom.IsNil(node) {
		return nil
	}

	for _, filter := range ff.FilterFunctions() {
		node = gedcom.Filter(node, filter)
	}

	return node
}
