package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html"
	"strings"
)

type individualCompare struct {
	comparison    gedcom.IndividualComparison
	includePlaces bool
	hideSame      bool
}

func newIndividualCompare(comparison gedcom.IndividualComparison, includePlaces, hideSame bool) *individualCompare {
	return &individualCompare{
		comparison:    comparison,
		includePlaces: includePlaces,
		hideSame:      hideSame,
	}
}

func getName(node *gedcom.IndividualNode) string {
	if node == nil {
		return ""
	}

	return node.Name().String()
}

func getTag(node *gedcom.IndividualNode, tags ...gedcom.Tag) string {
	if node == nil {
		return ""
	}

	if n := gedcom.First(gedcom.NodesWithTagPath(node, tags...)); n != nil {
		return n.Value()
	}

	return ""
}

func (c *individualCompare) String() string {
	left := c.comparison.Left
	right := c.comparison.Right

	name := ""
	if n := left; n != nil {
		name = n.Name().String()
	}
	if n := right; name == "" && n != nil {
		name = n.Name().String()
	}

	diffRows := []struct {
		name string
		tags []gedcom.Tag
	}{
		{"Birth Date", []gedcom.Tag{gedcom.TagBirth, gedcom.TagDate}},
		{"Birth Place", []gedcom.Tag{gedcom.TagBirth, gedcom.TagPlace}},
		{"Baptism Date", []gedcom.Tag{gedcom.TagChristening, gedcom.TagDate}},
		{"Baptism Place", []gedcom.Tag{gedcom.TagChristening, gedcom.TagPlace}},
		{"Death Date", []gedcom.Tag{gedcom.TagDeath, gedcom.TagDate}},
		{"Death Place", []gedcom.Tag{gedcom.TagDeath, gedcom.TagPlace}},
		{"Burial Date", []gedcom.Tag{gedcom.TagBurial, gedcom.TagDate}},
		{"Burial Place", []gedcom.Tag{gedcom.TagBurial, gedcom.TagPlace}},
	}

	tableRows := []fmt.Stringer{
		newDiffRow("Name", getName(left), getName(right), c.hideSame),
	}
	for _, diffRow := range diffRows {
		if !c.includePlaces && strings.HasSuffix(diffRow.name, " Place") {
			continue
		}

		tableRows = append(tableRows, newDiffRow(diffRow.name,
			getTag(left, diffRow.tags...),
			getTag(right, diffRow.tags...),
			c.hideSame,
		))
	}

	// Parents
	fatherLeft := ""
	motherLeft := ""
	if c.comparison.Left != nil {
		if p := c.comparison.Left.Parents(); len(p) > 0 {
			if n := p[0].Husband(); n != nil && n.Name() != nil {
				fatherLeft = n.Name().String()
			}
			if n := p[0].Wife(); n != nil && n.Name() != nil {
				motherLeft = n.Name().String()
			}
		}
	}

	fatherRight := ""
	motherRight := ""
	if c.comparison.Right != nil {
		if p := c.comparison.Right.Parents(); len(p) > 0 {
			if n := p[0].Husband(); n != nil && n.Name() != nil {
				fatherRight = n.Name().String()
			}
			if n := p[0].Wife(); n != nil && n.Name() != nil {
				motherRight = n.Name().String()
			}
		}
	}

	tableRows = append(tableRows,
		newDiffRow("Father", fatherLeft, fatherRight, c.hideSame),
		newDiffRow("Mother", motherLeft, motherRight, c.hideSame),
	)

	return html.NewComponents(
		html.NewBigTitle(name),
		html.NewSpace(),
		html.NewTable("", tableRows...),
	).String()
}
