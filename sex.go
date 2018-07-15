package gedcom

type Sex string

const (
	SexMale    = Sex("M")
	SexFemale  = Sex("F")
	SexUnknown = Sex("U")
)

func (s Sex) String() string {
	switch s {
	case SexMale:
		return "Male"
	case SexFemale:
		return "Female"
	default: // case SexUnknown:
		return "Unknown"
	}
}
