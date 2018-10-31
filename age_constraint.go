package gedcom

// AgeConstraint is used to describe if the individual was living during the
// calculated age.
//
// See the AgeConstraint constants for more information.
type AgeConstraint int

const (
	// The constraint is not known. This is the case if none of the other
	// options can be determined to be true.
	AgeConstraintUnknown AgeConstraint = iota

	// The age represents a time before the known birth. This would be
	// represented as a negative age.
	AgeConstraintBeforeBirth

	// The age represents the number of years since the known birth of the
	// individual. Even if the age is an approximation or a wide range it can
	// still be considered as fully or partly within their lifetime.
	AgeConstraintLiving

	// The age is after the known death of the individual. Like the "before
	// birth" constraint it can be sometimes useful to know the age of the
	// individual if they were still living or subtract their death age to see
	// how many years after their death the event may have occurred.
	//
	// You can use IndividualNode.Age() to fetch the maximum living age of a now
	// deceased individual.
	AgeConstraintAfterDeath
)

// Strings returns a human-readable form of the constant, like "After Death".
func (ac AgeConstraint) String() string {
	switch ac {
	case AgeConstraintUnknown:
		// Do nothing, fall through to the final return.

	case AgeConstraintBeforeBirth:
		return "Before Birth"

	case AgeConstraintLiving:
		return "Living"

	case AgeConstraintAfterDeath:
		return "After Death"
	}

	return "Unknown"
}
