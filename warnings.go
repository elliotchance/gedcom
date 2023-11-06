package gedcom

import (
	"io"
	"sort"

	"github.com/elliotchance/gedcom/v39/html/core"
)

type Warnings []Warning

func (ws Warnings) Strings() (ss []string) {
	for _, w := range ws {
		ss = append(ss, w.String())
	}

	return
}

func (ws Warnings) WriteHTMLTo(w io.Writer) (int64, error) {
	var data [][]string

	for _, warning := range ws {
		data = append(data, []string{
			warning.Name(),
			warning.Context().String(),
			warning.String(),
		})
	}

	sort.Slice(data, func(i, j int) bool {
		a := data[i][1]
		b := data[j][1]

		return a < b
	})

	rows := []core.Component{
		core.NewTableHead(
			"#",
			"Name",
			"Context",
			"Description",
		),
	}

	for i, row := range data {
		rows = append(rows, core.NewTableRow(
			core.NewTableCell(core.NewNumber(i+1)),
			core.NewTableCell(core.NewText(row[0])),
			core.NewTableCell(core.NewText(row[1])),
			core.NewTableCell(core.NewText(row[2])),
		))
	}

	return core.NewTable("", rows...).WriteHTMLTo(w)
}
