package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

var placeTests = []struct {
	node    *gedcom.PlaceNode
	name    string // Name()
	county  string // County()
	state   string // State()
	country string // Country()
}{
	{
		node:    nil,
		name:    "",
		county:  "",
		state:   "",
		country: "",
	},
	{
		node:    gedcom.NewPlaceNode(nil, "Waterloo", "", nil),
		name:    "Waterloo",
		county:  "",
		state:   "",
		country: "",
	},
	{
		node:    gedcom.NewPlaceNode(nil, "Waterloo, NSW", "", nil),
		name:    "Waterloo, NSW",
		county:  "",
		state:   "",
		country: "",
	},
	{
		node:    gedcom.NewPlaceNode(nil, "Cove,Cache,Utah,USA.", "", nil),
		name:    "Cove",
		county:  "Cache",
		state:   "Utah",
		country: "USA.",
	},
	{
		node:    gedcom.NewPlaceNode(nil, "Cove, Cache, Utah, USA.", "", nil),
		name:    "Cove",
		county:  "Cache",
		state:   "Utah",
		country: "USA.",
	},
	{
		node:    gedcom.NewPlaceNode(nil, "United States", "", nil),
		name:    "United States",
		county:  "",
		state:   "",
		country: "United States",
	},
	{
		node:    gedcom.NewPlaceNode(nil, "Foo, australia.", "", nil),
		name:    "Foo, australia.",
		county:  "",
		state:   "",
		country: "Australia",
	},
	{
		node:    gedcom.NewPlaceNode(nil, "Bar, Nashville, USA", "", nil),
		name:    "Bar, Nashville, USA",
		county:  "",
		state:   "",
		country: "USA",
	},
	{
		node:    gedcom.NewPlaceNode(nil, "Hobbitown, New zealand ", "", nil),
		name:    "Hobbitown, New zealand",
		county:  "",
		state:   "",
		country: "New Zealand",
	},
}

func TestNewPlaceNode(t *testing.T) {
	doc := gedcom.NewDocument()
	child := gedcom.NewNameNode(doc, "", "", nil)
	node := gedcom.NewPlaceNode(doc, "foo", "bar", []gedcom.Node{child})

	assert.NotNil(t, node)
	assert.IsType(t, node, (*gedcom.PlaceNode)(nil))
	assert.Equal(t, gedcom.TagPlace, node.Tag())
	assert.Equal(t, []gedcom.Node{child}, node.Nodes())
	assert.Equal(t, doc, node.Document())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "bar", node.Pointer())
}

func TestPlaceNode_Name(t *testing.T) {
	Name := tf.Function(t, (*gedcom.PlaceNode).Name)

	for _, test := range placeTests {
		Name(test.node).Returns(test.name)
	}
}

func TestPlaceNode_County(t *testing.T) {
	County := tf.Function(t, (*gedcom.PlaceNode).County)

	for _, test := range placeTests {
		County(test.node).Returns(test.county)
	}
}

func TestPlaceNode_State(t *testing.T) {
	State := tf.Function(t, (*gedcom.PlaceNode).State)

	for _, test := range placeTests {
		State(test.node).Returns(test.state)
	}
}

func TestPlaceNode_Country(t *testing.T) {
	Country := tf.Function(t, (*gedcom.PlaceNode).Country)

	for _, test := range placeTests {
		Country(test.node).Returns(test.country)
	}
}

func TestPlaceNode_Format(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		node     *gedcom.PlaceNode
		expected *gedcom.FormatNode
	}{
		{
			node:     nil,
			expected: nil,
		},
		{
			node:     gedcom.NewPlaceNode(nil, "", "", nil),
			expected: nil,
		},
		{
			node:     gedcom.NewPlaceNode(nil, "", "", []gedcom.Node{}),
			expected: nil,
		},
		{
			node: gedcom.NewPlaceNode(nil, "", "", []gedcom.Node{
				gedcom.NewFormatNode(nil, "", "", []gedcom.Node{}),
			}),
			expected: gedcom.NewFormatNode(nil, "", "", []gedcom.Node{}),
		},
		{
			node: gedcom.NewPlaceNode(nil, "", "", []gedcom.Node{
				gedcom.NewNameNode(nil, "", "", []gedcom.Node{}),
			}),
			expected: nil,
		},
		{
			node: gedcom.NewPlaceNode(nil, "", "", []gedcom.Node{
				gedcom.NewNameNode(nil, "", "", []gedcom.Node{}),
				gedcom.NewFormatNode(nil, "1", "", []gedcom.Node{}),
				gedcom.NewFormatNode(nil, "2", "", []gedcom.Node{}),
			}),
			expected: gedcom.NewFormatNode(nil, "1", "", []gedcom.Node{}),
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Format(), test.expected)
		})
	}
}

func TestPlaceNode_PhoneticVariations(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		node     *gedcom.PlaceNode
		expected []*gedcom.PhoneticVariationNode
	}{
		{
			node:     nil,
			expected: []*gedcom.PhoneticVariationNode{},
		},
		{
			node:     gedcom.NewPlaceNode(nil, "", "", nil),
			expected: []*gedcom.PhoneticVariationNode{},
		},
		{
			node:     gedcom.NewPlaceNode(nil, "", "", []gedcom.Node{}),
			expected: []*gedcom.PhoneticVariationNode{},
		},
		{
			node: gedcom.NewPlaceNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewPhoneticVariationNode(nil, "", "", []gedcom.Node{}),
			}),
			expected: []*gedcom.PhoneticVariationNode{
				gedcom.NewPhoneticVariationNode(nil, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewPlaceNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewPhoneticVariationNode(nil, "", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
			}),
			expected: []*gedcom.PhoneticVariationNode{
				gedcom.NewPhoneticVariationNode(nil, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewPlaceNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewPhoneticVariationNode(nil, "foo", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
				gedcom.NewPhoneticVariationNode(nil, "bar", "", []gedcom.Node{}),
			}),
			expected: []*gedcom.PhoneticVariationNode{
				gedcom.NewPhoneticVariationNode(nil, "foo", "", []gedcom.Node{}),
				gedcom.NewPhoneticVariationNode(nil, "bar", "", []gedcom.Node{}),
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.PhoneticVariations(), test.expected)
		})
	}
}

func TestPlaceNode_RomanizedVariations(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		node     *gedcom.PlaceNode
		expected []*gedcom.RomanizedVariationNode
	}{
		{
			node:     nil,
			expected: []*gedcom.RomanizedVariationNode{},
		},
		{
			node:     gedcom.NewPlaceNode(nil, "", "", nil),
			expected: []*gedcom.RomanizedVariationNode{},
		},
		{
			node:     gedcom.NewPlaceNode(nil, "", "", []gedcom.Node{}),
			expected: []*gedcom.RomanizedVariationNode{},
		},
		{
			node: gedcom.NewPlaceNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewRomanizedVariationNode(nil, "", "", []gedcom.Node{}),
			}),
			expected: []*gedcom.RomanizedVariationNode{
				gedcom.NewRomanizedVariationNode(nil, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewPlaceNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewRomanizedVariationNode(nil, "", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
			}),
			expected: []*gedcom.RomanizedVariationNode{
				gedcom.NewRomanizedVariationNode(nil, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewPlaceNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewRomanizedVariationNode(nil, "foo", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
				gedcom.NewRomanizedVariationNode(nil, "bar", "", []gedcom.Node{}),
			}),
			expected: []*gedcom.RomanizedVariationNode{
				gedcom.NewRomanizedVariationNode(nil, "foo", "", []gedcom.Node{}),
				gedcom.NewRomanizedVariationNode(nil, "bar", "", []gedcom.Node{}),
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.RomanizedVariations(), test.expected)
		})
	}
}

func TestPlaceNode_Map(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		node     *gedcom.PlaceNode
		expected *gedcom.MapNode
	}{
		{
			node:     nil,
			expected: nil,
		},
		{
			node:     gedcom.NewPlaceNode(nil, "", "", nil),
			expected: nil,
		},
		{
			node:     gedcom.NewPlaceNode(nil, "", "", []gedcom.Node{}),
			expected: nil,
		},
		{
			node: gedcom.NewPlaceNode(nil, "", "", []gedcom.Node{
				gedcom.NewMapNode(nil, "", "", []gedcom.Node{}),
			}),
			expected: gedcom.NewMapNode(nil, "", "", []gedcom.Node{}),
		},
		{
			node: gedcom.NewPlaceNode(nil, "", "", []gedcom.Node{
				gedcom.NewNameNode(nil, "", "", []gedcom.Node{}),
			}),
			expected: nil,
		},
		{
			node: gedcom.NewPlaceNode(nil, "", "", []gedcom.Node{
				gedcom.NewNameNode(nil, "", "", []gedcom.Node{}),
				gedcom.NewMapNode(nil, "1", "", []gedcom.Node{}),
				gedcom.NewMapNode(nil, "2", "", []gedcom.Node{}),
			}),
			expected: gedcom.NewMapNode(nil, "1", "", []gedcom.Node{}),
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Map(), test.expected)
		})
	}
}

func TestPlaceNode_Notes(t *testing.T) {
	// ghost:ignore
	var tests = []struct {
		node     *gedcom.PlaceNode
		expected []*gedcom.NoteNode
	}{
		{
			node:     nil,
			expected: []*gedcom.NoteNode{},
		},
		{
			node:     gedcom.NewPlaceNode(nil, "", "", nil),
			expected: []*gedcom.NoteNode{},
		},
		{
			node:     gedcom.NewPlaceNode(nil, "", "", []gedcom.Node{}),
			expected: []*gedcom.NoteNode{},
		},
		{
			node: gedcom.NewPlaceNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNoteNode(nil, "", "", []gedcom.Node{}),
			}),
			expected: []*gedcom.NoteNode{
				gedcom.NewNoteNode(nil, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewPlaceNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNoteNode(nil, "", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
			}),
			expected: []*gedcom.NoteNode{
				gedcom.NewNoteNode(nil, "", "", []gedcom.Node{}),
			},
		},
		{
			node: gedcom.NewPlaceNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewNoteNode(nil, "foo", "", []gedcom.Node{}),
				gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{}),
				gedcom.NewNoteNode(nil, "bar", "", []gedcom.Node{}),
			}),
			expected: []*gedcom.NoteNode{
				gedcom.NewNoteNode(nil, "foo", "", []gedcom.Node{}),
				gedcom.NewNoteNode(nil, "bar", "", []gedcom.Node{}),
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Notes(), test.expected)
		})
	}
}
