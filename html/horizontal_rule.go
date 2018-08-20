package html

// HorizontalRule is a dividing line.
type HorizontalRule struct{}

func NewHorizontalRule() *HorizontalRule {
	return &HorizontalRule{}
}

func (c *HorizontalRule) String() string {
	return "<hr/>"
}
