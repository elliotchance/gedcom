package html

import "io"

// FooterRow appears on all pages at the bottom.
type FooterRow struct{}

func NewFooterRow() *FooterRow {
	return &FooterRow{}
}

func (c *FooterRow) WriteTo(w io.Writer) (int64, error) {
	link := NewLink(
		NewText("github.com/elliotchance/gedcom"),
		"https://github.com/elliotchance/gedcom",
	)

	content := NewComponents(
		NewText("Generated with "),
		link,
	)

	return NewComponents(
		NewHorizontalRuleRow(),
		NewRow(NewColumn(EntireRow, NewDiv("text-center", content))),
		NewSpace(),
	).WriteTo(w)
}
