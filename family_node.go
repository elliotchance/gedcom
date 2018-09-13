package gedcom

// FamilyNode represents a family.
type FamilyNode struct {
	*SimpleNode
	cachedHusband, cachedWife bool
	husband, wife             *IndividualNode
}

func NewFamilyNode(document *Document, pointer string, children []Node) *FamilyNode {
	return &FamilyNode{
		NewSimpleNode(document, TagFamily, "", pointer, children),
		false, false, nil, nil,
	}
}

// If the node is nil the result will also be nil.
func (node *FamilyNode) Husband() (husband *IndividualNode) {
	if node == nil {
		return nil
	}

	if node.cachedHusband {
		return node.husband
	}

	defer func() {
		node.husband = husband
		node.cachedHusband = true
	}()

	return node.partner(TagHusband)
}

// If the node is nil the result will also be nil.
func (node *FamilyNode) Wife() (wife *IndividualNode) {
	if node == nil {
		return nil
	}

	if node.cachedWife {
		return node.wife
	}

	defer func() {
		node.wife = wife
		node.cachedWife = true
	}()

	return node.partner(TagWife)
}

func (node *FamilyNode) partner(tag Tag) *IndividualNode {
	tags := NodesWithTag(node, tag)
	if len(tags) == 0 {
		return nil
	}

	pointer := valueToPointer(tags[0].Value())
	individual := node.document.NodeByPointer(pointer)
	if individual == nil {
		return nil
	}

	return individual.(*IndividualNode)
}

// TODO: Needs tests
//
// If the node is nil the result will also be nil.
func (node *FamilyNode) Children() IndividualNodes {
	if node == nil {
		return nil
	}

	children := IndividualNodes{}

	for _, n := range NodesWithTag(node, TagChild) {
		pointer := node.document.NodeByPointer(valueToPointer(n.Value()))
		child := pointer.(*IndividualNode)
		children = append(children, child)
	}

	return children
}

// TODO: Needs tests
//
// If the node is nil the result will also be nil.
func (node *FamilyNode) HasChild(individual *IndividualNode) bool {
	if node == nil {
		return false
	}

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
// The options.MaxYears allows the error margin on dates to be adjusted. See
// DefaultMaxYearsForSimilarity for more information.
func (node *FamilyNode) Similarity(other *FamilyNode, depth int, options *SimilarityOptions) float64 {
	if depth != 0 {
		panic("depth can only be 0")
	}

	// It does not matter if any of the partners are nil, Similarity will handle
	// these gracefully.
	husband := node.Husband().Similarity(other.Husband(), options)
	wife := node.Wife().Similarity(other.Wife(), options)

	return (husband + wife) / 2
}
