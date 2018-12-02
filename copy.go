// Copying
//
// All nodes (since they implement the Node interface) also implement the
// NodeCopier interface which provides the ShallowCopy() function.
//
// A shallow copy returns a new node with all the same properties, but no
// children.
//
// On the other hand there is a DeepCopy function which returns a new node with
// all recursive children also copied. This ensures that the new returned node
// can be manipulated without affecting the original node or any of its
// children.
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
