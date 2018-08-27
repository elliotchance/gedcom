package gedcom

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
// Note: You should not use this constructor for general use. Instead use
// NewNode which will return a *SimpleNode if a more appropriate node type
// exists for the tag.
func NewSimpleNode(document *Document, tag Tag, value, pointer string, children []Node) *SimpleNode {
	return &SimpleNode{
		document: document,
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

func (node *SimpleNode) Document() *Document {
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
	if node == nil || IsNil(node2) {
		return false
	}

	return node.tag == node2.Tag() &&
		node.value == node2.Value() &&
		node.pointer == node2.Pointer()
}

func (node *SimpleNode) SetDocument(document *Document) {
	node.document = document

	for _, child := range node.children {
		child.SetDocument(document)
	}
}

func (node *SimpleNode) Nodes() []Node {
	return node.children
}

func (node *SimpleNode) AddNode(n Node) {
	node.children = append(node.children, n)
}

func (node *SimpleNode) String() string {
	return GedcomLine(0, node)
}
