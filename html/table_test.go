package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestTable_WriteTo(t *testing.T) {
	c := testComponent(t, "Table")

	c(html.NewTable("")).
		Returns("<table class=\"table \"></table>")

	c(html.NewTable("colorful stuff")).
		Returns("<table class=\"table colorful stuff\"></table>")

	c(html.NewTable("c", html.NewTableRow())).
		Returns("<table class=\"table c\"><tr></tr></table>")

	c(html.NewTable("c", html.NewTableRow(html.NewTableCell(html.NewText("ok"))))).
		Returns("<table class=\"table c\"><tr><td scope=\"col\">ok</td></tr></table>")

	c(html.NewTable("c", html.NewTableRow(
		html.NewTableCell(html.NewText("ok")),
		html.NewTableCell(html.NewText("cool")),
	))).
		Returns("<table class=\"table c\"><tr><td scope=\"col\">ok</td><td scope=\"col\">cool</td></tr></table>")

	c(html.NewTable("c",
		html.NewTableRow(
			html.NewTableCell(html.NewText("ok")),
		),
		html.NewTableRow(
			html.NewTableCell(html.NewText("cool")),
		),
	)).
		Returns("<table class=\"table c\"><tr><td scope=\"col\">ok</td></tr><tr><td scope=\"col\">cool</td></tr></table>")
}
