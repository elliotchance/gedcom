package gedcom

// Indicates the method used in transforming the text to a romanized variation.
//
// These constants can be used for RomanizedVariationNode.Type.Value. The value
// is not limited to these constants. Any user defined value is also valid.
const (
	RomanizedVariationTypePinyin    = "pinyin"
	RomanizedVariationTypeRomaji    = "romaji"
	RomanizedVariationTypeWadegiles = "wadegiles"
)

// RomanizedVariationNode represents a romanized variation of a superior text
// string.
//
// New in Gedcom 5.5.1.
type RomanizedVariationNode struct {
	*SimpleNode
}

// NewRomanizedVariationNode creates a new ROMN node.
func NewRomanizedVariationNode(document *Document, value, pointer string, children []Node) *RomanizedVariationNode {
	return &RomanizedVariationNode{
		newSimpleNode(document, TagRomanized, value, pointer, children),
	}
}

func (node *RomanizedVariationNode) Type() *TypeNode {
	return getType(node)
}

func getType(node Node) *TypeNode {
	n := First(NodesWithTag(node, TagType))

	if IsNil(n) {
		return nil
	}

	return n.(*TypeNode)
}
