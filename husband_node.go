package gedcom

import "fmt"

// HusbandNode is an individual in the family role of a married man or father.
type HusbandNode struct {
	*SimpleNode
	family *FamilyNode
}

func newHusbandNode(family *FamilyNode, value string, children ...Node) *HusbandNode {
	return &HusbandNode{
		newSimpleNode(TagHusband, value, "", children...),
		family,
	}
}

func newHusbandNodeWithIndividual(family *FamilyNode, individual *IndividualNode) *HusbandNode {
	// TODO: check individual belongs to the same document as family
	return newHusbandNode(family, fmt.Sprintf("@%s@", individual.Pointer()))
}

func (node *HusbandNode) Family() *FamilyNode {
	return node.family
}

func (node *HusbandNode) Individual() *IndividualNode {
	if node == nil {
		return nil
	}

	n := node.family.document.NodeByPointer(valueToPointer(node.value))

	if IsNil(n) {
		return nil
	}

	return n.(*IndividualNode)
}

func (node *HusbandNode) Similarity(other *HusbandNode, options SimilarityOptions) float64 {
	if node == nil || other == nil {
		return 0.5
	}

	return node.Individual().Similarity(other.Individual(), options)
}

func (node *HusbandNode) IsIndividual(node2 *IndividualNode) bool {
	if node == nil || node2 == nil {
		return false
	}

	return node.Individual().Is(node2)
}
