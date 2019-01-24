package gedcom_test

import (
	"strings"
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
)

var tests = map[string]*gedcom.Document{
	"":     gedcom.NewDocumentWithNodes(nil),
	"\n\n": gedcom.NewDocumentWithNodes(nil),
	"0 HEAD": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNode(nil, gedcom.TagHeader, "", ""),
	}),
	"0 HEAD\n1 CHAR UTF-8": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagHeader, "", "", []gedcom.Node{
			gedcom.NewNode(nil, gedcom.TagCharacterSet, "UTF-8", ""),
		}),
	}),
	"0 HEAD\n\n1 CHAR UTF-8\n": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagHeader, "", "", []gedcom.Node{
			gedcom.NewNode(nil, gedcom.TagCharacterSet, "UTF-8", ""),
		}),
	}),
	"0 HEAD\n1 CHAR UTF-8\n1 SOUR Ancestry.com Family Trees": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagHeader, "", "", []gedcom.Node{
			gedcom.NewNode(nil, gedcom.TagCharacterSet, "UTF-8", ""),
			gedcom.NewSourceNode(nil, "Ancestry.com Family Trees", "", nil),
		}),
	}),
	"0 HEAD\n1 CHAR UTF-8\n1 CHAR UTF-8": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagHeader, "", "", []gedcom.Node{
			gedcom.NewNode(nil, gedcom.TagCharacterSet, "UTF-8", ""),
			gedcom.NewNode(nil, gedcom.TagCharacterSet, "UTF-8", ""),
		}),
	}),
	"0 HEAD\n1 SOUR Ancestry.com Family Trees": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagHeader, "", "", []gedcom.Node{
			gedcom.NewSourceNode(nil, "Ancestry.com Family Trees", "", nil),
		}),
	}),
	"0 HEAD\n1 BIRT": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagHeader, "", "", []gedcom.Node{
			gedcom.NewBirthNode(nil, "", "", nil),
		}),
	}),
	"0 HEAD\n1 GEDC\n2 VERS (2010.3)": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagHeader, "", "", []gedcom.Node{
			gedcom.NewNodeWithChildren(nil, gedcom.TagGedcomInformation, "", "", []gedcom.Node{
				gedcom.NewNode(nil, gedcom.TagVersion, "(2010.3)", ""),
			}),
		}),
	}),
	"0 HEAD\n1 GEDC\n2 VERS 5.5": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagHeader, "", "", []gedcom.Node{
			gedcom.NewNodeWithChildren(nil, gedcom.TagGedcomInformation, "", "", []gedcom.Node{
				gedcom.NewNode(nil, gedcom.TagVersion, "5.5", ""),
			}),
		}),
	}),
	"0 HEAD\n1 GEDC\n2 FORM LINEAGE-LINKED": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagHeader, "", "", []gedcom.Node{
			gedcom.NewNodeWithChildren(nil, gedcom.TagGedcomInformation, "", "", []gedcom.Node{
				gedcom.NewFormatNode(nil, "LINEAGE-LINKED", "", nil),
			}),
		}),
	}),
	"0 HEAD\n1 BIRT\n2 PLAC Camperdown, Nsw, Australia": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagHeader, "", "", []gedcom.Node{
			gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
				gedcom.NewPlaceNode(nil, "Camperdown, Nsw, Australia", "", nil),
			}),
		}),
	}),
	"0 HEAD\n1 NAME Elliot Rupert de Peyster /Chance/": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagHeader, "", "", []gedcom.Node{
			gedcom.NewNameNode(nil, "Elliot Rupert de Peyster /Chance/", "", nil),
		}),
	}),
	"0 HEAD\n0 @P1@ INDI": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagHeader, "", "", nil),
		gedcom.NewIndividualNode(nil, "", "P1", nil),
	}),
	"0 HEAD\n1 SEX M\n0 @P1@ INDI": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagHeader, "", "", []gedcom.Node{
			gedcom.NewNode(nil, gedcom.TagSex, "M", ""),
		}),
		gedcom.NewIndividualNode(nil, "", "P1", nil),
	}),
	"0 HEAD\n1 SEX M": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagHeader, "", "", []gedcom.Node{
			gedcom.NewNode(nil, gedcom.TagSex, "M", ""),
		}),
	}),
	"0 HEAD\n1 BIRT\n2 PLAC Camperdown, Nsw, Australia\n1 SEX M": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagHeader, "", "", []gedcom.Node{
			gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
				gedcom.NewPlaceNode(nil, "Camperdown, Nsw, Australia", "", nil),
			}),
			gedcom.NewNode(nil, gedcom.TagSex, "M", ""),
		}),
	}),
	"0 HEAD\n0 @P1@ INDI\n1 BIRT": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNode(nil, gedcom.TagHeader, "", ""),
		gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
			gedcom.NewBirthNode(nil, "", "", nil),
		}),
	}),
	"0 HEAD\n1 GEDC\n2 FORM LINEAGE-LINKED\n0 @P1@ INDI": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagHeader, "", "", []gedcom.Node{
			gedcom.NewNodeWithChildren(nil, gedcom.TagGedcomInformation, "", "", []gedcom.Node{
				gedcom.NewFormatNode(nil, "LINEAGE-LINKED", "", nil),
			}),
		}),
		gedcom.NewIndividualNode(nil, "", "P1", nil),
	}),
	"0 HEAD0\n1 HEAD1\n2 HEAD2\n3 HEAD3\n0 HEAD00": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagFromString("HEAD0"), "", "", []gedcom.Node{
			gedcom.NewNodeWithChildren(nil, gedcom.TagFromString("HEAD1"), "", "", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagFromString("HEAD2"), "", "", []gedcom.Node{
					gedcom.NewNode(nil, gedcom.TagFromString("HEAD3"), "", ""),
				}),
			}),
		}),
		gedcom.NewNode(nil, gedcom.TagFromString("HEAD00"), "", ""),
	}),
	"0 HEAD0\n1 HEAD1\n2 HEAD2\n3 HEAD3\n1 HEAD10": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagFromString("HEAD0"), "", "", []gedcom.Node{
			gedcom.NewNodeWithChildren(nil, gedcom.TagFromString("HEAD1"), "", "", []gedcom.Node{
				gedcom.NewNodeWithChildren(nil, gedcom.TagFromString("HEAD2"), "", "", []gedcom.Node{
					gedcom.NewNode(nil, gedcom.TagFromString("HEAD3"), "", ""),
				}),
			}),
			gedcom.NewNode(nil, gedcom.TagFromString("HEAD10"), "", ""),
		}),
	}),
	"0 HEAD0\r1 HEAD1": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagFromString("HEAD0"), "", "", []gedcom.Node{
			gedcom.NewNode(nil, gedcom.TagFromString("HEAD1"), "", ""),
		}),
	}),
	"0 HEAD0\r\n1 HEAD1": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagFromString("HEAD0"), "", "", []gedcom.Node{
			gedcom.NewNode(nil, gedcom.TagFromString("HEAD1"), "", ""),
		}),
	}),
	"0 HEAD0\n1 HEAD1\n1 HEAD10\n2 HEAD2": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagFromString("HEAD0"), "", "", []gedcom.Node{
			gedcom.NewNode(nil, gedcom.TagFromString("HEAD1"), "", ""),
			gedcom.NewNodeWithChildren(nil, gedcom.TagFromString("HEAD10"), "", "", []gedcom.Node{
				gedcom.NewNode(nil, gedcom.TagFromString("HEAD2"), "", ""),
			}),
		}),
	}),
	"0 HEAD\n1 BIRT ": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNodeWithChildren(nil, gedcom.TagHeader, "", "", []gedcom.Node{
			gedcom.NewBirthNode(nil, "", "", nil),
		}),
	}),
	"0 @P221@ INDI\n1 BIRT\n2 DATE 1851\n1 DEAT\n2 DATE 1856": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewIndividualNode(nil, "", "P221", []gedcom.Node{
			gedcom.NewBirthNode(nil, "", "", []gedcom.Node{
				gedcom.NewDateNode(nil, "1851", "", nil),
			}),
			gedcom.NewNodeWithChildren(nil, gedcom.TagDeath, "", "", []gedcom.Node{
				gedcom.NewDateNode(nil, "1856", "", nil),
			}),
		}),
	}),
	"0 @F1@ FAM\n1 HUSB @P2@\n1 WIFE @P3@": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewFamilyNode(nil, "F1", []gedcom.Node{
			gedcom.NewNode(nil, gedcom.TagHusband, "@P2@", ""),
			gedcom.NewNode(nil, gedcom.TagWife, "@P3@", ""),
		}),
	}),
	"0 DATE 1856": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewDateNode(nil, "1856", "", nil),
	}),
	"0 NAME κόσμε": gedcom.NewDocumentWithNodes([]gedcom.Node{
		gedcom.NewNameNode(nil, "κόσμε", "", nil),
	}),
}

func TestDecoder_Decode(t *testing.T) {
	for ged, expected := range tests {
		t.Run("", func(t *testing.T) {
			decoder := gedcom.NewDecoder(strings.NewReader(ged))

			actual, err := decoder.Decode()
			assert.NoError(t, err, ged)

			expected.MaxLivingAge = gedcom.DefaultMaxLivingAge

			for _, n := range expected.Nodes() {
				n.SetDocument(expected)
			}

			assertDocumentEqual(t, expected, actual, ged)
		})
	}

	t.Run("BOM", func(t *testing.T) {
		ged := "\xEF\xBB\xBF0 HEAD\n1 CHAR UTF-8"
		decoder := gedcom.NewDecoder(strings.NewReader(ged))
		expected := gedcom.NewDocumentWithNodes([]gedcom.Node{
			gedcom.NewNodeWithChildren(nil, gedcom.TagHeader, "", "", []gedcom.Node{
				gedcom.NewNode(nil, gedcom.TagCharacterSet, "UTF-8", ""),
			}),
		})
		expected.HasBOM = true

		actual, err := decoder.Decode()
		assert.NoError(t, err, ged)

		expected.MaxLivingAge = gedcom.DefaultMaxLivingAge

		for _, n := range expected.Nodes() {
			n.SetDocument(expected)
		}

		assertDocumentEqual(t, expected, actual, ged)
	})
}

func trimSpaces(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, "\r", "\n", -1)
	s = strings.Replace(s, "\n\n", "\n", -1)

	return s
}

func TestDocument_String(t *testing.T) {
	for expected, actual := range tests {
		t.Run("", func(t *testing.T) {
			newGed := actual.String()
			assert.Equal(t, trimSpaces(expected), trimSpaces(newGed), expected)
		})
	}
}

func TestDocument_GEDCOMString(t *testing.T) {
	for expected, actual := range tests {
		t.Run("", func(t *testing.T) {
			newGed := actual.GEDCOMString(0)
			assert.Equal(t, trimSpaces(expected), trimSpaces(newGed), expected)
		})
	}
}

func TestNewNode(t *testing.T) {
	const p = "pointer"
	const v = "value"

	for _, test := range []struct {
		tag      gedcom.Tag
		expected gedcom.Node
	}{
		{gedcom.TagBaptism, gedcom.NewBaptismNode(nil, v, p, nil)},
		{gedcom.TagBirth, gedcom.NewBirthNode(nil, v, p, nil)},
		{gedcom.TagBurial, gedcom.NewBurialNode(nil, v, p, nil)},
		{gedcom.TagDate, gedcom.NewDateNode(nil, v, p, nil)},
		{gedcom.TagDeath, gedcom.NewDeathNode(nil, v, p, nil)},
		{gedcom.TagEvent, gedcom.NewEventNode(nil, v, p, nil)},
		{gedcom.TagFamily, gedcom.NewFamilyNode(nil, p, nil)},
		{gedcom.UnofficialTagFamilySearchID1, gedcom.NewFamilySearchIDNode(nil, gedcom.UnofficialTagFamilySearchID1, v)},
		{gedcom.UnofficialTagFamilySearchID2, gedcom.NewFamilySearchIDNode(nil, gedcom.UnofficialTagFamilySearchID2, v)},
		{gedcom.TagFormat, gedcom.NewFormatNode(nil, v, p, nil)},
		{gedcom.TagIndividual, gedcom.NewIndividualNode(nil, v, p, nil)},
		{gedcom.TagLatitude, gedcom.NewLatitudeNode(nil, v, p, nil)},
		{gedcom.TagLongitude, gedcom.NewLongitudeNode(nil, v, p, nil)},
		{gedcom.TagMap, gedcom.NewMapNode(nil, v, p, nil)},
		{gedcom.TagName, gedcom.NewNameNode(nil, v, p, nil)},
		{gedcom.TagNote, gedcom.NewNoteNode(nil, v, p, nil)},
		{gedcom.TagNickname, gedcom.NewNicknameNode(nil, v, p, nil)},
		{gedcom.TagPhonetic, gedcom.NewPhoneticVariationNode(nil, v, p, nil)},
		{gedcom.TagPlace, gedcom.NewPlaceNode(nil, v, p, nil)},
		{gedcom.TagResidence, gedcom.NewResidenceNode(nil, v, p, nil)},
		{gedcom.TagRomanized, gedcom.NewRomanizedVariationNode(nil, v, p, nil)},
		{gedcom.TagSource, gedcom.NewSourceNode(nil, v, p, nil)},
		{gedcom.TagType, gedcom.NewTypeNode(nil, v, p, nil)},
		{gedcom.TagVersion, gedcom.NewNode(nil, gedcom.TagVersion, v, p)},
		{gedcom.UnofficialTagUniqueID, gedcom.NewUniqueIDNode(nil, v, p, nil)},
	} {
		t.Run(test.tag.String(), func(t *testing.T) {
			assert.Equal(t, test.expected, gedcom.NewNode(nil, test.tag, v, p))
		})
	}

	t.Run("ImplementsGEDCOMStringer", func(t *testing.T) {
		node := gedcom.NewDateNode(nil, "", "", nil)
		assert.Implements(t, (*gedcom.GEDCOMStringer)(nil), node)
	})

	t.Run("ImplementsGEDCOMLiner", func(t *testing.T) {
		node := gedcom.NewDateNode(nil, "", "", nil)
		assert.Implements(t, (*gedcom.GEDCOMLiner)(nil), node)
	})
}
