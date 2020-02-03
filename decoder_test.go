package gedcom_test

import (
	"github.com/elliotchance/gedcom/tag"
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
			doc.AddNode(gedcom.NewNode(tag.TagHeader, "", ""))
		},
	},
	"OneChild1": {
		"0 HEAD\n1 CHAR UTF-8",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagHeader, "", "",
				gedcom.NewNode(tag.TagCharacterSet, "UTF-8", ""),
			))
		},
	},
	"OneChild2": {
		"0 HEAD\n\n1 CHAR UTF-8\n",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagHeader, "", "",
				gedcom.NewNode(tag.TagCharacterSet, "UTF-8", ""),
			))
		},
	},
	"TwoChildren": {
		"0 HEAD\n1 CHAR UTF-8\n1 @S1@ SOUR Ancestry.com Family Trees",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagHeader, "", "",
				gedcom.NewNode(tag.TagCharacterSet, "UTF-8", ""),
				gedcom.NewSourceNode("Ancestry.com Family Trees", "S1"),
			))
		},
	},
	"TwoIdenticalChildren": {
		"0 HEAD\n1 CHAR UTF-8\n1 CHAR UTF-8",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagHeader, "", "",
				gedcom.NewNode(tag.TagCharacterSet, "UTF-8", ""),
				gedcom.NewNode(tag.TagCharacterSet, "UTF-8", ""),
			))
		},
	},
	"SpaceInValue": {
		"0 HEAD\n1 @S1@ SOUR Ancestry.com Family Trees",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagHeader, "", "",
				gedcom.NewSourceNode("Ancestry.com Family Trees", "S1"),
			))
		},
	},
	"SimpleChild": {
		"0 HEAD\n1 BIRT",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagHeader, "", "",
				gedcom.NewBirthNode(""),
			))
		},
	},
	"ThreeDeep1": {
		"0 HEAD\n1 GEDC\n2 VERS (2010.3)",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagHeader, "", "",
				gedcom.NewNode(tag.TagGedcomInformation, "", "",
					gedcom.NewNode(tag.TagVersion, "(2010.3)", ""),
				),
			))
		},
	},
	"ThreeDeep2": {
		"0 HEAD\n1 GEDC\n2 VERS 5.5",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagHeader, "", "",
				gedcom.NewNode(tag.TagGedcomInformation, "", "",
					gedcom.NewNode(tag.TagVersion, "5.5", ""),
				),
			))
		},
	},
	"FORM": {
		"0 HEAD\n1 GEDC\n2 FORM LINEAGE-LINKED",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagHeader, "", "",
				gedcom.NewNode(tag.TagGedcomInformation, "", "",
					gedcom.NewFormatNode("LINEAGE-LINKED"),
				),
			))
		},
	},
	"PLAC": {
		"0 HEAD\n1 BIRT\n2 PLAC Camperdown, Nsw, Australia",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagHeader, "", "",
				gedcom.NewBirthNode("",
					gedcom.NewPlaceNode("Camperdown, Nsw, Australia"),
				),
			))
		},
	},
	"NAME": {
		"0 HEAD\n1 NAME Elliot Rupert de Peyster /Chance/",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagHeader, "", "",
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/"),
			))
		},
	},
	"Pointer": {
		"0 HEAD\n0 @P1@ INDI",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagHeader, "", ""))
			doc.AddIndividual("P1")
		},
	},
	"NestedPointer": {
		"0 HEAD\n1 SEX M\n0 @P1@ INDI",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagHeader, "", "",
				gedcom.NewNode(tag.TagSex, "M", ""),
			))
			doc.AddIndividual("P1")
		},
	},
	"SEX": {
		"0 HEAD\n1 SEX M",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagHeader, "", "",
				gedcom.NewNode(tag.TagSex, "M", ""),
			))
		},
	},
	"Nested1": {
		"0 HEAD\n1 BIRT\n2 PLAC Camperdown, Nsw, Australia\n1 SEX M",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagHeader, "", "",
				gedcom.NewBirthNode("",
					gedcom.NewPlaceNode("Camperdown, Nsw, Australia"),
				),
				gedcom.NewNode(tag.TagSex, "M", ""),
			))
		},
	},
	"MultipleRoots": {
		"0 HEAD\n0 @P1@ INDI\n1 BIRT",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagHeader, "", ""))
			doc.AddIndividual("P1",
				gedcom.NewBirthNode(""),
			)
		},
	},
	"Nested2": {
		"0 HEAD\n1 GEDC\n2 FORM LINEAGE-LINKED\n0 @P1@ INDI",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagHeader, "", "",
				gedcom.NewNode(tag.TagGedcomInformation, "", "",
					gedcom.NewFormatNode("LINEAGE-LINKED"),
				),
			))
			doc.AddIndividual("P1")
		},
	},
	"Nested3": {
		"0 HEAD0\n1 HEAD1\n2 HEAD2\n3 HEAD3\n0 HEAD00",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagFromString("HEAD0"), "", "",
				gedcom.NewNode(tag.TagFromString("HEAD1"), "", "",
					gedcom.NewNode(tag.TagFromString("HEAD2"), "", "",
						gedcom.NewNode(tag.TagFromString("HEAD3"), "", ""),
					),
				),
			))
			doc.AddNode(gedcom.NewNode(tag.TagFromString("HEAD00"), "", ""))
		},
	},
	"Nested4": {
		"0 HEAD0\n1 HEAD1\n2 HEAD2\n3 HEAD3\n1 HEAD10",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagFromString("HEAD0"), "", "",
				gedcom.NewNode(tag.TagFromString("HEAD1"), "", "",
					gedcom.NewNode(tag.TagFromString("HEAD2"), "", "",
						gedcom.NewNode(tag.TagFromString("HEAD3"), "", ""),
					),
				),
				gedcom.NewNode(tag.TagFromString("HEAD10"), "", ""),
			))
		},
	},
	"BadTag1": {
		"0 HEAD0\r1 HEAD1",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagFromString("HEAD0"), "", "",
				gedcom.NewNode(tag.TagFromString("HEAD1"), "", ""),
			))
		},
	},
	"WindowsLineEnding": {
		"0 HEAD0\r\n1 HEAD1",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagFromString("HEAD0"), "", "",
				gedcom.NewNode(tag.TagFromString("HEAD1"), "", ""),
			))
		},
	},
	"BadTag2": {
		"0 HEAD0\n1 HEAD1\n1 HEAD10\n2 HEAD2",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagFromString("HEAD0"), "", "",
				gedcom.NewNode(tag.TagFromString("HEAD1"), "", ""),
				gedcom.NewNode(tag.TagFromString("HEAD10"), "", "",
					gedcom.NewNode(tag.TagFromString("HEAD2"), "", ""),
				),
			))
		},
	},
	"SpaceWithoutValue": {
		"0 HEAD\n1 BIRT ",
		func(doc *gedcom.Document) {
			doc.AddNode(gedcom.NewNode(tag.TagHeader, "", "",
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
				gedcom.NewNode(tag.TagDeath, "", "",
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
			doc.AddNode(gedcom.NewNode(tag.TagName, "oh@ok", "uh"))
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
		doc.AddNode(gedcom.NewNode(tag.TagFromString("_PUBLISH"), "", ""))

		assertDocumentEqual(t, doc, actual)
	})

	t.Run("BOM", func(t *testing.T) {
		ged := "\xEF\xBB\xBF0 HEAD\n1 CHAR UTF-8"
		decoder := gedcom.NewDecoder(strings.NewReader(ged))
		expected := gedcom.NewDocumentWithNodes(gedcom.Nodes{
			gedcom.NewNode(tag.TagHeader, "", "",
				gedcom.NewNode(tag.TagCharacterSet, "UTF-8", ""),
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
			doc.AddNode(gedcom.NewNode(tag.TagFromString("TEXT"), "Karen Marie McMillan<br>Gender: Female<br>Birth: Aug 21 1957 - \n\n\t&nbsp;Texas, USA", ""))

			assertDocumentEqual(t, doc, actual)
		}
	})

	t.Run("IndentTooBig", func(t *testing.T) {
		decoder := gedcom.NewDecoder(strings.NewReader("0 @I59238932@ INDI\n2 NPFX Mrs William Cornens\n1 SEX F"))

		assert.PanicsWithValue(t, "indent is too large - missing parent? at line 2: 2 NPFX Mrs William Cornens", func() {
			_, _ = decoder.Decode()
		})
	})

	t.Run("IndentTooBigAllowed", func(t *testing.T) {
		decoder := gedcom.NewDecoder(strings.NewReader("0 @I59238932@ INDI\n2 NPFX Mrs William Cornens\n1 SEX F"))
		decoder.AllowInvalidIndents = true

		actual, err := decoder.Decode()
		if assert.NoError(t, err) {
			doc := gedcom.NewDocument()
			doc.AddIndividual("I59238932",
				gedcom.NewNode(tag.TagFromString("NPFX"), "Mrs William Cornens", ""),
				gedcom.NewNode(tag.TagFromString("SEX"), "F", ""))

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
		tag      tag.Tag
		expected gedcom.Node
	}{
		{tag.TagBaptism, gedcom.NewBaptismNode(v)},
		{tag.TagBirth, gedcom.NewBirthNode(v)},
		{tag.TagBurial, gedcom.NewBurialNode(v)},
		{tag.TagDate, gedcom.NewDateNode(v)},
		{tag.TagDeath, gedcom.NewDeathNode(v)},
		{tag.TagEvent, gedcom.NewEventNode(v)},
		{tag.UnofficialTagFamilySearchID1, gedcom.NewFamilySearchIDNode(tag.UnofficialTagFamilySearchID1, v)},
		{tag.UnofficialTagFamilySearchID2, gedcom.NewFamilySearchIDNode(tag.UnofficialTagFamilySearchID2, v)},
		{tag.TagFormat, gedcom.NewFormatNode(v)},
		{tag.TagLatitude, gedcom.NewLatitudeNode(v)},
		{tag.TagLongitude, gedcom.NewLongitudeNode(v)},
		{tag.TagMap, gedcom.NewMapNode(v)},
		{tag.TagName, gedcom.NewNameNode(v)},
		{tag.TagNote, gedcom.NewNoteNode(v)},
		{tag.TagNickname, gedcom.NewNicknameNode(v)},
		{tag.TagPhonetic, gedcom.NewPhoneticVariationNode(v)},
		{tag.TagPlace, gedcom.NewPlaceNode(v)},
		{tag.TagResidence, gedcom.NewResidenceNode(v)},
		{tag.TagRomanized, gedcom.NewRomanizedVariationNode(v)},
		{tag.TagSource, gedcom.NewSourceNode(v, p)},
		{tag.TagType, gedcom.NewTypeNode(v)},
		{tag.TagVersion, gedcom.NewNode(tag.TagVersion, v, p)},
		{tag.UnofficialTagUniqueID, gedcom.NewUniqueIDNode(v)},
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
