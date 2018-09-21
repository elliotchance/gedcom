package util_test

import (
	"github.com/elliotchance/gedcom/util"
	"github.com/elliotchance/tf"
	"testing"
)

func TestStringSliceContains(t *testing.T) {
	StringSliceContains := tf.Function(t, util.StringSliceContains)

	StringSliceContains(nil, "foo").Returns(false)
	StringSliceContains([]string{}, "foo").Returns(false)
	StringSliceContains([]string{"foo"}, "foo").Returns(true)
	StringSliceContains([]string{"foo", "bar"}, "bar").Returns(true)
	StringSliceContains([]string{"foo", "bar"}, "baz").Returns(false)
}
