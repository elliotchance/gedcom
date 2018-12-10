package gedcom

// NodeSet is a collection of unique nodes.
//
// The nodes are considered unique only by their address, not by any other
// measure of equality.
type NodeSet map[Node]bool

// Add a node to the set.
//
// The nodes are considered unique only by their address, not by any other
// measure of equality.
//
// Attempting to add a node that already exists will have no effect.
func (ns NodeSet) Add(node Node) {
	ns[node] = true
}

// Has returns true if a node exists in the set.
func (ns NodeSet) Has(node Node) bool {
	_, ok := ns[node]

	return ok
}
