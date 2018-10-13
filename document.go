package gedcom

import (
	"bytes"
	"os"
	"strings"
)

// DefaultMaxLivingAge is used when creating a new document. See
// Document.MaxLivingAge for a full description.
const DefaultMaxLivingAge = 100.0

// Document represents a whole GEDCOM document. It is possible for a
// Document to contain zero Nodes, this means the GEDCOM file was empty. It
// may also (and usually) contain several Nodes.
//
// You should not instantiate a Document yourself because there are sensible
// defaults and cache that need to be setup. Use one of the NewDocument
// constructors instead.
type Document struct {
	// nodes is private because we need to track changes.
	nodes []Node

	// pointerCache is setup once when the document is created.
	pointerCache map[string]Node

	families []*FamilyNode

	// HasBOM controls if the encoded stream will start with the Byte Order
	// Mark.
	//
	// This is not recommended by the UTF-8 standard and many applications will
	// have problems reading the data. However, streams that were decoded
	// containing the BOM will retain it so that the re-encoded stream is as
	// compatible and similar to the original stream as possible.
	//
	// Also see Decoder.consumeOptionalBOM().
	HasBOM bool

	// MaxLivingAge is used by Individual.IsLiving to determine if an individual
	// without a DeathNode should be considered living.
	//
	// 100 is chosen as a reasonable default. If you set it to 0 then an
	// individual will never be considered dead without a DeathNode.
	MaxLivingAge float64
}

// String will render the entire GEDCOM document.
func (doc *Document) String() string {
	buf := bytes.NewBufferString("")

	encoder := NewEncoder(buf, doc)
	err := encoder.Encode()
	if err != nil {
		panic(err)
	}

	return buf.String()
}

// Individuals returns all of the people in the document.
func (doc *Document) Individuals() IndividualNodes {
	individuals := IndividualNodes{}

	for _, node := range doc.Nodes() {
		if n, ok := node.(*IndividualNode); ok {
			individuals = append(individuals, n)
		}
	}

	return individuals
}

// NodeByPointer returns the Node for a pointer value.
//
// If the pointer does not exist nil is returned.
func (doc *Document) NodeByPointer(ptr string) Node {
	// Build the cache once.
	if doc.pointerCache == nil {
		doc.pointerCache = map[string]Node{}

		for _, node := range doc.Nodes() {
			if node.Pointer() != "" {
				doc.pointerCache[node.Pointer()] = node
			}
		}
	}

	return doc.pointerCache[ptr]
}

// Families returns the family entities in the document.
func (doc *Document) Families() (families []*FamilyNode) {
	if doc.families != nil {
		return doc.families
	}

	defer func() {
		doc.families = families
	}()

	families = []*FamilyNode{}

	for _, node := range doc.Nodes() {
		if n, ok := node.(*FamilyNode); ok {
			families = append(families, n)
		}
	}

	return families
}

// TODO: needs tests
func (doc *Document) Places() map[*PlaceNode]Node {
	places := map[*PlaceNode]Node{}

	for _, node := range doc.Nodes() {
		extractPlaces(node, places)
	}

	return places
}

func extractPlaces(n Node, dest map[*PlaceNode]Node) {
	for _, node := range n.Nodes() {
		if place, ok := node.(*PlaceNode); ok {
			// The place points to the parent node which is the thing that the
			// place is describing.
			dest[place] = n
		} else {
			extractPlaces(node, dest)
		}
	}
}

// TODO: Needs tests
func (doc *Document) Sources() []*SourceNode {
	sources := []*SourceNode{}

	for _, node := range doc.Nodes() {
		if n, ok := node.(*SourceNode); ok {
			sources = append(sources, n)
		}
	}

	return sources
}

// AddNode appends a node to the document.
//
// If the node is nil this function has no effect.
//
// If the node already exists it will be added again. This will cause problems
// with duplicate references.
func (doc *Document) AddNode(node Node) *Document {
	if !IsNil(node) {
		doc.nodes = append(doc.nodes, node)
		doc.pointerCache = nil
	}

	return doc
}

// Nodes returns the root nodes for the document.
//
// It is important that the slice returned is not manually manipulated (such as
// appending) because it may cause the internal cache to all out of sync. You
// may manipulate the nodes themselves.
func (doc *Document) Nodes() []Node {
	return doc.nodes
}

// NewDocumentFromGEDCOMFile returns a decoded Document from the provided file.
//
// If the file does not exist, be read or parse then an error is returned and
// the document will be nil.
func NewDocumentFromGEDCOMFile(path string) (*Document, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	decoder := NewDecoder(file)
	return decoder.Decode()
}

// NewDocumentFromString creates a document from a string containing GEDCOM
// data.
//
// An error is returned if a line cannot be parsed.
func NewDocumentFromString(gedcom string) (*Document, error) {
	decoder := NewDecoder(strings.NewReader(gedcom))

	return decoder.Decode()
}

// NewDocument returns an empty document.
func NewDocument() *Document {
	return &Document{
		MaxLivingAge: DefaultMaxLivingAge,
	}
}

// NewDocumentWithNodes creates a new document with the provided root nodes.
func NewDocumentWithNodes(nodes []Node) *Document {
	document := NewDocument()
	document.nodes = nodes

	return document
}
