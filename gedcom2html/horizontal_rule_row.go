package main

// horizontalRuleRow is a dividing line.
type horizontalRuleRow struct{}

func newHorizontalRuleRow() *horizontalRuleRow {
	return &horizontalRuleRow{}
}

func (c *horizontalRuleRow) String() string {
	return newRow(newColumn(entireRow, newHorizontalRule())).String()
}
