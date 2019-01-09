package html

import "io"

// HorizontalRuleRow is a dividing line.
type HorizontalRuleRow struct{}

func NewHorizontalRuleRow() *HorizontalRuleRow {
	return &HorizontalRuleRow{}
}

func (c *HorizontalRuleRow) WriteTo(w io.Writer) (int64, error) {
	return NewRow(NewColumn(EntireRow, NewHorizontalRule())).WriteTo(w)
}
