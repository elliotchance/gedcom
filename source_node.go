package gedcom

// SourceNode represents a source.
type SourceNode struct {
	*SimpleNode
}

func NewSourceNode(value, pointer string, children []Node) *SourceNode {
	return &SourceNode{
		NewSimpleNode(TagSource, value, pointer, children),
	}
}

func (node *SourceNode) Title() string {
	if n := First(NodesWithTag(node, TagTitle)); n != nil {
		return n.Value()
	}

	return ""
}
