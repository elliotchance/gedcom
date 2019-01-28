package gedcom_test

import (
	"errors"
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

var uniqueIDNodeTests = map[string]struct {
	node     *gedcom.UniqueIDNode
	uuid     gedcom.UUID
	uuidErr  error
	checksum string
}{
	"Empty": {
		node:     gedcom.NewUniqueIDNode(nil, "", "", nil),
		uuid:     gedcom.UUID(""),
		uuidErr:  errors.New("invalid UUID: "),
		checksum: "",
	},
	"UUIDOnly": {
		node:     gedcom.NewUniqueIDNode(nil, "EE13561DDB204985BFFDEEBF82A5226C", "", nil),
		uuid:     gedcom.UUID("ee13561d-db20-4985-bffd-eebf82a5226c"),
		checksum: "",
	},
	"UUIDAndChecksum": {
		node:     gedcom.NewUniqueIDNode(nil, "EE13561DDB204985BFFDEEBF82A5226C5B2E", "", nil),
		uuid:     gedcom.UUID("ee13561d-db20-4985-bffd-eebf82a5226c"),
		checksum: "5B2E",
	},
	"GUIDOnly": {
		node:     gedcom.NewUniqueIDNode(nil, "{EE13561D-DB20-4985-BFFD-EEBF82A5226C}", "", nil),
		uuid:     gedcom.UUID("ee13561d-db20-4985-bffd-eebf82a5226c"),
		checksum: "",
	},
}

func TestNewUniqueIDNode(t *testing.T) {
	doc := gedcom.NewDocument()
	child := gedcom.NewNameNode(doc, "", "", nil)
	node := gedcom.NewUniqueIDNode(doc, "foo", "bar", []gedcom.Node{child})

	assert.NotNil(t, node)
	assert.IsType(t, node, (*gedcom.UniqueIDNode)(nil))
	assert.Equal(t, gedcom.UnofficialTagUniqueID, node.Tag())
	assert.Equal(t, []gedcom.Node{child}, node.Nodes())
	assert.Equal(t, doc, node.Document())
	assert.Equal(t, "foo", node.Value())
	assert.Equal(t, "bar", node.Pointer())
}

func TestUniqueIDNode_UUID(t *testing.T) {
	for testName, test := range uniqueIDNodeTests {
		t.Run(testName, func(t *testing.T) {
			uuid, err := test.node.UUID()

			if test.uuidErr != nil {
				assert.Equal(t, test.uuidErr, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.uuid, uuid)
			}
		})
	}
}

func TestUniqueIDNode_Checksum(t *testing.T) {
	for testName, test := range uniqueIDNodeTests {
		t.Run(testName, func(t *testing.T) {
			assert.Equal(t, test.checksum, test.node.Checksum())
		})
	}
}

func TestUniqueIDNode_Equals(t *testing.T) {
	for testName, test := range map[string]struct {
		n1, n2   *gedcom.UniqueIDNode
		expected bool
	}{
		"Equal1": {
			gedcom.NewUniqueIDNode(nil, "EE13561DDB204985BFFDEEBF82A5226C", "", nil),
			gedcom.NewUniqueIDNode(nil, "EE13561DDB204985BFFDEEBF82A5226C", "", nil),
			true,
		},
		"Equal2": {
			gedcom.NewUniqueIDNode(nil, "EE13561DDB204985BFFDEEBF82A5226C", "", nil),
			gedcom.NewUniqueIDNode(nil, "EE13561DDB204985BFFDEEBF82A5226C5B2E", "", nil),
			true,
		},
		"Equal3": {
			gedcom.NewUniqueIDNode(nil, "EE13561DDB204985BFFDEEBF82A5226C5B2E", "", nil),
			gedcom.NewUniqueIDNode(nil, "EE13561DDB204985BFFDEEBF82A5226C5B2E", "", nil),
			true,
		},
		"NotEqual1": {
			gedcom.NewUniqueIDNode(nil, "AE13561DDB204985BFFDEEBF82A5226C", "", nil),
			gedcom.NewUniqueIDNode(nil, "EE13561DDB204985BFFDEEBF82A5226C", "", nil),
			false,
		},
		"NotEqual2": {
			gedcom.NewUniqueIDNode(nil, "AE13561DDB204985BFFDEEBF82A5226C", "", nil),
			gedcom.NewUniqueIDNode(nil, "EE13561DDB204985BFFDEEBF82A5226C5B2E", "", nil),
			false,
		},
		"NotEqual3": {
			gedcom.NewUniqueIDNode(nil, "AE13561DDB204985BFFDEEBF82A5226C5B2E", "", nil),
			gedcom.NewUniqueIDNode(nil, "EE13561DDB204985BFFDEEBF82A5226C5B2E", "", nil),
			false,
		},
	} {
		t.Run(testName, func(t *testing.T) {
			assert.Equal(t, test.expected, test.n1.Equals(test.n2))
		})
	}
}
