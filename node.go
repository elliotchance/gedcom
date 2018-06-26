package gedcom

import "fmt"

type Node interface {
	fmt.Stringer

	Tag() Tag
	Value() string
	Pointer() string
	ChildNodes() []Node
	AddChildNode(node Node)
}
