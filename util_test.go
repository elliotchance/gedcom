package gedcom

import (
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
			assert.Equal(t, test.want, Atoi(test.s))
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
			assert.Equal(t, test.want, CleanSpace(test.s))
		})
	}
}
