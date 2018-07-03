package gedcom

// FamilyNode represents a family.
type FamilyNode struct {
	*SimpleNode
}

func NewFamilyNode(pointer string, children []Node) *FamilyNode {
	return &FamilyNode{
		NewSimpleNode(TagFamily, "", pointer, children),
	}
}

func (node *FamilyNode) Husband(document *Document) *IndividualNode {
	return node.partner(document, TagHusband)
}

func (node *FamilyNode) Wife(document *Document) *IndividualNode {
	return node.partner(document, TagWife)
}

func (node *FamilyNode) partner(document *Document, tag Tag) *IndividualNode {
	tags := node.NodesWithTag(tag)
	if len(tags) == 0 {
		return nil
	}

	pointer := valueToPointer(tags[0].Value())
	individual := document.NodeByPointer(pointer)
	if individual == nil {
		return nil
	}

	return individual.(*IndividualNode)
}
