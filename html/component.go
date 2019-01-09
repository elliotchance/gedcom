package html

import "io"

type Component interface {
	io.WriterTo
}
