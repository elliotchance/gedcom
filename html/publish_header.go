package html

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
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

// PublishHeader is the tabbed section at the top of each page. The PublishHeader will be the
// same on all pages except that some pages will use an extra tab for that page.
type PublishHeader struct {
	document    *gedcom.Document
	extraTab    string
	selectedTab string
	options     *PublishShowOptions
}

func NewPublishHeader(document *gedcom.Document, extraTab string, selectedTab string, options *PublishShowOptions) *PublishHeader {
	return &PublishHeader{
		document:    document,
		extraTab:    extraTab,
		selectedTab: selectedTab,
		options:     options,
	}
}

func (c *PublishHeader) WriteHTMLTo(w io.Writer) (int64, error) {
	letters := GetIndexLetters(c.document, c.options.LivingVisibility)

	items := []*core.NavItem{}

	if c.options.ShowIndividuals {
		badge := core.NewCountBadge(len(c.document.Individuals()))
		title := core.NewComponents(core.NewText("Individuals "), badge)
		item := core.NewNavItem(
			title,
			c.selectedTab == selectedIndividualsTab,
			PageIndividuals(letters[0]),
		)
		items = append(items, item)
	}

	if c.options.ShowPlaces {
		badge := core.NewCountBadge(len(GetPlaces(c.document)))
		item := core.NewNavItem(
			core.NewComponents(core.NewText("Places "), badge),
			c.selectedTab == selectedPlacesTab,
			PagePlaces(),
		)
		items = append(items, item)
	}

	if c.options.ShowFamilies {
		badge := core.NewCountBadge(len(c.document.Families()))
		item := core.NewNavItem(
			core.NewComponents(core.NewText("Families "), badge),
			c.selectedTab == selectedFamiliesTab,
			PageFamilies(),
		)
		items = append(items, item)
	}

	if c.options.ShowSurnames {
		badge := core.NewCountBadge(getSurnames(c.document).Len())
		item := core.NewNavItem(
			core.NewComponents(core.NewText("Surnames "), badge),
			c.selectedTab == selectedSurnamesTab,
			PageSurnames(),
		)
		items = append(items, item)
	}

	if c.options.ShowSources {
		badge := core.NewCountBadge(len(c.document.Sources()))
		item := core.NewNavItem(
			core.NewComponents(core.NewText("Sources "), badge),
			c.selectedTab == selectedSourcesTab,
			PageSources(),
		)
		items = append(items, item)
	}

	if c.options.ShowStatistics {
		item := core.NewNavItem(
			core.NewText("Statistics"),
			c.selectedTab == selectedStatisticsTab,
			PageStatistics(),
		)
		items = append(items, item)
	}

	if c.extraTab != "" {
		item := core.NewNavItem(
			core.NewText(c.extraTab),
			c.selectedTab == selectedExtraTab,
			"#",
		)
		items = append(items, item)
	}

	return core.NewComponents(
		core.NewSpace(),
		core.NewNavTabs(items),
		core.NewSpace(),
	).WriteHTMLTo(w)
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
