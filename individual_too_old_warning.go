package gedcom

import "fmt"

type IndividualTooOldWarning struct {
	SimpleWarning
	Individual      *IndividualNode
	YearsOldAtDeath float64
}

func NewIndividualTooOldWarning(individual *IndividualNode, yearsOldAtDeath float64) *IndividualTooOldWarning {
	return &IndividualTooOldWarning{
		Individual:      individual,
		YearsOldAtDeath: yearsOldAtDeath,
	}
}

func (w *IndividualTooOldWarning) Name() string {
	return "IndividualTooOld"
}

func (w *IndividualTooOldWarning) String() string {
	return fmt.Sprintf("%s was %.0f years old at the time of %s death.",
		w.Individual.String(), w.YearsOldAtDeath,
		w.Individual.Sex().OwnershipWord())
}
