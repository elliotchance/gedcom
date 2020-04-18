// Filtering and Tree Walking
//
// The Filter function recursively removes or manipulates nodes with a
// FilterFunction:
//
//   newNodes := gedcom.Filter(node, func (node gedcom.Node) (gedcom.Node, bool) {
//     if node.Tag().Is(gedcom.TagIndividual) {
//       // false means it will not traverse children, since an
//       // individual can never be inside of another individual.
//       return node, false
//     }
//
//     return nil, false
//   })
//
//   // Remove all tags that are not official.
//   newNodes := gedcom.Filter(node, gedcom.OfficialTagFilter())
//
// Some examples of Filter functions include BlacklistTagFilter,
// OfficialTagFilter, SimpleNameFilter and WhitelistTagFilter.
package gedcom

// FilterFunction is used with the Filter function.
//
// The node (as passed in through the parameter) will be replaced with newNode.
// You should return the same node argument if you do not want the node to be
// changed.
//
// The traverseChildren argument decides if the traversal should continue
// through the children of node. If the traversal continues (traverseChildren is
// true) the FilterFunction will always receive the children of the node, and
// not the children (if any) of newNode.
//
// If the newNode is nil then it will be removed and the children will not be
// traversed, regardless of the traverseChildren value.
type FilterFunction func(node Node) (newNode Node, traverseChildren bool)

// Filter returns a new nest node structure by recursively filtering all
// children based on a callback FilterFunction.
//
// See FilterFunction for the implementation details of fn.
//
// There are several other functions that can be used as filters including;
// WhitelistTagFilter, BlacklistTagFilter and OfficialTagFilter.
//
// Some nodes, such as an IndividualNode or FamilyNode cannot be created without
// being attached to a document. Filter will attach these to a new document
// which can be accessed through their respective Document() method.
func Filter(root Node, document *Document, fn FilterFunction) Node {
	entityMap := entityMap{}

	return filter(root, fn, entityMap, document, nil)
}

func filter(root Node, fn FilterFunction, entityMap entityMap, document *Document, family *FamilyNode) Node {
	newRoot, keepTraversing := fn(root)
	if IsNil(newRoot) {
		return nil
	}

	if familyNoder, ok := newRoot.(FamilyNoder); ok {
		fam := familyNoder.Family()

		family = entityMap.GetOrAssign(fam, func() interface{} {
			return document.AddFamily(fam.Pointer())
		}).(*FamilyNode)
	}

	result := shallowCopyNode(newRoot, document, family)

	if keepTraversing {
		for _, child := range root.Nodes() {
			newNode := filter(child, fn, entityMap, document, family)
			if newNode != nil {
				result.AddNode(newNode)
			}
		}
	} else {
		for _, child := range newRoot.Nodes() {
			result.AddNode(child)
		}
	}

	return result
}

// Copy a node without children.
//
// Some nodes require a document (such as an IndividualNode) or family (such as
// a ChildNode) to be created. Since we don't want to attach these to the
// existing documents, families, etc. new entities will have to be passed in.
//
// One important thing to note here is that we don't want to create new
// documents, etc for every single node we copy because that will leave the new
// nodes totally fractured and not in a state that we would expect to traverse.
// Be careful to reuse the document and other entities in a reasonable way.
func shallowCopyNode(node Node, document *Document, family *FamilyNode) Node {
	tag := node.Tag()
	value := node.Value()
	pointer := node.Pointer()

	return newNode(document, family, tag, value, pointer)
}

// WhitelistTagFilter returns any node that is one of the provided tags.
//
// This is the opposite of BlacklistTagFilter.
//
// See the Filter() function.
func WhitelistTagFilter(tags ...Tag) FilterFunction {
	return func(node Node) (Node, bool) {
		for _, tag := range tags {
			if tag.Is(node.Tag()) {
				return node, true
			}
		}

		return nil, false
	}
}

// BlacklistTagFilter returns any node that is not one of the provided tags.
//
// This is the opposite of WhitelistTagFilter.
//
// See the Filter function.
func BlacklistTagFilter(tags ...Tag) FilterFunction {
	return func(node Node) (Node, bool) {
		for _, tag := range tags {
			if tag.Is(node.Tag()) {
				return nil, false
			}
		}

		return node, true
	}
}

// OfficialTagFilter returns any node that is official tag. See Tag.IsOfficial
// for more information.
//
// See the Filter function.
func OfficialTagFilter() FilterFunction {
	return func(node Node) (Node, bool) {
		isOfficial := node.Tag().IsOfficial()

		return NodeCondition(isOfficial, node, nil), isOfficial
	}
}

// SimpleNameFilter flattens NAME nodes with the provided format.
//
// This is useful for comparing names when the components of the name (title,
// suffix, etc) are less important than the name itself.
func SimpleNameFilter(format NameFormat) FilterFunction {
	return func(node Node) (Node, bool) {
		if name, ok := node.(*NameNode); ok {
			newNode := newNode(
				nil,
				nil,
				TagName,
				name.Format(format),
				name.Pointer(),
			)

			return newNode, false
		}

		return node, true
	}
}

// OnlyVitalsTagFilter removes all tags except for vital individual information.
//
// The vital nodes are (or multiples in the same individual of): Name, birth,
// baptism, death and burial. Within these only the date and place is retained.
func OnlyVitalsTagFilter() FilterFunction {
	return WhitelistTagFilter(
		// Level 0: We have to allow this for the children.
		TagIndividual,

		// Level 1.
		TagName, TagBirth, TagBaptism, TagDeath, TagBurial,

		// Level 2: These should only ever appear as direct children of the tags
		// above.
		TagGivenName, TagSurname, TagSurnamePrefix, TagNamePrefix,
		TagNameSuffix, TagTitle, TagDate, TagPlace,
	)
}

// RemoveEmptyDeathTagFilter removes any Death (DEAT) events that do not have
// any child nodes (which would otherwise be information like the date or place.
//
// This is because some applications use the death tag as a marker without any
// further information which can cause problems when comparing individuals.
func RemoveEmptyDeathTagFilter() FilterFunction {
	return func(node Node) (Node, bool) {
		if death, ok := node.(*DeathNode); ok && len(death.Nodes()) == 0 {
			return nil, false
		}

		return node, true
	}
}

func RemoveDuplicateNamesFilter() FilterFunction {
	return func(node Node) (Node, bool) {
		if individual, ok := node.(*IndividualNode); ok {
			newIndividual := newIndividualNode(individual.Document(),
				individual.Pointer())
			names := map[string]bool{}

			for _, n := range individual.children {
				if name, isName := n.(*NameNode); isName {
					nameString := name.String()
					if names[nameString] == true {
						continue
					}

					names[nameString] = true
				}
				newIndividual.AddNode(n)
			}

			return newIndividual, false
		}

		// Individuals can only exist on the root level so there's no need to
		// recurse.
		return node, true
	}
}
