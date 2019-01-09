package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestTag_WriteTo(t *testing.T) {
	c := testComponent(t, "Tag")

	c(html.NewTag("a", nil, html.NewText("hi"))).Returns(`<a>hi</a>`)

	c(html.NewTag("a", map[string]string{
		"foo": "bar",
	}, html.NewText("there"))).Returns(`<a foo="bar">there</a>`)

	c(html.NewTag("a", map[string]string{
		"foo": "bar",
		"bar": "baz",
	}, html.NewText("there"))).Returns("<a bar=\"baz\" foo=\"bar\">there</a>")

	c(html.NewTag("a", map[string]string{
		"bar": "baz",
	}, html.NewText("there"))).Returns(`<a bar="baz">there</a>`)
}
