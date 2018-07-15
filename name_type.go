package gedcom

type NameType string

const (
	NameTypeMarriedName = NameType("married")
	NameTypeAlsoKnownAs = NameType("aka")
	NameTypeMaidenName  = NameType("maiden")
)

// Is usually blank for the primary name
func (t NameType) String() string {
	switch t {
	case NameTypeMarriedName:
		return "Married Name"

	case NameTypeAlsoKnownAs:
		return "Also Known As"

	case NameTypeMaidenName:
		return "Maiden Name"
	}

	return string(t)
}
