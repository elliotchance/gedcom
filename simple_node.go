package gedcom

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"sync"
)

// SimpleNode is used as the default node type when there is no more appropriate
// or specific type to use.
type SimpleNode struct {
	tag      Tag
	value    string
	pointer  string
	children Nodes
}

// newSimpleNode creates a non-specific node.
//
// Unlike all of the other node types this constructor is not public because it
// is used internally by NewNode if a specific node type can not be determined.
func newSimpleNode(tag Tag, value, pointer string, children ...Node) *SimpleNode {
	return &SimpleNode{
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

// Identifier returns the identifier for the node; empty string if the node is nil
func (node *SimpleNode) Identifier() string {
	if node == nil {
		return ""
	}

	return fmt.Sprintf("@%s@", node.pointer)
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

	useAncestrySourceMatching := flag.Lookup("ancestry-source-matching").Value.String() //indexes a map CommandLine.formal
	//if both Ancestry sources, only check if their _APID is the same
	if useAncestrySourceMatching == "true" && node.Tag().String() == "Source" && tag.String() == "Source" {
		if node.Value() == node2.Value() { //if they have the same source id, then no need to check the apid
			return true
		}
		for _, leftNode := range node.Nodes() {
			for _, rightNode := range node2.Nodes() {
				if leftNode.Tag().String() == "_APID" &&
					rightNode.Tag().String() == "_APID" &&
					rightNode.Value() == leftNode.Value() {
					return true
				}
			}
		}
	}
	value := node2.Value()
	if node.value != value {
		return false
	}

	return node.pointer == node2.Pointer()
}

// If the node is nil the result will also be nil.
func (node *SimpleNode) Nodes() Nodes {
	if node == nil {
		return nil
	}

	return node.children
}

func (node *SimpleNode) AddNode(n Node) {
	node.children = append(node.children, n)

	// This is pretty crude and nasty. I'm sorry if your workflow is to switch
	// between small changes and large sweeping reads but this will do for now.
	//
	// We can't simply remove this node because we would have to make sure we
	// work our way up the chain which we have no easy way of doing right now.
	nodeCache = &sync.Map{}
}

func (node *SimpleNode) DeleteNode(n Node) (didDelete bool) {
	node.children, didDelete = node.children.deleteNode(n)

	return
}

// If the node is nil the result be an empty string.
func (node *SimpleNode) String() string {
	if node == nil {
		return ""
	}

	return node.value
}

func (node *SimpleNode) MarshalJSON() ([]byte, error) {
	m := node.ObjectMap()

	return json.Marshal(m)
}

func (node *SimpleNode) ObjectMap() map[string]interface{} {
	m := map[string]interface{}{
		"Tag": node.Tag().Tag(),
	}

	if node.Value() != "" {
		m["Value"] = node.Value()
	}

	if node.Pointer() != "" {
		m["Pointer"] = node.Pointer()
	}

	nodes := node.Nodes()
	if len(nodes) > 0 {
		m["Nodes"] = nodes
	}

	return m
}

// ShallowCopy returns a new node that has the same properties as the input node
// without any children.
//
// If the input node is nil then nil is also returned.
func (node *SimpleNode) ShallowCopy() Node {
	if IsNil(node) {
		return nil
	}

	tag := node.Tag()
	value := node.Value()
	pointer := node.Pointer()

	return NewNode(tag, value, pointer)
}

// GEDCOMString is the recursive version of GEDCOMLine. It will render a node
// and all of its children (if any) as a multi-line GEDCOM string.
//
// GEDCOMString will not work with a nil value. You can use the package
// GEDCOMString function to gracefully handle nils.
//
// The indent will only be included if it is at least 0. If you want to use
// GEDCOMString to compare the string values of nodes or exclude the indent you
// should use the NoIndent constant.
func (node *SimpleNode) GEDCOMString(indent int) string {
	document := NewDocumentWithNodes(Nodes{node})

	return document.GEDCOMString(indent)
}

// GEDCOMLine converts a node into its single line GEDCOM value. It is used
// several places including the actual Encoder.
//
// GEDCOMLine, as the name would suggest, does not handle children. You can use
// GEDCOMString if you need the child nodes as well.
//
// GEDCOMLine will not work with a nil value. You can use the package GEDCOMLine
// function to gracefully handle nils.
//
// The indent will only be included if it is at least 0. If you want to use
// GEDCOMLine to compare the string values of nodes or exclude the indent you
// should use the NoIndent constant.
func (node *SimpleNode) GEDCOMLine(indent int) string {
	buf := bytes.NewBufferString("")

	if indent >= 0 {
		buf.WriteString(fmt.Sprintf("%d ", indent))
	}

	if p := node.Pointer(); p != "" {
		buf.WriteString(fmt.Sprintf("@%s@ ", p))
	}

	buf.WriteString(node.Tag().Tag())

	if v := node.Value(); v != "" {
		buf.WriteByte(' ')
		buf.WriteString(v)
	}

	return buf.String()
}

// SetNodes replaces all of the child nodes.
//
// You can use SetNodes(nil) to remove all child nodes.
func (node *SimpleNode) SetNodes(nodes Nodes) {
	node.children = nodes
}

func (node *SimpleNode) RawSimpleNode() *SimpleNode {
	return node
}
