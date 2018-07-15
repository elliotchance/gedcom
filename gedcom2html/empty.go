package main

// empty is used a placeholder for a component where nothing should be visible.
type empty struct{}

func newEmpty() *empty {
	return &empty{}
}

func (c *empty) String() string {
	return "&nbsp;"
}
