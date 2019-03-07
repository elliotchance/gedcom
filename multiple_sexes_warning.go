package gedcom

import (
	"fmt"
	"strings"
)

type MultipleSexesWarning struct {
	SimpleWarning
	Individual *IndividualNode

	// Sexes should contain at least two elements. The elements may have the
	// same values.
	Sexes []*SexNode
}

func NewMultipleSexesWarning(individual *IndividualNode, sexes []*SexNode) *MultipleSexesWarning {
	return &MultipleSexesWarning{
		Individual: individual,
		Sexes:      sexes,
	}
}

func (w *MultipleSexesWarning) Name() string {
	return "MultipleSexes"
}

func (w *MultipleSexesWarning) String() string {
	var sexes []string

	for _, sex := range w.Sexes {
		sexes = append(sexes, sex.String())
	}

	return fmt.Sprintf("%s has multiple sexes; %s and %s.",
		w.Individual.String(), strings.Join(sexes[:len(sexes)-1], ", "),
		sexes[len(sexes)-1])
}
