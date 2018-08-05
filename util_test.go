package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
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
		[]gedcom.Node{gedcom.NewNameNode("a", "", nil)},
		gedcom.NewNameNode("a", "", nil),
		gedcom.NewNameNode("a", "", nil),
	},
	{
		[]gedcom.Node{nil, gedcom.NewNameNode("a", "", nil)},
		nil,
		gedcom.NewNameNode("a", "", nil),
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
		{gedcom.NewSimpleNode(gedcom.TagVersion, "foo", "", nil), "foo"},
		{gedcom.NewNameNode("foo bar", "", nil), "foo bar"},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.want, gedcom.Value(test.node))
		})
	}
}

func TestCompound(t *testing.T) {
	n1 := gedcom.NewNameNode("Joe /Bloggs/", "", []gedcom.Node{})
	n2 := gedcom.NewNameNode("Jane /Doe/", "", []gedcom.Node{})
	n3 := gedcom.NewNameNode("John /Smith/", "", []gedcom.Node{})

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
