package gedcom

type NodeCopier interface {
	// ShallowCopy returns a new node with all the same properties, but no
	// children.
	//
	// You should assume that it is not safe to use ShallowCopy() on a nil
	// value.
	//
	// See DeepCopy() for copying nodes recursively.
	ShallowCopy() Node
}

// DeepCopy returns a new node with all recursive children also duplicated.
func DeepCopy(node Node) Node {
	if IsNil(node) {
		return nil
	}

	return Filter(node, func(node Node) (newNode Node, traverseChildren bool) {
		return node.ShallowCopy(), true
	})
}
