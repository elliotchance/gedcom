package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestLink_WriteTo(t *testing.T) {
	c := testComponent(t, "Link")

	c(html.NewLink(html.NewText("hi"), "dest")).
		Returns("<a href=\"dest\">hi</a>")

	c(html.NewLink(html.NewText("a&b"), "#foo")).
		Returns("<a href=\"#foo\">a&amp;b</a>")
}
