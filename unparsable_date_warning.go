package gedcom

import (
	"fmt"
)

type UnparsableDateWarning struct {
	Date    *DateNode
	Context WarningContext
}

func NewUnparsableDateWarning(date *DateNode) *UnparsableDateWarning {
	return &UnparsableDateWarning{
		Date: date,
	}
}

func (w *UnparsableDateWarning) Name() string {
	return "UnparsableDate"
}

func (w *UnparsableDateWarning) String() string {
	return fmt.Sprintf(`Unparsable date "%s"`, w.Date.Value())
}

func (w *UnparsableDateWarning) SetContext(context WarningContext) {
	w.Context = context
}

func (w *UnparsableDateWarning) MarshalQ() interface{} {
	return map[string]interface{}{
		"String":  w.String(),
		"Name":    w.Name(),
		"Context": w.Context.MarshalQ(),

		"Date": w.Date.Value(),
	}
}
