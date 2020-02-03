package gedcom

import "github.com/elliotchance/gedcom/tag"

// FormatNode represents an assigned name given to a consistent format in which
// information can be conveyed.
type FormatNode struct {
	*SimpleNode
}

// NewFormatNode creates a new FORM node.
func NewFormatNode(value string, children ...Node) *FormatNode {
	return &FormatNode{
		newSimpleNode(tag.TagFormat, value, "", children...),
	}
}
