package main

import (
	"github.com/elliotchance/gedcom"
)

const (
	selectedIndividualsTab = "individuals"
	selectedPlacesTab      = "places"
	selectedFamiliesTab    = "families"
	selectedSourcesTab     = "sources"
	selectedStatisticsTab  = "statistics"
	selectedExtraTab       = "extra"
)

// header is the tabbed section at the top of each page. The header will be the
// same on all pages except that some pages will use an extra tab for that page.
type header struct {
	document    *gedcom.Document
	extraTab    string
	selectedTab string
}

func newHeader(document *gedcom.Document, extraTab string, selectedTab string) *header {
	return &header{
		document:    document,
		extraTab:    extraTab,
		selectedTab: selectedTab,
	}
}

func (c *header) String() string {
	items := []*navItem{
		newNavItem(
			"Individuals "+newCountBadge(len(c.document.Individuals())).String(),
			c.selectedTab == selectedIndividualsTab,
			pageIndividuals(symbolLetter),
		),
		newNavItem(
			"Places "+newCountBadge(len(getPlaces(c.document))).String(),
			c.selectedTab == selectedPlacesTab,
			pagePlaces(),
		),
		newNavItem(
			"Families "+newCountBadge(len(c.document.Families())).String(),
			c.selectedTab == selectedFamiliesTab,
			pageFamilies(),
		),
		newNavItem(
			"Sources "+newCountBadge(len(c.document.Sources())).String(),
			c.selectedTab == selectedSourcesTab,
			pageSources(),
		),
		newNavItem(
			"Statistics",
			c.selectedTab == selectedStatisticsTab,
			pageStatistics(),
		),
	}

	if c.extraTab != "" {
		items = append(items, newNavItem(
			c.extraTab,
			c.selectedTab == selectedExtraTab,
			"#",
		))
	}

	return newComponents(
		newSpace(),
		newNavTabs(items),
		newSpace(),
	).String()
}
