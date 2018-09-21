package gedcom

// NoteNode represents additional information provided by the submitter for
// understanding the enclosing data.
type NoteNode struct {
	*SimpleNode
}

// NewNoteNode creates a new NOTE node.
func NewNoteNode(document *Document, value, pointer string, children []Node) *NoteNode {
	return &NoteNode{
		newSimpleNode(document, TagNote, value, pointer, children),
	}
}
