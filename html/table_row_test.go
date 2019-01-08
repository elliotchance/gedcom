package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestTableRow_WriteTo(t *testing.T) {
	c := testComponent(t, "TableRow")

	c(html.NewTableRow()).Returns("<tr></tr>")

	c(html.NewTableRow(html.NewTableCell(html.NewText("hi")))).
		Returns("<tr><td scope=\"col\">hi</td></tr>")

	c(html.NewTableRow(
		html.NewTableCell(html.NewText("hi")),
		html.NewTableCell(html.NewText("there")),
	)).Returns("<tr><td scope=\"col\">hi</td><td scope=\"col\">there</td></tr>")
}
