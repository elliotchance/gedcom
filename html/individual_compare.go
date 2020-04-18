package html

import (
	"bytes"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
)

type IndividualCompare struct {
	comparison     *gedcom.IndividualComparison
	filterFlags    *gedcom.FilterFlags
	progress       chan gedcom.Progress
	compareOptions *gedcom.IndividualNodesCompareOptions
	visibility     LivingVisibility
	cache          []byte
	cacheErr       error
}

func NewIndividualCompare(comparison *gedcom.IndividualComparison, filterFlags *gedcom.FilterFlags, progress chan gedcom.Progress, compareOptions *gedcom.IndividualNodesCompareOptions, visibility LivingVisibility) *IndividualCompare {
	return &IndividualCompare{
		comparison:     comparison,
		filterFlags:    filterFlags,
		progress:       progress,
		compareOptions: compareOptions,
		visibility:     visibility,
	}
}

func (c *IndividualCompare) appendChildren(nd *gedcom.NodeDiff, prefix string) []core.Component {
	title := prefix + nd.Tag().String()
	tableRows := []core.Component{}

	row := NewDiffRow(title, nd, c.filterFlags.HideEqual)
	tableRows = c.appendDiffRow(tableRows, row)

	for _, child := range nd.Children {
		children := c.appendChildren(child, prefix+"&nbsp;&nbsp;&nbsp;&nbsp;")
		tableRows = append(tableRows, children...)
	}

	return tableRows
}

func (c *IndividualCompare) appendDiffRow(rows []core.Component, row *DiffRow) []core.Component {
	if row.isEmpty() {
		return rows
	}

	return append(rows, row)
}

func (c *IndividualCompare) isEmpty() bool {
	// Trigger cache.
	buf := bytes.NewBuffer(nil)
	n, _ := c.WriteHTMLTo(buf)

	return n == 0
}

func (c *IndividualCompare) addProgress() {
	if c.progress != nil {
		c.progress <- gedcom.Progress{
			Add: 1,
		}
	}
}

func (c *IndividualCompare) WriteHTMLTo(w io.Writer) (int64, error) {
	if c.cache == nil {
		buf := bytes.NewBuffer(nil)
		_, c.cacheErr = c.writeHTMLTo(buf)
		c.cache = buf.Bytes()
	}

	if c.cacheErr != nil {
		return 0, c.cacheErr
	}

	n, err := w.Write(c.cache)

	return int64(n), err
}

func (c *IndividualCompare) writeHTMLTo(w io.Writer) (int64, error) {
	left := c.comparison.Left
	right := c.comparison.Right

	c.addProgress()

	var name core.Component = nil

	if n := left; n != nil {
		name = NewIndividualNameAndDates(n, c.visibility, "")
	}

	if n := right; name == nil && n != nil {
		name = NewIndividualNameAndDates(n, c.visibility, "")
	}

	if name == nil {
		name = core.NewEmpty()
	}

	// We don't want the filters below to modify the original nodes in any way.
	doc := gedcom.NewDocument()

	if !gedcom.IsNil(left) {
		left = c.filterFlags.Filter(left, doc).(*gedcom.IndividualNode)
	}

	if !gedcom.IsNil(right) {
		right = c.filterFlags.Filter(right, doc).(*gedcom.IndividualNode)
	}

	diff := gedcom.CompareNodes(left, right)

	diff.Sort()

	tableRows := c.appendChildren(diff, "")

	// Parents
	leftParents := gedcom.IndividualNodes{}
	if !gedcom.IsNil(left) {
		for _, parents := range left.Parents() {
			if parent := parents.Husband(); parent != nil {
				leftParents = append(leftParents, parent.Individual())
			}
			if parent := parents.Wife(); parent != nil {
				leftParents = append(leftParents, parent.Individual())
			}
		}
	}

	rightParents := gedcom.IndividualNodes{}
	if !gedcom.IsNil(right) {
		for _, parents := range right.Parents() {
			if parent := parents.Husband(); parent != nil {
				rightParents = append(rightParents, parent.Individual())
			}
			if parent := parents.Wife(); parent != nil {
				rightParents = append(rightParents, parent.Individual())
			}
		}
	}

	for _, parents := range leftParents.Compare(rightParents, c.compareOptions) {
		var row *DiffRow
		name := "Parent"

		switch {
		case !gedcom.IsNil(parents.Left) && !gedcom.IsNil(parents.Right):
			row = NewDiffRow(name, &gedcom.NodeDiff{
				Left:  parents.Left,
				Right: parents.Right,
			}, c.filterFlags.HideEqual)

		case !gedcom.IsNil(parents.Left):
			row = NewDiffRow(name, &gedcom.NodeDiff{
				Left: parents.Left,
			}, c.filterFlags.HideEqual)

		case !gedcom.IsNil(parents.Right):
			row = NewDiffRow(name, &gedcom.NodeDiff{
				Right: parents.Right,
			}, c.filterFlags.HideEqual)
		}

		tableRows = c.appendDiffRow(tableRows, row)
	}

	// Spouses
	switch {
	case !gedcom.IsNil(left) && !gedcom.IsNil(right):
		for _, spouse := range left.Spouses().Compare(right.Spouses(), c.compareOptions) {
			nodeDiff := &gedcom.NodeDiff{}

			if spouse.Left != nil {
				nodeDiff.Left = spouse.Left
			}

			if spouse.Right != nil {
				nodeDiff.Right = spouse.Right
			}

			row := NewDiffRow("Spouse", nodeDiff, c.filterFlags.HideEqual)

			tableRows = c.appendDiffRow(tableRows, row)
		}

	case !gedcom.IsNil(left):
		for _, spouse := range left.Spouses() {
			row := NewDiffRow("Spouse", &gedcom.NodeDiff{
				Left: spouse,
			}, c.filterFlags.HideEqual)

			tableRows = c.appendDiffRow(tableRows, row)
		}

	case !gedcom.IsNil(right):
		for _, spouse := range right.Spouses() {
			row := NewDiffRow("Spouse", &gedcom.NodeDiff{
				Right: spouse,
			}, c.filterFlags.HideEqual)

			tableRows = c.appendDiffRow(tableRows, row)
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

	// We should not show the header if the content would be blank.
	if len(tableRows) == 0 {
		return writeNothing()
	}

	return core.NewComponents(
		core.NewAnchor(leftAnchor),
		core.NewAnchor(rightAnchor),
		core.NewCard(name, core.CardNoBadgeCount,
			core.NewTable("", tableRows...)),
	).WriteHTMLTo(w)
}
