package gedcom

import (
	"bufio"
	"bytes"
	"io"
	"regexp"
	"strconv"
)

// Decoder represents a GEDCOM decoder.
type Decoder struct {
	r *bufio.Reader
}

// Create a new decoder to parse a reader that contain GEDCOM data.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r: bufio.NewReader(r),
	}
}

// Decode will parse the entire GEDCOM stream (until EOF is reached) and return
// a Document. If the GEDCOM stream is not valid then the document node will
// be nil and the error is returned.
//
// A blank GEDCOM or a GEDCOM that only contains empty lines is valid and a
// Document will be returned with zero nodes.
func (dec *Decoder) Decode() (*Document, error) {
	document := &Document{
		Nodes: []Node{},
	}
	indents := []Node{}

	finished := false
	for !finished {
		line, err := dec.readLine()
		if err != nil {
			if err != io.EOF {
				return nil, err
			}

			finished = true
		}

		// Skip blank lines.
		if line == "" {
			continue
		}

		node, indent := parseLine(line)

		// Add a root node to the document.
		if indent == 0 {
			document.Nodes = append(document.Nodes, node)

			// There can be multiple root nodes so make sure we always reset all
			// indent pointers.
			indents = []Node{node}

			continue
		}

		i := indents[indent-1]

		switch {
		case indent >= len(indents):
			// Descending one level. It is not valid for a child to have an
			// indent that is more than one greater than the parent. This would
			// be a corrupt GEDCOM and lead to a panic.
			indents = append(indents, node)

		case indent < len(indents)-1:
			// Moving back to a parent. It is possible for this leap to be
			// greater than one so trim the indent levels back as many times as
			// needed to represent the new indent level.
			indents = indents[:indent+1]
			indents[indent] = node

		default:
			// This case would be "indent == len(indents)-1" (the indent does
			// not change from the previous line). However, since it is the only
			// other logical possibility there is no need to evaluate it for the
			// case condition.
			//
			// Make sure we update the current indent with the new node so that
			// children get place on this node and not the previous one.
			indents[indent] = node
		}

		i.AddNode(node)
	}

	return document, nil
}

func (dec *Decoder) readLine() (string, error) {
	buf := new(bytes.Buffer)

	for {
		b, err := dec.r.ReadByte()
		if err != nil {
			return string(buf.Bytes()), err
		}

		// The line endings in the GEDCOM files can be different. A newline and
		// carriage return are both considered to be the end of the line and empty
		// lines are ignored so we can treat both of these characters as independent
		// line terminators.
		if b == '\n' || b == '\r' {
			break
		}

		buf.WriteByte(b)
	}

	return string(buf.Bytes()), nil
}

func parseLine(line string) (Node, int) {
	parts := regexp.
		MustCompile(`^(\d) (@\w+@ )?(\w+)( .*)?$`).
		FindStringSubmatch(line)

	indent := 0
	if len(parts) > 1 {
		indent, _ = strconv.Atoi(parts[1])
	}

	pointer := ""
	if len(parts) > 2 && len(parts[2]) > 4 {
		pointer = parts[2][1 : len(parts[2])-2]
	}

	tag := Tag("")
	if len(parts) > 3 {
		tag = Tag(parts[3])
	}

	value := ""
	if len(parts) > 4 && len(parts[4]) > 0 {
		value = parts[4][1:]
	}

	switch tag {
	case TagFamily:
		return NewFamilyNode(pointer, []Node{}), indent

	case TagIndividual:
		return NewIndividualNode(value, pointer, []Node{}), indent

	case TagName:
		return NewNameNode(value, pointer, []Node{}), indent

	case TagPlace:
		return NewPlaceNode(value, pointer, []Node{}), indent

	default:
		return NewSimpleNode(tag, value, pointer, []Node{}), indent
	}
}
