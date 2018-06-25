package gedcom

type Node interface {
	Indent() int
	Tag() string
	Value() string
	Pointer() string
	ChildNodes() []Node
	AddChildNode(node Node)
}
