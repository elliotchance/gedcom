package gedcom

import "fmt"

type SiblingsBornTooCloseWarning struct {
	Sibling1, Sibling2 *ChildNode
	Context            WarningContext
}

func NewSiblingsBornTooCloseWarning(sibling1, sibling2 *ChildNode) *SiblingsBornTooCloseWarning {
	return &SiblingsBornTooCloseWarning{
		Sibling1: sibling1,
		Sibling2: sibling2,
	}
}

func (w *SiblingsBornTooCloseWarning) Name() string {
	return "SiblingsBornTooClose"
}

func (w *SiblingsBornTooCloseWarning) String() string {
	birth1, _ := w.Sibling1.Individual().Birth()
	birth2, _ := w.Sibling2.Individual().Birth()
	min, _, _ := birth1.Sub(birth2)

	return fmt.Sprintf("The siblings %s and %s were born within %s of each other.",
		w.Sibling1, w.Sibling2, Duration(min).String())
}

func (w *SiblingsBornTooCloseWarning) SetContext(context WarningContext) {
	w.Context = context
}

