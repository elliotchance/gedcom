package gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func LoadingScreen() fyne.CanvasObject {
	return widget.NewVBox(
		widget.NewLabel("Loading..."),
		widget.NewProgressBarInfinite(),
	)
}
