package gedcom

const (
	SexMale    = "M"
	SexFemale  = "F"
	SexUnknown = "U"
)

// Indicates the sex of an individual--male or female.
type SexNode struct {
	*SimpleNode
}

// NewSexNode returns a node that represents a known or unknown gender. You
// should use one of the constants; SexMale, SexFemale or SexUnknown.
func NewSexNode(value string) *SexNode {
	return &SexNode{
		newSimpleNode(TagSex, value, ""),
	}
}

func (node *SexNode) IsMale() bool {
	if node == nil {
		return false
	}

	return node.Value() == SexMale
}

func (node *SexNode) IsFemale() bool {
	if node == nil {
		return false
	}

	return node.Value() == SexFemale
}

func (node *SexNode) IsUnknown() bool {
	if node == nil {
		return true
	}

	return !node.IsMale() && !node.IsFemale()
}

func (node *SexNode) String() string {
	if node == nil {
		goto unknown
	}

	switch {
	case node.IsMale():
		return "Male"

	case node.IsFemale():
		return "Female"
	}

unknown:
	// case SexUnknown:
	return "Unknown"
}

func (node *SexNode) OwnershipWord() string {
	if node == nil {
		goto unknown
	}

	switch {
	case node.IsMale():
		return "his"

	case node.IsFemale():
		return "her"
	}

unknown:
	return "their"
}
