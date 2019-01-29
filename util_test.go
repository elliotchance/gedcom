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
	nodes       gedcom.Nodes
	first, last gedcom.Node
}{
	{gedcom.Nodes{}, nil, nil},
	{gedcom.Nodes{nil}, nil, nil},
	{
		gedcom.Nodes{gedcom.NewNameNode("a")},
		gedcom.NewNameNode("a"),
		gedcom.NewNameNode("a"),
	},
	{
		gedcom.Nodes{nil, gedcom.NewNameNode("a")},
		gedcom.NewNameNode("a"),
		gedcom.NewNameNode("a"),
	},
	{
		gedcom.Nodes{
			nil,
			gedcom.NewNameNode("a"),
			gedcom.NewNameNode("b"),
			nil,
		},
		gedcom.NewNameNode("a"),
		gedcom.NewNameNode("b"),
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
		{gedcom.NewNode(gedcom.TagVersion, "foo", ""), "foo"},
		{gedcom.NewNameNode("foo bar"), "foo bar"},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.want, gedcom.Value(test.node))
		})
	}
}

func TestCompound(t *testing.T) {
	n1 := gedcom.NewNameNode("Joe /Bloggs/")
	n2 := gedcom.NewNameNode("Jane /Doe/")
	n3 := gedcom.NewNameNode("John /Smith/")

	tests := []struct {
		inputs []interface{}
		want   gedcom.Nodes
	}{
		{[]interface{}{}, gedcom.Nodes{}},
		{[]interface{}{nil}, gedcom.Nodes{}},
		{[]interface{}{n1}, gedcom.Nodes{n1}},
		{[]interface{}{n1, n1}, gedcom.Nodes{n1, n1}},
		{[]interface{}{n1, n2}, gedcom.Nodes{n1, n2}},
		{[]interface{}{gedcom.Nodes{n1, n2}}, gedcom.Nodes{n1, n2}},
		{[]interface{}{gedcom.Nodes{n1, n2}, n3}, gedcom.Nodes{n1, n2, n3}},
		{[]interface{}{gedcom.Nodes{nil, n2}, n3}, gedcom.Nodes{n2, n3}},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.want, gedcom.Compound(test.inputs...))
		})
	}
}

func TestNodeCondition(t *testing.T) {
	NodeCondition := tf.Function(t, gedcom.NodeCondition)

	bob := gedcom.NewNameNode("Bob")
	sally := gedcom.NewNameNode("Sally")

	NodeCondition(true, bob, sally).Returns(bob)
	NodeCondition(false, bob, sally).Returns(sally)
}

func TestPointer(t *testing.T) {
	tests := []struct {
		node gedcom.Node
		want string
	}{
		{nil, ""},
		{gedcom.NewNode(gedcom.TagVersion, "foo", "a"), "a"},
		{gedcom.NewDocument().AddIndividual("b"), "b"},
		{gedcom.NewDocument().AddFamily("c"), "c"},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.want, gedcom.Pointer(test.node))
		})
	}
}

func TestDateAndPlace(t *testing.T) {
	date3Sep1953 := gedcom.NewDateNode("3 Sep 1953")
	date1Sep1953 := gedcom.NewDateNode("1 Sep 1953")
	place1 := gedcom.NewPlaceNode("Australia")
	place2 := gedcom.NewPlaceNode("United Kingdom")

	// ghost:ignore
	for _, test := range []struct {
		nodes gedcom.Nodes
		date  *gedcom.DateNode
		place *gedcom.PlaceNode
	}{
		{
			nodes: gedcom.Nodes{},
			date:  nil,
			place: nil,
		},
		{
			nodes: gedcom.Nodes{
				gedcom.NewBirthNode("", date3Sep1953),
			},
			date:  date3Sep1953,
			place: nil,
		},
		{
			nodes: gedcom.Nodes{
				gedcom.NewBirthNode("", date3Sep1953),
				gedcom.NewBirthNode("", date1Sep1953),
			},
			date:  date3Sep1953,
			place: nil,
		},
		{
			nodes: gedcom.Nodes{
				gedcom.NewBirthNode("", place1),
			},
			date:  nil,
			place: place1,
		},
		{
			nodes: gedcom.Nodes{
				gedcom.NewBirthNode("", date3Sep1953, place1),
			},
			date:  date3Sep1953,
			place: place1,
		},
		{
			nodes: gedcom.Nodes{
				gedcom.NewBirthNode("", date3Sep1953),
				gedcom.NewBirthNode("", place1),
			},
			date:  date3Sep1953,
			place: place1,
		},
		{
			nodes: gedcom.Nodes{
				gedcom.NewBirthNode("", date3Sep1953),
				gedcom.NewBirthNode("", place1),
				gedcom.NewBirthNode("", place2, date1Sep1953),
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
