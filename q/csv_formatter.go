package q

import (
	"errors"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"
	"unicode"

	"github.com/elliotchance/gedcom/v39"
)

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
	commaIndex := strings.Index(s, ",")
	quoteIndex := strings.Index(s, `"`)

	if commaIndex < 0 && quoteIndex < 0 {
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
