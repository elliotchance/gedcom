// Encoding a Document
//
//   buf := bytes.NewBufferString("")
//
//   encoder := NewEncoder(buf, doc)
//   err := encoder.Encode()
//   if err != nil {
//     panic(err)
//   }
//
// If you need the GEDCOM data as a string you can simply using fmt.Stringer:
//
//   data := document.String()
//
package gedcom

import (
	"io"
)

// Encoder represents a GEDCOM encoder.
type Encoder struct {
	w           io.Writer
	document    *Document
	startIndent int
}

// Create a new encoder to generate GEDCOM data.
func NewEncoder(w io.Writer, document *Document) *Encoder {
	return &Encoder{
		w:        w,
		document: document,
	}
}

func (enc *Encoder) renderNode(indent int, node Node) error {
	gedcomLine := node.GEDCOMLine(indent) + "\n"
	_, err := enc.w.Write([]byte(gedcomLine))
	if err != nil {
		return err
	}

	for _, child := range node.Nodes() {
		nextIndent := indent + 1
		if indent == NoIndent {
			nextIndent = NoIndent
		}

		err = enc.renderNode(nextIndent, child)
		if err != nil {
			return err
		}
	}

	return nil
}

// Encode will write the GEDCOM document to the Writer.
func (enc *Encoder) Encode() (err error) {
	err = enc.restoreOptionalBOM()

	for _, node := range enc.document.Nodes() {
		err = enc.renderNode(enc.startIndent, node)
		if err != nil {
			return
		}
	}

	return
}

// See Decoder.consumeOptionalBOM for more information.
func (enc *Encoder) restoreOptionalBOM() (err error) {
	if enc.document.HasBOM {
		_, err = enc.w.Write(byteOrderMark)
	}

	return
}
