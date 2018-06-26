package gedcom

import (
	"fmt"
	"bytes"
)

// SimpleNode is used as the default node type when there is no more appropriate
// or specific type to use.
type SimpleNode struct {
	tag      Tag
	value    string
	pointer  string
	children []Node
}

func NewSimpleNode(tag Tag, value, pointer string, children []Node) *SimpleNode {
	return &SimpleNode{
		tag:      tag,
		value:    value,
		pointer:  pointer,
		children: children,
	}
}

func (node *SimpleNode) Tag() Tag {
	return node.tag
}

func (node *SimpleNode) Value() string {
	return node.value
}

func (node *SimpleNode) Pointer() string {
	return node.pointer
}

func (node *SimpleNode) ChildNodes() []Node {
	return node.children
}

func (node *SimpleNode) AddChildNode(n Node) {
	node.children = append(node.children, n)
}

func (node *SimpleNode) String() string {
	buf := bytes.NewBufferString("")

	if node.pointer != "" {
		buf.WriteString(fmt.Sprintf("@%s@ ", node.pointer))
	}

	buf.WriteString(string(node.tag))

	if node.value != "" {
		buf.WriteByte(' ')
		buf.WriteString(node.value)
	}

	return buf.String()
}
