package core_test

import (
	"github.com/elliotchance/gedcom/html/core"
	"testing"
)

func TestTable_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "Table")

	c(core.NewTable("")).
		Returns("<table class=\"table \"></table>")

	c(core.NewTable("colorful stuff")).
		Returns("<table class=\"table colorful stuff\"></table>")

	c(core.NewTable("c", core.NewTableRow())).
		Returns("<table class=\"table c\"><tr></tr></table>")

	c(core.NewTable("c", core.NewTableRow(core.NewTableCell(core.NewText("ok"))))).
		Returns("<table class=\"table c\"><tr><td scope=\"col\">ok</td></tr></table>")

	c(core.NewTable("c", core.NewTableRow(
		core.NewTableCell(core.NewText("ok")),
		core.NewTableCell(core.NewText("cool")),
	))).
		Returns("<table class=\"table c\"><tr><td scope=\"col\">ok</td><td scope=\"col\">cool</td></tr></table>")

	c(core.NewTable("c",
		core.NewTableRow(
			core.NewTableCell(core.NewText("ok")),
		),
		core.NewTableRow(
			core.NewTableCell(core.NewText("cool")),
		),
	)).
		Returns("<table class=\"table c\"><tr><td scope=\"col\">ok</td></tr><tr><td scope=\"col\">cool</td></tr></table>")
}
