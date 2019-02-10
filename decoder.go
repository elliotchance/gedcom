// Decoding a Document
//
// Decoding a GEDCOM stream:
//
//   ged := "0 HEAD\n1 CHAR UTF-8"
//
//   decoder := gedcom.NewDecoder(strings.NewReader(ged))
//   document, err := decoder.Decode()
//   if err != nil {
//     panic(err)
//   }
//
// If you are reading from a file you can use NewDocumentFromGEDCOMFile:
//
//   document, err := gedcom.NewDocumentFromGEDCOMFile("family.ged")
//   if err != nil {
//       panic(err)
//   }
//
package gedcom

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

// See Decoder.consumeOptionalBOM().
var byteOrderMark = []byte{0xef, 0xbb, 0xbf}

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
	document := NewDocument()
	indents := Nodes{}
	var family *FamilyNode

	document.HasBOM = dec.consumeOptionalBOM()

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

		node, indent, err := parseLine(line, document, family)
		if err != nil {
			return nil, fmt.Errorf("line %d: %s", lineNumber, err)
		}

		// Families cannot be nested so any children that appear after this node
		// will be attached to the most recently seen family. We do not need to
		// set this back to nil after we exit the family node.
		if f, ok := node.(*FamilyNode); ok {
			family = f
		}

		// Add a root node to the document.
		if indent == 0 {
			document.AddNode(node)

			// There can be multiple root nodes so make sure we always reset all
			// indent pointers.
			indents = Nodes{node}

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

	// Build the cache once.
	document.buildPointerCache()

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
		// carriage return are both considered to be the end of the line and
		// empty lines are ignored so we can treat both of these characters as
		// independent line terminators.
		if b == '\n' || b == '\r' {
			break
		}

		buf.WriteByte(b)
	}

	return string(buf.Bytes()), nil
}

var lineRegexp = regexp.MustCompile(`^(\d) (@[^@]+@ )?(\w+) ?(.*)?$`)

func parseLine(line string, document *Document, family *FamilyNode) (Node, int, error) {
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

	return newNode(document, family, tag, value, pointer), indent, nil
}

// NewNode creates a node with no children. It is also the correct way to
// create a shallow copy of a node.
//
// If the node tag is recognised as a more specific type, such as *DateNode then
// that will be returned. Otherwise a *SimpleNode will be used.
func NewNode(tag Tag, value, pointer string, children ...Node) Node {
	return newNodeWithChildren(nil, nil, tag, value, pointer, children)
}

func newNode(document *Document, family *FamilyNode, tag Tag, value, pointer string) Node {
	return newNodeWithChildren(document, family, tag, value, pointer, nil)
}

func newNodeWithChildren(document *Document, family *FamilyNode, tag Tag, value, pointer string, children Nodes) Node {
	var node Node

	switch tag {
	case TagBaptism:
		node = NewBaptismNode(value, children...)

	case TagBirth:
		node = NewBirthNode(value, children...)

	case TagBurial:
		node = NewBurialNode(value, children...)

	case TagChild:
		needsFamily(family, tag)

		node = newChildNode(family, value, children...)

	case TagDate:
		node = NewDateNode(value, children...)

	case TagDeath:
		node = NewDeathNode(value, children...)

	case TagEvent:
		node = NewEventNode(value, children...)

	case TagFamily:
		needsDocument(document, tag)

		node = newFamilyNode(document, pointer, children...)

	case UnofficialTagFamilySearchID1, UnofficialTagFamilySearchID2:
		node = NewFamilySearchIDNode(tag, value, children...)

	case TagFormat:
		node = NewFormatNode(value, children...)

	case TagHusband:
		needsFamily(family, tag)

		node = newHusbandNode(family, value, children...)

	case TagIndividual:
		needsDocument(document, tag)

		node = newIndividualNode(document, pointer, children...)

	case TagLatitude:
		node = NewLatitudeNode(value, children...)

	case TagLongitude:
		node = NewLongitudeNode(value, children...)

	case TagMap:
		node = NewMapNode(value, children...)

	case TagName:
		node = NewNameNode(value, children...)

	case TagNickname:
		node = NewNicknameNode(value, children...)

	case TagNote:
		node = NewNoteNode(value, children...)

	case TagPhonetic:
		node = NewPhoneticVariationNode(value, children...)

	case TagPlace:
		node = NewPlaceNode(value, children...)

	case TagResidence:
		node = NewResidenceNode(value, children...)

	case TagRomanized:
		node = NewRomanizedVariationNode(value, children...)

	case TagSex:
		node = NewSexNode(value)

	case TagSource:
		node = NewSourceNode(value, pointer, children...)

	case TagType:
		node = NewTypeNode(value, children...)

	case UnofficialTagUniqueID:
		node = NewUniqueIDNode(value, children...)

	case TagWife:
		needsFamily(family, tag)

		node = newWifeNode(family, value, children...)
	}

	if IsNil(node) {
		node = newSimpleNode(tag, value, pointer, children...)
	} else {
		simpleNode := node.RawSimpleNode()
		simpleNode.pointer = pointer
	}

	return node
}

func needsDocument(document *Document, tag Tag) {
	if document == nil {
		panic(fmt.Sprintf("cannot create %s without a document", tag))
	}
}

func needsFamily(family *FamilyNode, tag Tag) {
	if family == nil {
		panic(fmt.Sprintf("cannot create %s without a family", tag))
	}
}

// consumeOptionalBOM will test and discard the Byte Order Mark at the start of
// the stream.
//
// In order to keep the original stream as intact as possible when encoding the
// BOM will be written back if it existed originally.
//
// Use of a BOM is neither required nor recommended for UTF-8, but may be
// encountered in contexts where UTF-8 data is converted from other encoding
// forms that use a BOM or where the BOM is used as a UTF-8 signature. See the
// “Byte Order Mark” subsection in Section 16.8, Specials, for more information.
// - 2.6 Encoding Schemes, http://www.unicode.org/versions/Unicode5.0.0/ch02.pdf
func (dec *Decoder) consumeOptionalBOM() bool {
	possibleBOM, _ := dec.r.Peek(3)
	hasBOM := bytes.Compare(possibleBOM, byteOrderMark) == 0

	if hasBOM {
		dec.r.Discard(3)
	}

	return hasBOM
}
