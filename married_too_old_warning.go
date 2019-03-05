package gedcom

import "fmt"

type MarriedTooOldWarning struct {
	SimpleWarning
	Family             *FamilyNode
	Spouse             *IndividualNode
	YearsOldAtMarriage float64
}

func NewMarriedTooOldWarning(family *FamilyNode, spouse *IndividualNode, yearsOldAtMarriage float64) *MarriedTooOldWarning {
	return &MarriedTooOldWarning{
		Family:             family,
		Spouse:             spouse,
		YearsOldAtMarriage: yearsOldAtMarriage,
	}
}

func (w *MarriedTooOldWarning) Name() string {
	return "MarriedTooOld"
}

func (w *MarriedTooOldWarning) String() string {
	partnerName := "husband"
	if w.Family.Wife().IsIndividual(w.Spouse) {
		partnerName = "wife"
	}

	return fmt.Sprintf("The %s %s married at %.0f years old.",
		partnerName, w.Spouse.String(), w.YearsOldAtMarriage)
}
