package gedcom

// FilterTraverseMode controls how Filter will handle child nodes. See constants
// for usage.
type FilterTraverseMode int

const (
	// FilterTraverseModeOldChildren is the default traverse mode. It will cause
	// Filter() to traverse the children of the original node passed into the
	// filter function and ignore any children returned by the new node.
	FilterTraverseModeOldChildren = FilterTraverseMode(iota)

	// FilterTraverseModeStop will not traverse any children. If the new node
	// contains any children they will be replaced on the original node.
	FilterTraverseModeStop = FilterTraverseMode(iota)

	// FilterTraverseModeNewChildren will replace the original node with the
	// newly returned node, including children. Then proceed to traverse the new
	// children.
	FilterTraverseModeNewChildren = FilterTraverseMode(iota)
)

// FilterResult is returned from a filtering function and provides state and
// instructions about how Filter() should proceed.
type FilterResult struct {
	// Node replaces the node as passed in to the filter function. You should
	// return the same node argument if you do not want the node to be changed.
	//
	// If Node is nil then it will be removed and the children will not be
	// traversed, regardless of the FilterTraverseMode value.
	Node

	// FilterTraverseMode decides how the traversal should continue. See
	// FilterTraverseMode constants for more information.
	FilterTraverseMode
}

// FilterFunction is used with the Filter function.
//
// See FilterResult for a full explanation.
type FilterFunction func(node Node) FilterResult

// Filter returns a new nest node structure by recursively filtering all
// children based on a callback FilterFunction.
//
// There are several other functions that can be used as filters including;
// WhitelistTagFilter, BlacklistTagFilter, OfficialTagFilter and more.
func Filter(root Node, fn FilterFunction) Node {
	result := fn(root)
	if result.Node == nil {
		return nil
	}

	nodesToTraverse := []Node(nil)

	switch result.FilterTraverseMode {
	case FilterTraverseModeStop:
		return result.Node

	case FilterTraverseModeOldChildren:
		nodesToTraverse = root.Nodes()

	case FilterTraverseModeNewChildren:
		nodesToTraverse = result.Node.Nodes()
	}

	newNode := shallowCopyNode(result.Node)
	for _, child := range nodesToTraverse {
		newNode := Filter(child, fn)
		if newNode != nil {
			newNode.AddNode(newNode)
		}
	}

	return newNode
}

func shallowCopyNode(node Node) Node {
	document := node.Document()
	tag := node.Tag()
	value := node.Value()
	pointer := node.Pointer()

	return NewNode(document, tag, value, pointer)
}

// WhitelistTagFilter returns any node that is one of the provided tags.
//
// This is the opposite of BlacklistTagFilter.
//
// See the Filter() function.
func WhitelistTagFilter(tags ...Tag) FilterFunction {
	return func(node Node) FilterResult {
		for _, tag := range tags {
			if tag.Is(node.Tag()) {
				return FilterResult{node, FilterTraverseModeOldChildren}
			}
		}

		return FilterResult{nil, FilterTraverseModeStop}
	}
}

// BlacklistTagFilter returns any node that is not one of the provided tags.
//
// This is the opposite of WhitelistTagFilter.
//
// See the Filter function.
func BlacklistTagFilter(tags ...Tag) FilterFunction {
	return func(node Node) FilterResult {
		for _, tag := range tags {
			if tag.Is(node.Tag()) {
				return FilterResult{nil, FilterTraverseModeStop}
			}
		}

		return FilterResult{node, FilterTraverseModeOldChildren}
	}
}

// OfficialTagFilter returns any node that is official tag. See Tag.IsOfficial
// for more information.
//
// See the Filter function.
func OfficialTagFilter() FilterFunction {
	return func(node Node) FilterResult {
		isOfficial := node.Tag().IsOfficial()
		newNode := NodeCondition(isOfficial, node, nil)

		if isOfficial {
			return FilterResult{newNode, FilterTraverseModeOldChildren}
		}

		return FilterResult{newNode, FilterTraverseModeStop}
	}
}

// SimpleNameFilter flattens NAME nodes.
//
// This is useful for comparing names when the components of the name (title,
// suffix, etc) are less important than the name itself.
//
// The new name nodes will have a value constructed with NameNode.GedcomName.
func SimpleNameFilter() FilterFunction {
	return func(node Node) FilterResult {
		if name, ok := node.(*NameNode); ok {
			newNode := NewNameNode(name.Document(), name.GedcomName(),
				name.Pointer(), nil)

			return FilterResult{newNode, FilterTraverseModeStop}
		}

		return FilterResult{node, FilterTraverseModeOldChildren}
	}
}

// SingleNameFilter will remove all but the first name found on an individuals.
func SingleNameFilter() FilterFunction {
	return func(node Node) FilterResult {
		if individual, ok := node.(*IndividualNode); ok {
			names := individual.Names()

			if len(names) != 0 && len(names) > 1 {
				// Start with the first name and add all the remaining non-name
				// nodes.
				newChildren := []Node{names[0]}

				for _, child := range individual.Nodes() {
					if !child.Tag().Is(TagName) {
						newChildren = append(newChildren, child)
					}
				}

				newIndividual := NewIndividualNode(individual.Document(),
					individual.Value(), individual.Pointer(), newChildren)

				return FilterResult{
					newIndividual, FilterTraverseModeNewChildren,
				}
			}

			return FilterResult{individual, FilterTraverseModeOldChildren}
		}

		return FilterResult{node, FilterTraverseModeOldChildren}
	}
}
