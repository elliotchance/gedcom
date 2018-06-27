package gedcom

// IndividualNode represents a person.
type IndividualNode struct {
	*SimpleNode
}

func NewIndividualNode(value, pointer string, children []Node) *IndividualNode {
	return &IndividualNode{
		NewSimpleNode(Name, value, pointer, children),
	}
}

func (node *IndividualNode) Names() []*NameNode {
	nameTags := node.NodesWithTag(Name)
	names := make([]*NameNode, len(nameTags))

	for i, name := range nameTags {
		names[i] = name.(*NameNode)
	}

	return names
}
