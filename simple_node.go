package gedcom

import "encoding/json"

// SimpleNode is used as the default node type when there is no more appropriate
// or specific type to use.
type SimpleNode struct {
	document *Document
	tag      Tag
	value    string
	pointer  string
	children []Node
}

// NewSimpleNode creates a non-specific node.
//
// Unlike all of the other node types this constructor is not public because it
// is used internally by NewNode if a specific node type can not be determined.
func newSimpleNode(document *Document, tag Tag, value, pointer string, children []Node) *SimpleNode {
	return &SimpleNode{
		document: document,
		tag:      tag,
		value:    value,
		pointer:  pointer,
		children: children,
	}
}

// If the node is nil the result will be an empty tag.
func (node *SimpleNode) Tag() Tag {
	if node == nil {
		return Tag{}
	}

	return node.tag
}

// If the node is nil the result will be an empty string.
func (node *SimpleNode) Value() string {
	if node == nil {
		return ""
	}

	return node.value
}

// If the node is nil the result will be an empty string.
func (node *SimpleNode) Pointer() string {
	if node == nil {
		return ""
	}

	return node.pointer
}

// If the node is nil the result will also be nil.
func (node *SimpleNode) Document() *Document {
	if node == nil {
		return nil
	}

	return node.document
}

// Equals compares two nodes for value equality.
//
// 1. If either or both nodes are nil then false is always returned.
// 2. Nodes are compared only by their root value (shallow) meaning any value
// for the child nodes is ignored.
// 3. The document the node belongs to is not taken into consideration to be
// able to compare nodes by value across different documents.
// 4. A node is considered to have the same value (and therefore be equal) is
// both nodes share the all of the same tag, value and pointer.
func (node *SimpleNode) Equals(node2 Node) bool {
	if node == nil {
		return false
	}

	if IsNil(node2) {
		return false
	}

	tag := node2.Tag()
	if node.tag != tag {
		return false
	}

	value := node2.Value()
	if node.value != value {
		return false
	}

	return node.pointer == node2.Pointer()
}

// If the node is nil the invocation will not have any effect.
func (node *SimpleNode) SetDocument(document *Document) {
	if node == nil {
		return
	}

	node.document = document

	for _, child := range node.children {
		child.SetDocument(document)
	}
}

// If the node is nil the result will also be nil.
func (node *SimpleNode) Nodes() []Node {
	if node == nil {
		return nil
	}

	return node.children
}

func (node *SimpleNode) AddNode(n Node) {
	node.children = append(node.children, n)
}

// If the node is nil the result be an empty string.
func (node *SimpleNode) String() string {
	if node == nil {
		return ""
	}

	return node.value
}

func (node *SimpleNode) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Tag": node.Tag().Tag(),
	}

	if node.Value() != "" {
		m["Value"] = node.Value()
	}

	if node.Pointer() != "" {
		m["Pointer"] = node.Pointer()
	}

	if len(node.Nodes()) > 0 {
		m["Nodes"] = node.Nodes()
	}

	return json.Marshal(m)
}
