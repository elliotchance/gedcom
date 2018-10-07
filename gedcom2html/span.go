package main

import "fmt"

type span struct {
	class, value string
}

func newSpan(class, value string) *span {
	return &span{
		value: value,
		class: class,
	}
}

func (c *span) String() string {
	return fmt.Sprintf(`<span class="%s">%s</span>`, c.class, c.value)
}
