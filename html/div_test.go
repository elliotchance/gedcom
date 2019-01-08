package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestDiv_WriteTo(t *testing.T) {
	c := testComponent(t, "Div")

	c(html.NewDiv("", html.NewText("foo"))).Returns(`<div>foo</div>`)

	c(html.NewDiv(`hi there`, html.NewText("foo"))).
		Returns("<div class=\"hi there\">foo</div>")
}
