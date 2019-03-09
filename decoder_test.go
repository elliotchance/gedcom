package gedcom_test

import (
	"strings"
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
)

var decoderTests = map[string]struct {
	ged      string
	expected func(*gedcom.Document)
}{
	"Empty": {
		"",
		func(doc *gedcom.Document) {},
	},
	"OnlyNewLine": {
		"\n\n",
		func(doc *gedcom.Document) {},
	},
	"OnlyRoot": {
		"0 HEAD",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagHeader, "", ""))
		},
	},
	"OneChild1": {
		"0 HEAD\n1 CHAR UTF-8",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagHeader, "", "",
				gedcom.NewNode(gedcom.TagCharacterSet, "UTF-8", ""),
			))
		},
	},
	"OneChild2": {
		"0 HEAD\n\n1 CHAR UTF-8\n",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagHeader, "", "",
				gedcom.NewNode(gedcom.TagCharacterSet, "UTF-8", ""),
			))
		},
	},
	"TwoChildren": {
		"0 HEAD\n1 CHAR UTF-8\n1 @S1@ SOUR Ancestry.com Family Trees",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagHeader, "", "",
				gedcom.NewNode(gedcom.TagCharacterSet, "UTF-8", ""),
				gedcom.NewSourceNode("Ancestry.com Family Trees", "S1"),
			))
		},
	},
	"TwoIdenticalChildren": {
		"0 HEAD\n1 CHAR UTF-8\n1 CHAR UTF-8",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagHeader, "", "",
				gedcom.NewNode(gedcom.TagCharacterSet, "UTF-8", ""),
				gedcom.NewNode(gedcom.TagCharacterSet, "UTF-8", ""),
			))
		},
	},
	"SpaceInValue": {
		"0 HEAD\n1 @S1@ SOUR Ancestry.com Family Trees",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagHeader, "", "",
				gedcom.NewSourceNode("Ancestry.com Family Trees", "S1"),
			))
		},
	},
	"SimpleChild": {
		"0 HEAD\n1 BIRT",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagHeader, "", "",
				gedcom.NewBirthNode(""),
			))
		},
	},
	"ThreeDeep1": {
		"0 HEAD\n1 GEDC\n2 VERS (2010.3)",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagHeader, "", "",
				gedcom.NewNode(gedcom.TagGedcomInformation, "", "",
					gedcom.NewNode(gedcom.TagVersion, "(2010.3)", ""),
				),
			))
		},
	},
	"ThreeDeep2": {
		"0 HEAD\n1 GEDC\n2 VERS 5.5",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagHeader, "", "",
				gedcom.NewNode(gedcom.TagGedcomInformation, "", "",
					gedcom.NewNode(gedcom.TagVersion, "5.5", ""),
				),
			))
		},
	},
	"FORM": {
		"0 HEAD\n1 GEDC\n2 FORM LINEAGE-LINKED",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagHeader, "", "",
				gedcom.NewNode(gedcom.TagGedcomInformation, "", "",
					gedcom.NewFormatNode("LINEAGE-LINKED"),
				),
			))
		},
	},
	"PLAC": {
		"0 HEAD\n1 BIRT\n2 PLAC Camperdown, Nsw, Australia",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagHeader, "", "",
				gedcom.NewBirthNode("",
					gedcom.NewPlaceNode("Camperdown, Nsw, Australia"),
				),
			))
		},
	},
	"NAME": {
		"0 HEAD\n1 NAME Elliot Rupert de Peyster /Chance/",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagHeader, "", "",
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
			))
		},
	},
	"Pointer": {
		"0 HEAD\n0 @P1@ INDI",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagHeader, "", ""))
			doc.AddIndividual("P1")
		},
	},
	"NestedPointer": {
		"0 HEAD\n1 SEX M\n0 @P1@ INDI",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagHeader, "", "",
				gedcom.NewNode(gedcom.TagSex, "M", ""),
			))
			doc.AddIndividual("P1")
		},
	},
	"SEX": {
		"0 HEAD\n1 SEX M",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagHeader, "", "",
				gedcom.NewNode(gedcom.TagSex, "M", ""),
			))
		},
	},
	"Nested1": {
		"0 HEAD\n1 BIRT\n2 PLAC Camperdown, Nsw, Australia\n1 SEX M",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagHeader, "", "",
				gedcom.NewBirthNode("",
					gedcom.NewPlaceNode("Camperdown, Nsw, Australia"),
				),
				gedcom.NewNode(gedcom.TagSex, "M", ""),
			))
		},
	},
	"MultipleRoots": {
		"0 HEAD\n0 @P1@ INDI\n1 BIRT",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagHeader, "", ""))
			doc.AddIndividual("P1",
				gedcom.NewBirthNode(""),
			)
		},
	},
	"Nested2": {
		"0 HEAD\n1 GEDC\n2 FORM LINEAGE-LINKED\n0 @P1@ INDI",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagHeader, "", "",
				gedcom.NewNode(gedcom.TagGedcomInformation, "", "",
					gedcom.NewFormatNode("LINEAGE-LINKED"),
				),
			))
			doc.AddIndividual("P1")
		},
	},
	"Nested3": {
		"0 HEAD0\n1 HEAD1\n2 HEAD2\n3 HEAD3\n0 HEAD00",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagFromString("HEAD0"), "", "",
				gedcom.NewNode(gedcom.TagFromString("HEAD1"), "", "",
					gedcom.NewNode(gedcom.TagFromString("HEAD2"), "", "",
						gedcom.NewNode(gedcom.TagFromString("HEAD3"), "", ""),
					),
				),
			))
			doc.AddNode(gedcom.NewNode(gedcom.TagFromString("HEAD00"), "", ""))
		},
	},
	"Nested4": {
		"0 HEAD0\n1 HEAD1\n2 HEAD2\n3 HEAD3\n1 HEAD10",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagFromString("HEAD0"), "", "",
				gedcom.NewNode(gedcom.TagFromString("HEAD1"), "", "",
					gedcom.NewNode(gedcom.TagFromString("HEAD2"), "", "",
						gedcom.NewNode(gedcom.TagFromString("HEAD3"), "", ""),
					),
				),
				gedcom.NewNode(gedcom.TagFromString("HEAD10"), "", ""),
			))
		},
	},
	"BadTag1": {
		"0 HEAD0\r1 HEAD1",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagFromString("HEAD0"), "", "",
				gedcom.NewNode(gedcom.TagFromString("HEAD1"), "", ""),
			))
		},
	},
	"WindowsLineEnding": {
		"0 HEAD0\r\n1 HEAD1",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagFromString("HEAD0"), "", "",
				gedcom.NewNode(gedcom.TagFromString("HEAD1"), "", ""),
			))
		},
	},
	"BadTag2": {
		"0 HEAD0\n1 HEAD1\n1 HEAD10\n2 HEAD2",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagFromString("HEAD0"), "", "",
				gedcom.NewNode(gedcom.TagFromString("HEAD1"), "", ""),
				gedcom.NewNode(gedcom.TagFromString("HEAD10"), "", "",
					gedcom.NewNode(gedcom.TagFromString("HEAD2"), "", ""),
				),
			))
		},
	},
	"SpaceWithoutValue": {
		"0 HEAD\n1 BIRT ",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagHeader, "", "",
				gedcom.NewBirthNode(""),
			))
		},
	},
	"Nested5": {
		"0 @P221@ INDI\n1 BIRT\n2 DATE 1851\n1 DEAT\n2 DATE 1856",
		func(doc *gedcom.Document) {
			doc.AddIndividual("P221",
				gedcom.NewBirthNode("",
					gedcom.NewDateNode("1851"),
				),
				gedcom.NewNode(gedcom.TagDeath, "", "",
					gedcom.NewDateNode("1856"),
				),
			)
		},
	},
	"HUSB_WIFE": {
		"0 @F1@ FAM\n1 HUSB @P2@\n1 WIFE @P3@",
		func(doc *gedcom.Document) {
			f1 := doc.AddFamily("F1")
			f1.SetHusbandPointer("P2")
			f1.SetWifePointer("P3")
		},
	},
	"RootDate": {
		"0 DATE 1856",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewDateNode("1856"))
		},
	},
	"UTF-8": {
		"0 NAME κόσμε",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNameNode("κόσμε"))
		},
	},
	"NonAlphanumericCharactersAreAllowed": {
		"0 @R-1577718385@ FAM",
		func(doc *gedcom.Document) {
			doc.AddFamily("R-1577718385")
		},
	},
	"CrazyLongPointerValuesAreAlsoFine": {
		"0 @SomeReallyLongPointerThatShouldProbablyNotBeUsedByWeShouldDemonstrateThatThereIsNoLimitToTheLength@ FAM",
		func(doc *gedcom.Document) {
			doc.AddFamily("SomeReallyLongPointerThatShouldProbablyNotBeUsedByWeShouldDemonstrateThatThereIsNoLimitToTheLength")
		},
	},
	"AnyPunctuationIsPermittedAsLongAsItIsntAn@": {
		"0 @~!-#$%^&*()@ FAM",
		func(doc *gedcom.Document) {
			doc.AddFamily("~!-#$%^&*()")
		},
	},
	"NonGreedyConsume": {
		"0 @uh@ NAME oh@ok",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(gedcom.TagName, "oh@ok", "uh"))
		},
	},
}

func TestDecoder_Decode(t *testing.T) {
	for testName, test := range decoderTests {
		t.Run(testName, func(t *testing.T) {
			decoder := gedcom.NewDecoder(strings.NewReader(test.ged))

			actual, err := decoder.Decode()
			assert.NoError(t, err, test.ged)

			doc := gedcom.NewDocument()
			test.expected(doc)

			assertDocumentEqual(t, doc, actual)
		})

		t.Run(testName+"WithAllowMultiLine", func(t *testing.T) {
			decoder := gedcom.NewDecoder(strings.NewReader(test.ged))
			decoder.AllowMultiLine = true

			actual, err := decoder.Decode()
			assert.NoError(t, err, test.ged)

			doc := gedcom.NewDocument()
			test.expected(doc)

			assertDocumentEqual(t, doc, actual)
		})
	}

	t.Run("DoubleSpace", func(t *testing.T) {
		// Issue #243

		decoder := gedcom.NewDecoder(strings.NewReader("0  _PUBLISH"))

		actual, err := decoder.Decode()
		assert.NoError(t, err)

		doc := gedcom.NewDocument()
		doc.AddNode(gedcom.NewNode(gedcom.TagFromString("_PUBLISH"), "", ""))

		assertDocumentEqual(t, doc, actual)
	})

	t.Run("BOM", func(t *testing.T) {
		ged := "\xEF\xBB\xBF0 HEAD\n1 CHAR UTF-8"
		decoder := gedcom.NewDecoder(strings.NewReader(ged))
		expected := gedcom.NewDocumentWithNodes(gedcom.Nodes{
			gedcom.NewNode(gedcom.TagHeader, "", "",
				gedcom.NewNode(gedcom.TagCharacterSet, "UTF-8", ""),
			),
		})
		expected.HasBOM = true

		actual, err := decoder.Decode()
		assert.NoError(t, err, ged)

		expected.MaxLivingAge = gedcom.DefaultMaxLivingAge

		assertDocumentEqual(t, expected, actual, ged)
	})

	t.Run("AllowMultiLine", func(t *testing.T) {
		// It is not valid for GEDCOM values to contain new lines or carriage
		// returns. However, some application dump data without correctly using
		// the CONT tags.
		//
		// Strictly speaking we should bail out with an error but there are too
		// many cases that are difficult to clean up for consumers so we offer
		// and option to permit it.

		decoder := gedcom.NewDecoder(strings.NewReader("0 TEXT Karen Marie McMillan<br>Gender: Female<br>Birth: Aug 21 1957 - \n\n\t&nbsp;Texas, USA"))
		decoder.AllowMultiLine = true

		actual, err := decoder.Decode()
		if assert.NoError(t, err) {
			doc := gedcom.NewDocument()
			doc.AddNode(gedcom.NewNode(gedcom.TagFromString("TEXT"), "Karen Marie McMillan<br>Gender: Female<br>Birth: Aug 21 1957 - \n\n\t&nbsp;Texas, USA", ""))

			assertDocumentEqual(t, doc, actual)
		}
	})
}

func trimSpaces(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, "\r", "\n", -1)
	s = strings.Replace(s, "\n\n", "\n", -1)

	return s
}

func TestDocument_String(t *testing.T) {
	for testName, test := range decoderTests {
		t.Run(testName, func(t *testing.T) {
			doc := gedcom.NewDocument()
			test.expected(doc)
			newGed := doc.String()
			assert.Equal(t, trimSpaces(test.ged), trimSpaces(newGed))
		})
	}
}

func TestDocument_GEDCOMString(t *testing.T) {
	for testName, test := range decoderTests {
		t.Run(testName, func(t *testing.T) {
			doc := gedcom.NewDocument()
			test.expected(doc)
			newGed := doc.GEDCOMString(0)
			assert.Equal(t, trimSpaces(test.ged), trimSpaces(newGed))
		})
	}
}

func TestNewNode(t *testing.T) {
	const p = "pointer"
	const v = "value"

	// Nodes that require a document like Individual and Family are not included
	// in this tests.
	for _, test := range []struct {
		tag      gedcom.Tag
		expected gedcom.Node
	}{
		{gedcom.TagBaptism, gedcom.NewBaptismNode(v)},
		{gedcom.TagBirth, gedcom.NewBirthNode(v)},
		{gedcom.TagBurial, gedcom.NewBurialNode(v)},
		{gedcom.TagDate, gedcom.NewDateNode(v)},
		{gedcom.TagDeath, gedcom.NewDeathNode(v)},
		{gedcom.TagEvent, gedcom.NewEventNode(v)},
		{gedcom.UnofficialTagFamilySearchID1, gedcom.NewFamilySearchIDNode(gedcom.UnofficialTagFamilySearchID1, v)},
		{gedcom.UnofficialTagFamilySearchID2, gedcom.NewFamilySearchIDNode(gedcom.UnofficialTagFamilySearchID2, v)},
		{gedcom.TagFormat, gedcom.NewFormatNode(v)},
		{gedcom.TagLatitude, gedcom.NewLatitudeNode(v)},
		{gedcom.TagLongitude, gedcom.NewLongitudeNode(v)},
		{gedcom.TagMap, gedcom.NewMapNode(v)},
		{gedcom.TagName, gedcom.NewNameNode(v)},
		{gedcom.TagNote, gedcom.NewNoteNode(v)},
		{gedcom.TagNickname, gedcom.NewNicknameNode(v)},
		{gedcom.TagPhonetic, gedcom.NewPhoneticVariationNode(v)},
		{gedcom.TagPlace, gedcom.NewPlaceNode(v)},
		{gedcom.TagResidence, gedcom.NewResidenceNode(v)},
		{gedcom.TagRomanized, gedcom.NewRomanizedVariationNode(v)},
		{gedcom.TagSource, gedcom.NewSourceNode(v, p)},
		{gedcom.TagType, gedcom.NewTypeNode(v)},
		{gedcom.TagVersion, gedcom.NewNode(gedcom.TagVersion, v, p)},
		{gedcom.UnofficialTagUniqueID, gedcom.NewUniqueIDNode(v)},
	} {
		t.Run(test.tag.String(), func(t *testing.T) {
			assertEqual(t, test.expected, gedcom.NewNode(test.tag, v, p))
		})
	}

	t.Run("ImplementsGEDCOMStringer", func(t *testing.T) {
		node := gedcom.NewDateNode("")
		assert.Implements(t, (*gedcom.GEDCOMStringer)(nil), node)
	})

	t.Run("ImplementsGEDCOMLiner", func(t *testing.T) {
		node := gedcom.NewDateNode("")
		assert.Implements(t, (*gedcom.GEDCOMLiner)(nil), node)
	})
}
