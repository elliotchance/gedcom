package gedcom

// Flatten returns a slice of nodes from a tree structure.
//
// The original nodes will not be modified, but they will also not be copied.
// Any changes to the result nodes will directly affect the original tree
// structure. This also means that the nodes returned will have the references
// to their children as before.
//
// Flatten makes no guarantees about the same node appearing more than once if
// it also appears more then once in the original structure. Be careful not to
// pass in structures that have circular references between nodes.
//
// If the input is nil then the result will also be nil.
//
// If you would like a completely new copy of the data you can use:
//
//   Flatten(DeepCopy(node))
//
func Flatten(node Node) []Node {
	if IsNil(node) {
		return nil
	}

	result := []Node{}

	Filter(node, func(node Node) (newNode Node, traverseChildren bool) {
		result = append(result, node)

		return node, true
	})

	return result
}

// FlattenAll works as Flatten with multiple inputs that are returned as a
// single slice.
//
// If any of the nodes are nil they will be ignored.
func FlattenAll(nodes []Node) (result []Node) {
	for _, node := range nodes {
		if IsNil(node) {
			continue
		}

		result = append(result, Flatten(node)...)
	}

	return
}
