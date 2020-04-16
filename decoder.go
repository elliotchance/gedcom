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
	"github.com/elliotchance/gedcom/tag"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// See Decoder.consumeOptionalBOM().
var byteOrderMark = []byte{0xef, 0xbb, 0xbf}

// Decoder represents a GEDCOM decoder.
type Decoder struct {
	r *bufio.Reader

	// It is not valid for GEDCOM values to contain new lines or carriage
	// returns. However, some application dump data without correctly using the
	// CONT tags.
	//
	// Strictly speaking we should bail out with an error but there are too many
	// cases that are difficult to clean up for consumers so we offer and option
	// to permit it.
	//
	// When enabled any line than cannot be parsed will be considered an
	// extension of the previous line (including the new line character).
	AllowMultiLine bool

	// AllowInvalidIndents allows a child node to have an indent greater than +1
	// of the parent. AllowInvalidIndents is disabled by default because if this
	// happens the GEDCOM file is broken in some possibly serious way and
	// certainly not a valid GEDCOM file.
	//
	// The biggest problem with having the indents wrongly aligned is that nodes
	// that are expected to be a certain depth (such as NPFX inside a NAME) will
	// probably break or interfere with a traversal algorithm that is not
	// expecting the node to be there/at that level.
	//
	// Another important thing to note is that the incorrect indent level will
	// not be retained when writing the Document back to a GEDCOM.
	AllowInvalidIndents bool
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
	mutIndents := Nodes{}
	var family *FamilyNode

	document.HasBOM = dec.consumeOptionalBOM()

	// Only used when AllowMultiLine is enabled.
	var previousNode Node

	for lineNumber, finished := 0, false; !finished; {
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
			if dec.AllowMultiLine && previousNode != nil {
				previousNode.RawSimpleNode().value += "\n"
			}

			continue
		}

		node, mutIndent, err := parseLine(line, document, family)
		if err != nil {
			if dec.AllowMultiLine && previousNode != nil {
				previousNode.RawSimpleNode().value += "\n" + line
				continue
			}

			return nil, fmt.Errorf("line %d: %s", lineNumber, err)
		}

		// Families cannot be nested so any children that appear after this node
		// will be attached to the most recently seen family. We do not need to
		// set this back to nil after we exit the family node.
		if f, ok := node.(*FamilyNode); ok {
			family = f
		}

		// Add a root node to the document.
		if mutIndent == 0 {
			dec.trimNodeValue(previousNode)
			document.AddNode(node)
			previousNode = node

			// There can be multiple root nodes so make sure we always reset all
			// mutIndent pointers.
			mutIndents = Nodes{node}

			continue
		}

		if mutIndent-1 >= len(mutIndents) {
			// This means the file is not valid. I have seen it in very rare
			// cases. See full explanation in AllowInvalidIndents.
			if dec.AllowInvalidIndents {
				mutIndent = len(mutIndents)
			} else {
				panic(fmt.Sprintf(
					"indent is too large - missing parent? at line %d: %s",
					lineNumber, line))
			}
		}

		i := mutIndents[mutIndent-1]

		switch {
		case mutIndent >= len(mutIndents):
			// Descending one level. It is not valid for a child to have an
			// mutIndent that is more than one greater than the parent. This would
			// be a corrupt GEDCOM and lead to a panic.
			mutIndents = append(mutIndents, node)

		case mutIndent < len(mutIndents)-1:
			// Moving back to a parent. It is possible for this leap to be
			// greater than one so trim the mutIndent levels back as many times as
			// needed to represent the new mutIndent level.
			mutIndents = mutIndents[:mutIndent+1]
			mutIndents[mutIndent] = node

		default:
			// This case would be "mutIndent == len(mutIndents)-1" (the mutIndent does
			// not change from the previous line). However, since it is the only
			// other logical possibility there is no need to evaluate it for the
			// case condition.
			//
			// Make sure we update the current mutIndent with the new node so that
			// children get place on this node and not the previous one.
			mutIndents[mutIndent] = node
		}

		dec.trimNodeValue(previousNode)
		i.AddNode(node)

		previousNode = node
	}

	dec.trimNodeValue(previousNode)

	// Build the cache once.
	document.buildPointerCache()

	return document, nil
}

func (dec *Decoder) trimNodeValue(previousNode Node) {
	// When AllowMultiLine is enabled we have to be careful to trim the
	// surrounding spaces off the value so it can be interpreted correct.
	//
	// Another solution would be to ignore blank lines entirely, but then we
	// would miss the paragraph separators in multiline text.
	if !IsNil(previousNode) {
		newValue := strings.TrimSpace(previousNode.RawSimpleNode().value)
		previousNode.RawSimpleNode().value = newValue
	}
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

var lineRegexp = regexp.MustCompile(`^(\d) +(@[^@]+@ )?(\w+) ?(.*)?$`)

func parseLine(line string, document *Document, family *FamilyNode) (Node, int, error) {
	parts := lineRegexp.FindStringSubmatch(line)

	if len(parts) == 0 {
		return nil, 0, fmt.Errorf("could not parse: %s", line)
	}

	// Indent (required).
	indent, _ := strconv.Atoi(parts[1])

	// Pointer (optional).
	pointer := trimPointer(parts[2])

	// Tag (required).
	tag := tag.TagFromString(parts[3])

	// Value (optional).
	value := parts[4]

	return newNode(document, family, tag, value, pointer), indent, nil
}

func trimPointer(p string) string {
	if p != "" {
		// Trim off the surrounding '@'.
		return p[1 : len(p)-2]
	}

	return ""
}

// NewNode creates a node with no children. It is also the correct way to
// create a shallow copy of a node.
//
// If the node tag is recognised as a more specific type, such as *DateNode then
// that will be returned. Otherwise a *SimpleNode will be used.
func NewNode(tag tag.Tag, value, pointer string, children ...Node) Node {
	return newNodeWithChildren(nil, nil, tag, value, pointer, children)
}

func newNode(document *Document, family *FamilyNode, tag tag.Tag, value, pointer string) Node {
	return newNodeWithChildren(document, family, tag, value, pointer, nil)
}

func newNodeWithChildren(document *Document, family *FamilyNode, t tag.Tag, value, pointer string, children Nodes) Node {
	var node Node

	switch t {
	case tag.TagBaptism:
		node = NewBaptismNode(value, children...)

	case tag.TagBirth:
		node = NewBirthNode(value, children...)

	case tag.TagBurial:
		node = NewBurialNode(value, children...)

	case tag.TagChild:
		needsFamily(family, t)

		node = newChildNode(family, value, children...)

	case tag.TagDate:
		node = NewDateNode(value, children...)

	case tag.TagDeath:
		node = NewDeathNode(value, children...)

	case tag.TagEvent:
		node = NewEventNode(value, children...)

	case tag.TagFamily:
		needsDocument(document, t)

		node = newFamilyNode(document, pointer, children...)

	case tag.UnofficialTagFamilySearchID1, tag.UnofficialTagFamilySearchID2:
		node = NewFamilySearchIDNode(t, value, children...)

	case tag.TagFormat:
		node = NewFormatNode(value, children...)

	case tag.TagHusband:
		needsFamily(family, t)

		node = newHusbandNode(family, value, children...)

	case tag.TagIndividual:
		needsDocument(document, t)

		node = newIndividualNode(document, pointer, children...)

	case tag.TagLatitude:
		node = NewLatitudeNode(value, children...)

	case tag.TagLongitude:
		node = NewLongitudeNode(value, children...)

	case tag.TagMap:
		node = NewMapNode(value, children...)

	case tag.TagName:
		node = NewNameNode(value, children...)

	case tag.TagNickname:
		node = NewNicknameNode(value, children...)

	case tag.TagNote:
		node = NewNoteNode(value, children...)

	case tag.TagPhonetic:
		node = NewPhoneticVariationNode(value, children...)

	case tag.TagPlace:
		node = NewPlaceNode(value, children...)

	case tag.TagResidence:
		node = NewResidenceNode(value, children...)

	case tag.TagRomanized:
		node = NewRomanizedVariationNode(value, children...)

	case tag.TagSex:
		node = NewSexNode(value)

	case tag.TagSource:
		node = NewSourceNode(value, pointer, children...)

	case tag.TagType:
		node = NewTypeNode(value, children...)

	case tag.UnofficialTagUniqueID:
		node = NewUniqueIDNode(value, children...)

	case tag.TagWife:
		needsFamily(family, t)

		node = newWifeNode(family, value, children...)
	}

	if IsNil(node) {
		node = newSimpleNode(t, value, pointer, children...)
	} else {
		simpleNode := node.RawSimpleNode()
		simpleNode.pointer = pointer
	}

	return node
}

func needsDocument(document *Document, tag tag.Tag) {
	if document == nil {
		panic(fmt.Sprintf("cannot create %s without a document", tag))
	}
}

func needsFamily(family *FamilyNode, tag tag.Tag) {
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
