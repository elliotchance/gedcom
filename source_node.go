package gedcom

// SourceNode represents a source.
type SourceNode struct {
	*SimpleNode
}

func NewSourceNode(value, pointer string, children ...Node) *SourceNode {
	return &SourceNode{
		newSimpleNode(TagSource, value, pointer, children...),
	}
}

// If the node is nil the result will be an empty string.
func (node *SourceNode) Title() string {
	if n := First(NodesWithTag(node, TagTitle)); n != nil {
		return n.Value()
	}

	return ""
}
