package gedcom

import "github.com/elliotchance/gedcom/tag"

// FamilySearchIDNode is the unique identifier for the person on
// FamilySearch.org. A FamilySearch ID always takes the form of:
//
//   LZDP-V9V
//
// There are several known tags that carry the FamilySearch ID:
//
//   _FID (UnofficialTagFamilySearchID1): Seen exported from MacFamilyFree.
//   _FSFTID (UnofficialTagFamilySearchID2): Some other applications.
//
type FamilySearchIDNode struct {
	*SimpleNode
}

func NewFamilySearchIDNode(tag tag.Tag, value string, children ...Node) *FamilySearchIDNode {
	return &FamilySearchIDNode{
		newSimpleNode(tag, value, "", children...),
	}
}

// FamilySearchIDNodeTags returns all of the known GEDCOM tags that can be
// represented by a FamilySearchIDNode.
func FamilySearchIDNodeTags() []tag.Tag {
	return []tag.Tag{
		tag.UnofficialTagFamilySearchID1,
		tag.UnofficialTagFamilySearchID2,
	}
}
