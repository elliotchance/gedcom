package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
)

var familySearchIDNodeTests = map[string]struct {
	node  *gedcom.FamilySearchIDNode
	tag   gedcom.Tag
	value string
}{
	"1": {
		node:  gedcom.NewFamilySearchIDNode(gedcom.UnofficialTagFamilySearchID1, "LZDP-V7V"),
		tag:   gedcom.UnofficialTagFamilySearchID1,
		value: "LZDP-V7V",
	},
	"2": {
		node:  gedcom.NewFamilySearchIDNode(gedcom.UnofficialTagFamilySearchID2, "ZZDP-V7V"),
		tag:   gedcom.UnofficialTagFamilySearchID2,
		value: "ZZDP-V7V",
	},
}

func TestNewFamilySearchIDNode(t *testing.T) {
	node := gedcom.NewFamilySearchIDNode(gedcom.UnofficialTagFamilySearchID2, "LZDP-V7V")

	assert.NotNil(t, node)
	assert.IsType(t, node, (*gedcom.FamilySearchIDNode)(nil))
	assert.Equal(t, gedcom.UnofficialTagFamilySearchID2, node.Tag())
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
	assert.Equal(t, []gedcom.Tag{
		gedcom.UnofficialTagFamilySearchID1,
		gedcom.UnofficialTagFamilySearchID2,
	}, gedcom.FamilySearchIDNodeTags())
}
