package html

// HorizontalRuleRow is a dividing line.
type HorizontalRuleRow struct{}

func NewHorizontalRuleRow() *HorizontalRuleRow {
	return &HorizontalRuleRow{}
}

func (c *HorizontalRuleRow) String() string {
	return NewRow(NewColumn(EntireRow, NewHorizontalRule())).String()
}
