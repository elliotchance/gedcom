package gedcom

const (
	SexMale = "M"
	SexFemale = "F"
	SexUnknown = "U"
)

// IndividualNode represents a person.
type IndividualNode struct {
	*SimpleNode
}

func NewIndividualNode(value, pointer string, children []Node) *IndividualNode {
	return &IndividualNode{
		NewSimpleNode(TagIndividual, value, pointer, children),
	}
}

func (node *IndividualNode) Names() []*NameNode {
	nameTags := node.NodesWithTag(TagName)
	names := make([]*NameNode, len(nameTags))

	for i, name := range nameTags {
		names[i] = name.(*NameNode)
	}

	return names
}

func (node *IndividualNode) Sex() string {
	sex := node.NodesWithTag(TagSex)
	if len(sex) == 0 {
		return SexUnknown
	}

	return sex[0].Value()
}
