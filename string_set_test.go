package gedcom_test

import (
	"testing"
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"sort"
)

func TestNewStringSet(t *testing.T) {
	ss := gedcom.NewStringSet("Foo", "foo", "Foo")
	assert.Equal(t, 2, ss.Len())
}

func TestStringSet_Add(t *testing.T) {
	ss := gedcom.NewStringSet()
	assert.Equal(t, 0, ss.Len())

	ss.Add("foo")
	assert.Equal(t, 1, ss.Len())

	ss.Add()
	assert.Equal(t, 1, ss.Len())

	ss.Add("bar", "baz", "bar")
	assert.Equal(t, 3, ss.Len())

	ss.Add("foo")
	assert.Equal(t, 3, ss.Len())
}

func TestStringSet_Has(t *testing.T) {
	ss := gedcom.NewStringSet("foo", "bar")

	assert.True(t, ss.Has("foo"))
	assert.True(t, ss.Has("bar"))
	assert.False(t, ss.Has("baz"))
}

func TestStringSet_Intersects(t *testing.T) {
	ss1 := gedcom.NewStringSet("foo", "bar")
	ss2 := gedcom.NewStringSet("bar", "baz")
	ss3 := gedcom.NewStringSet("qux")

	assert.True(t, ss1.Intersects(ss2))
	assert.False(t, ss1.Intersects(ss3))
	assert.True(t, ss2.Intersects(ss1))
	assert.False(t, ss2.Intersects(ss3))
	assert.False(t, ss3.Intersects(ss1))
	assert.False(t, ss3.Intersects(ss2))
}

func TestStringSet_Iterate(t *testing.T) {
	ss := gedcom.NewStringSet("foo", "bar")

	t.Run("AllItems", func(t *testing.T) {
		var result []string
		ss.Iterate(func(s string) bool {
			result = append(result, s)

			return true
		})

		sort.Strings(result)

		assert.Equal(t, []string{"bar", "foo"}, result)
	})

	t.Run("Stop", func(t *testing.T) {
		var result []string
		ss.Iterate(func(s string) bool {
			result = append(result, s)

			return false
		})

		assert.Len(t, result, 1)
	})
}

func TestStringSet_Len(t *testing.T) {
	ss1 := gedcom.NewStringSet("foo", "bar")
	ss2 := gedcom.NewStringSet()
	ss3 := gedcom.NewStringSet("qux")

	assert.Equal(t, 2, ss1.Len())
	assert.Equal(t, 0, ss2.Len())
	assert.Equal(t, 1, ss3.Len())
}

func TestStringSet_Strings(t *testing.T) {
	ss := gedcom.NewStringSet("foo", "bar")

	assert.Equal(t, []string{"bar", "foo"}, ss.Strings())
}

func TestStringSet_String(t *testing.T) {
	ss := gedcom.NewStringSet("foo", "bar")

	assert.Equal(t, "(bar,foo)", ss.String())
}
