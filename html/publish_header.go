package html

import (
	"github.com/elliotchance/gedcom"
	"io"
)

const (
	selectedIndividualsTab = "individuals"
	selectedPlacesTab      = "places"
	selectedFamiliesTab    = "families"
	selectedSurnamesTab    = "surnames"
	selectedSourcesTab     = "sources"
	selectedStatisticsTab  = "statistics"
	selectedExtraTab       = "extra"
)

type PublishShowOptions struct {
	ShowIndividuals bool
	ShowPlaces      bool
	ShowFamilies    bool
	ShowSurnames    bool
	ShowSources     bool
	ShowStatistics  bool
	Checksum        bool
}

// PublishHeader is the tabbed section at the top of each page. The PublishHeader will be the
// same on all pages except that some pages will use an extra tab for that page.
type PublishHeader struct {
	document    *gedcom.Document
	extraTab    string
	selectedTab string
	options     PublishShowOptions
}

func NewPublishHeader(document *gedcom.Document, extraTab string, selectedTab string, options PublishShowOptions) *PublishHeader {
	return &PublishHeader{
		document:    document,
		extraTab:    extraTab,
		selectedTab: selectedTab,
		options:     options,
	}
}

func (c *PublishHeader) WriteTo(w io.Writer) (int64, error) {
	letters := GetIndexLetters(c.document)

	items := []*NavItem{}

	if c.options.ShowIndividuals {
		var badge Component = NewEmpty()
		if !c.options.Checksum {
			badge = NewCountBadge(len(c.document.Individuals()))
		}

		title := NewComponents(NewText("Individuals "), badge)
		item := NewNavItem(
			title,
			c.selectedTab == selectedIndividualsTab,
			PageIndividuals(letters[0]),
		)
		items = append(items, item)
	}

	if c.options.ShowPlaces {
		var badge Component = NewEmpty()
		if !c.options.Checksum {
			badge = NewCountBadge(len(GetPlaces(c.document)))
		}

		item := NewNavItem(
			NewComponents(NewText("Places "), badge),
			c.selectedTab == selectedPlacesTab,
			PagePlaces(),
		)
		items = append(items, item)
	}

	if c.options.ShowFamilies {
		var badge Component = NewEmpty()
		if !c.options.Checksum {
			badge = NewCountBadge(len(c.document.Families()))
		}

		item := NewNavItem(
			NewComponents(NewText("Families "), badge),
			c.selectedTab == selectedFamiliesTab,
			PageFamilies(),
		)
		items = append(items, item)
	}

	if c.options.ShowSurnames {
		var badge Component = NewEmpty()
		if !c.options.Checksum {
			badge = NewCountBadge(getSurnames(c.document).Len())
		}

		item := NewNavItem(
			NewComponents(NewText("Surnames "), badge),
			c.selectedTab == selectedSurnamesTab,
			PageSurnames(),
		)
		items = append(items, item)
	}

	if c.options.ShowSources {
		var badge Component = NewEmpty()
		if !c.options.Checksum {
			badge = NewCountBadge(len(c.document.Sources()))
		}

		item := NewNavItem(
			NewComponents(NewText("Sources "), badge),
			c.selectedTab == selectedSourcesTab,
			PageSources(),
		)
		items = append(items, item)
	}

	if c.options.ShowStatistics {
		item := NewNavItem(
			NewText("Statistics"),
			c.selectedTab == selectedStatisticsTab,
			PageStatistics(),
		)
		items = append(items, item)
	}

	if c.extraTab != "" {
		item := NewNavItem(
			NewText(c.extraTab),
			c.selectedTab == selectedExtraTab,
			"#",
		)
		items = append(items, item)
	}

	return NewComponents(
		NewSpace(),
		NewNavTabs(items),
		NewSpace(),
	).WriteTo(w)
}

var surnames = gedcom.NewStringSet()

func getSurnames(document *gedcom.Document) *gedcom.StringSet {
	if surnames.Len() == 0 {
		for _, individual := range document.Individuals() {
			surname := individual.Name().Surname()
			if surname != "" {
				surnames.Add(surname)
			}
		}
	}

	return surnames
}
