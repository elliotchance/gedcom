package gedcom

import "bytes"

// DocumentNode represents a whole GEDCOM document. It is possible for a
// DocumentNode to contain zero Nodes, this means the GEDCOM file was empty. It
// may also (and usually) contain several Nodes.
type DocumentNode struct {
	Nodes []Node
}

// String will render the entire GEDCOM document.
func (node *DocumentNode) String() string {
	buf := bytes.NewBufferString("")

	for i, child := range node.Nodes {
		if i > 0 {
			buf.WriteByte('\n')
		}
		buf.WriteString(child.String())
	}

	return buf.String()
}
