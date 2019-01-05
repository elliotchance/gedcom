package html_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
	"github.com/elliotchance/gedcom/util"
	"github.com/stretchr/testify/assert"
)

var (
	elliot = individual("P1", "Elliot /Chance/", "4 Jan 1843", "17 Mar 1907")
	john   = individual("P2", "John /Smith/", "4 Jan 1803", "17 Mar 1877")
	jane   = individual("P3", "Jane /Doe/", "3 Mar 1803", "14 June 1877")
)

func individual(pointer, fullName, birth, death string) *gedcom.IndividualNode {
	nodes := []gedcom.Node{}

	if fullName != "" {
		nodes = append(nodes, name(fullName))
	}

	if birth != "" {
		nodes = append(nodes, born(birth))
	}

	if death != "" {
		nodes = append(nodes, died(death))
	}

	return gedcom.NewIndividualNode(nil, "", pointer, nodes)
}

func name(value string) gedcom.Node {
	return gedcom.NewNameNode(nil, value, "", nil)
}

func born(value string) *gedcom.BirthNode {
	return gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, value, "", []gedcom.Node{}),
	})
}

func died(value string) gedcom.Node {
	return gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{
		gedcom.NewDateNode(nil, value, "", []gedcom.Node{}),
	})
}

func TestDiffPage_String(t *testing.T) {
	doc := gedcom.NewDocument()
	jane.SetDocument(doc)
	elliot.SetDocument(doc)
	john.SetDocument(doc)

	comparisons := gedcom.IndividualComparisons{
		{jane, jane, gedcom.NewSurroundingSimilarity(0.5, 1.0, 1.0, 1.0)},
		{elliot, nil, nil},
		{nil, john, nil},
	}
	filterFlags := &util.FilterFlags{}
	googleAnalyticsID := ""

	component := html.NewDiffPage(comparisons, filterFlags, googleAnalyticsID,
		html.DiffPageShowAll, html.DiffPageSortHighestSimilarity)

	assert.Contains(t, component.String(), "<html>")
	assert.Contains(t, component.String(), "<title>Comparison</title>")
	assert.Contains(t, component.String(),
		"Jane Doe (<em>b.</em> 3 Mar 1803&nbsp;&nbsp;&nbsp;<em>d.</em> 14 Jun 1877)")
	assert.Contains(t, component.String(),
		"Elliot Chance (<em>b.</em> 4 Jan 1843&nbsp;&nbsp;&nbsp;<em>d.</em> 17 Mar 1907)")
	assert.Contains(t, component.String(),
		"John Smith (<em>b.</em> 4 Jan 1803&nbsp;&nbsp;&nbsp;<em>d.</em> 17 Mar 1877)")
	assert.Contains(t, component.String(), "</html>")
}
