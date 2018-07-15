package main

// footer appears on all pages at the bottom.
type footer struct{}

func newFooter() *footer {
	return &footer{}
}

func (c *footer) String() string {
	return newComponents(
		newHorizontalRuleRow(),
		newRow(newColumn(entireRow, newDiv("text-center", newText(`Generated with
		<a href="https://github.com/elliotchance/gedcom">github.com/elliotchance/gedcom</a>.`)))),
		newSpace(),
	).String()
}
