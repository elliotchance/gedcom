package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39"
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
		node:    gedcom.NewPlaceNode("Waterloo"),
		name:    "Waterloo",
		county:  "",
		state:   "",
		country: "",
	},
	{
		node:    gedcom.NewPlaceNode("Waterloo, NSW"),
		name:    "Waterloo, NSW",
		county:  "",
		state:   "",
		country: "",
	},
	{
		node:    gedcom.NewPlaceNode("Cove,Cache,Utah,USA."),
		name:    "Cove",
		county:  "Cache",
		state:   "Utah",
		country: "USA.",
	},
	{
		node:    gedcom.NewPlaceNode("Cove, Cache, Utah, USA."),
		name:    "Cove",
		county:  "Cache",
		state:   "Utah",
		country: "USA.",
	},
	{
		node:    gedcom.NewPlaceNode("United States"),
		name:    "United States",
		county:  "",
		state:   "",
		country: "United States",
	},
	{
		node:    gedcom.NewPlaceNode("Foo, australia."),
		name:    "Foo, australia.",
		county:  "",
		state:   "",
		country: "Australia",
	},
	{
		node:    gedcom.NewPlaceNode("Bar, Nashville, USA"),
		name:    "Bar, Nashville, USA",
		county:  "",
		state:   "",
		country: "USA",
	},
	{
		node:    gedcom.NewPlaceNode("Hobbitown, New zealand "),
		name:    "Hobbitown, New zealand",
		county:  "",
		state:   "",
		country: "New Zealand",
	},
}

func TestNewPlaceNode(t *testing.T) {
	child := gedcom.NewNameNode("")
	node := gedcom.NewPlaceNode("foo", child)

	assert.NotNil(t, node)
	assert.IsType(t, node, (*gedcom.PlaceNode)(nil))
	assert.Equal(t, gedcom.TagPlace, node.Tag())
	assert.Equal(t, gedcom.Nodes{child}, node.Nodes())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "", node.Pointer())
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
			node:     gedcom.NewPlaceNode(""),
			expected: nil,
		},
		{
			node:     gedcom.NewPlaceNode(""),
			expected: nil,
		},
		{
			node: gedcom.NewPlaceNode("",
				gedcom.NewFormatNode(""),
			),
			expected: gedcom.NewFormatNode(""),
		},
		{
			node: gedcom.NewPlaceNode("",
				gedcom.NewNameNode(""),
			),
			expected: nil,
		},
		{
			node: gedcom.NewPlaceNode("",
				gedcom.NewNameNode(""),
				gedcom.NewFormatNode("1"),
				gedcom.NewFormatNode("2"),
			),
			expected: gedcom.NewFormatNode("1"),
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
			node:     gedcom.NewPlaceNode(""),
			expected: []*gedcom.PhoneticVariationNode{},
		},
		{
			node:     gedcom.NewPlaceNode(""),
			expected: []*gedcom.PhoneticVariationNode{},
		},
		{
			node: gedcom.NewNode(gedcom.TagPlace, "", "P1",
				gedcom.NewPhoneticVariationNode(""),
			).(*gedcom.PlaceNode),
			expected: []*gedcom.PhoneticVariationNode{
				gedcom.NewPhoneticVariationNode(""),
			},
		},
		{
			node: gedcom.NewNode(gedcom.TagPlace, "", "P1",
				gedcom.NewPhoneticVariationNode(""),
				gedcom.NewNode(gedcom.TagDeath, "", ""),
			).(*gedcom.PlaceNode),
			expected: []*gedcom.PhoneticVariationNode{
				gedcom.NewPhoneticVariationNode(""),
			},
		},
		{
			node: gedcom.NewNode(gedcom.TagPlace, "", "P1",
				gedcom.NewPhoneticVariationNode("foo"),
				gedcom.NewNode(gedcom.TagDeath, "", ""),
				gedcom.NewPhoneticVariationNode("bar"),
			).(*gedcom.PlaceNode),
			expected: []*gedcom.PhoneticVariationNode{
				gedcom.NewPhoneticVariationNode("foo"),
				gedcom.NewPhoneticVariationNode("bar"),
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
			node:     gedcom.NewPlaceNode(""),
			expected: []*gedcom.RomanizedVariationNode{},
		},
		{
			node:     gedcom.NewPlaceNode(""),
			expected: []*gedcom.RomanizedVariationNode{},
		},
		{
			node: gedcom.NewNode(gedcom.TagPlace, "", "P1",
				gedcom.NewRomanizedVariationNode(""),
			).(*gedcom.PlaceNode),
			expected: []*gedcom.RomanizedVariationNode{
				gedcom.NewRomanizedVariationNode(""),
			},
		},
		{
			node: gedcom.NewNode(gedcom.TagPlace, "", "P1",
				gedcom.NewRomanizedVariationNode(""),
				gedcom.NewNode(gedcom.TagDeath, "", ""),
			).(*gedcom.PlaceNode),
			expected: []*gedcom.RomanizedVariationNode{
				gedcom.NewRomanizedVariationNode(""),
			},
		},
		{
			node: gedcom.NewNode(gedcom.TagPlace, "", "P1",
				gedcom.NewRomanizedVariationNode("foo"),
				gedcom.NewNode(gedcom.TagDeath, "", ""),
				gedcom.NewRomanizedVariationNode("bar"),
			).(*gedcom.PlaceNode),
			expected: []*gedcom.RomanizedVariationNode{
				gedcom.NewRomanizedVariationNode("foo"),
				gedcom.NewRomanizedVariationNode("bar"),
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
			node:     gedcom.NewPlaceNode(""),
			expected: nil,
		},
		{
			node:     gedcom.NewPlaceNode(""),
			expected: nil,
		},
		{
			node: gedcom.NewPlaceNode("",
				gedcom.NewMapNode(""),
			),
			expected: gedcom.NewMapNode(""),
		},
		{
			node: gedcom.NewPlaceNode("",
				gedcom.NewNameNode(""),
			),
			expected: nil,
		},
		{
			node: gedcom.NewPlaceNode("",
				gedcom.NewNameNode(""),
				gedcom.NewMapNode("1"),
				gedcom.NewMapNode("2"),
			),
			expected: gedcom.NewMapNode("1"),
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
			node:     gedcom.NewPlaceNode(""),
			expected: []*gedcom.NoteNode{},
		},
		{
			node:     gedcom.NewPlaceNode(""),
			expected: []*gedcom.NoteNode{},
		},
		{
			node: gedcom.NewNode(gedcom.TagPlace, "", "P1",
				gedcom.NewNoteNode(""),
			).(*gedcom.PlaceNode),
			expected: []*gedcom.NoteNode{
				gedcom.NewNoteNode(""),
			},
		},
		{
			node: gedcom.NewNode(gedcom.TagPlace, "", "P1",
				gedcom.NewNoteNode(""),
				gedcom.NewNode(gedcom.TagDeath, "", ""),
			).(*gedcom.PlaceNode),
			expected: []*gedcom.NoteNode{
				gedcom.NewNoteNode(""),
			},
		},
		{
			node: gedcom.NewNode(gedcom.TagPlace, "", "P1",
				gedcom.NewNoteNode("foo"),
				gedcom.NewNode(gedcom.TagDeath, "", ""),
				gedcom.NewNoteNode("bar"),
			).(*gedcom.PlaceNode),
			expected: []*gedcom.NoteNode{
				gedcom.NewNoteNode("foo"),
				gedcom.NewNoteNode("bar"),
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Notes(), test.expected)
		})
	}
}
