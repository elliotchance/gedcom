package gedcom

import (
	"bytes"
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
	_, err := enc.w.Write([]byte(GedcomLine(indent, node) + "\n"))
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

// GedcomLine converts a node into its single line GEDCOM value. It is used
// several places including the actual Encoder.
//
// GedcomLine, as the name would suggest, does not handle children. You should
// use the proper Encoder instead.
//
// GedcomLine will handle nil nodes gracefully by returning an empty string.
//
// The indent will only be included if it is at least 0. If you want to use
// GedcomLine to compare the string values of nodes or exclude the indent you
// should pass -1 as the indent.
func GedcomLine(indent int, node Node) string {
	if IsNil(node) {
		return ""
	}

	buf := bytes.NewBufferString("")

	if indent >= 0 {
		buf.WriteString(fmt.Sprintf("%d ", indent))
	}

	if p := node.Pointer(); p != "" {
		buf.WriteString(fmt.Sprintf("@%s@ ", p))
	}

	buf.WriteString(node.Tag().Tag())

	if v := node.Value(); v != "" {
		buf.WriteByte(' ')
		buf.WriteString(v)
	}

	return buf.String()
}
