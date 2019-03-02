package core_test

import (
	"github.com/elliotchance/gedcom/html/core"
	"testing"
)

func TestColumn_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "Column")

	c(core.NewColumn(12, core.NewText("foo"))).
		Returns("<div class=\"col-12\">foo</div>")

	c(core.NewColumn(6, core.NewText("bar"))).
		Returns("<div class=\"col-6\">bar</div>")
}
