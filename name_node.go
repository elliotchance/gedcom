package gedcom

import (
	"regexp"
	"strings"
	"unicode"
)

// NameFormat constants can be used with NameNode.Format.
const (
	// This is the written format, also used by String().
	NameFormatWritten = "%t %p %f %m %l %s"

	// This is the style used in GEDCOM NAME nodes. It is used in GedcomName().
	//
	// It should be noted that while the formatted name is valid GEDCOM, it
	// cannot be reverse back into its individual name parts.
	NameFormatGEDCOM = "%t %p %f %m /%l/ %s"

	// NameFormatIndex is appropriate for showing names that are indexed by
	// their surname, such as "Smith, Bob"
	NameFormatIndex = "%m %l, %t %p %f %s"
)

// NameNode represents all the parts that make up a single name. An individual
// may have more than one name, each one would be represented by a NameNode.
type NameNode struct {
	*SimpleNode
}

func NewNameNode(document *Document, value, pointer string, children []Node) *NameNode {
	return &NameNode{
		NewSimpleNode(document, TagName, value, pointer, children),
	}
}

var nameRegexp = regexp.MustCompile("([^/]*)(/[^/]*/)?(.*)")

func (node *NameNode) parts() []string {
	return nameRegexp.FindStringSubmatch(node.value)
}

// GivenName is the given or earned name used for official identification of a
// person. It is also commonly known as the "first name".
func (node *NameNode) GivenName() string {
	if node == nil {
		return ""
	}

	// GivenName is the proper first name.
	givenNames := NodesWithTag(node, TagGivenName)
	if len(givenNames) > 0 {
		return CleanSpace(givenNames[0].Value())
	}

	// Fall back to trying to extract the first name from NAME tag.
	return CleanSpace(node.parts()[1])
}

// Surname is a family name passed on or used by members of a family.
func (node *NameNode) Surname() string {
	if node == nil {
		return ""
	}

	// Surname is the proper last name.
	surnames := NodesWithTag(node, TagSurname)
	if len(surnames) > 0 {
		return CleanSpace(surnames[0].Value())
	}

	// Fallback to trying to extract the surname from the NAME tag.
	lastName := CleanSpace(node.parts()[2])
	if lastName == "" {
		return ""
	}

	// The surname (if provided) will be wrapped within //.
	return lastName[1 : len(lastName)-1]
}

func (node *NameNode) Prefix() string {
	if node == nil {
		return ""
	}

	// NamePrefix is the proper name prefix. If it is not provided then no
	// prefix should be returned.
	namePrefixes := NodesWithTag(node, TagNamePrefix)
	if len(namePrefixes) > 0 {
		return CleanSpace(namePrefixes[0].Value())
	}

	return ""
}

func (node *NameNode) Suffix() string {
	if node == nil {
		return ""
	}

	// NameSuffix is the proper name suffix.
	nameSuffixes := NodesWithTag(node, TagNameSuffix)
	if len(nameSuffixes) > 0 {
		return CleanSpace(nameSuffixes[0].Value())
	}

	// Otherwise fallback to trying to extract it from the NAME.
	return CleanSpace(node.parts()[3])
}

func (node *NameNode) SurnamePrefix() string {
	if node == nil {
		return ""
	}

	// SurnameSuffix is the proper surname prefix.
	surnamePrefixes := NodesWithTag(node, TagSurnamePrefix)
	if len(surnamePrefixes) > 0 {
		return CleanSpace(surnamePrefixes[0].Value())
	}

	// Otherwise return nothing.
	return ""
}

func (node *NameNode) Title() string {
	if node == nil {
		return ""
	}

	// Title is the proper individual title.
	titles := NodesWithTag(node, TagTitle)
	if len(titles) > 0 {
		return CleanSpace(titles[0].Value())
	}

	// Otherwise return nothing.
	return ""
}

// String returns all name components in the format that would be written like
// "Grand Duke Bob Smith Esq.". It specifically uses NameFormatWritten.
func (node *NameNode) String() string {
	return node.Format(NameFormatWritten)
}

func (node *NameNode) Type() NameType {
	if node == nil {
		return NameTypeNormal
	}

	if nameType := First(NodesWithTag(node, TagType)); nameType != nil {
		return NameType(nameType.Value())
	}

	// Otherwise return nothing.
	return NameTypeNormal
}

// GedcomName returns the simplified GEDCOM name often used also as the value
// for the NAME node.
//
// The only difference between this as String() is that the surname is
// encapsulated inside forward slashes like:
//
//   Sir Elliot Rupert /Chance/ Sr
//
// Even this uses the NameFormatGEDCOM it may return a different value from
// Format(NameFormatGEDCOM) because any empty surnames will be removed.
func (node *NameNode) GedcomName() string {
	name := node.Format(NameFormatGEDCOM)

	return CleanSpace(strings.Replace(name, "//", "", -1))
}

// Format returns a formatted name.
//
// There are some common formats described with the NameFormat constants.
//
// Format works similar to Printf where placeholders represent different
// components of the name:
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
// case also.
//
// Whitespace before, after and between name components will be removed.
//
// Examples:
//
//   name.Format("%l, %f")     // Smith, Bob
//   name.Format("%f %L")      // Bob SMITH
//   name.Format("%f %m (%l)") // Bob (Smith)
//
func (node *NameNode) Format(format string) string {
	result := ""

	for i := 0; i < len(format); i++ {
		if format[i] == '%' && i < len(format)-1 {
			nextLetter := format[i+1]

			switch nextLetter {
			case '%':
				result += "%"

			case 'f', 'F':
				result += renderNameComponent(nextLetter, node.GivenName())

			case 'l', 'L':
				result += renderNameComponent(nextLetter, node.Surname())

			case 'm', 'M':
				result += renderNameComponent(nextLetter, node.SurnamePrefix())

			case 'p', 'P':
				result += renderNameComponent(nextLetter, node.Prefix())

			case 's', 'S':
				result += renderNameComponent(nextLetter, node.Suffix())

			case 't', 'T':
				result += renderNameComponent(nextLetter, node.Title())

			default:
				result += "%" + string(nextLetter)
			}

			i++
		} else {
			result += string(format[i])
		}
	}

	return CleanSpace(result)
}

func renderNameComponent(letter byte, namePart string) string {
	isUpper := unicode.IsUpper(rune(letter))

	return conditionalUpperCase(namePart, isUpper)
}

func conditionalUpperCase(s string, upperCase bool) string {
	if upperCase {
		return strings.ToUpper(s)
	}

	return s
}
