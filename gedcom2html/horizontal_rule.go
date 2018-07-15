package main

// horizontalRule is a dividing line.
type horizontalRule struct{}

func newHorizontalRule() *horizontalRule {
	return &horizontalRule{}
}

func (c *horizontalRule) String() string {
	return "<hr/>"
}
