package html

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
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
	filterFlags       *gedcom.FilterFlags
	googleAnalyticsID string
	sort              string
	show              string
	progress          chan gedcom.Progress
	compareOptions    *gedcom.IndividualNodesCompareOptions
	visibility        LivingVisibility
	leftGedcomPath    string
	rightGedcomPath   string
}

func NewDiffPage(comparisons gedcom.IndividualComparisons, filterFlags *gedcom.FilterFlags, googleAnalyticsID string, show, sort string, progress chan gedcom.Progress, compareOptions *gedcom.IndividualNodesCompareOptions, visibility LivingVisibility, leftGedcomPath string, rightGedcomPath string) *DiffPage {
	return &DiffPage{
		comparisons:       comparisons,
		filterFlags:       filterFlags,
		googleAnalyticsID: googleAnalyticsID,
		show:              show,
		sort:              sort,
		progress:          progress,
		compareOptions:    compareOptions,
		visibility:        visibility,
		leftGedcomPath: leftGedcomPath,
		rightGedcomPath: rightGedcomPath,
	}
}

func (c *DiffPage) sortByWrittenName(comparisons []*IndividualCompare, i, j int) bool {
	a := comparisons[i].comparison.Left
	b := comparisons[j].comparison.Left

	if a == nil {
		a = comparisons[i].comparison.Right
	}

	if b == nil {
		b = comparisons[j].comparison.Right
	}

	aName := a.Name().String()
	bName := b.Name().String()

	return aName < bName
}

func (c *DiffPage) sortByHighestSimilarity(comparisons []*IndividualCompare, i, j int) bool {
	a := c.weightedSimilarity(comparisons[i].comparison)
	b := c.weightedSimilarity(comparisons[j].comparison)

	if a != b {
		// Greater than because we want the highest matches up the top.
		return a > b
	}

	// Fallback to sorting by name for non-matches
	return c.sortByWrittenName(comparisons, i, j)
}

func (c *DiffPage) sortComparisons(comparisons []*IndividualCompare) {
	sortFns := map[string]func(*DiffPage, []*IndividualCompare, int, int) bool{
		DiffPageSortWrittenName:       (*DiffPage).sortByWrittenName,
		DiffPageSortHighestSimilarity: (*DiffPage).sortByHighestSimilarity,
	}

	sortFn := sortFns[c.sort]
	sort.SliceStable(comparisons, func(i, j int) bool {
		return sortFn(c, comparisons, i, j)
	})
}

func (c *DiffPage) weightedSimilarity(comparison *gedcom.IndividualComparison) float64 {
	s := comparison.Similarity

	if s != nil {
		return s.WeightedSimilarity()
	}

	return 0.0
}

func (c *DiffPage) createJobs() chan *IndividualCompare {
	jobs := make(chan *IndividualCompare, 10)

	go func() {
		for _, comparison := range c.comparisons {
			jobs <- NewIndividualCompare(comparison,
				c.filterFlags, c.progress, c.compareOptions, c.visibility)
		}

		close(jobs)
	}()

	return jobs
}

func (c *DiffPage) processJobs(jobs chan *IndividualCompare) chan *IndividualCompare {
	results := make(chan *IndividualCompare, 10)

	go func() {
		util.WorkerPool(c.compareOptions.ConcurrentJobs(), func(_ int) {
			for job := range jobs {
				if c.shouldSkip(job) {
					continue
				}

				results <- job
			}
		})

		close(results)
	}()

	return results
}

func (c *DiffPage) sortResults(in chan *IndividualCompare) chan *IndividualCompare {
	out := make(chan *IndividualCompare, 10)

	go func() {
		// We have to read all results before they can be sorted.
		all := []*IndividualCompare{}
		for comparison := range in {
			all = append(all, comparison)
		}

		c.sortComparisons(all)

		// Send all results back. We expect there is only one receiver for this
		// to work.
		for _, item := range all {
			out <- item
		}

		close(out)
	}()

	return out
}

func (c *DiffPage) WriteHTMLTo(w io.Writer) (int64, error) {
	if c.progress != nil {
		c.progress <- gedcom.Progress{
			Total: int64(len(c.comparisons)),
		}
	}

	jobs := c.createJobs()
	results := c.processJobs(jobs)
	results = c.sortResults(results)

	precalculatedComparisons := []*IndividualCompare{}

	for comparison := range results {
		precalculatedComparisons = append(precalculatedComparisons, comparison)
	}

	// The index at the top of the page.
	var rows []core.Component
	numOnlyLeft := 0
	numOnlyRight := 0
	numSimilar := 0
	numEqual := 0
	for _, comparison := range precalculatedComparisons {
		weightedSimilarity := c.weightedSimilarity(comparison.comparison)

		leftClass := ""
		rightClass := ""

		switch {
		case comparison.comparison.Left != nil && comparison.comparison.Right == nil: //right is missing
			leftClass = "bg-warning"
			numOnlyLeft++

		case comparison.comparison.Left == nil && comparison.comparison.Right != nil: //left is missing
			rightClass = "bg-primary"
			numOnlyRight++

		case weightedSimilarity < 1: //neither are missing, but they aren't identical
			leftClass = "bg-info"
			rightClass = "bg-info"
			numSimilar++

		case c.filterFlags.HideEqual: //are identical, but user said to hide equals
			numEqual++
			continue
		default:
			numEqual++
		}
		rows = append(rows, c.getRow(comparison, leftClass, rightClass, weightedSimilarity))
	}

	leftHeader := fmt.Sprint(c.leftGedcomPath, " (", numOnlyLeft, " only in left)")
	rightHeader := fmt.Sprint(c.rightGedcomPath, " (", numOnlyRight, " only in right)")
	class := "text-center"
	attr := map[string]string{}
	headerTag := "h5"
	wereHidden := ""
	if c.filterFlags.HideEqual {
		wereHidden = " - were hidden"
	}
	middleHeader := fmt.Sprint("Similarity score", " (", numSimilar, " similar, and ", numEqual, " equal", wereHidden, ")")
	header := []core.Component{core.NewTableRow(
		core.NewTableCell(
			core.NewTag(headerTag, attr, core.NewText(leftHeader))).Class(class),
		core.NewTableCell(
			core.NewTag(headerTag, attr, core.NewText(middleHeader))).Class(class),
		core.NewTableCell(
			core.NewTag(headerTag, attr, core.NewText(rightHeader))).Class(class))}

	// Individual pages
	components := []core.Component{
		core.NewSpace(),
		core.NewCard(core.NewText("Individuals"), core.CardNoBadgeCount,
			core.NewTable("",  append(header, rows...)...)),
		core.NewSpace(),
	}
	for _, comparison := range precalculatedComparisons {
		components = append(components, comparison, core.NewSpace())
	}

	return core.NewPage(
		"Comparison",
		core.NewRow(core.NewColumn(core.EntireRow, core.NewComponents(components...))),
		c.googleAnalyticsID,
	).WriteHTMLTo(w)
}

func (c *DiffPage) getRow(comparison *IndividualCompare, leftClass string, rightClass string, weightedSimilarity float64) *core.TableRow {

	leftNameAndDates := NewIndividualNameAndDatesLink(comparison.comparison.Left, c.visibility, "")
	rightNameAndDates := NewIndividualNameAndDatesLink(comparison.comparison.Right, c.visibility, "")

	left := core.NewTableCell(leftNameAndDates).Class(leftClass)
	right := core.NewTableCell(rightNameAndDates).Class(rightClass)

	middle := core.NewTableCell(core.NewText(""))
	if weightedSimilarity != 0 {
		similarityString := fmt.Sprintf("%.2f%%", weightedSimilarity*100)
		middle = core.NewTableCell(core.NewText(similarityString)).
			Class("text-center " + leftClass)
	}

	return core.NewTableRow(left, middle, right)
}

func (c *DiffPage) shouldSkip(comparison *IndividualCompare) bool {
	switch c.show {
	case DiffPageShowAll:
		// Do nothing, we want to show all.

	case DiffPageShowSubset:
		if gedcom.IsNil(comparison.comparison.Right) {
			return true
		}

	case DiffPageShowOnlyMatches:
		if gedcom.IsNil(comparison.comparison.Left) || gedcom.IsNil(comparison.comparison.Right) {
			return true
		}
	}

	return comparison.isEmpty()
}
