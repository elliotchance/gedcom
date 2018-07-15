package gedcom

import (
	"fmt"
	"io"
)

// Encoder represents a GEDCOM encoder.
type Encoder struct {
	w        io.Writer
	document *Document
}

// Create a new encoder to generate GEDCOM data.
func NewEncoder(w io.Writer, document *Document) *Encoder {
	return &Encoder{
		w:        w,
		document: document,
	}
}

func (enc *Encoder) renderNode(indent int, node Node) error {
	_, err := enc.w.Write([]byte(fmt.Sprintf("%d %s\n", indent, node.gedcomLine())))
	if err != nil {
		return err
	}

	for _, child := range node.Nodes() {
		err = enc.renderNode(indent+1, child)
		if err != nil {
			return err
		}
	}

	return nil
}

// Encode will write the GEDCOM document to the Writer.
func (enc *Encoder) Encode() error {
	for _, node := range enc.document.Nodes {
		err := enc.renderNode(0, node)
		if err != nil {
			return err
		}
	}

	return nil
}
