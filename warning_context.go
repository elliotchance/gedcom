package gedcom

// A WarningContext describes the entity the warning is attached to.
//
// For example if the warning was an unparsable date for the birth of an
// individual then the Individual would be set, but not the Family.
type WarningContext struct {
	Individual *IndividualNode
	Family     *FamilyNode
}

func (context WarningContext) String() string {
	switch {
	case context.Individual != nil:
		return context.Individual.String()

	case context.Family != nil:
		return context.Family.String()

	default:
		return "None"
	}
}
