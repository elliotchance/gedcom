package gedcom

import "fmt"

type Node interface {
	fmt.Stringer

	// The node itself.
	Tag() Tag
	Value() string
	Pointer() string

	// Child nodes.
	Nodes() []Node
	AddNode(node Node)
	NodesWithTag(tag Tag) []Node
	FirstNodeWithTag(tag Tag) Node
	FirstNodeWithTagPath(tagPath ...Tag) Node
	HasNestedChild(node Node) bool

	// gedcomLine is for rendering the GEDCOM lines.
	gedcomLine() string
}
