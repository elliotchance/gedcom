package core_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/html/core"
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
