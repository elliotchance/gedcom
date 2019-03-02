package gedcom

import (
	"fmt"
	"strings"
)

// IncorrectEventOrderWarning produces warnings where events for an individual
// are in an incorrect order. For example a burial before a death event, or a
// baptism before a birth.
type IncorrectEventOrderWarning struct {
	SimpleWarning
	FirstEvent, SecondEvent         Node
	FirstDateRange, SecondDateRange DateRange
}

func NewIncorrectEventOrderWarning(firstEvent Node, firstDateRange DateRange, secondEvent Node, secondDateRange DateRange) *IncorrectEventOrderWarning {
	return &IncorrectEventOrderWarning{
		FirstEvent:      firstEvent,
		FirstDateRange:  firstDateRange,
		SecondEvent:     secondEvent,
		SecondDateRange: secondDateRange,
	}
}

func (w *IncorrectEventOrderWarning) Name() string {
	return "IncorrectEventOrder"
}

func (w *IncorrectEventOrderWarning) String() string {
	return fmt.Sprintf(`The %s (%s) was before the %s (%s) of %s.`,
		strings.ToLower(w.FirstEvent.Tag().String()), w.FirstDateRange,
		strings.ToLower(w.SecondEvent.Tag().String()), w.SecondDateRange,
		w.Context.Individual)
}

func (w *IncorrectEventOrderWarning) MarshalQ() interface{} {
	return map[string]interface{}{
		"String":  w.String(),
		"Name":    w.Name(),
		"Context": w.Context.MarshalQ(),

		"FirstEvent":  w.FirstEvent,
		"SecondEvent": w.SecondEvent,
	}
}
