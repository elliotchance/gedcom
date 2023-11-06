package core_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/html/core"
)

func TestDiv_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "Div")

	c(core.NewDiv("", core.NewText("foo"))).Returns(`<div>foo</div>`)

	c(core.NewDiv(`hi there`, core.NewText("foo"))).
		Returns("<div class=\"hi there\">foo</div>")
}
