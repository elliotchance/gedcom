package q

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/html/core"
	"io"
	"os"
	"reflect"
)

type HTMLFormatter struct {
	Writer io.Writer
}

func (f *HTMLFormatter) Write(result interface{}) error {
	pageTitle := "gedcom"

	// Nil should be treated as a blank document.
	if gedcom.IsNil(result) {
		_, err := core.NewPage(pageTitle, core.NewSpace(), "").
			WriteHTMLTo(os.Stdout)

		return err
	}

	if x, ok := result.(core.Component); ok {
		row := core.NewRow(core.NewColumn(core.EntireRow, x))
		_, err := core.NewPage(pageTitle, row, "").WriteHTMLTo(os.Stdout)

		return err
	}

	t := reflect.ValueOf(result)
	if t.Kind() == reflect.Slice {
		for i := 0; i < t.Len(); i++ {
			err := f.Write(t.Index(i).Interface())
			if err != nil {
				return err
			}
		}

		return nil
	}

	fallbackFormatter := &PrettyJSONFormatter{
		Writer: f.Writer,
	}

	f.Writer.Write([]byte("<pre>"))
	err := fallbackFormatter.Write(result)
	f.Writer.Write([]byte("\n</pre>"))

	return err
}
