package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var tests = map[string]*gedcom.Document{
	"": {
		Nodes: []gedcom.Node{},
	},
	"\n\n": {
		Nodes: []gedcom.Node{},
	},
	"0 HEAD": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagHeader, "", "", []gedcom.Node{}),
		},
	},
	"0 HEAD\n1 CHAR UTF-8": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagHeader, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagCharacterSet, "UTF-8", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD\n\n1 CHAR UTF-8\n": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagHeader, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagCharacterSet, "UTF-8", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD\n1 CHAR UTF-8\n1 SOUR Ancestry.com Family Trees": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagHeader, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagCharacterSet, "UTF-8", "", []gedcom.Node{}),
				gedcom.NewSourceNode(nil, "Ancestry.com Family Trees", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD\n1 CHAR UTF-8\n1 CHAR UTF-8": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagHeader, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagCharacterSet, "UTF-8", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(nil, gedcom.TagCharacterSet, "UTF-8", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD\n1 SOUR Ancestry.com Family Trees": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagHeader, "", "", []gedcom.Node{
				gedcom.NewSourceNode(nil, "Ancestry.com Family Trees", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD\n1 BIRT": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagHeader, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagBirth, "", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD\n1 GEDC\n2 VERS (2010.3)": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagHeader, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagGedcomInformation, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(nil, gedcom.TagVersion, "(2010.3)", "", []gedcom.Node{}),
				}),
			}),
		},
	},
	"0 HEAD\n1 GEDC\n2 VERS 5.5": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagHeader, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagGedcomInformation, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(nil, gedcom.TagVersion, "5.5", "", []gedcom.Node{}),
				}),
			}),
		},
	},
	"0 HEAD\n1 GEDC\n2 FORM LINEAGE-LINKED": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagHeader, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagGedcomInformation, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(nil, gedcom.TagFormat, "LINEAGE-LINKED", "", []gedcom.Node{}),
				}),
			}),
		},
	},
	"0 HEAD\n1 BIRT\n2 PLAC Camperdown, Nsw, Australia": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagHeader, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagBirth, "", "", []gedcom.Node{
					gedcom.NewPlaceNode(nil, "Camperdown, Nsw, Australia", "", []gedcom.Node{}),
				}),
			}),
		},
	},
	"0 HEAD\n1 NAME Elliot Rupert de Peyster /Chance/": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagHeader, "", "", []gedcom.Node{
				gedcom.NewNameNode(nil, "Elliot Rupert de Peyster /Chance/", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD\n0 @P1@ INDI": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagHeader, "", "", []gedcom.Node{}),
			gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{}),
		},
	},
	"0 HEAD\n1 SEX M\n0 @P1@ INDI": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagHeader, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagSex, "M", "", []gedcom.Node{}),
			}),
			gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{}),
		},
	},
	"0 HEAD\n1 SEX M": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagHeader, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagSex, "M", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD\n1 BIRT\n2 PLAC Camperdown, Nsw, Australia\n1 SEX M": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagHeader, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagBirth, "", "", []gedcom.Node{
					gedcom.NewPlaceNode(nil, "Camperdown, Nsw, Australia", "", []gedcom.Node{}),
				}),
				gedcom.NewSimpleNode(nil, gedcom.TagSex, "M", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD\n0 @P1@ INDI\n1 BIRT": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagHeader, "", "", []gedcom.Node{}),
			gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagBirth, "", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD\n1 GEDC\n2 FORM LINEAGE-LINKED\n0 @P1@ INDI": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagHeader, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagGedcomInformation, "", "", []gedcom.Node{
					gedcom.NewSimpleNode(nil, gedcom.TagFormat, "LINEAGE-LINKED", "", []gedcom.Node{}),
				}),
			}),
			gedcom.NewIndividualNode(nil, "", "P1", []gedcom.Node{}),
		},
	},
	"0 HEAD0\n1 HEAD1\n2 HEAD2\n3 HEAD3\n0 HEAD00": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagFromString("HEAD0"), "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagFromString("HEAD1"), "", "", []gedcom.Node{
					gedcom.NewSimpleNode(nil, gedcom.TagFromString("HEAD2"), "", "", []gedcom.Node{
						gedcom.NewSimpleNode(nil, gedcom.TagFromString("HEAD3"), "", "", []gedcom.Node{}),
					}),
				}),
			}),
			gedcom.NewSimpleNode(nil, gedcom.TagFromString("HEAD00"), "", "", []gedcom.Node{}),
		},
	},
	"0 HEAD0\n1 HEAD1\n2 HEAD2\n3 HEAD3\n1 HEAD10": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagFromString("HEAD0"), "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagFromString("HEAD1"), "", "", []gedcom.Node{
					gedcom.NewSimpleNode(nil, gedcom.TagFromString("HEAD2"), "", "", []gedcom.Node{
						gedcom.NewSimpleNode(nil, gedcom.TagFromString("HEAD3"), "", "", []gedcom.Node{}),
					}),
				}),
				gedcom.NewSimpleNode(nil, gedcom.TagFromString("HEAD10"), "", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD0\r1 HEAD1": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagFromString("HEAD0"), "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagFromString("HEAD1"), "", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD0\r\n1 HEAD1": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagFromString("HEAD0"), "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagFromString("HEAD1"), "", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD0\n1 HEAD1\n1 HEAD10\n2 HEAD2": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagFromString("HEAD0"), "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagFromString("HEAD1"), "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(nil, gedcom.TagFromString("HEAD10"), "", "", []gedcom.Node{
					gedcom.NewSimpleNode(nil, gedcom.TagFromString("HEAD2"), "", "", []gedcom.Node{}),
				}),
			}),
		},
	},
	"0 HEAD\n1 BIRT ": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagHeader, "", "", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagBirth, "", "", []gedcom.Node{}),
			}),
		},
	},
	"0 @P221@ INDI\n1 BIRT\n2 DATE 1851\n1 DEAT\n2 DATE 1856": {
		Nodes: []gedcom.Node{
			gedcom.NewIndividualNode(nil, "", "P221", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagBirth, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "1851", "", []gedcom.Node{}),
				}),
				gedcom.NewSimpleNode(nil, gedcom.TagDeath, "", "", []gedcom.Node{
					gedcom.NewDateNode(nil, "1856", "", []gedcom.Node{}),
				}),
			}),
		},
	},
	"0 @F1@ FAM\n1 HUSB @P2@\n1 WIFE @P3@": {
		Nodes: []gedcom.Node{
			gedcom.NewFamilyNode(nil, "F1", []gedcom.Node{
				gedcom.NewSimpleNode(nil, gedcom.TagHusband, "@P2@", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(nil, gedcom.TagWife, "@P3@", "", []gedcom.Node{}),
			}),
		},
	},
	"0 DATE 1856": {
		Nodes: []gedcom.Node{
			gedcom.NewDateNode(nil, "1856", "", []gedcom.Node{}),
		},
	},
}

func TestDecoder_Decode(t *testing.T) {
	for ged, expected := range tests {
		t.Run("", func(t *testing.T) {
			decoder := gedcom.NewDecoder(strings.NewReader(ged))
			actual, err := decoder.Decode()

			assert.NoError(t, err, ged)

			for _, n := range expected.Nodes {
				n.SetDocument(expected)
			}
			assert.Equal(t, expected, actual, ged)
		})
	}
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
