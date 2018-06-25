package gedcom

import (
	"fmt"
	"bytes"
)

// SimpleNode is used as the default node type when there is no more appropriate
// or specific type to use.
type SimpleNode struct {
	indent   int
	tag      string
	value    string
	pointer  string
	children []Node
}

func NewSimpleNode(indent int, tag, value, pointer string, children []Node) *SimpleNode {
	return &SimpleNode{
		indent:   indent,
		tag:      tag,
		value:    value,
		pointer:  pointer,
		children: children,
	}
}

func (node *SimpleNode) Indent() int {
	return node.indent
}

func (node *SimpleNode) Tag() string {
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
	buf := bytes.NewBufferString(fmt.Sprintf("%d", node.indent))

	if node.pointer != "" {
		buf.WriteString(fmt.Sprintf(" @%s@", node.pointer))
	}

	buf.WriteByte(' ')
	buf.WriteString(node.tag)

	if node.value != "" {
		buf.WriteByte(' ')
		buf.WriteString(node.value)
	}

	for _, child := range node.children {
		buf.WriteByte('\n')
		buf.WriteString(child.String())
	}

	return buf.String()
}
