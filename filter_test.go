package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	root := gedcom.NewDocument().AddIndividual("P1",
		gedcom.NewNameNode("Elliot /Chance/"),
		gedcom.NewBirthNode("",
			gedcom.NewDateNode("6 MAY 1989"),
		),
	)

	for _, test := range []struct {
		filter   gedcom.FilterFunction
		expected string
	}{
		{
			filter: func(node gedcom.Node) (gedcom.Node, bool) {
				return nil, false
			},
			expected: "",
		},
		{
			filter: func(node gedcom.Node) (gedcom.Node, bool) {
				return node, true
			},
			expected: `0 @P1@ INDI
1 NAME Elliot /Chance/
1 BIRT
2 DATE 6 MAY 1989
`,
		},
		{
			filter: func(node gedcom.Node) (gedcom.Node, bool) {
				return node.ShallowCopy(), false
			},
			expected: `0 @P1@ INDI
`,
		},
		{
			filter: func(node gedcom.Node) (gedcom.Node, bool) {
				return node, false
			},
			expected: `0 @P1@ INDI
1 NAME Elliot /Chance/
1 BIRT
2 DATE 6 MAY 1989
`,
		},
		{
			filter: func(node gedcom.Node) (gedcom.Node, bool) {
				if node.Tag().Is(gedcom.TagIndividual) {
					// false means it will not traverse children, since an
					// individual can never be inside of another individual.
					return node.ShallowCopy(), false
				}

				return nil, false
			},
			expected: `0 @P1@ INDI
`,
		},
		{
			filter: func(node gedcom.Node) (gedcom.Node, bool) {
				if node.Tag().Is(gedcom.TagIndividual) {
					// false means it will not traverse children, since an
					// individual can never be inside of another individual.
					return node, false
				}

				return nil, false
			},
			expected: `0 @P1@ INDI
1 NAME Elliot /Chance/
1 BIRT
2 DATE 6 MAY 1989
`,
		},
		{
			filter: func(node gedcom.Node) (gedcom.Node, bool) {
				t := node.Tag()
				return gedcom.NodeCondition(
					t.Is(gedcom.TagIndividual) || t.Is(gedcom.TagName),
					node,
					nil,
				), true
			},
			expected: `0 @P1@ INDI
1 NAME Elliot /Chance/
`,
		},
		{
			filter: func(node gedcom.Node) (gedcom.Node, bool) {
				t := node.Tag()
				return gedcom.NodeCondition(
					t.Is(gedcom.TagIndividual) || t.Is(gedcom.TagDate),
					node,
					nil,
				), true
			},
			expected: `0 @P1@ INDI
`,
		},
		{
			filter: func(node gedcom.Node) (gedcom.Node, bool) {
				t := node.Tag()
				return gedcom.NodeCondition(
					t.Is(gedcom.TagIndividual) || t.Is(gedcom.TagBirth),
					node,
					nil,
				), true
			},
			expected: `0 @P1@ INDI
1 BIRT
`,
		},
		{
			filter: func(node gedcom.Node) (gedcom.Node, bool) {
				t := node.Tag()
				return gedcom.NodeCondition(
					t.Is(gedcom.TagIndividual) || t.Is(gedcom.TagBirth) || t.Is(gedcom.TagDate),
					node,
					nil,
				), true
			},
			expected: `0 @P1@ INDI
1 BIRT
2 DATE 6 MAY 1989
`,
		},
		{
			filter: func(node gedcom.Node) (gedcom.Node, bool) {
				if node.Tag().Is(gedcom.TagName) {
					return gedcom.NewDateNode("1 APR 1943"), true
				}

				return node, true
			},
			expected: `0 @P1@ INDI
1 DATE 1 APR 1943
1 BIRT
2 DATE 6 MAY 1989
`,
		},
	} {
		t.Run("", func(t *testing.T) {
			doc := gedcom.NewDocument()
			node := gedcom.Filter(root, doc, test.filter)
			result := gedcom.GEDCOMString(node, 0)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestWhitelistTagFilter(t *testing.T) {
	root := gedcom.NewDocument().AddIndividual("P1",
		gedcom.NewNameNode("Elliot /Chance/"),
		gedcom.NewBirthNode("",
			gedcom.NewDateNode("6 MAY 1989"),
		),
	)

	for _, test := range []struct {
		tags     []gedcom.Tag
		expected string
	}{
		{
			tags:     []gedcom.Tag{},
			expected: ``,
		},
		{
			tags: []gedcom.Tag{gedcom.TagIndividual},
			expected: `0 @P1@ INDI
`,
		},
		{
			tags:     []gedcom.Tag{gedcom.TagBirth},
			expected: ``,
		},
		{
			tags: []gedcom.Tag{gedcom.TagBirth, gedcom.TagIndividual},
			expected: `0 @P1@ INDI
1 BIRT
`,
		},
	} {
		t.Run("", func(t *testing.T) {
			filter := gedcom.WhitelistTagFilter(test.tags...)
			doc := gedcom.NewDocument()
			result := gedcom.GEDCOMString(gedcom.Filter(root, doc, filter), 0)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestBlacklistTagFilter(t *testing.T) {
	root := gedcom.NewDocument().AddIndividual("P1",
		gedcom.NewNameNode("Elliot /Chance/"),
		gedcom.NewBirthNode("",
			gedcom.NewDateNode("6 MAY 1989"),
		),
	)

	for _, test := range []struct {
		tags     []gedcom.Tag
		expected string
	}{
		{
			tags: []gedcom.Tag{},
			expected: `0 @P1@ INDI
1 NAME Elliot /Chance/
1 BIRT
2 DATE 6 MAY 1989
`,
		},
		{
			tags:     []gedcom.Tag{gedcom.TagIndividual},
			expected: ``,
		},
		{
			tags: []gedcom.Tag{gedcom.TagBirth},
			expected: `0 @P1@ INDI
1 NAME Elliot /Chance/
`,
		},
		{
			tags:     []gedcom.Tag{gedcom.TagBirth, gedcom.TagIndividual},
			expected: ``,
		},
	} {
		t.Run("", func(t *testing.T) {
			filter := gedcom.BlacklistTagFilter(test.tags...)
			doc := gedcom.NewDocument()
			result := gedcom.GEDCOMString(gedcom.Filter(root, doc, filter), 0)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestOfficialTagFilter(t *testing.T) {
	root := gedcom.NewDocument().AddIndividual("P1",
		gedcom.NewNode(gedcom.UnofficialTagCreated, "Elliot /Chance/", "",
			gedcom.NewDateNode("3 Mar 2007"),
		),
		gedcom.NewBirthNode("",
			gedcom.NewDateNode("6 MAY 1989"),
		),
	)

	for _, test := range []struct {
		expected string
	}{
		{
			expected: `0 @P1@ INDI
1 BIRT
2 DATE 6 MAY 1989
`,
		},
	} {
		t.Run("", func(t *testing.T) {
			filter := gedcom.OfficialTagFilter()
			doc := gedcom.NewDocument()
			result := gedcom.GEDCOMString(gedcom.Filter(root, doc, filter), 0)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSimpleNameFilter(t *testing.T) {
	// ghost:ignore
	for _, test := range []struct {
		root     gedcom.Node
		format   gedcom.NameFormat
		expected string
	}{
		{
			root: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot /Chance/"),
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("6 MAY 1989"),
				),
			),
			format: gedcom.NameFormatGEDCOM,
			expected: `0 @P1@ INDI
1 NAME Elliot /Chance/
1 BIRT
2 DATE 6 MAY 1989
`,
		},
		{
			root: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("6 MAY 1989"),
				),
				gedcom.NewNameNode("Elliot /Chance/",
					gedcom.NewNode(gedcom.TagSurname, "Smith", ""),
				),
			),
			format: gedcom.NameFormatGEDCOM,
			expected: `0 @P1@ INDI
1 BIRT
2 DATE 6 MAY 1989
1 NAME Elliot /Smith/
`,
		},
		{
			root: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("",
					gedcom.NewNode(gedcom.TagGivenName, "Bob", ""),
					gedcom.NewNode(gedcom.TagSurname, "Smith", ""),
				),
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("6 MAY 1989"),
				),
			),
			format: gedcom.NameFormatGEDCOM,
			expected: `0 @P1@ INDI
1 NAME Bob /Smith/
1 BIRT
2 DATE 6 MAY 1989
`,
		},
		{
			root: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot /Chance/"),
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("6 MAY 1989"),
				),
			),
			format: gedcom.NameFormatWritten,
			expected: `0 @P1@ INDI
1 NAME Elliot Chance
1 BIRT
2 DATE 6 MAY 1989
`,
		},
		{
			root: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("",
					gedcom.NewNode(gedcom.TagGivenName, "Bob", ""),
					gedcom.NewNode(gedcom.TagSurname, "Smith", ""),
				),
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("6 MAY 1989"),
				),
			),
			format: gedcom.NameFormatIndex,
			expected: `0 @P1@ INDI
1 NAME Smith, Bob
1 BIRT
2 DATE 6 MAY 1989
`,
		},
	} {
		t.Run("", func(t *testing.T) {
			filter := gedcom.SimpleNameFilter(test.format)
			doc := gedcom.NewDocument()
			result := gedcom.GEDCOMString(gedcom.Filter(test.root, doc, filter), 0)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestOnlyVitalsTagFilter(t *testing.T) {
	// ghost:ignore
	for testName, test := range map[string]struct {
		root     gedcom.Node
		expected string
	}{
		"SimpleNameAndBirthDate": {
			root: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot /Chance/"),
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("6 MAY 1989"),
				),
			),
			expected: `0 @P1@ INDI
1 NAME Elliot /Chance/
1 BIRT
2 DATE 6 MAY 1989
`,
		},
		"ComplexName1AndDeathDate": {
			root: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewDeathNode("",
					gedcom.NewDateNode("6 MAY 1989"),
				),
				gedcom.NewNameNode("Elliot /Chance/",
					gedcom.NewNode(gedcom.TagSurname, "Smith", ""),
				),
			),
			expected: `0 @P1@ INDI
1 DEAT
2 DATE 6 MAY 1989
1 NAME Elliot /Chance/
2 SURN Smith
`,
		},
		"ComplexName2AndBirthPlace": {
			root: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("",
					gedcom.NewNode(gedcom.TagGivenName, "Bob", ""),
					gedcom.NewNode(gedcom.TagSurname, "Smith", ""),
				),
				gedcom.NewBirthNode("",
					gedcom.NewPlaceNode("Sydney, Australia"),
				),
			),
			expected: `0 @P1@ INDI
1 NAME
2 GIVN Bob
2 SURN Smith
1 BIRT
2 PLAC Sydney, Australia
`,
		},
		"Source": {
			root: gedcom.NewSourceNode("", "P1",
				gedcom.NewNameNode("",
					gedcom.NewNode(gedcom.TagGivenName, "Bob", ""),
					gedcom.NewNode(gedcom.TagSurname, "Smith", ""),
				),
				gedcom.NewBirthNode("",
					gedcom.NewPlaceNode("Sydney, Australia"),
				),
			),
			expected: ``,
		},
		"IndividualNote": {
			root: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("",
					gedcom.NewNode(gedcom.TagGivenName, "Bob", ""),
					gedcom.NewNode(gedcom.TagSurname, "Smith", ""),
				),
				gedcom.NewBirthNode("",
					gedcom.NewPlaceNode("Sydney, Australia"),
				),
				gedcom.NewNode(gedcom.TagNote, "foo", ""),
			),
			expected: `0 @P1@ INDI
1 NAME
2 GIVN Bob
2 SURN Smith
1 BIRT
2 PLAC Sydney, Australia
`,
		},
		"Burial": {
			root: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("",
					gedcom.NewNode(gedcom.TagGivenName, "Bob", ""),
					gedcom.NewNode(gedcom.TagTitle, "Smith", ""),
				),
				gedcom.NewBurialNode("",
					gedcom.NewPlaceNode("6 MAY 1989"),
					gedcom.NewPlaceNode("Sydney, Australia"),
				),
			),
			expected: `0 @P1@ INDI
1 NAME
2 GIVN Bob
2 TITL Smith
1 BURI
2 PLAC 6 MAY 1989
2 PLAC Sydney, Australia
`,
		},
		"Baptism": {
			root: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("",
					gedcom.NewNode(gedcom.TagNameSuffix, "Bob", ""),
					gedcom.NewNode(gedcom.TagSurnamePrefix, "Smith", ""),
				),
				gedcom.NewBaptismNode("",
					gedcom.NewPlaceNode("6 MAY 1989"),
					gedcom.NewPlaceNode("7 MAY 1989"),
				),
			),
			expected: `0 @P1@ INDI
1 NAME
2 NSFX Bob
2 SPFX Smith
1 BAPM
2 PLAC 6 MAY 1989
2 PLAC 7 MAY 1989
`,
		},
	} {
		t.Run(testName, func(t *testing.T) {
			filter := gedcom.OnlyVitalsTagFilter()
			doc := gedcom.NewDocument()
			result := gedcom.GEDCOMString(gedcom.Filter(test.root, doc, filter), 0)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestRemoveEmptyDeathTagFilter(t *testing.T) {
	// ghost:ignore
	for testName, test := range map[string]struct {
		root     gedcom.Node
		expected string
	}{
		"WithoutDeath": {
			root: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot /Chance/"),
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("6 MAY 1989"),
				),
			),
			expected: `0 @P1@ INDI
1 NAME Elliot /Chance/
1 BIRT
2 DATE 6 MAY 1989
`,
		},
		"WithDeath": {
			root: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot /Chance/"),
				gedcom.NewDeathNode("",
					gedcom.NewDateNode("6 MAY 1989"),
				),
			),
			expected: `0 @P1@ INDI
1 NAME Elliot /Chance/
1 DEAT
2 DATE 6 MAY 1989
`,
		},
		"WithEmptyDeath1": {
			root: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot /Chance/"),
				gedcom.NewDeathNode(""),
			),
			expected: `0 @P1@ INDI
1 NAME Elliot /Chance/
`,
		},
		"WithEmptyDeath2": {
			root: gedcom.NewDocument().AddIndividual("P1",
				gedcom.NewNameNode("Elliot /Chance/"),
				gedcom.NewDeathNode("Y"),
			),
			expected: `0 @P1@ INDI
1 NAME Elliot /Chance/
`,
		},
	} {
		t.Run(testName, func(t *testing.T) {
			filter := gedcom.RemoveEmptyDeathTagFilter()
			doc := gedcom.NewDocument()
			result := gedcom.GEDCOMString(gedcom.Filter(test.root, doc, filter), 0)
			assert.Equal(t, test.expected, result)
		})
	}
}
