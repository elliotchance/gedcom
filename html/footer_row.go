package html

// FooterRow appears on all pages at the bottom.
type FooterRow struct{}

func NewFooterRow() *FooterRow {
	return &FooterRow{}
}

func (c *FooterRow) String() string {
	content := NewComponents(
		NewText("Generated with "),
		NewLink("github.com/elliotchance/gedcom", "https://github.com/elliotchance/gedcom"),
	)

	return NewComponents(
		NewHorizontalRuleRow(),
		NewRow(NewColumn(EntireRow, NewDiv("text-center", content))),
		NewSpace(),
	).String()
}
