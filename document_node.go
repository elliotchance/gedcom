package gedcom

// DocumentNode represents a whole GEDCOM document. It is possible for a
// DocumentNode to contain zero Nodes, this means the GEDCOM file was empty. It
// may also (and usually) contain several Nodes.
type DocumentNode struct {
	Nodes []Node
}
