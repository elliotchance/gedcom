package gedcom

import (
	"fmt"
	"regexp"
	"strings"
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
	// NamePrefix is the proper name prefix. If it is not provided then no
	// prefix should be returned.
	namePrefixes := NodesWithTag(node, TagNamePrefix)
	if len(namePrefixes) > 0 {
		return CleanSpace(namePrefixes[0].Value())
	}

	return ""
}

func (node *NameNode) Suffix() string {
	// NameSuffix is the proper name suffix.
	nameSuffixes := NodesWithTag(node, TagNameSuffix)
	if len(nameSuffixes) > 0 {
		return CleanSpace(nameSuffixes[0].Value())
	}

	// Otherwise fallback to trying to extract it from the NAME.
	return CleanSpace(node.parts()[3])
}

func (node *NameNode) SurnamePrefix() string {
	// SurnameSuffix is the proper surname prefix.
	surnamePrefixes := NodesWithTag(node, TagSurnamePrefix)
	if len(surnamePrefixes) > 0 {
		return CleanSpace(surnamePrefixes[0].Value())
	}

	// Otherwise return nothing.
	return ""
}

func (node *NameNode) Title() string {
	// Title is the proper individual title.
	titles := NodesWithTag(node, TagTitle)
	if len(titles) > 0 {
		return CleanSpace(titles[0].Value())
	}

	// Otherwise return nothing.
	return ""
}

func (node *NameNode) String() string {
	return CleanSpace(fmt.Sprintf("%s %s %s %s %s %s", node.Title(),
		node.Prefix(), node.GivenName(), node.SurnamePrefix(), node.Surname(),
		node.Suffix()))
}

func (node *NameNode) Type() NameType {
	if nameType := First(NodesWithTag(node, TagType)); nameType != nil {
		return NameType(nameType.Value())
	}

	// Otherwise return nothing.
	return ""
}

// GedcomName returns the simplified GEDCOM name often used also as the value
// for the NAME node.
//
// The only difference between this as String() is that the surname is
// encapsulated inside forward slashes like:
//
//   Sir Elliot Rupert /Chance/ Sr
//
func (node *NameNode) GedcomName() string {
	name := fmt.Sprintf("%s %s %s %s /%s/ %s", node.Title(),
		node.Prefix(), node.GivenName(), node.SurnamePrefix(), node.Surname(),
		node.Suffix())

	return CleanSpace(strings.Replace(name, "//", "", -1))
}
