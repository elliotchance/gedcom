package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
	"github.com/elliotchance/gedcom/util"
	"sort"
	"strings"
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
	letters := getIndexLetters(c.document)

	items := []*navItem{}

	if !optionNoIndividuals {
		item := newNavItem(
			"Individuals "+newCountBadge(len(c.document.Individuals())).String(),
			c.selectedTab == selectedIndividualsTab,
			pageIndividuals(letters[0]),
		)
		items = append(items, item)
	}

	if !optionNoPlaces {
		item := newNavItem(
			"Places "+newCountBadge(len(getPlaces(c.document))).String(),
			c.selectedTab == selectedPlacesTab,
			pagePlaces(),
		)
		items = append(items, item)
	}

	if !optionNoFamilies {
		item := newNavItem(
			"Families "+newCountBadge(len(c.document.Families())).String(),
			c.selectedTab == selectedFamiliesTab,
			pageFamilies(),
		)
		items = append(items, item)
	}

	if !optionNoSurnames {
		item := newNavItem(
			"Surnames "+newCountBadge(len(getSurnames(c.document))).String(),
			c.selectedTab == selectedSurnamesTab,
			pageSurnames(),
		)
		items = append(items, item)
	}

	if !optionNoSources {
		item := newNavItem(
			"Sources "+newCountBadge(len(c.document.Sources())).String(),
			c.selectedTab == selectedSourcesTab,
			pageSources(),
		)
		items = append(items, item)
	}

	if !optionNoStatistics {
		item := newNavItem(
			"Statistics",
			c.selectedTab == selectedStatisticsTab,
			pageStatistics(),
		)
		items = append(items, item)
	}

	if c.extraTab != "" {
		item := newNavItem(
			c.extraTab,
			c.selectedTab == selectedExtraTab,
			"#",
		)
		items = append(items, item)
	}

	return html.NewComponents(
		html.NewSpace(),
		newNavTabs(items),
		html.NewSpace(),
	).String()
}

var surnames = []string{}

func getSurnames(document *gedcom.Document) []string {
	if len(surnames) == 0 {
		for _, individual := range document.Individuals() {
			surname := individual.Name().Surname()
			if surname == "" || util.StringSliceContains(surnames, surname) {
				continue
			}

			surnames = append(surnames, surname)
		}

		sort.SliceStable(surnames, func(i, j int) bool {
			return strings.ToLower(surnames[i]) < strings.ToLower(surnames[j])
		})
	}

	return surnames
}
