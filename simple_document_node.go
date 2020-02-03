package gedcom

import "github.com/elliotchance/gedcom/tag"

type simpleDocumentNode struct {
	*SimpleNode
	document *Document
}

func newSimpleDocumentNode(document *Document, tag tag.Tag, value, pointer string, children ...Node) *simpleDocumentNode {
	return &simpleDocumentNode{
		SimpleNode: newSimpleNode(tag, value, pointer, children...),
		document:   document,
	}
}

func (node *simpleDocumentNode) Document() *Document {
	return node.document
}

func (node *simpleDocumentNode) ShallowCopy() Node {
	if IsNil(node) {
		return nil
	}

	document := node.Document()
	tag := node.Tag()
	value := node.Value()
	pointer := node.Pointer()

	newNode := newSimpleDocumentNode(document, tag, value, pointer)
	document.AddNode(newNode)

	return newNode
}
