package html

// Empty is used a placeholder for a component where nothing should be visible.
type Empty struct{}

func NewEmpty() *Empty {
	return &Empty{}
}

func (c *Empty) String() string {
	return "&nbsp;"
}
