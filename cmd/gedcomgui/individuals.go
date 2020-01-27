package main

import (
	"github.com/elliotchance/gedcom"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"strconv"
	"strings"
	"time"
)

type IndividualFilter struct {
	filter, birth, death string
}

func filterDate(expected string, date *gedcom.DateNode, individual *gedcom.IndividualNode) bool {
	if expected != "" {
		if date == nil {
			return true
		}

		dateParts := strings.Split(expected, "-")
		expectedDate := expected
		if len(dateParts) == 2 {
			if dateParts[0] == "" {
				dateParts[0] = "0"
			}

			if dateParts[1] == "" {
				dateParts[1] = strconv.Itoa(time.Now().Year())
			}

			expectedDate = "Between " + dateParts[0] + " and " + dateParts[1]
		}

		dateRange := gedcom.NewDateRangeWithString(expectedDate)
		if !date.DateRange().Compare(dateRange).IsPartiallyEqual() {
			return true
		}
	}

	return false
}

func newIndividualsTab(individuals gedcom.IndividualNodes) *widgets.QWidget {
	columns := TableColumns{
		{"Name", true},
		{"Birth Date", true},
		{"Birth Place", false},
		{"Death Date", true},
		{"Death Place", false},
	}

	filter := &IndividualFilter{}

	model := NewTableModel(columns, func(columns TableColumns) (rows [][]string) {
		for _, individual := range individuals {
			name := strings.ToLower(gedcom.String(individual.Name()))
			if filter.filter != "" && !strings.Contains(name, filter.filter) {
				continue
			}

			birthDate, _ := individual.EstimatedBirthDate()
			if filterDate(filter.birth, birthDate, individual) {
				continue
			}

			deathDate, _ := individual.EstimatedDeathDate()
			if filterDate(filter.death, deathDate, individual) {
				continue
			}

			var row []string
			for _, column := range columns {
				switch column.Name {
				case "Name":
					name := individual.Name()
					if name != nil {
						row = append(row, name.Format(gedcom.NameFormatIndex))
					} else {
						row = append(row, "")
					}
				case "Birth Date":
					date, _ := individual.Birth()
					row = append(row, gedcom.String(date))
				case "Birth Place":
					_, place := individual.Birth()
					row = append(row, gedcom.String(place))
				case "Death Date":
					date, _ := individual.Death()
					row = append(row, gedcom.String(date))
				case "Death Place":
					_, place := individual.Death()
					row = append(row, gedcom.String(place))
				}
			}

			rows = append(rows, row)
		}

		return
	})

	tableView := model.TableView()

	searchField := widgets.NewQLineEdit(nil)
	searchField.SetPlaceholderText("Filter name")
	searchField.ConnectTextChanged(func(text string) {
		filter.filter = strings.ToLower(text)
		tableView.SetModel(model.Model())
	})

	birthDate := widgets.NewQLineEdit(nil)
	birthDate.SetPlaceholderText("Birth")
	birthDate.SetFixedWidth(100)
	birthDate.ConnectTextChanged(func(text string) {
		filter.birth = text
		tableView.SetModel(model.Model())
	})

	deathDate := widgets.NewQLineEdit(nil)
	deathDate.SetPlaceholderText("Death")
	deathDate.SetFixedWidth(100)
	deathDate.ConnectTextChanged(func(text string) {
		filter.death = text
		tableView.SetModel(model.Model())
	})

	optionsWindow := newIndividualViewOptions(columns, func() {
		tableView.SetModel(model.Model())
	})

	viewOptionsButton := widgets.NewQPushButton2("View Options", nil)
	viewOptionsButton.ConnectClicked(func(_ bool) {
		optionsWindow.Show()
	})

	layout := widgets.NewQGridLayout(nil)
	layout.AddWidget2(searchField, 0, 0, 0)
	layout.AddWidget2(birthDate, 0, 1, 0)
	layout.AddWidget2(deathDate, 0, 2, 0)
	layout.AddWidget2(viewOptionsButton, 0, 3, 0)
	layout.AddWidget3(tableView, 1, 0, 1, 4, 0)

	centralWidget := widgets.NewQWidget(nil, 0)
	centralWidget.SetLayout(layout)

	return centralWidget
}

func newIndividualViewOptions(columns TableColumns, didChange func()) *widgets.QMainWindow {
	model := gui.NewQStandardItemModel2(len(columns.AllNames()), 1, nil)

	for index, field := range columns {
		item := gui.NewQStandardItem2(field.Name)
		item.SetCheckable(true)
		if field.Show {
			item.SetCheckState(core.Qt__Checked)
		} else {
			item.SetCheckState(core.Qt__Unchecked)
		}
		model.SetItem(index, 0, item)
	}

	listView := widgets.NewQListView(nil)
	listView.SetModel(model)
	listView.ConnectClicked(func(index *core.QModelIndex) {
		columns[index.Row()].Show = !columns[index.Row()].Show
		didChange()
	})

	layout := widgets.NewQGridLayout(nil)
	layout.AddWidget2(listView, 0, 0, 0)

	centralWidget := widgets.NewQWidget(nil, 0)
	centralWidget.SetLayout(layout)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetCentralWidget(listView)
	window.SetWindowTitle("View Options")

	return window
}
