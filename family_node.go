package gedcom

import (
	"fmt"
	"time"
)

// FamilyNode represents a family.
type FamilyNode struct {
	*simpleDocumentNode
	cachedHusband, cachedWife bool
	husband                   *HusbandNode
	wife                      *WifeNode
}

func newFamilyNode(document *Document, pointer string, children ...Node) *FamilyNode {
	return &FamilyNode{
		newSimpleDocumentNode(document, TagFamily, "", pointer, children...),
		false, false, nil, nil,
	}
}

// If the node is nil the result will also be nil.
func (node *FamilyNode) Husband() (husband *HusbandNode) {
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

	possibleHusband := First(NodesWithTag(node, TagHusband))

	if IsNil(possibleHusband) {
		return nil
	}

	return possibleHusband.(*HusbandNode)
}

// If the node is nil the result will also be nil.
func (node *FamilyNode) Wife() (wife *WifeNode) {
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

	possibleWife := First(NodesWithTag(node, TagWife))

	if IsNil(possibleWife) {
		return nil
	}

	return possibleWife.(*WifeNode)
}

// TODO: Needs tests
//
// If the node is nil the result will also be nil.
func (node *FamilyNode) Children() ChildNodes {
	if node == nil {
		return nil
	}

	children := ChildNodes{}

	for _, n := range NodesWithTag(node, TagChild) {
		children = append(children, n.(*ChildNode))
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
func (node *FamilyNode) Similarity(other *FamilyNode, depth int, options SimilarityOptions) float64 {
	if depth != 0 {
		panic("depth can only be 0")
	}

	// It does not matter if any of the partners are nil, Similarity will handle
	// these gracefully.
	husband := node.Husband().Similarity(other.Husband(), options)
	wife := node.Wife().Similarity(other.Wife(), options)

	return (husband + wife) / 2
}

func (node *FamilyNode) addChild(value string) *ChildNode {
	n := newChildNode(node, value)
	node.AddNode(n)

	return n
}

func (node *FamilyNode) AddChild(individual *IndividualNode) *ChildNode {
	n := newChildNodeWithIndividual(node, individual)
	node.AddNode(n)

	return n
}

func (node *FamilyNode) SetHusband(individual *IndividualNode) *FamilyNode {
	if individual == nil {
		DeleteNodesWithTag(node, TagHusband)

		return node
	}

	return node.SetHusbandPointer(individual.Pointer())
}

func (node *FamilyNode) SetWife(individual *IndividualNode) *FamilyNode {
	if individual == nil {
		DeleteNodesWithTag(node, TagWife)

		return node
	}

	return node.SetWifePointer(individual.Pointer())
}

func (node *FamilyNode) SetWifePointer(pointer string) *FamilyNode {
	wife := node.Wife()
	value := fmt.Sprintf("@%s@", pointer)
	if wife != nil {
		wife.value = value
	}

	node.AddNode(newNode(nil, node, TagWife, value, ""))
	node.cachedWife = false

	return node
}

func (node *FamilyNode) SetHusbandPointer(pointer string) *FamilyNode {
	husband := node.Husband()
	value := fmt.Sprintf("@%s@", pointer)
	if husband != nil {
		husband.value = value
	}

	husbandNode := newNode(nil, node, TagHusband, value, "")
	node.AddNode(husbandNode)
	node.cachedHusband = false

	return node
}

func (node *FamilyNode) resetCache() {
	node.cachedHusband = false
	node.cachedWife = false
	node.husband = nil
	node.wife = nil
}

func (node *FamilyNode) childrenBornBeforeParentsWarnings() (warnings Warnings) {
	fatherBirth, _ := node.Husband().Individual().Birth()
	motherBirth, _ := node.Wife().Individual().Birth()

	for _, child := range node.Children() {
		childBirth, _ := child.Individual().Birth()
		if !childBirth.IsValid() {
			continue
		}

		if fatherBirth.IsValid() && childBirth.IsBefore(fatherBirth) {
			warning := NewChildBornBeforeParentWarning(
				node.Husband().Individual(),
				child,
			)
			warnings = append(warnings, warning)
		}

		if motherBirth.IsValid() && childBirth.IsBefore(motherBirth) {
			warning := NewChildBornBeforeParentWarning(
				node.Wife().Individual(),
				child,
			)
			warnings = append(warnings, warning)
		}
	}

	return
}

func (node *FamilyNode) siblingsBornTooCloseWarnings() (warnings Warnings) {
	pairs := IndividualNodePairs{}

	for _, child1 := range node.Children() {
		child1Birth, _ := child1.Individual().Birth()
		for _, child2 := range node.Children() {
			// Exclude matching siblings to themselves. Technically we do not
			// need to do this check because children born on the same day would
			// be considered twins. However, its better to have it here for
			// completeness.
			if child1.Individual().Is(child2.Individual()) {
				continue
			}

			child2Birth, _ := child2.Individual().Birth()
			min, max, err := child1Birth.Sub(child2Birth)
			if err != nil {
				continue
			}

			nineMonths := Duration(274 * 24 * time.Hour)
			if min < nineMonths || max < nineMonths {
				pair := &IndividualNodePair{
					Left:  child1.Individual(),
					Right: child2.Individual(),
				}
				if !pairs.Has(pair) {
					warning := NewSiblingsBornTooCloseWarning(
						child1,
						child2,
					)
					warnings = append(warnings, warning)

					pairs = append(pairs, pair)
				}
			}
		}
	}

	return
}

func (node *FamilyNode) Warnings() (warnings Warnings) {
	warnings = append(warnings, node.childrenBornBeforeParentsWarnings()...)
	warnings = append(warnings, node.siblingsBornTooCloseWarnings()...)

	return
}
