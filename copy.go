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

// DeepCopy returns a new node with all recursive children also duplicated. The
// document provided will be where the new node will be attached. This can be
// the same document, but it must not be nil.
func DeepCopy(node Node, document *Document) Node {
	if IsNil(node) {
		return nil
	}

	// We must track the last family seen for nodes that require a family. For
	// example, husband, wife and child nodes.
	var family *FamilyNode

	return Filter(node, document, func(node Node) (newNode Node, traverseChildren bool) {
		if fam, ok := node.(*FamilyNode); ok {
			family = fam
		}

		return shallowCopyNode(node, document, family), true
	})
}
