package gedcom

import "fmt"

type ChildBornBeforeParentWarning struct {
	SimpleWarning
	Parent *IndividualNode
	Child  *ChildNode
}

func NewChildBornBeforeParentWarning(parent *IndividualNode, child *ChildNode) *ChildBornBeforeParentWarning {
	return &ChildBornBeforeParentWarning{
		Parent: parent,
		Child:  child,
	}
}

func (w *ChildBornBeforeParentWarning) Name() string {
	return "ChildBornBeforeParent"
}

func (w *ChildBornBeforeParentWarning) String() string {
	relationship := w.relationshipName()

	return fmt.Sprintf("The child %s was born before %s %s %s.",
		w.Child, w.Child.Individual().Sex().OwnershipWord(), relationship, w.Parent)
}

func (w *ChildBornBeforeParentWarning) relationshipName() string {
	if w.Child.Father().IsIndividual(w.Parent) {
		return "father"
	}

	if w.Child.Mother().IsIndividual(w.Parent) {
		return "mother"
	}

	return "parent"
}
