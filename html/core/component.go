package core

import "io"

type Component interface {
	WriteHTMLTo(w io.Writer) (n int64, err error)
}
