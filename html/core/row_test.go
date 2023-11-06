package core_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/html/core"
)

func TestRow_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "Row")

	c(core.NewRow(core.NewColumn(12, core.NewText("hi")))).
		Returns("<div class=\"row\"><div class=\"col-12\">hi</div></div>")

	c(core.NewRow(
		core.NewColumn(6, core.NewText("hi")),
		core.NewColumn(6, core.NewText("there")),
	)).Returns("<div class=\"row\"><div class=\"col-6\">hi</div><div class=\"col-6\">there</div></div>")
}
