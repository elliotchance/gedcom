package gedcom

import "fmt"

type InverseSpousesWarning struct {
	SimpleWarning
	Family        *FamilyNode
	Husband, Wife *IndividualNode
}

func NewInverseSpousesWarning(family *FamilyNode, husband, wife *IndividualNode) *InverseSpousesWarning {
	return &InverseSpousesWarning{
		Family:  family,
		Husband: husband,
		Wife:    wife,
	}
}

func (w *InverseSpousesWarning) Name() string {
	return "InverseSpouses"
}

func (w *InverseSpousesWarning) String() string {
	husbandText := "husband"
	wifeText := "wife"

	if len(w.Family.Children()) > 0 {
		husbandText = "father"
		wifeText = "mother"
	}

	return fmt.Sprintf("%s (%s) is the %s and %s (%s) is the %s.",
		w.Husband.String(), w.Husband.Sex().String(), husbandText,
		w.Wife.String(), w.Wife.Sex().String(), wifeText)
}
