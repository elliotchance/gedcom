package core

import "io"
import "time"

// FooterRow appears on all pages at the bottom.
type FooterRow struct{}

func NewFooterRow() *FooterRow {
	return &FooterRow{}
}

func (c *FooterRow) WriteHTMLTo(w io.Writer) (int64, error) {
	link := NewLink(
		NewText("github.com/elliotchance/gedcom"),
		"https://github.com/elliotchance/gedcom",
	)

  curTime := time.Now()
  layout := "Jan 2, 2006, 3.04 pm"
	content := NewComponents(
		NewText("Generated with "),
		link,
    NewText(" on "),
    NewText(curTime.Format(layout)),
	)

	return NewComponents(
		NewHorizontalRuleRow(),
		NewRow(NewColumn(EntireRow, NewDiv("text-center", content))),
		NewSpace(),
	).WriteHTMLTo(w)
}
