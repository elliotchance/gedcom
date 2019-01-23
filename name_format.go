package gedcom

// NameFormat describes how an individuals name should be formatted.
//
// The format is a string that contains placeholders and works similar to
// fmt.Printf where placeholders represent different components of the name:
//
//   %% "%"
//   %f GivenName
//   %l Surname
//   %m SurnamePrefix
//   %p Prefix
//   %s Suffix
//   %t Title
//
// Each of the letters may be in upper case to convert the name part to upper
// case also. Whitespace before, after and between name components will be
// removed:
//
//   name.Format("%l, %f")     // "Smith, Bob"
//   name.Format("%f %L")      // "Bob SMITH"
//   name.Format("%f %m (%l)") // "Bob (Smith)"
//
type NameFormat string

// NameFormat constants can be used with NameNode.Format.
const (
	// This is the written format, also used by String().
	NameFormatWritten NameFormat = "%t %p %f %m %l %s"

	// This is the style used in GEDCOM NAME nodes. It is used in GedcomName().
	//
	// It should be noted that while the formatted name is valid GEDCOM, it
	// cannot be reverse back into its individual name parts.
	NameFormatGEDCOM NameFormat = "%t %p %f %m /%l/ %s"

	// NameFormatIndex is appropriate for showing names that are indexed by
	// their surname, such as "Smith, Bob"
	NameFormatIndex NameFormat = "%m %l, %t %p %f %s"
)

// NewNameFormatByName returns one of the NameFormat constants.
//
// The name is the lowercase name of the constant without the prefix. For
// example, "NameFormatWritten" would have the name "written".
//
// It will return the original name and false if the name is not known. This
// allows custom formats to be passed through like:
//
//   NewNameFormatByName("gedcom") // "%t %p %f %m /%l/ %s", true
//   NewNameFormatByName("%L %f")  // "%L %f", false
//
func NewNameFormatByName(name string) (NameFormat, bool) {
	switch name {
	case "written":
		return NameFormatWritten, true

	case "gedcom":
		return NameFormatGEDCOM, true

	case "index":
		return NameFormatIndex, true
	}

	return NameFormat(name), false
}
