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
	NoCensuses   bool
	ExcludeTags  []string

	// Only official tags.
	OnlyOfficial bool

	// Condense NAME nodes to a simple string.
	SimpleNames bool

	// Only use the first name.
	SingleName bool
}

func NewFilterFlags() *FilterFlags {
	return &FilterFlags{}
}

func (ff *FilterFlags) SetupCLI() {
	flag.BoolVar(&ff.NoPlaces, "no-places", false, CLIDescription(`
		Exclude places. This is the same as "-exclude PLAC".`))
	flag.BoolVar(&ff.NoEvents, "no-events", false, CLIDescription(`
		Exclude events. This is the same as "-exclude EVEN".`))
	flag.BoolVar(&ff.NoResidences, "no-residences", false, CLIDescription(`
		Exclude residence events. This is the same as "-exclude RESI".`))
	flag.BoolVar(&ff.NoSources, "no-sources", false, CLIDescription(`
		Exclude sources. This is the same as "-exclude SOUR".`))
	flag.BoolVar(&ff.NoMaps, "no-maps", false, CLIDescription(`
		Exclude maps (locations). This is the same as "-exclude MAP".`))
	flag.BoolVar(&ff.NoChanges, "no-changes", false, CLIDescription(`
		Exclude change timestamps. This is the same as "-exclude CHAN".`))
	flag.BoolVar(&ff.NoObjects, "no-objects", false, CLIDescription(`
		Exclude objects. This is the same as "-exclude OBJ".`))
	flag.BoolVar(&ff.NoLabels, "no-labels", false, CLIDescription(`
		Exclude labels. This is the same as "-exclude LABL".`))
	flag.BoolVar(&ff.NoCensuses, "no-censuses", false, CLIDescription(`
		Exclude censuses. This is the same as "-exclude CENS".`))

	flag.BoolVar(&ff.OnlyOfficial, "only-official", false, CLIDescription(`
		Only include official GEDCOM tags.`))

	flag.BoolVar(&ff.SimpleNames, "simple-names", false, CLIDescription(`
		Simplify names.`))
	flag.BoolVar(&ff.SingleName, "single-name", false, CLIDescription(`
		Only use the first name for an individual.`))
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
		&ff.NoCensuses:   gedcom.TagCensus,
	}

	blacklistTags := []gedcom.Tag{}

	for _, tagName := range ff.ExcludeTags {
		tag := gedcom.TagFromString(tagName)

		blacklistTags = append(blacklistTags, tag)
	}

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

	if ff.SingleName {
		filters = append(filters, gedcom.SingleNameFilter())
	}

	return filters
}

func (ff *FilterFlags) AddExcludeTag(tag string) {
	if StringSliceContains(ff.ExcludeTags, tag) {
		return
	}

	ff.ExcludeTags = append(ff.ExcludeTags, tag)
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
