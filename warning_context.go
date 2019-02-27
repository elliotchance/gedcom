package gedcom

// A WarningContext describes the entity the warning is attached to.
//
// For example if the warning was an unparsable date for the birth of an
// individual then the Individual would be set, but not the Family.
type WarningContext struct {
	Individual *IndividualNode
	Family     *FamilyNode
}

func (context WarningContext) MarshalQ() interface{} {
	m := map[string]string{}

	if context.Individual != nil {
		m["Individual"] = context.Individual.Pointer()
	}

	if context.Family != nil {
		m["Family"] = context.Family.Pointer()
	}

	return m
}
