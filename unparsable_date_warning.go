package gedcom

import "fmt"

type UnparsableDateWarning struct {
	Date *DateNode
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
