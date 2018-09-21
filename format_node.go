package gedcom

// FormatNode represents an assigned name given to a consistent format in which
// information can be conveyed.
type FormatNode struct {
	*SimpleNode
}

// NewFormatNode creates a new FORM node.
func NewFormatNode(document *Document, value, pointer string, children []Node) *FormatNode {
	return &FormatNode{
		NewSimpleNode(document, TagFormat, value, pointer, children),
	}
}
