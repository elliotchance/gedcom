package core

import (
	"fmt"
	"io"
)

func write(w io.Writer, data []byte) (int64, error) {
	n, err := w.Write(data)

	return int64(n), err
}

func writeString(w io.Writer, data string) (int64, error) {
	return write(w, []byte(data))
}

func appendString(w io.Writer, data string) int64 {
	n, err := writeString(w, data)
	if err != nil {
		panic(err)
	}

	return n
}

func appendComponent(w io.Writer, component Component) int64 {
	n, err := component.WriteHTMLTo(w)
	if err != nil {
		panic(err)
	}

	return n
}

func writeSprintf(w io.Writer, format string, args ...interface{}) (int64, error) {
	return writeString(w, fmt.Sprintf(format, args...))
}

func appendSprintf(w io.Writer, format string, args ...interface{}) int64 {
	n, err := writeSprintf(w, format, args...)
	if err != nil {
		panic(err)
	}

	return n
}

func writeNothing() (int64, error) {
	return 0, nil
}
