package gedcom

// IndividualNode represents a person.
type IndividualNode struct {
	*SimpleNode
}

func NewIndividualNode(value, pointer string, children []Node) *IndividualNode {
	return &IndividualNode{
		NewSimpleNode(TagIndividual, value, pointer, children),
	}
}

// TODO: Needs tests
func (node *IndividualNode) Name() *NameNode {
	nameTag := First(NodesWithTag(node, TagName))
	if nameTag != nil {
		return nameTag.(*NameNode)
	}

	return nil
}

func (node *IndividualNode) Names() []*NameNode {
	nameTags := NodesWithTag(node, TagName)
	names := make([]*NameNode, len(nameTags))

	for i, name := range nameTags {
		names[i] = name.(*NameNode)
	}

	return names
}

func (node *IndividualNode) Sex() Sex {
	sex := NodesWithTag(node, TagSex)
	if len(sex) == 0 {
		return SexUnknown
	}

	return Sex(sex[0].Value())
}

// TODO: needs tests
func (node *IndividualNode) Spouses(doc *Document) []*IndividualNode {
	spouses := []*IndividualNode{}

	for _, family := range doc.Families() {
		husband := family.Husband(doc)
		wife := family.Wife(doc)

		// We only care about families that have both parties (otherwise there
		// is no spouse to add).
		if husband == nil || wife == nil {
			continue
		}

		if husband.Pointer() == node.Pointer() {
			spouses = append(spouses, wife)
		}

		if wife.Pointer() == node.Pointer() {
			spouses = append(spouses, husband)
		}
	}

	return spouses
}

// TODO: needs tests
func (node *IndividualNode) Families(doc *Document) []*FamilyNode {
	families := []*FamilyNode{}

	for _, family := range doc.Families() {
		if family.HasChild(doc, node) || family.Husband(doc).Is(node) || family.Wife(doc).Is(node) {
			families = append(families, family)
		}
	}

	return families
}

// TODO: needs tests
func (node *IndividualNode) Is(individual *IndividualNode) bool {
	return node != nil && individual != nil && node.Pointer() == individual.Pointer()
}

// TODO: needs tests
func (node *IndividualNode) FamilyWithSpouse(doc *Document, spouse *IndividualNode) *FamilyNode {
	for _, family := range doc.Families() {
		a := family.Husband(doc).Is(node) && family.Wife(doc).Is(spouse)
		b := family.Wife(doc).Is(node) && family.Husband(doc).Is(spouse)

		if a || b {
			return family
		}
	}

	return nil
}

// TODO: needs tests
func (node *IndividualNode) FamilyWithUnknownSpouse(doc *Document) *FamilyNode {
	for _, family := range doc.Families() {
		a := family.Husband(doc).Is(node) && family.Wife(doc) == nil
		b := family.Wife(doc).Is(node) && family.Husband(doc) == nil

		if a || b {
			return family
		}
	}

	return nil
}

// TODO: needs tests
func (node *IndividualNode) IsLiving() bool {
	return len(NodesWithTag(node, TagDeath)) == 0
}

// Births returns zero or more birth events for the individual.
func (node *IndividualNode) Births() []Node {
	return NodesWithTag(node, TagBirth)
}

// Baptisms returns zero or more baptism events for the individual. The baptisms
// do not include LDS baptisms.
func (node *IndividualNode) Baptisms() []Node {
	return NodesWithTag(node, TagBaptism)
}

// Deaths returns zero or more death events for the individual. It is common for
// individuals to not have a death event if the death date is not known. If you
// need to check if an individual is living you should use IsLiving().
func (node *IndividualNode) Deaths() []Node {
	return NodesWithTag(node, TagDeath)
}

// Burials returns zero or more burial events for the individual.
func (node *IndividualNode) Burials() []Node {
	return NodesWithTag(node, TagBurial)
}

// Descent collects the immediate relationships of an individual.
func (node *IndividualNode) Descent(doc *Document) *Descent {
	descent := &Descent{
		Parents:        []*FamilyNode{},
		Individual:     node,
		SpouseChildren: map[*IndividualNode][]*IndividualNode{},
	}

	for _, family := range node.Families(doc) {
		if family.HasChild(doc, node) {
			descent.Parents = append(descent.Parents, family)
		} else {
			var spouse *IndividualNode

			if family.Husband(doc).Is(node) {
				spouse = family.Wife(doc)
			} else {
				spouse = family.Husband(doc)
			}

			familyWithSpouse := node.FamilyWithSpouse(doc, spouse)
			var children []*IndividualNode
			if familyWithSpouse != nil {
				children = familyWithSpouse.Children(doc)
			}
			descent.SpouseChildren[spouse] = children

			// Find children with unknown spouse.
			unknownSpouseFamily := node.FamilyWithUnknownSpouse(doc)
			if unknownSpouseFamily != nil {
				descent.SpouseChildren[nil] = unknownSpouseFamily.Children(doc)
			}
		}
	}

	return descent
}
