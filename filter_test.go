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
			filter: func(node gedcom.Node) gedcom.FilterResult {
				return gedcom.FilterResult{}
			},
			expected: "",
		},
		{
			filter: func(node gedcom.Node) gedcom.FilterResult {
				return gedcom.FilterResult{node, gedcom.FilterTraverseModeOldChildren}
			},
			expected: `0 @P1@ INDI
1 NAME Elliot /Chance/
1 BIRT
2 DATE 6 MAY 1989
`,
		},
		{
			filter: func(node gedcom.Node) gedcom.FilterResult {
				return gedcom.FilterResult{node, gedcom.FilterTraverseModeStop}
			},
			expected: `0 @P1@ INDI
`,
		},
		{
			filter: func(node gedcom.Node) gedcom.FilterResult {
				if node.Tag().Is(gedcom.TagIndividual) {
					// false means it will not traverse children, since an
					// individual can never be inside of another individual.
					return gedcom.FilterResult{node, gedcom.FilterTraverseModeStop}
				}

				return gedcom.FilterResult{nil, gedcom.FilterTraverseModeStop}
			},
			expected: `0 @P1@ INDI
`,
		},
		{
			filter: func(node gedcom.Node) gedcom.FilterResult {
				t := node.Tag()
				newNode := gedcom.NodeCondition(
					t.Is(gedcom.TagIndividual) || t.Is(gedcom.TagName),
					node,
					nil,
				)

				return gedcom.FilterResult{newNode, gedcom.FilterTraverseModeOldChildren}
			},
			expected: `0 @P1@ INDI
1 NAME Elliot /Chance/
`,
		},
		{
			filter: func(node gedcom.Node) gedcom.FilterResult {
				t := node.Tag()
				newNode := gedcom.NodeCondition(
					t.Is(gedcom.TagIndividual) || t.Is(gedcom.TagDate),
					node,
					nil,
				)

				return gedcom.FilterResult{newNode, gedcom.FilterTraverseModeOldChildren}
			},
			expected: `0 @P1@ INDI
`,
		},
		{
			filter: func(node gedcom.Node) gedcom.FilterResult {
				t := node.Tag()
				newNode := gedcom.NodeCondition(
					t.Is(gedcom.TagIndividual) || t.Is(gedcom.TagBirth),
					node,
					nil,
				)

				return gedcom.FilterResult{newNode, gedcom.FilterTraverseModeOldChildren}
			},
			expected: `0 @P1@ INDI
1 BIRT
`,
		},
		{
			filter: func(node gedcom.Node) gedcom.FilterResult {
				t := node.Tag()
				newNode := gedcom.NodeCondition(
					t.Is(gedcom.TagIndividual) || t.Is(gedcom.TagBirth) || t.Is(gedcom.TagDate),
					node,
					nil,
				)

				return gedcom.FilterResult{newNode, gedcom.FilterTraverseModeOldChildren}
			},
			expected: `0 @P1@ INDI
1 BIRT
2 DATE 6 MAY 1989
`,
		},
		{
			filter: func(node gedcom.Node) gedcom.FilterResult {
				if node.Tag().Is(gedcom.TagName) {
					newNode := gedcom.NewDateNode(nil, "1 APR 1943", "", nil)
					return gedcom.FilterResult{newNode, gedcom.FilterTraverseModeOldChildren}
				}

				return gedcom.FilterResult{node, gedcom.FilterTraverseModeOldChildren}
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
		gedcom.NewNodeWithChildren(nil, gedcom.UnofficialTagCreated, "Elliot /Chance/", "", []gedcom.Node{
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

func TestSimpleNameFilter(t *testing.T) {
	// ghost:ignore
	for _, test := range []struct {
		root     gedcom.Node
		expected string
	}{
		{
			root: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "6 MAY 1989", "", nil),
				}),
			}),
			expected: `0 @P1@ INDI
1 NAME Elliot /Chance/
1 BIRT
2 DATE 6 MAY 1989
`,
		},
		{
			root: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "6 MAY 1989", "", nil),
				}),
				gedcom.NewNameNode(nil, "Elliot /Chance/", "", []gedcom.Node{
					gedcom.NewNodeWithChildren(nil, gedcom.TagSurname, "Smith", "", nil),
				}),
			}),
			expected: `0 @P1@ INDI
1 BIRT
2 DATE 6 MAY 1989
1 NAME Elliot /Smith/
`,
		},
		{
			root: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNameNode(nil, "", "", []gedcom.Node{
					gedcom.NewNodeWithChildren(nil, gedcom.TagGivenName, "Bob", "", nil),
					gedcom.NewNodeWithChildren(nil, gedcom.TagSurname, "Smith", "", nil),
				}),
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "6 MAY 1989", "", nil),
				}),
			}),
			expected: `0 @P1@ INDI
1 NAME Bob /Smith/
1 BIRT
2 DATE 6 MAY 1989
`,
		},
	} {
		t.Run("", func(t *testing.T) {
			filter := gedcom.SimpleNameFilter()
			result := gedcom.NodeGedcom(gedcom.Filter(test.root, filter))
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSingleNameFilter(t *testing.T) {
	// ghost:ignore
	for _, test := range []struct {
		root     gedcom.Node
		expected string
	}{
		//		{
		//			root: gedcom.NewIndividualNode(nil, "", "P1", nil),
		//			expected: `0 @P1@ INDI
		//`,
		//		},
		{
			root: gedcom.NewIndividualNode(nil, "", "P2", []gedcom.Node{
				gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
				gedcom.NewNameNode(nil, "Elliot R /Chance/", "", nil),
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "6 MAY 1989", "", nil),
				}),
			}),
			expected: `0 @P2@ INDI
1 NAME Elliot /Chance/
1 BIRT
2 DATE 6 MAY 1989
`,
		},
		//		{
		//			root: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
		//				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
		//					gedcom.NewDateNode(nil, "6 MAY 1989", "", nil),
		//				}),
		//				gedcom.NewNameNode(nil, "Elliot /Chance/", "", []gedcom.Node{
		//					gedcom.NewNodeWithChildren(nil, gedcom.TagSurname, "Smith", "", nil),
		//				}),
		//			}),
		//			expected: `0 @P1@ INDI
		//1 BIRT
		//2 DATE 6 MAY 1989
		//1 NAME Elliot /Smith/
		//`,
		//		},
		//		{
		//			root: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
		//				gedcom.NewNameNode(nil, "", "", []gedcom.Node{
		//					gedcom.NewNodeWithChildren(nil, gedcom.TagGivenName, "Bob", "", nil),
		//					gedcom.NewNodeWithChildren(nil, gedcom.TagSurname, "Smith", "", nil),
		//				}),
		//				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
		//					gedcom.NewDateNode(nil, "6 MAY 1989", "", nil),
		//				}),
		//			}),
		//			expected: `0 @P1@ INDI
		//1 NAME Bob /Smith/
		//1 BIRT
		//2 DATE 6 MAY 1989
		//`,
		//		},
	} {
		t.Run("", func(t *testing.T) {
			filter := gedcom.SingleNameFilter()
			result := gedcom.NodeGedcom(gedcom.Filter(test.root, filter))
			assert.Equal(t, test.expected, result)
		})
	}
}
