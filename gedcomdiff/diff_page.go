package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
	"sort"
)

type diffPage struct {
	comparisons   []gedcom.IndividualComparison
	options       *gedcom.SimilarityOptions
	includePlaces bool
	hideSame      bool
}

func newDiffPage(comparisons []gedcom.IndividualComparison, options *gedcom.SimilarityOptions, includePlaces, hideSame bool) *diffPage {
	return &diffPage{
		comparisons:   comparisons,
		options:       options,
		includePlaces: includePlaces,
		hideSame:      hideSame,
	}
}

func (c *diffPage) String() string {
	rows := []fmt.Stringer{}

	sort.SliceStable(c.comparisons, func(i, j int) bool {
		a := c.comparisons[i].Left
		b := c.comparisons[j].Left

		if a == nil {
			a = c.comparisons[i].Right
		}

		if b == nil {
			b = c.comparisons[j].Right
		}

		return a.Name().String() < b.Name().String()
	})

	for _, comparison := range c.comparisons {
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
		}

		rows = append(rows, html.NewTableRow(
			html.NewTableCell(leftClass,
				html.NewIndividualNameAndDates(comparison.Left, true, "")),
			html.NewTableCell(rightClass,
				html.NewIndividualNameAndDates(comparison.Right, true, "")),
		))
	}

	// Individual pages
	components := []fmt.Stringer{
		html.NewBigTitle("Individuals"),
		html.NewSpace(),
		html.NewTable("", rows...),
	}
	for _, comparison := range c.comparisons {
		components = append(components,
			newIndividualCompare(comparison, c.includePlaces, c.hideSame))
	}

	return html.NewPage(
		"Comparison",
		html.NewComponents(components...),
	).String()
}
