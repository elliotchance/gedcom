package gedcom

type NameType string

const (
	NameTypeNormal      = NameType("")
	NameTypeMarriedName = NameType("married")
	NameTypeAlsoKnownAs = NameType("aka")
	NameTypeMaidenName  = NameType("maiden")
	NameTypeNickname    = NameType("nick")
)

// Is usually blank for the primary name
func (t NameType) String() string {
	switch t {
	case NameTypeNormal:
		return "Normal"

	case NameTypeMarriedName:
		return "Married Name"

	case NameTypeAlsoKnownAs:
		return "Also Known As"

	case NameTypeMaidenName:
		return "Maiden Name"

	case NameTypeNickname:
		return "Nickname"
	}

	return string(t)
}
