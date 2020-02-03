package gedcom

import "github.com/elliotchance/gedcom/tag"

// NoteNode represents additional information provided by the submitter for
// understanding the enclosing data.
type NoteNode struct {
	*SimpleNode
}

// NewNoteNode creates a new NOTE node.
func NewNoteNode(value string, children ...Node) *NoteNode {
	return &NoteNode{
		newSimpleNode(tag.TagNote, value, "", children...),
	}
}
