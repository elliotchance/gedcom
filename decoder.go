package gedcom

import (
	"bufio"
	"bytes"
	"fmt"
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
	lineNumber := 0
	for !finished {
		lineNumber++

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

		node, indent, err := parseLine(document, line)
		if err != nil {
			return nil, fmt.Errorf("line %d: %s", lineNumber, err)
		}

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

var lineRegexp = regexp.MustCompile(`^(\d) (@\w+@ )?(\w+) ?(.*)?$`)

func parseLine(document *Document, line string) (Node, int, error) {
	parts := lineRegexp.FindStringSubmatch(line)

	if len(parts) == 0 {
		return nil, 0, fmt.Errorf("could not parse: %s", line)
	}

	// Indent (required).
	indent, _ := strconv.Atoi(parts[1])

	// Pointer (optional).
	pointer := ""
	if parts[2] != "" {
		// Trim off the surrounding '@'.
		pointer = parts[2][1 : len(parts[2])-2]
	}

	// Tag (required).
	tag := TagFromString(parts[3])

	// Value (optional).
	value := parts[4]

	return NewNode(document, tag, value, pointer), indent, nil
}

// NewNode creates a node with no children. It is also the correct way to
// create a shallow copy of a node.
//
// If the node tag is recognised as a more specific type, such as *DateNode then
// that will be returned. Otherwise a *SimpleNode will be used.
func NewNode(document *Document, tag Tag, value, pointer string) Node {
	switch tag {
	case TagBirth:
		return NewBirthNode(document, value, pointer, nil)

	case TagDate:
		return NewDateNode(document, value, pointer, nil)

	case TagEvent:
		return NewEventNode(document, value, pointer, nil)

	case TagFamily:
		return NewFamilyNode(document, pointer, nil)

	case TagFormat:
		return NewFormatNode(document, value, pointer, nil)

	case TagIndividual:
		return NewIndividualNode(document, value, pointer, nil)

	case TagLatitude:
		return NewLatitudeNode(document, value, pointer, nil)

	case TagLongitude:
		return NewLongitudeNode(document, value, pointer, nil)

	case TagMap:
		return NewMapNode(document, value, pointer, nil)

	case TagName:
		return NewNameNode(document, value, pointer, nil)

	case TagNote:
		return NewNoteNode(document, value, pointer, nil)

	case TagPhonetic:
		return NewPhoneticVariationNode(document, value, pointer, nil)

	case TagPlace:
		return NewPlaceNode(document, value, pointer, nil)

	case TagResidence:
		return NewResidenceNode(document, value, pointer, nil)

	case TagRomanized:
		return NewRomanizedVariationNode(document, value, pointer, nil)

	case TagSource:
		return NewSourceNode(document, value, pointer, nil)

	case TagType:
		return NewTypeNode(document, value, pointer, nil)
	}

	return NewSimpleNode(document, tag, value, pointer, nil)
}
