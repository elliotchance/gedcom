package gedcom

// IndividualNode represents a person.
type IndividualNode struct {
	*SimpleNode
}

// SpouseChildren connects a single spouse to a set of children. The children
// may appear under multiple spouses. This is only useful when in combination
// with an individual (that would be the other spouse).
//
// The spouse can be nil, indicating that the spouse it not known for the
// assigned children. You should not assume that you can also recover the other
// spouse from one of the keys in this map as the map is valid to be empty or to
// only contain a nil key.
type SpouseChildren map[*IndividualNode]IndividualNodes

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
func (node *IndividualNode) Spouses(doc *Document) IndividualNodes {
	spouses := IndividualNodes{}

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

// Parents returns the families for which this individual is a child. There may
// be zero or more parents for an individual. The families returned will all
// reference this individual as child. However the father, mother or both may
// not exist.
//
// It is also possible to have duplicate families. That is, families that have
// the same husband and wife combinations if these families are defined in the
// GEDCOM file.
func (node *IndividualNode) Parents(doc *Document) []*FamilyNode {
	parents := []*FamilyNode{}

	for _, family := range node.Families(doc) {
		if family.HasChild(doc, node) {
			parents = append(parents, family)
		}
	}

	return parents
}

// SpouseChildren maps the known spouses to their children. The spouse will be
// nil if the other parent is not known for some or all of the children.
// Children can appear under multiple spouses.
func (node *IndividualNode) SpouseChildren(doc *Document) SpouseChildren {
	spouseChildren := map[*IndividualNode]IndividualNodes{}

	for _, family := range node.Families(doc) {
		if !family.HasChild(doc, node) {
			var spouse *IndividualNode

			if family.Husband(doc).Is(node) {
				spouse = family.Wife(doc)
			} else {
				spouse = family.Husband(doc)
			}

			familyWithSpouse := node.FamilyWithSpouse(doc, spouse)
			var children IndividualNodes
			if familyWithSpouse != nil {
				children = familyWithSpouse.Children(doc)
			}
			spouseChildren[spouse] = children

			// Find children with unknown spouse.
			unknownSpouseFamily := node.FamilyWithUnknownSpouse(doc)
			if unknownSpouseFamily != nil {
				spouseChildren[nil] = unknownSpouseFamily.Children(doc)
			}
		}
	}

	return spouseChildren
}

// LDSBaptisms returns zero or more LDS baptism events for the individual. These
// are not to be confused with Baptisms().
func (node *IndividualNode) LDSBaptisms() []Node {
	return NodesWithTag(node, TagLDSBaptism)
}

// EstimatedBirthDate attempts to find the exact or approximate birth date of an
// individual. It does this by looking at the births, baptisms and LDS baptisms.
// If any of them contain a date then the lowest date value is returned based on
// the Years() value which takes in account the full date range.
//
// This logic is loosely based off the idea that if the birth date is not known
// that a baptism usually happens when the individual is quite young (and
// therefore close to the their birth date).
//
// It is worth noting that since EstimatedBirthDate returns the lowest possible
// date that an exact birth date will be ignored if another event happens in a
// range before that. For example, if an individual has a birth date of
// "9 Feb 1983" but the Baptism was "9 Jan 1983" then the Baptism is returned.
// This data must be wrong in either case but EstimatedBirthDate cannot make a
// sensible decision in this case so it always returned the earliest date.
//
// EstimatedBirthDate is useful when comparing individuals where the exact dates
// are less important that attempting to serve approximate information for
// comparison. You almost certainly do not want to use the EstimatedBirthDate
// value for anything meaningful aside from comparisons.
func (node *IndividualNode) EstimatedBirthDate() *DateNode {
	potentialNodes :=
		Compound(node.Births(), node.Baptisms(), node.LDSBaptisms())

	bestMatch := (*DateNode)(nil)

	for _, potentialNode := range potentialNodes {
		for _, potentialDateNode := range NodesWithTag(potentialNode, TagDate) {
			node := potentialDateNode.(*DateNode)
			if bestMatch == nil || node.Years() < bestMatch.Years() {
				bestMatch = node
			}
		}
	}

	return bestMatch
}

// EstimatedDeathDate attempts to find the exact or approximate death date of an
// individual. It does this by returning the earliest death date based on the
// value of Years(). If there are no death dates then it will attempt to return
// the minimum burial date.
//
// This logic is loosely based off the idea that if the death date is not known
// that a burial usually happens a short time after the death of the individual.
//
// It is worth noting that EstimatedDeathDate will always return a death date if
// one is present before falling back to a possibly more specific burial date.
// One example of this might be a death date that has a large range such as
// "1983 - 1993". The burial may be a much more specific date like "Apr 1985".
// This almost certainly indicates that the death date was around early 1985,
// however the larger death date range will still be returned.
//
// EstimatedDeathDate is useful when comparing individuals where the exact dates
// are less important that attempting to serve approximate information for
// comparison. You almost certainly do not want to use the EstimatedDeathDate
// value for anything meaningful aside from comparisons.
func (node *IndividualNode) EstimatedDeathDate() *DateNode {
	// Try to return the earliest the death date first.
	bestMatch := (*DateNode)(nil)

	for _, potentialNode := range node.Deaths() {
		for _, potentialDateNode := range NodesWithTag(potentialNode, TagDate) {
			node := potentialDateNode.(*DateNode)
			if bestMatch == nil || node.Years() < bestMatch.Years() {
				bestMatch = node
			}
		}
	}

	if bestMatch != nil {
		return bestMatch
	}

	// Fall back to the earliest burial date.
	for _, potentialNode := range node.Burials() {
		for _, potentialDateNode := range NodesWithTag(potentialNode, TagDate) {
			node := potentialDateNode.(*DateNode)
			if bestMatch == nil || node.Years() < bestMatch.Years() {
				bestMatch = node
			}
		}
	}

	// bestMatch will be nil if there were no date nodes found.
	return bestMatch
}

// Similarity calculates how similar two individuals are. The returned value
// will be between 0.0 and 1.0 where 1.0 means an exact match.
//
// The similarity is based off three equally weighted components, the
// individuals name, estimated birth and estimated death date and is calculated
// as follows:
//
//   similarity = (nameSimilarity + birthSimilarity + deathSimilarity) / 3
//
// Individual names are compared with the StringSimilarity function that does
// not consider the punctuation and extra spacing.
//
// An individual may have more than one name, if this is the case then each name
// is checked and the highest matching combination is used.
//
// The birth and death dates use the EstimatedBirthDate and EstimatedDeathDate
// functions respectively. These functions are allowed to make some estimates
// when critical information like the birth date does not exist so there is more
// data to include in the comparison.
//
// Both dates are compared with the DateNode.Similarity function, which also
// returns a value of 0.0 to 1.0. To put simply the dates must existing within
// an error margin (for example, 10 years in either direction). Higher scores
// are awarded to dates that are more relatively closer to each other on a
// parabola. See DateNode.Similarity for a full explanation of how it deals with
// approximate dates and date ranges.
func (node *IndividualNode) Similarity(other *IndividualNode) float64 {
	// Compare the matrix of names.
	nameSimilarity := 0.0

	for _, name1 := range node.Names() {
		for _, name2 := range other.Names() {
			similarity := StringSimilarity(name1.String(), name2.String())

			if similarity > nameSimilarity {
				nameSimilarity = similarity
			}
		}
	}

	// Compare the dates.
	birthSimilarity := node.EstimatedBirthDate().
		Similarity(other.EstimatedBirthDate(), DefaultMaxYearsForSimilarity)

	deathSimilarity := node.EstimatedDeathDate().
		Similarity(other.EstimatedDeathDate(), DefaultMaxYearsForSimilarity)

	// Final calculation.
	return (nameSimilarity + birthSimilarity + deathSimilarity) / 3.0
}
