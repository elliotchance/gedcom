package gedcom

import (
	"regexp"
	"strings"
)

// NameNode represents all the parts that make up a single name. An individual
// may have more than one name, each one would be represented by a NameNode.
type NameNode struct {
	*SimpleNode
}

func NewNameNode(value, pointer string, children []Node) *NameNode {
	return &NameNode{
		&SimpleNode{
			tag:      Name,
			value:    value,
			pointer:  pointer,
			children: children,
		},
	}
}

func (node *NameNode) parts() []string {
	return regexp.MustCompile("([^/]*)(/[^/]*/)?(.*)").
		FindStringSubmatch(node.value)
}

func (node *NameNode) trimSpaces(s string) string {
	// Run this twice to make sure we reduce odd numbers of spaces down.
	s = strings.Replace(s, "  ", " ", -1)
	s = strings.Replace(s, "  ", " ", -1)

	return strings.TrimSpace(s)
}

// GivenName is the given or earned name used for official identification of a
// person. It is also commonly known as the "first name".
func (node *NameNode) GivenName() string {
	// GivenName is the proper first name.
	givenNames := node.NodesWithTag(GivenName)
	if len(givenNames) > 0 {
		return node.trimSpaces(givenNames[0].Value())
	}

	// Fall back to trying to extract the first name from NAME tag.
	return node.trimSpaces(node.parts()[1])
}

// Surname is a family name passed on or used by members of a family.
func (node *NameNode) Surname() string {
	// Surname is the proper last name.
	surnames := node.NodesWithTag(Surname)
	if len(surnames) > 0 {
		return node.trimSpaces(surnames[0].Value())
	}

	// Fallback to trying to extract the surname from the NAME tag.
	lastName := node.trimSpaces(node.parts()[2])
	if lastName == "" {
		return ""
	}

	// The surname (if provided) will be wrapped within //.
	return lastName[1 : len(lastName)-1]
}

func (node *NameNode) Prefix() string {
	// NamePrefix is the proper name prefix. If it is not provided then no
	// prefix should be returned.
	namePrefixes := node.NodesWithTag(NamePrefix)
	if len(namePrefixes) > 0 {
		return node.trimSpaces(namePrefixes[0].Value())
	}

	return ""
}

func (node *NameNode) Suffix() string {
	// NameSuffix is the proper name suffix.
	nameSuffixes := node.NodesWithTag(NameSuffix)
	if len(nameSuffixes) > 0 {
		return node.trimSpaces(nameSuffixes[0].Value())
	}

	// Otherwise fallback to trying to extract it from the NAME.
	return node.trimSpaces(node.parts()[3])
}

func (node *NameNode) SurnamePrefix() string {
	// SurnameSuffix is the proper surname prefix.
	surnamePrefixes := node.NodesWithTag(SurnamePrefix)
	if len(surnamePrefixes) > 0 {
		return node.trimSpaces(surnamePrefixes[0].Value())
	}

	// Otherwise return nothing.
	return ""
}

func (node *NameNode) Title() string {
	// Title is the proper individual title.
	titles := node.NodesWithTag(Title)
	if len(titles) > 0 {
		return node.trimSpaces(titles[0].Value())
	}

	// Otherwise return nothing.
	return ""
}
