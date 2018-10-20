package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
	"github.com/elliotchance/gedcom/util"
	"sort"
)

type diffPage struct {
	comparisons       []gedcom.IndividualComparison
	options           *gedcom.SimilarityOptions
	filterFlags       *util.FilterFlags
	googleAnalyticsID string
}

func newDiffPage(comparisons []gedcom.IndividualComparison, options *gedcom.SimilarityOptions, filterFlags *util.FilterFlags, googleAnalyticsID string) *diffPage {
	return &diffPage{
		comparisons:       comparisons,
		options:           options,
		filterFlags:       filterFlags,
		googleAnalyticsID: googleAnalyticsID,
	}
}

func (c *diffPage) sortByWrittenName(i, j int) bool {
	a := c.comparisons[i].Left
	b := c.comparisons[j].Left

	if a == nil {
		a = c.comparisons[i].Right
	}

	if b == nil {
		b = c.comparisons[j].Right
	}

	return a.Name().String() < b.Name().String()
}

func (c *diffPage) sortByHighestSimilarity(i, j int) bool {
	a := c.comparisons[i].Similarity.WeightedSimilarity(c.options)
	b := c.comparisons[j].Similarity.WeightedSimilarity(c.options)

	if a != b {
		// Greater than because we want the highest matches up the top.
		return a > b
	}

	// Fallback to sorting by name for non-matches
	return c.sortByWrittenName(i, j)
}

func (c *diffPage) sortComparisons() {
	sortFns := map[string]func(*diffPage, int, int) bool{
		optionSortWrittenName:       (*diffPage).sortByWrittenName,
		optionSortHighestSimilarity: (*diffPage).sortByHighestSimilarity,
	}

	sortFn := sortFns[optionSort]
	sort.SliceStable(c.comparisons, func(i, j int) bool {
		return sortFn(c, i, j)
	})
}

func (c *diffPage) String() string {
	rows := []fmt.Stringer{}

	c.sortComparisons()

	for _, comparison := range c.comparisons {
		if shouldSkip(comparison) {
			continue
		}

		weightedSimilarity := comparison.Similarity.WeightedSimilarity(c.options)

		leftClass := ""
		rightClass := ""

		switch {
		case comparison.Left != nil && comparison.Right == nil:
			leftClass = "bg-warning"

		case comparison.Left == nil && comparison.Right != nil:
			rightClass = "bg-primary"

		case weightedSimilarity < 1:
			leftClass = "bg-info"
			rightClass = "bg-info"

		case c.filterFlags.HideEqual:
			continue
		}

		leftNameAndDates := html.NewIndividualNameAndDatesLink(comparison.Left, true, "")
		rightNameAndDates := html.NewIndividualNameAndDatesLink(comparison.Right, true, "")

		left := html.NewTableCell(leftNameAndDates).Class(leftClass)
		right := html.NewTableCell(rightNameAndDates).Class(rightClass)

		middle := html.NewTableCell(html.NewText(""))
		if weightedSimilarity != 0 {
			similarityString := fmt.Sprintf("%.2f%%", weightedSimilarity*100)
			middle = html.NewTableCell(html.NewText(similarityString)).
				Class("text-center " + leftClass)
		}

		tableRow := html.NewTableRow(left, middle, right)

		rows = append(rows, tableRow)
	}

	// Individual pages
	components := []fmt.Stringer{
		html.NewBigTitle(1, "Individuals"),
		html.NewSpace(),
		html.NewTable("", rows...),
	}
	for _, comparison := range c.comparisons {
		if shouldSkip(comparison) {
			continue
		}

		compare := newIndividualCompare(comparison, c.filterFlags)
		components = append(components, compare)
	}

	return html.NewPage(
		"Comparison",
		html.NewComponents(components...),
		c.googleAnalyticsID,
	).String()
}

func shouldSkip(comparison gedcom.IndividualComparison) bool {
	switch optionShow {
	case optionShowAll:
		// Do nothing, we want to show all.

	case optionShowSubset:
		if gedcom.IsNil(comparison.Right) {
			return true
		}

	case optionShowOnlyMatches:
		if gedcom.IsNil(comparison.Left) || gedcom.IsNil(comparison.Right) {
			return true
		}
	}

	return false
}
