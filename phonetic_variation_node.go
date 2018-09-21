package gedcom

// Indicates the method used in transforming the text to the phonetic variation.
//
// These constants can be used for PhoneticVariationNode.Type.Value. The value
// is not limited to these constants. Any user defined value is also valid.
const (
	// Phonetic method for sounding Korean glifs.
	PhoneticVariationTypeHangul = "hangul"

	// Hiragana and/or Katakana characters were used in sounding the Kanji
	// character used by japanese.
	PhoneticVariationTypeKana = "kana"
)

// PhoneticVariationNode represents a phonetic variation of a superior text
// string.
//
// New in Gedcom 5.5.1.
type PhoneticVariationNode struct {
	*SimpleNode
}

// NewPhoneticVariationNode creates a new FONE node.
func NewPhoneticVariationNode(document *Document, value, pointer string, children []Node) *PhoneticVariationNode {
	return &PhoneticVariationNode{
		NewSimpleNode(document, TagPhonetic, value, pointer, children),
	}
}

func (node *PhoneticVariationNode) Type() *TypeNode {
	return getType(node)
}
