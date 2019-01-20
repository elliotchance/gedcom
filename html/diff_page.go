package html

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/util"
	"io"
	"sort"
)

// These are used for optionShow. If you update these options you will also
// need to adjust validateOptions.
const (
	DiffPageShowAll         = "all" // default
	DiffPageShowOnlyMatches = "only-matches"
	DiffPageShowSubset      = "subset"
)

// These are used for optionSort. If you update these options you will also
// need to adjust validateOptions.
const (
	DiffPageSortWrittenName       = "written-name" // default
	DiffPageSortHighestSimilarity = "highest-similarity"
)

type DiffPage struct {
	comparisons       gedcom.IndividualComparisons
	filterFlags       *util.FilterFlags
	googleAnalyticsID string
	sort              string
	show              string
	visibility        LivingVisibility
}

func NewDiffPage(comparisons gedcom.IndividualComparisons, filterFlags *util.FilterFlags, googleAnalyticsID string, show, sort string, visibility LivingVisibility) *DiffPage {
	return &DiffPage{
		comparisons:       comparisons,
		filterFlags:       filterFlags,
		googleAnalyticsID: googleAnalyticsID,
		show:              show,
		sort:              sort,
		visibility:        visibility,
	}
}

func (c *DiffPage) sortByWrittenName(i, j int) bool {
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

func (c *DiffPage) sortByHighestSimilarity(i, j int) bool {
	a := c.weightedSimilarity(c.comparisons[i])
	b := c.weightedSimilarity(c.comparisons[j])

	if a != b {
		// Greater than because we want the highest matches up the top.
		return a > b
	}

	// Fallback to sorting by name for non-matches
	return c.sortByWrittenName(i, j)
}

func (c *DiffPage) sortComparisons() {
	sortFns := map[string]func(*DiffPage, int, int) bool{
		DiffPageSortWrittenName:       (*DiffPage).sortByWrittenName,
		DiffPageSortHighestSimilarity: (*DiffPage).sortByHighestSimilarity,
	}

	sortFn := sortFns[c.sort]
	sort.SliceStable(c.comparisons, func(i, j int) bool {
		return sortFn(c, i, j)
	})
}

func (c *DiffPage) weightedSimilarity(comparison *gedcom.IndividualComparison) float64 {
	s := comparison.Similarity

	if s != nil {
		return s.WeightedSimilarity()
	}

	return 0.0
}

func (c *DiffPage) WriteTo(w io.Writer) (int64, error) {
	rows := []Component{}

	c.sortComparisons()

	for _, comparison := range c.comparisons {
		if c.shouldSkip(comparison) {
			continue
		}

		weightedSimilarity := c.weightedSimilarity(comparison)

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

		leftNameAndDates := NewIndividualNameAndDatesLink(comparison.Left, c.visibility, "")
		rightNameAndDates := NewIndividualNameAndDatesLink(comparison.Right, c.visibility, "")

		left := NewTableCell(leftNameAndDates).Class(leftClass)
		right := NewTableCell(rightNameAndDates).Class(rightClass)

		middle := NewTableCell(NewText(""))
		if weightedSimilarity != 0 {
			similarityString := fmt.Sprintf("%.2f%%", weightedSimilarity*100)
			middle = NewTableCell(NewText(similarityString)).
				Class("text-center " + leftClass)
		}

		tableRow := NewTableRow(left, middle, right)

		rows = append(rows, tableRow)
	}

	// Individual pages
	components := []Component{
		NewBigTitle(1, NewText("Individuals")),
		NewSpace(),
		NewTable("", rows...),
	}
	for _, comparison := range c.comparisons {
		if c.shouldSkip(comparison) {
			continue
		}

		compare := NewIndividualCompare(comparison, c.filterFlags, c.visibility)
		components = append(components, compare)
	}

	return NewPage(
		"Comparison",
		NewComponents(components...),
		c.googleAnalyticsID,
	).WriteTo(w)
}

func (c *DiffPage) shouldSkip(comparison *gedcom.IndividualComparison) bool {
	switch c.show {
	case DiffPageShowAll:
		// Do nothing, we want to show all.

	case DiffPageShowSubset:
		if gedcom.IsNil(comparison.Right) {
			return true
		}

	case DiffPageShowOnlyMatches:
		if gedcom.IsNil(comparison.Left) || gedcom.IsNil(comparison.Right) {
			return true
		}
	}

	return false
}
