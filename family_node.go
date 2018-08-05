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
	tags := NodesWithTag(node, tag)
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

// TODO: Needs tests
func (node *FamilyNode) Children(document *Document) IndividualNodes {
	children := IndividualNodes{}

	for _, n := range NodesWithTag(node, TagChild) {
		pointer := document.NodeByPointer(valueToPointer(n.Value()))
		child := pointer.(*IndividualNode)
		children = append(children, child)
	}

	return children
}

// TODO: Needs tests
func (node *FamilyNode) HasChild(document *Document, individual *IndividualNode) bool {
	for _, n := range NodesWithTag(node, TagChild) {
		if n.Value() == "@"+individual.Pointer()+"@" {
			return true
		}
	}

	return false
}
