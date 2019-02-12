package gedcom

import "fmt"

// ChildNode is the natural, adopted, or sealed (LDS) child of a father and a
// mother.
type ChildNode struct {
	*SimpleNode
	family *FamilyNode
}

func newChildNode(family *FamilyNode, value string, children ...Node) *ChildNode {
	return &ChildNode{
		newSimpleNode(TagChild, value, "", children...),
		family,
	}
}

func newChildNodeWithIndividual(family *FamilyNode, individual *IndividualNode) *ChildNode {
	// TODO: check individual belongs to the same document as family
	return newChildNode(family, fmt.Sprintf("@%s@", individual.Pointer()))
}

func (node *ChildNode) Family() *FamilyNode {
	return node.family
}

func (node *ChildNode) Individual() *IndividualNode {
	n := node.family.document.NodeByPointer(valueToPointer(node.value))

	// TODO: may not exist
	return n.(*IndividualNode)
}

func (node *ChildNode) Father() *HusbandNode {
	return node.family.Husband()
}

func (node *ChildNode) Mother() *WifeNode {
	return node.family.Wife()
}

func (node *ChildNode) String() string {
	return node.Individual().String()
}
