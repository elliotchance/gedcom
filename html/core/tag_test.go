package core_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/html/core"
)

func TestTag_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "Tag")

	c(core.NewTag("a", nil, core.NewText("hi"))).Returns(`<a>hi</a>`)

	c(core.NewTag("a", map[string]string{
		"foo": "bar",
	}, core.NewText("there"))).Returns(`<a foo="bar">there</a>`)

	c(core.NewTag("a", map[string]string{
		"foo": "bar",
		"bar": "baz",
	}, core.NewText("there"))).Returns("<a bar=\"baz\" foo=\"bar\">there</a>")

	c(core.NewTag("a", map[string]string{
		"bar": "baz",
	}, core.NewText("there"))).Returns(`<a bar="baz">there</a>`)
}
