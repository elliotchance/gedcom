package gedcom

import "github.com/elliotchance/gedcom/tag"

// NicknameNode is a descriptive or familiar that is used instead of, or in
// addition to, one's proper name.
type NicknameNode struct {
	*SimpleNode
}

// NewNicknameNode creates a new NICK node.
func NewNicknameNode(value string, children ...Node) *NicknameNode {
	return &NicknameNode{
		newSimpleNode(tag.TagNickname, value, "", children...),
	}
}
