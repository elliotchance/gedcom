package gedcom

import (
	"fmt"
	"github.com/elliotchance/gedcom/tag"
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
		newSimpleDocumentNode(document, tag.TagFamily, "", pointer, children...),
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

	possibleHusband := First(NodesWithTag(node, tag.TagHusband))

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

	possibleWife := First(NodesWithTag(node, tag.TagWife))

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

	for _, n := range NodesWithTag(node, tag.TagChild) {
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

	for _, n := range NodesWithTag(node, tag.TagChild) {
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
		DeleteNodesWithTag(node, tag.TagHusband)

		return node
	}

	return node.SetHusbandPointer(individual.Pointer())
}

func (node *FamilyNode) SetWife(individual *IndividualNode) *FamilyNode {
	if individual == nil {
		DeleteNodesWithTag(node, tag.TagWife)

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

	node.AddNode(newNode(nil, node, tag.TagWife, value, ""))
	node.cachedWife = false

	return node
}

func (node *FamilyNode) SetHusbandPointer(pointer string) *FamilyNode {
	husband := node.Husband()
	value := fmt.Sprintf("@%s@", pointer)
	if husband != nil {
		husband.value = value
	}

	husbandNode := newNode(nil, node, tag.TagHusband, value, "")
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
	nineMonths := time.Duration(274 * 24 * time.Hour)
	twoDays := time.Duration(2 * 24 * time.Hour)

	for _, child1 := range node.Children() {
		child1Birth, _ := child1.Individual().Birth()

		// If the date range is greater than 9 months we do not have enough
		// accuracy, so bail out.
		if child1Birth.DateRange().Duration().Duration >= nineMonths {
			continue
		}

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

			// If the date range is greater than 9 months we do not have enough
			// accuracy, so bail out.
			if child2Birth.DateRange().Duration().Duration >= nineMonths {
				continue
			}

			// Twins or greater multiples may be born in the same day. We allow
			// for two days to compensate for rounding. Also it's possible for
			// multiple children to be born on either side of the midnight
			// barrier.
			if min.Duration < twoDays {
				continue
			}

			if min.Duration < nineMonths || max.Duration < nineMonths {
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

func (node *FamilyNode) appendMarriedOutOfRange(mutWarnings Warnings, age Age, spouse *IndividualNode) Warnings {
	if age.IsKnown && age.Years() < DefaultMinMarriageAge {
		warning := NewMarriedOutOfRangeWarning(
			node,
			spouse,
			age.Years(),
			"young",
		)
		mutWarnings = append(mutWarnings, warning)
	}

	if age.Years() > DefaultMaxMarriageAge {
		warning := NewMarriedOutOfRangeWarning(
			node,
			spouse,
			age.Years(),
			"old",
		)
		mutWarnings = append(mutWarnings, warning)
	}

	return mutWarnings
}

func (node *FamilyNode) marriedOutOfRange() (warnings Warnings) {
	marriages := NodesWithTag(node, tag.TagMarriage)

	for _, marriage := range marriages {
		if husband := node.Husband().Individual(); husband != nil {
			_, maxAge := husband.AgeAt(marriage)
			warnings = node.appendMarriedOutOfRange(warnings, maxAge, husband)
		}

		if wife := node.Wife().Individual(); wife != nil {
			_, maxAge := wife.AgeAt(marriage)
			warnings = node.appendMarriedOutOfRange(warnings, maxAge, wife)
		}
	}

	return
}

func (node *FamilyNode) inversePartnerWarnings() (warnings Warnings) {
	husband := node.Husband().Individual()
	wife := node.Wife().Individual()

	// We only consider the case when both spouses exist, have sexes and they
	// are exactly opposites. We do not want to catch same sex partnerships, of
	// which GEDCOM has no reasonable way to encode this.
	switch {
	case husband.Sex().IsFemale() && wife.Sex().IsMale():
		warning := NewInverseSpousesWarning(node, husband, wife)
		warnings = append(warnings, warning)
	}

	return
}

func (node *FamilyNode) Warnings() (mutWarnings Warnings) {
	mutWarnings = append(mutWarnings, node.childrenBornBeforeParentsWarnings()...)
	mutWarnings = append(mutWarnings, node.siblingsBornTooCloseWarnings()...)
	mutWarnings = append(mutWarnings, node.marriedOutOfRange()...)
	mutWarnings = append(mutWarnings, node.inversePartnerWarnings()...)

	return
}

func (node *FamilyNode) String() string {
	return fmt.Sprintf("%s %s %s", node.Husband().String(),
		node.symbol(), node.Wife().String())
}

func (node *FamilyNode) symbol() string {
	switch {
	case len(NodesWithTag(node, tag.TagDivorce)) > 0:
		return "⚮"

	case len(NodesWithTag(node, tag.TagMarriage)) > 0:
		return "⚭"
	}

	return "-"
}
