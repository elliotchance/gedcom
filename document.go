package gedcom

import "bytes"

// Document represents a whole GEDCOM document. It is possible for a
// Document to contain zero Nodes, this means the GEDCOM file was empty. It
// may also (and usually) contain several Nodes.
type Document struct {
	Nodes        []Node
	pointerCache map[string]Node
}

// String will render the entire GEDCOM document.
func (doc *Document) String() string {
	buf := bytes.NewBufferString("")

	encoder := NewEncoder(buf, doc)
	err := encoder.Encode()
	if err != nil {
		panic(err)
	}

	return buf.String()
}

func (doc *Document) Individuals() []*IndividualNode {
	individuals := []*IndividualNode{}

	for _, node := range doc.Nodes {
		if n, ok := node.(*IndividualNode); ok {
			individuals = append(individuals, n)
		}
	}

	return individuals
}

func (doc *Document) NodeByPointer(ptr string) Node {
	// Build the cache once.
	if doc.pointerCache == nil {
		doc.pointerCache = map[string]Node{}

		for _, node := range doc.Nodes {
			if node.Pointer() != "" {
				doc.pointerCache[node.Pointer()] = node
			}
		}
	}

	return doc.pointerCache[ptr]
}

func (doc *Document) Families() []*FamilyNode {
	families := []*FamilyNode{}

	for _, node := range doc.Nodes {
		if n, ok := node.(*FamilyNode); ok {
			families = append(families, n)
		}
	}

	return families
}

// TODO: needs tests
func (doc *Document) Places() map[*PlaceNode]Node {
	places := map[*PlaceNode]Node{}

	for _, node := range doc.Nodes {
		extractPlaces(node, places)
	}

	return places
}

func extractPlaces(n Node, dest map[*PlaceNode]Node) {
	for _, node := range n.Nodes() {
		if place, ok := node.(*PlaceNode); ok {
			// The place points to the parent node which is the thing that the
			// place is describing.
			dest[place] = n
		} else {
			extractPlaces(node, dest)
		}
	}
}

// TODO: Needs tests
func (doc *Document) Sources() []*SourceNode {
	sources := []*SourceNode{}

	for _, node := range doc.Nodes {
		if n, ok := node.(*SourceNode); ok {
			sources = append(sources, n)
		}
	}

	return sources
}
