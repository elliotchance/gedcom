package gedcom

import "fmt"

// MarriedOutOfRangeWarning is raised when the age of an individual at the time
// they were married is less than 16 (DefaultMinMarriageAge) or greater than 100
// (DefaultMaxMarriageAge).
type MarriedOutOfRangeWarning struct {
	SimpleWarning
	Family             *FamilyNode
	Spouse             *IndividualNode
	YearsOldAtMarriage float64

	// Either "young" or "old" to indicate which boundary was breached.
	Boundary string
}

func NewMarriedOutOfRangeWarning(family *FamilyNode, spouse *IndividualNode, yearsOldAtMarriage float64, boundary string) *MarriedOutOfRangeWarning {
	return &MarriedOutOfRangeWarning{
		Family:             family,
		Spouse:             spouse,
		YearsOldAtMarriage: yearsOldAtMarriage,
		Boundary:           boundary,
	}
}

func (w *MarriedOutOfRangeWarning) Name() string {
	return "MarriedOutOfRange"
}

func (w *MarriedOutOfRangeWarning) String() string {
	partnerName := "husband"
	if w.Family.Wife().IsIndividual(w.Spouse) {
		partnerName = "wife"
	}

	return fmt.Sprintf("The %s %s married too %s at %.0f years old.",
		partnerName, w.Spouse.String(), w.Boundary, w.YearsOldAtMarriage)
}
