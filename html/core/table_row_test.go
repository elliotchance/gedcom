package core_test

import (
	"github.com/elliotchance/gedcom/html/core"
	"testing"
)

func TestTableRow_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "TableRow")

	c(core.NewTableRow()).Returns("<tr></tr>")

	c(core.NewTableRow(core.NewTableCell(core.NewText("hi")))).
		Returns("<tr><td scope=\"col\">hi</td></tr>")

	c(core.NewTableRow(
		core.NewTableCell(core.NewText("hi")),
		core.NewTableCell(core.NewText("there")),
	)).Returns("<tr><td scope=\"col\">hi</td><td scope=\"col\">there</td></tr>")
}
