package html

// Footer appears on all pages at the bottom.
type Footer struct{}

func NewFooter() *Footer {
	return &Footer{}
}

func (c *Footer) String() string {
	return NewComponents(
		NewHorizontalRuleRow(),
		NewRow(NewColumn(EntireRow, NewDiv("text-center", NewText(`Generated with
		<a href="https://github.com/elliotchance/gedcom">github.com/elliotchance/gedcom</a>.`)))),
		NewSpace(),
	).String()
}
