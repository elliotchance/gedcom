package core_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/html/core"
)

func TestLink_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "Link")

	c(core.NewLink(core.NewText("hi"), "dest")).
		Returns("<a href=\"dest\">hi</a>")

	c(core.NewLink(core.NewText("a&b"), "#foo")).
		Returns("<a href=\"#foo\">a&amp;b</a>")
}
