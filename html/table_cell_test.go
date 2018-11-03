package html

import (
	"github.com/elliotchance/tf"
	"testing"
)

func TestTableCell_String(t *testing.T) {
	String := tf.Function(t, (*TableCell).String)

	c := NewText("foo bar")

	String(NewTableCell(c)).Returns(`<td scope="col">foo bar</td>`)
}

func TestTableCell_Header(t *testing.T) {
	String := tf.Function(t, (*TableCell).String)

	c := NewText("foo bar")

	String(NewTableCell(c).Header()).Returns(`<th scope="col">foo bar</th>`)
	String(NewTableCell(c).Header().Header()).Returns(`<th scope="col">foo bar</th>`)
}

func TestTableCell_NoWrap(t *testing.T) {
	String := tf.Function(t, (*TableCell).String)

	c := NewText("foo bar")

	String(NewTableCell(c).NoWrap()).Returns(`<td scope="col" nowrap="nowrap">foo bar</td>`)
	String(NewTableCell(c).NoWrap().NoWrap()).Returns(`<td scope="col" nowrap="nowrap">foo bar</td>`)
}

func TestTableCell_Class(t *testing.T) {
	String := tf.Function(t, (*TableCell).String)

	c := NewText("foo bar")

	String(NewTableCell(c).Class("")).Returns(`<td scope="col">foo bar</td>`)
	String(NewTableCell(c).Class("dot")).Returns(`<td scope="col" class="dot">foo bar</td>`)
	String(NewTableCell(c).Class("dot").Class("line")).Returns(`<td scope="col" class="line">foo bar</td>`)
}

func TestTableCell_Style(t *testing.T) {
	String := tf.Function(t, (*TableCell).String)

	c := NewText("foo bar")

	String(NewTableCell(c).Style("")).Returns(`<td scope="col">foo bar</td>`)
	String(NewTableCell(c).Style("dot")).Returns(`<td scope="col" style="dot">foo bar</td>`)
	String(NewTableCell(c).Style("dot").Style("line")).Returns(`<td scope="col" style="line">foo bar</td>`)
}
