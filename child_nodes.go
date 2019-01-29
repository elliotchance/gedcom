package gedcom

type ChildNodes []*ChildNode

func (nodes ChildNodes) Individuals() (individuals IndividualNodes) {
	for _, child := range nodes {
		pointer := valueToPointer(child.Value())
		individual := nodes[0].Family().Document().NodeByPointer(pointer)

		individuals = append(individuals, individual.(*IndividualNode))
	}

	return
}

func (nodes ChildNodes) Similarity(other ChildNodes, options SimilarityOptions) float64 {
	return nodes.Individuals().Similarity(other.Individuals(), options)
}

func (nodes ChildNodes) ByPointer(pointer string) *ChildNode {
	for _, node := range nodes {
		if node.Individual().Pointer() == pointer {
			return node
		}
	}

	return nil
}

func (nodes ChildNodes) IndividualByPointer(pointer string) *IndividualNode {
	for _, node := range nodes {
		if node.Individual().Pointer() == pointer {
			return node.Individual()
		}
	}

	return nil
}
