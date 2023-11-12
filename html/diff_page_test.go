package html_test

import (
	"testing"

	"bytes"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/gedcom/v39/html"
	"github.com/stretchr/testify/assert"
)

func individual(doc *gedcom.Document, pointer, fullName, birth, death string) *gedcom.IndividualNode {
	individual := doc.AddIndividual(pointer)

	if fullName != "" {
		individual.AddNode(gedcom.NewNameNode(fullName))
	}

	if birth != "" {
		individual.AddNode(gedcom.NewBirthNode("", gedcom.NewDateNode(birth)))
	}

	if death != "" {
		individual.AddNode(gedcom.NewDeathNode("", gedcom.NewDateNode(death)))
	}

	return individual
}

func TestDiffPage_WriteHTMLTo(t *testing.T) {
	doc := gedcom.NewDocument()
	elliot := individual(doc, "P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907")
	john := individual(doc, "P2", "John /Smith/", "4 Jan 1803", "17 Mar 1877")
	jane := individual(doc, "P3", "Jane /Doe/", "3 Mar 1803", "14 June 1877")

	comparisons := gedcom.IndividualComparisons{
		gedcom.NewIndividualComparison(jane, jane, gedcom.NewSurroundingSimilarity(0.5, 1.0, 1.0, 1.0)),
		gedcom.NewIndividualComparison(elliot, nil, nil),
		gedcom.NewIndividualComparison(nil, john, nil),
	}
	filterFlags := &gedcom.FilterFlags{}
	googleAnalyticsID := ""

	compareOptions := gedcom.NewIndividualNodesCompareOptions()
	component := html.NewDiffPage(comparisons, filterFlags, googleAnalyticsID,
		html.DiffPageShowAll, html.DiffPageSortHighestSimilarity, nil,
		compareOptions, html.LivingVisibilityPlaceholder, "left", "right")

	buf := bytes.NewBuffer(nil)
	component.WriteHTMLTo(buf)
	s := string(buf.Bytes())

	assert.Contains(t, s, "<html>")
	assert.Contains(t, s, "<title>Comparison</title>")
	assert.Contains(t, s,
		"Jane Doe (<em>b.</em> 3 Mar 1803&nbsp;&nbsp;&nbsp;<em>d.</em> 14 Jun 1877)")
	assert.Contains(t, s,
		"Elliot Chance (<em>b.</em> 4 Jan 1843&nbsp;&nbsp;&nbsp;<em>d.</em> 17 Mar 1907)")
	assert.Contains(t, s,
		"John Smith (<em>b.</em> 4 Jan 1803&nbsp;&nbsp;&nbsp;<em>d.</em> 17 Mar 1877)")
	assert.Contains(t, s, "</html>")
}
