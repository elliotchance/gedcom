package gedcom

// WifeNode is an individual in the role as a mother and/or married woman.
type WifeNode struct {
	*SimpleNode
	family *FamilyNode
}

func newWifeNode(family *FamilyNode, value string, children ...Node) *WifeNode {
	return &WifeNode{
		newSimpleNode(TagWife, value, "", children...),
		family,
	}
}

func (node *WifeNode) Family() *FamilyNode {
	return node.family
}

func (node *WifeNode) Individual() *IndividualNode {
	if node == nil {
		return nil
	}

	n := node.family.document.NodeByPointer(valueToPointer(node.value))

	if IsNil(n) {
		return nil
	}

	return n.(*IndividualNode)
}

func (node *WifeNode) Similarity(other *WifeNode, options SimilarityOptions) float64 {
	if node == nil || other == nil {
		return 0.5
	}

	return node.Individual().Similarity(other.Individual(), options)
}

func (node *WifeNode) IsIndividual(node2 *IndividualNode) bool {
	if node == nil || node2 == nil {
		return false
	}

	return node.Individual().Is(node2)
}
