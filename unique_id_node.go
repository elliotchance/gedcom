package gedcom

import (
	"github.com/elliotchance/gedcom/tag"
	"strings"
)

// UniqueIDNode represents a _UID.
//
// The _UID tag is what I call a Common GEDCOM Extension; a GEDCOM extension
// that's common to many different applications. The _UID tag is a legal GEDCOM
// extension; it starts with an underscore, like any vendor-defined extension
// should. However, it would be wrong to think of the _UID tag as a vendor- or
// product-specific tag. The _UID tag is not specific to any particular product
// or vendor. It is a common GEDCOM extension, supported by quite a few
// different products from different vendors.
//
// Applications that support the _UID tag include:
//
//   - Daub Ages!
//   - Ancestral Quest (AQ)
//   - Family Historian (FH)
//   - Family Origins for Windows (FOW)
//   - MyHeritage Family Tree Builder (FTB)
//   - Family Tree Heritage
//   - Family Tree Legends (FTL)
//   - Genbox Family History
//   - GenoPro
//   - Legacy Family Tree
//   - ohmiGene
//   - Personal Ancestral File (PAF)
//   - Reunion
//   - RootsMagic (RM)
//
// The above list was created by googling for GEDCOM files that contain the
// _UID tag. It probably isn't complete, but it does show that support for the
// _UID tag is not merely common, but also that it is supported by some of the
// best-known and most popular applications, such as PAF, RootsMagic and Legacy.
//
// — https://www.tamurajones.net/The_UIDTag.xhtml
type UniqueIDNode struct {
	*SimpleNode
}

func NewUniqueIDNode(value string, children ...Node) *UniqueIDNode {
	return &UniqueIDNode{
		newSimpleNode(tag.UnofficialTagUniqueID, value, "", children...),
	}
}

// UUID returns the UUID component of the UniqueID.
//
// The cross-reference identifiers are unique with a single GEDCOM file, but
// different GEDCOM files, are very likely to use the same identifiers for
// different records.
//
// For some purposes it would be nice to have truly unique identifiers. That
// requires two things; some way to create globally unique identifiers, and a
// GEDCOM tag to carry that identifier. The Universally Unique ID (UUID) is that
// globally unique identifier, and _UID is the tag that carries it.
//
// Microsoft developers known UUID as Globally Unique Identifier (GUID). What
// makes UUIDs so suitable is that they were developed so precisely so that
// everyone can generate UUIDs, without coordinating with anyone else, and still
// be practically sure that the generated identifier is unique.
//
// A UUID is an 128-bit (16-byte) number, generally represented by 32
// hexadecimal digits, divided into five groups, separated by hyphens, like
// this: 12345678-1234-1234-1234-123456789ABC.
//
// — https://www.tamurajones.net/The_UIDTag.xhtml
func (node *UniqueIDNode) UUID() (UUID, error) {
	value := node.cleanValue()

	if len(value) >= 32 {
		return NewUUIDFromString(value[:32])
	}

	return NewUUIDFromString(value)
}

// Checksum returns the optional last 4 digits after the UUID.
//
// This is what a UUID looks like in a PAF GEDCOM:
//
//   92FF8B766F327F48A256C3AE6DAE50D3A114
//
// Notice that the _UID value is not divided into groups separated by hyphens,
// but represented as one long hexadecimal number. What is not immediately
// obvious because of its length is that the number shown is not a UUID. It
// cannot be an UUID, because an UUID is 32 hexadecimal digits long, and the
// _UID value is 36 hexadecimal digits long. When we hyphenate the number for
// readability, we find that there are four extra digits:
//
//   92FF8B76-6F32-7F48-A256-C3AE6DAE50D3-A114
//
// The _UID value is a UUID followed by a checksum.
//
// That little fact, essential to making sense of the _UID value, used to be
// undocumented. I once figured that out myself, but you don't have to do so,
// nor take my word for it. Nowadays, FamilySearch documents it, if you know
// where to look. The FamilySearch document GEDCOM Unique Identifiers states
// that it provides guidelines for the use of UUIDs, but it actually documents
// the format of the PAF _UID value; a 32-hexadigit UUID value followed by a
// 4-hexadigit checksum. That brief document includes some Windows C code,
// possibly the actual PAF source code. That code shows how to calculate the
// checksum.
//
// PAF is a fork of Ancestral Quest, so it is no wonder that Ancestral Quest
// uses the same format. But it is not just Ancestral Quest that uses this
// format. This UUID format is the most popular one. Other applications that use
// the same UUID format for their _UID tag are Family Origins for Windows,
// Family Tree Heritage, Family Tree Legends, Genbox Family History, Legacy
// Family Tree, Reunion and RootsMagic.
//
// — https://www.tamurajones.net/The_UIDTag.xhtml
func (node *UniqueIDNode) Checksum() string {
	value := node.cleanValue()

	if len(value) > 32 {
		return value[32:]
	}

	return ""
}

// Equals returns true if both nodes have the same UUID value. The checksum (if
// any) is ignored.
//
// If either nodes are nil (or both) the result will always be false.
func (node *UniqueIDNode) Equals(node2 Node) bool {
	if IsNil(node) {
		return false
	}

	if IsNil(node2) {
		return false
	}

	if n2, ok := node2.(*UniqueIDNode); ok {
		u1, err1 := node.UUID()
		u2, err2 := n2.UUID()

		if err1 != nil || err2 != nil {
			return false
		}

		return u1.Equals(u2)
	}

	return false
}

func (node UniqueIDNode) cleanValue() string {
	mutS := node.Value()
	mutS = strings.Replace(mutS, "{", "", -1)
	mutS = strings.Replace(mutS, "}", "", -1)
	mutS = strings.Replace(mutS, "-", "", -1)

	return mutS
}
