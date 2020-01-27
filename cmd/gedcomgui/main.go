package main

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"github.com/therecipe/qt/widgets"
	"os"
	"sort"
	"strings"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetCentralWidget(newLoading())
	window.SetWindowTitle("gedcom")
	window.Show()

	gedcomFile := widgets.QFileDialog_GetOpenFileName(
		nil, "Open GEDCOM file", "", "", "", 0)
	fmt.Println(gedcomFile)

	document, err := gedcom.NewDocumentFromGEDCOMFile(gedcomFile)
	if err != nil {
		panic(err)
	}

	window.SetCentralWidget(newWindow(document))
	window.Resize2(1000, 500)

	widgets.QApplication_Exec()
}

func newWindow(document *gedcom.Document) *widgets.QWidget {
	individuals := document.Individuals()
	sort.Slice(individuals, func(i, j int) bool {
		a := individuals[i].Name().Format(gedcom.NameFormatIndex)
		b := individuals[j].Name().Format(gedcom.NameFormatIndex)

		return strings.ToLower(a) < strings.ToLower(b)
	})

	families := document.Families()
	sources := document.Sources()
	warnings := document.Warnings()

	var places []*gedcom.PlaceNode
	for place := range document.Places() {
		places = append(places, place)
	}

	tabs := widgets.NewQTabWidget(nil)
	tabs.AddTab(newIndividualsTab(individuals),
		fmt.Sprintf("Individuals (%d)", len(individuals)))
	tabs.AddTab(newFamiliesTab(families),
		fmt.Sprintf("Families (%d)", len(families)))
	tabs.AddTab(newPlacesTab(places),
		fmt.Sprintf("Places (%d)", len(places)))
	tabs.AddTab(newSourcesTab(sources),
		fmt.Sprintf("Sources (%d)", len(sources)))
	tabs.AddTab(newWarningsTab(warnings),
		fmt.Sprintf("Warnings (%d)", len(warnings)))

	layout := widgets.NewQGridLayout2()
	layout.AddWidget2(tabs, 0, 0, 0)

	centralWidget := widgets.NewQWidget(nil, 0)
	centralWidget.SetLayout(layout)

	return centralWidget
}

func newFamiliesTab(families gedcom.FamilyNodes) *widgets.QWidget {
	columns := TableColumns{
		{"Husband", true},
		{"Married", true},
		{"Wife", true},
	}
	model := NewTableModel(columns, func(columns TableColumns) (rows [][]string) {
		for _, family := range families {
			var row []string
			for _, column := range columns {
				switch column.Name {
				case "Husband":
					row = append(row, gedcom.String(family.Husband()))
				case "Married":
					dates := gedcom.Dates(gedcom.NodesWithTag(family, gedcom.TagMarriage)...)
					if len(dates) > 0 {
						row = append(row, dates[0].String())
					} else {
						row = append(row, "")
					}
				case "Wife":
					row = append(row, gedcom.String(family.Wife()))
				}
			}

			rows = append(rows, row)
		}

		return
	})

	return model.Widget()
}

func newPlacesTab(places []*gedcom.PlaceNode) *widgets.QWidget {
	columns := TableColumns{
		{"Place", true},
	}
	model := NewTableModel(columns, func(columns TableColumns) (rows [][]string) {
		for _, place := range places {
			var row []string
			for range columns {
				row = append(row, place.String())
			}

			rows = append(rows, row)
		}

		return
	})

	return model.Widget()
}

func newSourcesTab(sources []*gedcom.SourceNode) *widgets.QWidget {
	columns := TableColumns{
		{"Title", true},
	}
	model := NewTableModel(columns, func(columns TableColumns) (rows [][]string) {
		for _, source := range sources {
			var row []string
			for range columns {
				row = append(row, source.Title())
			}

			rows = append(rows, row)
		}

		return
	})

	return model.Widget()
}

func newWarningsTab(warnings gedcom.Warnings) *widgets.QWidget {
	columns := TableColumns{
		{"Description", true},
	}
	model := NewTableModel(columns, func(columns TableColumns) (rows [][]string) {
		for _, warning := range warnings {
			var row []string
			for range columns {
				row = append(row, fmt.Sprintf("%s for %s",
					warning.String(), warning.Context()))
			}

			rows = append(rows, row)
		}

		return
	})

	return model.Widget()
}

func newLoading() *widgets.QWidget {
	label := widgets.NewQLabel2("Loading", nil, 0)

	layout := widgets.NewQGridLayout2()
	layout.AddWidget2(label, 0, 0, 0)

	centralWidget := widgets.NewQWidget(nil, 0)
	centralWidget.SetLayout(layout)

	return centralWidget
}
