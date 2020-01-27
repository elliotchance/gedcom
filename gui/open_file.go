package gui

import (
	"github.com/sqweek/dialog"
)

func openFile() (string, error) {
	return dialog.File().
		//Filter("Mp3 audio file", "mp3").
		Load()
}
