package q

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/elliotchance/gedcom"
	"io"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

// Formatter is used to write the result to stream.
type Formatter interface {
	Write(result interface{}) error
}

type JSONFormatter struct {
	Writer io.Writer
}

func (f *JSONFormatter) Write(result interface{}) error {
	data, err := json.Marshal(result)
	if err != nil {
		return err
	}

	_, err = f.Writer.Write(append(data, '\n'))

	return err
}

type PrettyJSONFormatter struct {
	Writer io.Writer
}

func (f *PrettyJSONFormatter) Write(result interface{}) error {
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}

	_, err = f.Writer.Write(append(data, '\n'))

	return err
}

type CSVFormatter struct {
	Writer io.Writer
}

func (f *CSVFormatter) Write(result interface{}) error {
	columns, err := f.Header(result)
	if err != nil {
		return err
	}

	f.writeLine(columns)

	v := reflect.ValueOf(result)
	if v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			line := []string{}
			value := f.prepareLine(v.Index(i).Interface())

			for _, name := range columns {
				line = append(line, fmt.Sprintf("%v", value[name]))
			}

			f.writeLine(line)
		}
	}

	return nil
}

func (f *CSVFormatter) writeLine(fields []string) {
	for i := 0; i < len(fields); i++ {
		if i > 0 {
			fmt.Fprintf(f.Writer, ",")
		}

		f.writeValue(fields[i])
	}

	fmt.Fprintf(f.Writer, "\n")
}

func (f *CSVFormatter) writeValue(s string) {
	s = strings.Replace(s, `"`, `""`, -1)

	if strings.Index(s, ",") < 0 && strings.Index(s, `"`) < 0 {
		fmt.Fprintf(f.Writer, "%s", s)
	} else {
		fmt.Fprintf(f.Writer, `"%s"`, s)
	}
}

func (f *CSVFormatter) prepareLine(line interface{}) map[string]interface{} {
	if m, ok := line.(gedcom.ObjectMapper); ok {
		return m.ObjectMap()
	}

	l := reflect.ValueOf(line)

	if l.Kind() == reflect.Ptr {
		l = l.Elem()
	}

	if l.Kind() == reflect.Map {
		m := map[string]interface{}{}
		for _, name := range l.MapKeys() {
			m[name.Interface().(string)] = l.MapIndex(name).Interface()
		}

		return m
	}

	if l.Kind() == reflect.Struct {
		t := l.Type()
		m := map[string]interface{}{}

		for i := 0; i < t.NumField(); i++ {
			name := t.Field(i).Name

			// Ignore unexported
			if !unicode.IsUpper(rune(name[0])) {
				continue
			}

			m[name] = l.FieldByName(name).Interface()
		}

		return m
	}

	return nil
}

func (f *CSVFormatter) Header(result interface{}) ([]string, error) {
	v := reflect.ValueOf(result)

	if v.Kind() != reflect.Slice {
		return nil, errors.New("not a slice")
	}

	if v.Len() == 0 {
		return nil, nil
	}

	line := f.prepareLine(v.Index(0).Interface())

	s := []string{}
	for column := range line {
		s = append(s, column)
	}

	sort.Strings(s)

	return s, nil
}

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
