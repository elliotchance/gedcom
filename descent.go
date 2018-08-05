package gedcom

// Descent describes the direct relationships around an individual. It is a
// useful structure to explain the parents, spouses and children connected to an
// individual.
type Descent struct {
	// Parents may contain zero or more elements for each set of parents.
	Parents []*FamilyNode

	// Individual that has this descent context.
	Individual *IndividualNode

	// SpouseChildren maps the known spouses to their children. The spouse will
	// be nil if the other parent is not known. Children can appear under
	// multiple spouses.
	SpouseChildren map[*IndividualNode]IndividualNodes
}
