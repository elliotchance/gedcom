package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/elliotchance/gedcom"
)

func RunApp() {
	app := app.New()

	w := app.NewWindow("gedcom")
	w.SetContent(LoadingScreen())

	go func() {
		gedcomFile, err := openFile()
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		document, err := gedcom.NewDocumentFromGEDCOMFile(gedcomFile)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		warnings := document.Warnings()

		tabs := widget.NewTabContainer(
			widget.NewTabItemWithIcon("Info", theme.InfoIcon(), InfoScreen(document, gedcomFile, warnings)),
			//widget.NewTabItemWithIcon("Individuals", theme.WarningIcon(), WarningsScreen(warnings)),
			widget.NewTabItemWithIcon("Warnings", theme.WarningIcon(), WarningsScreen(warnings)),
		)
		tabs.SetTabLocation(widget.TabLocationLeading)
		w.Resize(fyne.NewSize(1000, 600))
		w.SetContent(tabs)
	}()

	w.ShowAndRun()
}
