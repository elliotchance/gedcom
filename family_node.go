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

// Similarity calculates the similarity between two families.
//
// The depth controls how many generations should be compared. A depth of 0 will
// only compare the husband/wife and not take into account any children. At the
// moment only a depth of 0 is supported. Any other depth will raise panic.
//
// doc1 and doc2 are used as the Documents for the current and other node
// respectively. If the two FamilyNodes come from the same Document you must
// specify the same Document for both values.
func (node *FamilyNode) Similarity(doc1, doc2 *Document, other *FamilyNode, depth int) float64 {
	if depth != 0 {
		panic("depth can only be 0")
	}

	// It does not matter if any of the partners are nil, Similarity will handle
	// these gracefully.
	husband := node.Husband(doc1).Similarity(other.Husband(doc2))
	wife := node.Wife(doc1).Similarity(other.Wife(doc2))

	return (husband + wife) / 2
}
