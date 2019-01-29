package gedcom

// Noder allows an instance to have child nodes.
type Noder interface {
	// Nodes returns any child nodes.
	Nodes() Nodes

	// AddNode will add a child to this node.
	//
	// There is no restriction on whether a node is not allow to have children
	// so you can expect that no error can occur.
	//
	// AddNode will always append the child at the end, even if there is is an
	// exact child that already exists. However, the order of node in a GEDCOM
	// file is almost always irrelevant.
	AddNode(node Node)

	// SetNodes replaces all of the child nodes.
	SetNodes(nodes Nodes)

	DeleteNode(node Node) bool
}
