package gedcom

// SurroundingSimilarity describes different aspects of the similarity of an
// individual by its immediate relationships; parents, spouses and children.
type SurroundingSimilarity struct {
	// ParentsSimilarity is the similarity of the fathers and mothers of the
	// individual. Each missing parent will be given 0.5. If both parents are
	// missing the parent similarity will also be 0.5.
	//
	// An individual can have zero or more pairs of parents, but only a single
	// ParentsSimilarity is provided. The ParentsSimilarity is the highest value
	// when each of the parents are compared with the other parents of the other
	// individual.
	ParentsSimilarity float64

	// IndividualSimilarity is the same as Individual.Similarity().
	IndividualSimilarity float64

	// SpousesSimilarity is the similarity of the spouses is compared with
	// IndividualNodes.Similarity() which is designed to compare several
	// individuals at once. It also handles comparing a different number of
	// individuals on either side.
	SpousesSimilarity float64

	// ChildrenSimilarity also uses IndividualNodes.Similarity() but without
	// respect to their parents (which in this case would be the current
	// individual and likely one of their spouses).
	//
	// It is done this way as to not skew the results if any particular parent
	// is unknown or the child is connected to a different spouse.
	ChildrenSimilarity float64
}
