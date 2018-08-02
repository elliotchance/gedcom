package gedcom_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/elliotchance/gedcom"
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

func TestCoalesce(t *testing.T) {
	tests := []struct {
		items []interface{}
		want  interface{}
	}{
		{[]interface{}{}, nil},
		{[]interface{}{nil}, nil},
		{[]interface{}{"A"}, "A"},
		{[]interface{}{"A", "B"}, "A"},
		{[]interface{}{nil, "B"}, "B"},
		{[]interface{}{nil, nil, "C"}, "C"},
		{[]interface{}{nil, "C", nil, "D"}, "C"},
		{[]interface{}{nil, nil, nil, nil}, nil},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.want, gedcom.Coalesce(test.items...))
		})
	}
}
