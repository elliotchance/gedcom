package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/elliotchance/gedcom"
	"strconv"
)

func InfoScreen(document *gedcom.Document, filePath string, warnings gedcom.Warnings) fyne.CanvasObject {
	form := widget.NewForm()
	form.Append("File", widget.NewLabel(filePath))
	form.Append("Individuals", widget.NewLabel(strconv.Itoa(len(document.Individuals()))))
	form.Append("Families", widget.NewLabel(strconv.Itoa(len(document.Families()))))
	form.Append("Sources", widget.NewLabel(strconv.Itoa(len(document.Sources()))))
	form.Append("Places", widget.NewLabel(strconv.Itoa(len(document.Places()))))
	form.Append("Warnings", widget.NewLabel(strconv.Itoa(len(warnings))))

	return form
}
