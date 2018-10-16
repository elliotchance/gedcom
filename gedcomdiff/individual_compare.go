package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
	"github.com/elliotchance/gedcom/util"
)

type individualCompare struct {
	comparison  gedcom.IndividualComparison
	filterFlags *util.FilterFlags
}

func newIndividualCompare(comparison gedcom.IndividualComparison, filterFlags *util.FilterFlags) *individualCompare {
	return &individualCompare{
		comparison:  comparison,
		filterFlags: filterFlags,
	}
}

func (c *individualCompare) appendChildren(nd *gedcom.NodeDiff, prefix string) []fmt.Stringer {
	title := prefix + nd.Tag().String()
	row := newDiffRow(title, nd, c.filterFlags.HideEqual)
	tableRows := []fmt.Stringer{row}

	for _, child := range nd.Children {
		children := c.appendChildren(child, prefix+"&nbsp;&nbsp;&nbsp;&nbsp;")
		tableRows = append(tableRows, children...)
	}

	return tableRows
}

func (c *individualCompare) String() string {
	left := c.comparison.Left
	right := c.comparison.Right

	name := ""
	if n := left; n != nil {
		name = html.NewIndividualNameAndDates(n, true, "").String()
	}
	if n := right; name == "" && n != nil {
		name = html.NewIndividualNameAndDates(n, true, "").String()
	}

	if !gedcom.IsNil(left) {
		left = c.filterFlags.Filter(left).(*gedcom.IndividualNode)
	}

	if !gedcom.IsNil(right) {
		right = c.filterFlags.Filter(right).(*gedcom.IndividualNode)
	}

	diff := gedcom.CompareNodes(left, right)

	diff.Sort()

	tableRows := c.appendChildren(diff, "")

	// Parents
	leftParents := gedcom.IndividualNodes{}
	if !gedcom.IsNil(left) {
		for _, parents := range left.Parents() {
			if parent := parents.Husband(); parent != nil {
				leftParents = append(leftParents, parent)
			}
			if parent := parents.Wife(); parent != nil {
				leftParents = append(leftParents, parent)
			}
		}
	}

	rightParents := gedcom.IndividualNodes{}
	if !gedcom.IsNil(right) {
		for _, parents := range right.Parents() {
			if parent := parents.Husband(); parent != nil {
				rightParents = append(rightParents, parent)
			}
			if parent := parents.Wife(); parent != nil {
				rightParents = append(rightParents, parent)
			}
		}
	}

	options := gedcom.NewSimilarityOptions()
	compareOptions := &gedcom.IndividualNodesCompareOptions{
		SimilarityOptions: options,
	}
	for _, parents := range leftParents.Compare(rightParents, compareOptions) {
		var row *diffRow
		name := "Parent"

		switch {
		case !gedcom.IsNil(parents.Left) && !gedcom.IsNil(parents.Right):
			row = newDiffRow(name, &gedcom.NodeDiff{
				Left:  parents.Left.Name(),
				Right: parents.Right.Name(),
			}, c.filterFlags.HideEqual)

		case !gedcom.IsNil(parents.Left):
			row = newDiffRow(name, &gedcom.NodeDiff{
				Left: parents.Left.Name(),
			}, c.filterFlags.HideEqual)

		case !gedcom.IsNil(parents.Right):
			row = newDiffRow(name, &gedcom.NodeDiff{
				Right: parents.Right.Name(),
			}, c.filterFlags.HideEqual)
		}

		tableRows = append(tableRows, row)
	}

	// Spouses
	switch {
	case !gedcom.IsNil(left) && !gedcom.IsNil(right):
		for _, spouse := range left.Spouses().Compare(right.Spouses(), compareOptions) {
			nodeDiff := &gedcom.NodeDiff{}

			if spouse.Left != nil {
				nodeDiff.Left = spouse.Left.Name()
			}

			if spouse.Right != nil {
				nodeDiff.Right = spouse.Right.Name()
			}

			row := newDiffRow("Spouse", nodeDiff, c.filterFlags.HideEqual)
			tableRows = append(tableRows, row)
		}

	case !gedcom.IsNil(left):
		for _, spouse := range left.Spouses() {
			row := newDiffRow("Spouse", &gedcom.NodeDiff{
				Left: spouse.Name(),
			}, c.filterFlags.HideEqual)
			tableRows = append(tableRows, row)
		}

	case !gedcom.IsNil(right):
		for _, spouse := range right.Spouses() {
			row := newDiffRow("Spouse", &gedcom.NodeDiff{
				Right: spouse.Name(),
			}, c.filterFlags.HideEqual)
			tableRows = append(tableRows, row)
		}
	}

	leftAnchor := ""
	rightAnchor := ""

	if c.comparison.Left != nil {
		leftAnchor = c.comparison.Left.Pointer()
	}

	if c.comparison.Right != nil {
		rightAnchor = c.comparison.Right.Pointer()
	}

	return html.NewComponents(
		html.NewAnchor(leftAnchor),
		html.NewAnchor(rightAnchor),
		html.NewBigTitle(1, name),
		html.NewSpace(),
		html.NewTable("", tableRows...),
	).String()
}
