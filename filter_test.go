package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	root := gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
		gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
		gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "6 MAY 1989", "", nil),
		}),
	})

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
				return node, false
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
					return gedcom.NewDateNode(nil, "1 APR 1943", "", nil), true
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
			result := gedcom.NodeGedcom(gedcom.Filter(root, test.filter))
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestWhitelistTagFilter(t *testing.T) {
	root := gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
		gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
		gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "6 MAY 1989", "", nil),
		}),
	})

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
			result := gedcom.NodeGedcom(gedcom.Filter(root, filter))
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestBlacklistTagFilter(t *testing.T) {
	root := gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
		gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
		gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "6 MAY 1989", "", nil),
		}),
	})

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
			result := gedcom.NodeGedcom(gedcom.Filter(root, filter))
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestOfficialTagFilter(t *testing.T) {
	root := gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
		gedcom.NewSimpleNode(nil, gedcom.UnofficialTagCreated, "Elliot /Chance/", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "3 Mar 2007", "", nil),
		}),
		gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
			gedcom.NewDateNode(nil, "6 MAY 1989", "", nil),
		}),
	})

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
			result := gedcom.NodeGedcom(gedcom.Filter(root, filter))
			assert.Equal(t, test.expected, result)
		})
	}
}
