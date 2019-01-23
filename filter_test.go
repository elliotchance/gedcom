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
			node := gedcom.Filter(root, test.filter)
			result := gedcom.GEDCOMString(node, 0)
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
			result := gedcom.GEDCOMString(gedcom.Filter(root, filter), 0)
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
			result := gedcom.GEDCOMString(gedcom.Filter(root, filter), 0)
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
			result := gedcom.GEDCOMString(gedcom.Filter(root, filter), 0)
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
			root: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "6 MAY 1989", "", nil),
				}),
			}),
			format: gedcom.NameFormatGEDCOM,
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
			format: gedcom.NameFormatGEDCOM,
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
			format: gedcom.NameFormatGEDCOM,
			expected: `0 @P1@ INDI
1 NAME Bob /Smith/
1 BIRT
2 DATE 6 MAY 1989
`,
		},
		{
			root: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "6 MAY 1989", "", nil),
				}),
			}),
			format: gedcom.NameFormatWritten,
			expected: `0 @P1@ INDI
1 NAME Elliot Chance
1 BIRT
2 DATE 6 MAY 1989
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
			result := gedcom.GEDCOMString(gedcom.Filter(test.root, filter), 0)
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
		"ComplexName1AndDeathDate": {
			root: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewDeathNode(nil, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "6 MAY 1989", "", nil),
				}),
				gedcom.NewNameNode(nil, "Elliot /Chance/", "", []gedcom.Node{
					gedcom.NewNodeWithChildren(nil, gedcom.TagSurname, "Smith", "", nil),
				}),
			}),
			expected: `0 @P1@ INDI
1 DEAT
2 DATE 6 MAY 1989
1 NAME Elliot /Chance/
2 SURN Smith
`,
		},
		"ComplexName2AndBirthPlace": {
			root: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNameNode(nil, "", "", []gedcom.Node{
					gedcom.NewNodeWithChildren(nil, gedcom.TagGivenName, "Bob", "", nil),
					gedcom.NewNodeWithChildren(nil, gedcom.TagSurname, "Smith", "", nil),
				}),
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
					gedcom.NewPlaceNode(nil, "Sydney, Australia", "", nil),
				}),
			}),
			expected: `0 @P1@ INDI
1 NAME
2 GIVN Bob
2 SURN Smith
1 BIRT
2 PLAC Sydney, Australia
`,
		},
		"Source": {
			root: gedcom.NewSourceNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNameNode(nil, "", "", []gedcom.Node{
					gedcom.NewNodeWithChildren(nil, gedcom.TagGivenName, "Bob", "", nil),
					gedcom.NewNodeWithChildren(nil, gedcom.TagSurname, "Smith", "", nil),
				}),
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
					gedcom.NewPlaceNode(nil, "Sydney, Australia", "", nil),
				}),
			}),
			expected: ``,
		},
		"IndividualNote": {
			root: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNameNode(nil, "", "", []gedcom.Node{
					gedcom.NewNodeWithChildren(nil, gedcom.TagGivenName, "Bob", "", nil),
					gedcom.NewNodeWithChildren(nil, gedcom.TagSurname, "Smith", "", nil),
				}),
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
					gedcom.NewPlaceNode(nil, "Sydney, Australia", "", nil),
				}),
				gedcom.NewNode(nil, gedcom.TagNote, "foo", ""),
			}),
			expected: `0 @P1@ INDI
1 NAME
2 GIVN Bob
2 SURN Smith
1 BIRT
2 PLAC Sydney, Australia
`,
		},
		"Burial": {
			root: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNameNode(nil, "", "", []gedcom.Node{
					gedcom.NewNodeWithChildren(nil, gedcom.TagGivenName, "Bob", "", nil),
					gedcom.NewNodeWithChildren(nil, gedcom.TagTitle, "Smith", "", nil),
				}),
				gedcom.NewBurialNode(nil, "", "", []gedcom.Node{
					gedcom.NewPlaceNode(nil, "6 MAY 1989", "", nil),
					gedcom.NewPlaceNode(nil, "Sydney, Australia", "", nil),
				}),
			}),
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
			root: gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNameNode(nil, "", "", []gedcom.Node{
					gedcom.NewNodeWithChildren(nil, gedcom.TagNameSuffix, "Bob", "", nil),
					gedcom.NewNodeWithChildren(nil, gedcom.TagSurnamePrefix, "Smith", "", nil),
				}),
				gedcom.NewBaptismNode(nil, "", "", []gedcom.Node{
					gedcom.NewPlaceNode(nil, "6 MAY 1989", "", nil),
					gedcom.NewPlaceNode(nil, "7 MAY 1989", "", nil),
				}),
			}),
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
			result := gedcom.GEDCOMString(gedcom.Filter(test.root, filter), 0)
			assert.Equal(t, test.expected, result)
		})
	}
}
