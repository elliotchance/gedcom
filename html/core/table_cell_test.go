package core_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39/html/core"
)

func TestTableCell_WriteHTMLTo(t *testing.T) {
	c := testComponent(t, "TableCell")

	text := core.NewText("foo bar")

	c(core.NewTableCell(text)).Returns(`<td scope="col">foo bar</td>`)
}

func TestTableCell_Header(t *testing.T) {
	c := testComponent(t, "TableCell_Header")

	text := core.NewText("foo bar")

	c(core.NewTableCell(text).Header()).Returns(`<th scope="col">foo bar</th>`)
	c(core.NewTableCell(text).Header().Header()).Returns(`<th scope="col">foo bar</th>`)
}

func TestTableCell_NoWrap(t *testing.T) {
	c := testComponent(t, "TableCell_NoWrap")

	text := core.NewText("foo bar")

	c(core.NewTableCell(text).NoWrap()).Returns(`<td scope="col" nowrap="nowrap">foo bar</td>`)
	c(core.NewTableCell(text).NoWrap().NoWrap()).Returns(`<td scope="col" nowrap="nowrap">foo bar</td>`)
}

func TestTableCell_Class(t *testing.T) {
	c := testComponent(t, "TableCell_Class")

	text := core.NewText("foo bar")

	c(core.NewTableCell(text).Class("")).Returns(`<td scope="col">foo bar</td>`)
	c(core.NewTableCell(text).Class("dot")).Returns(`<td scope="col" class="dot">foo bar</td>`)
	c(core.NewTableCell(text).Class("dot").Class("line")).Returns(`<td scope="col" class="line">foo bar</td>`)
}

func TestTableCell_Style(t *testing.T) {
	c := testComponent(t, "TableCell_Class")

	text := core.NewText("foo bar")

	c(core.NewTableCell(text).Style("")).Returns(`<td scope="col">foo bar</td>`)
	c(core.NewTableCell(text).Style("dot")).Returns(`<td scope="col" style="dot">foo bar</td>`)
	c(core.NewTableCell(text).Style("dot").Style("line")).Returns(`<td scope="col" style="line">foo bar</td>`)
}
