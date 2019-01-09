package html_test

import (
	"github.com/elliotchance/gedcom/html"
	"testing"
)

func TestStyledTableCell_WriteTo(t *testing.T) {
	c := testComponent(t, "StyledTableCell")

	c(html.NewStyledTableCell("color:black", "", html.NewText("hi"))).
		Returns("<td scope=\"col\" style=\"color:black\">hi</td>")

	c(html.NewStyledTableCell("color:red", "c1 c2", html.NewText("ok"))).
		Returns("<td class=\"c1 c2\" scope=\"col\" style=\"color:red\">ok</td>")
}
