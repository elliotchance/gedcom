package gedcom

import "fmt"

type Node interface {
	fmt.Stringer

	Indent() int
	Tag() string
	Value() string
	Pointer() string
	ChildNodes() []Node
	AddChildNode(node Node)
}
