package gui

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/elliotchance/gedcom"
)

func WarningsScreen(warnings gedcom.Warnings) fyne.CanvasObject {
	renderedWarnings := widget.NewVBox()

	if len(warnings) == 0 {
		renderedWarnings.Append(widget.NewLabel("No warnings. Hooray!"))
	}

	for j := 0; j < 100; j++ {
		for i, warning := range warnings {
			description := fmt.Sprintf("%d. %s for %s",
				i+1, warning.String(), warning.Context())
			renderedWarnings.Append(widget.NewLabel(description))
		}
	}

	return widget.NewScrollContainer(renderedWarnings)
}
