package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestAtoi(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"123", 123},
		{"0", 0},
		{"", 0},
		{"FOO", 0},
		{"12F", 0},
		{"F20", 0},
		{"0023", 23},
		{"002310", 2310},
		{"  2317 ", 2317},
		{" 0231 ", 231},
	}

	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			assert.Equal(t, test.want, gedcom.Atoi(test.s))
		})
	}
}

func TestCleanSpace(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"Foo", "Foo"},
		{"foo bar", "foo bar"},
		{"foo  bar baz", "foo bar baz"},
		{"foo   bar baz", "foo bar baz"},
		{"   bar bar", "bar bar"},
		{"bar baz  ", "bar baz"},
		{"  foo   bar  baz  ", "foo bar baz"},
	}

	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			assert.Equal(t, test.want, gedcom.CleanSpace(test.s))
		})
	}
}

var firtLastTests = []struct {
	nodes       []gedcom.Node
	first, last gedcom.Node
}{
	{[]gedcom.Node{}, nil, nil},
	{[]gedcom.Node{nil}, nil, nil},
	{
		[]gedcom.Node{gedcom.NewNameNode(nil, "a", "", nil)},
		gedcom.NewNameNode(nil, "a", "", nil),
		gedcom.NewNameNode(nil, "a", "", nil),
	},
	{
		[]gedcom.Node{nil, gedcom.NewNameNode(nil, "a", "", nil)},
		gedcom.NewNameNode(nil, "a", "", nil),
		gedcom.NewNameNode(nil, "a", "", nil),
	},
	{
		[]gedcom.Node{
			nil,
			gedcom.NewNameNode(nil, "a", "", nil),
			gedcom.NewNameNode(nil, "b", "", nil),
			nil,
		},
		gedcom.NewNameNode(nil, "a", "", nil),
		gedcom.NewNameNode(nil, "b", "", nil),
	},
}

func TestFirst(t *testing.T) {
	for _, test := range firtLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.first, gedcom.First(test.nodes))
		})
	}
}

func TestLast(t *testing.T) {
	for _, test := range firtLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.last, gedcom.Last(test.nodes))
		})
	}
}

func TestValue(t *testing.T) {
	tests := []struct {
		node gedcom.Node
		want string
	}{
		{nil, ""},
		{gedcom.NewNodeWithChildren(nil, gedcom.TagVersion, "foo", "", nil), "foo"},
		{gedcom.NewNameNode(nil, "foo bar", "", nil), "foo bar"},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.want, gedcom.Value(test.node))
		})
	}
}

func TestCompound(t *testing.T) {
	n1 := gedcom.NewNameNode(nil, "Joe /Bloggs/", "", []gedcom.Node{})
	n2 := gedcom.NewNameNode(nil, "Jane /Doe/", "", []gedcom.Node{})
	n3 := gedcom.NewNameNode(nil, "John /Smith/", "", []gedcom.Node{})

	tests := []struct {
		inputs []interface{}
		want   []gedcom.Node
	}{
		{[]interface{}{}, []gedcom.Node{}},
		{[]interface{}{nil}, []gedcom.Node{}},
		{[]interface{}{n1}, []gedcom.Node{n1}},
		{[]interface{}{n1, n1}, []gedcom.Node{n1, n1}},
		{[]interface{}{n1, nil, n2}, []gedcom.Node{n1, n2}},
		{[]interface{}{[]gedcom.Node{n1, n2}}, []gedcom.Node{n1, n2}},
		{[]interface{}{[]gedcom.Node{n1, n2}, n3}, []gedcom.Node{n1, n2, n3}},
		{[]interface{}{[]gedcom.Node{nil, n2}, n3}, []gedcom.Node{n2, n3}},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.want, gedcom.Compound(test.inputs...))
		})
	}
}

func TestNodeCondition(t *testing.T) {
	NodeCondition := tf.Function(t, gedcom.NodeCondition)

	bob := gedcom.NewNameNode(nil, "Bob", "", nil)
	sally := gedcom.NewNameNode(nil, "Sally", "", nil)

	NodeCondition(true, bob, sally).Returns(bob)
	NodeCondition(false, bob, sally).Returns(sally)
}

func TestPointer(t *testing.T) {
	tests := []struct {
		node gedcom.Node
		want string
	}{
		{nil, ""},
		{gedcom.NewNodeWithChildren(nil, gedcom.TagVersion, "foo", "a", nil), "a"},
		{gedcom.NewNameNode(nil, "foo bar", "b", nil), "b"},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.want, gedcom.Pointer(test.node))
		})
	}
}

func TestDateAndPlace(t *testing.T) {
	date3Sep1953 := gedcom.NewDateNode(nil, "3 Sep 1953", "", nil)
	date1Sep1953 := gedcom.NewDateNode(nil, "1 Sep 1953", "", nil)
	place1 := gedcom.NewPlaceNode(nil, "Australia", "", nil)
	place2 := gedcom.NewPlaceNode(nil, "United Kingdom", "", nil)

	// ghost:ignore
	for _, test := range []struct {
		nodes []gedcom.Node
		date  *gedcom.DateNode
		place *gedcom.PlaceNode
	}{
		{
			nodes: []gedcom.Node{},
			date:  nil,
			place: nil,
		},
		{
			nodes: []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{date3Sep1953}),
			},
			date:  date3Sep1953,
			place: nil,
		},
		{
			nodes: []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{date3Sep1953}),
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{date1Sep1953}),
			},
			date:  date3Sep1953,
			place: nil,
		},
		{
			nodes: []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{place1}),
			},
			date:  nil,
			place: place1,
		},
		{
			nodes: []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{date3Sep1953, place1}),
			},
			date:  date3Sep1953,
			place: place1,
		},
		{
			nodes: []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{date3Sep1953}),
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{place1}),
			},
			date:  date3Sep1953,
			place: place1,
		},
		{
			nodes: []gedcom.Node{
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{date3Sep1953}),
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{place1}),
				gedcom.NewBirthNode(nil, "", "", []gedcom.Node{place2, date1Sep1953}),
			},
			date:  date3Sep1953,
			place: place1,
		},
	} {
		t.Run("", func(t *testing.T) {
			date, place := gedcom.DateAndPlace(test.nodes...)

			assert.Equal(t, test.date, date)
			assert.Equal(t, test.place, place)
		})
	}
}
