package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type TableModel struct {
	columns TableColumns
	getData func(columns TableColumns) [][]string
}

func NewTableModel(columns TableColumns, getData func(columns TableColumns) [][]string) *TableModel {
	return &TableModel{
		columns: columns,
		getData: getData,
	}
}

func (m *TableModel) Widget() *widgets.QWidget {
	layout := widgets.NewQGridLayout2()
	layout.AddWidget2(m.TableView(), 0, 0, 0)

	centralWidget := widgets.NewQWidget(nil, 0)
	centralWidget.SetLayout(layout)

	return centralWidget
}

func (m *TableModel) TableView() *widgets.QTableView {
	table := widgets.NewQTableView(nil)
	table.SetModel(m.Model())
	table.HorizontalHeader().SetSectionResizeMode(widgets.QHeaderView__ResizeToContents)

	return table
}

func (m *TableModel) Model() *core.QAbstractTableModel {
	model := core.NewQAbstractTableModel(nil)

	visibleColumns := m.columns.VisibleColumns()
	data := m.getData(visibleColumns)

	model.ConnectHeaderData(func(section int, orientation core.Qt__Orientation, role int) *core.QVariant {
		if role != int(core.Qt__DisplayRole) || orientation == core.Qt__Vertical {
			return model.HeaderDataDefault(section, orientation, role)
		}

		return core.NewQVariant1(m.columns.VisibleNames()[section])
	})

	model.ConnectRowCount(func(parent *core.QModelIndex) int {
		return len(data)
	})

	model.ConnectColumnCount(func(parent *core.QModelIndex) int {
		return len(visibleColumns)
	})

	model.ConnectData(func(index *core.QModelIndex, role int) *core.QVariant {
		if role != int(core.Qt__DisplayRole) {
			return core.NewQVariant()
		}

		return core.NewQVariant1(data[index.Row()][index.Column()])
	})

	return model
}

type TableColumn struct {
	Name string
	Show bool
}

type TableColumns []*TableColumn

func (fields TableColumns) AllNames() (names []string) {
	for _, field := range fields {
		names = append(names, field.Name)
	}

	return
}

func (fields TableColumns) VisibleNames() (names []string) {
	for _, field := range fields {
		if field.Show {
			names = append(names, field.Name)
		}
	}

	return
}

func (fields TableColumns) VisibleColumns() (f TableColumns) {
	for _, field := range fields {
		if field.Show {
			f = append(f, field)
		}
	}

	return
}
