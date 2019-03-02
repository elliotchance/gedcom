package q

import (
	"fmt"
	"github.com/elliotchance/gedcom"
	"io"
	"reflect"
)

type GEDCOMFormatter struct {
	Writer io.Writer
}

func (f *GEDCOMFormatter) Write(result interface{}) error {
	// Nil should be treated as a blank document.
	if gedcom.IsNil(result) {
		return nil
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

	if x, ok := result.(gedcom.GEDCOMStringer); ok {
		f.Writer.Write([]byte(x.GEDCOMString(0)))

		return nil
	}

	return fmt.Errorf("%s does not implement gedcom.GEDCOMStringer", t.Type())
}
