package gedcom

// NicknameNode is a descriptive or familiar that is used instead of, or in
// addition to, one's proper name.
type NicknameNode struct {
	*SimpleNode
}

// NewNicknameNode creates a new NICK node.
func NewNicknameNode(document *Document, value, pointer string, children []Node) *NicknameNode {
	return &NicknameNode{
		newSimpleNode(document, TagNickname, value, pointer, children),
	}
}
