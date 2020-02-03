package gedcom_test

import (
	"github.com/elliotchance/gedcom/tag"
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
)

var familySearchIDNodeTests = map[string]struct {
	node  *gedcom.FamilySearchIDNode
	tag   tag.Tag
	value string
}{
	"1": {
		node:  gedcom.NewFamilySearchIDNode(tag.UnofficialTagFamilySearchID1, "LZDP-V7V"),
		tag:   tag.UnofficialTagFamilySearchID1,
		value: "LZDP-V7V",
	},
	"2": {
		node:  gedcom.NewFamilySearchIDNode(tag.UnofficialTagFamilySearchID2, "ZZDP-V7V"),
		tag:   tag.UnofficialTagFamilySearchID2,
		value: "ZZDP-V7V",
	},
}

func TestNewFamilySearchIDNode(t *testing.T) {
	node := gedcom.NewFamilySearchIDNode(tag.UnofficialTagFamilySearchID2, "LZDP-V7V")

	assert.NotNil(t, node)
	assert.IsType(t, node, (*gedcom.FamilySearchIDNode)(nil))
	assert.Equal(t, tag.UnofficialTagFamilySearchID2, node.Tag())
	assert.Equal(t, gedcom.Nodes(nil), node.Nodes())
	assert.Equal(t, "LZDP-V7V", node.Value())
	assert.Equal(t, "", node.Pointer())
}

func TestNewFamilySearchIDNode_String(t *testing.T) {
	for testName, test := range familySearchIDNodeTests {
		t.Run(testName, func(t *testing.T) {
			assert.Equal(t, test.value, test.node.String())
		})
	}
}

func TestFamilySearchIDNodeTags(t *testing.T) {
	assert.Equal(t, []tag.Tag{
		tag.UnofficialTagFamilySearchID1,
		tag.UnofficialTagFamilySearchID2,
	}, gedcom.FamilySearchIDNodeTags())
}
