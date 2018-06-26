package gedcom_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/elliotchance/gedcom"
	"strings"
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
			gedcom.NewSimpleNode("HEAD", "", "", []gedcom.Node{}),
		},
	},
	"0 HEAD\n1 CHAR UTF-8": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode("HEAD", "", "", []gedcom.Node{
				gedcom.NewSimpleNode("CHAR", "UTF-8", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD\n\n1 CHAR UTF-8\n": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode("HEAD", "", "", []gedcom.Node{
				gedcom.NewSimpleNode("CHAR", "UTF-8", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD\n1 CHAR UTF-8\n1 SOUR Ancestry.com Family Trees": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode("HEAD", "", "", []gedcom.Node{
				gedcom.NewSimpleNode("CHAR", "UTF-8", "", []gedcom.Node{}),
				gedcom.NewSimpleNode("SOUR", "Ancestry.com Family Trees", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD\n1 CHAR UTF-8\n1 CHAR UTF-8": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode("HEAD", "", "", []gedcom.Node{
				gedcom.NewSimpleNode("CHAR", "UTF-8", "", []gedcom.Node{}),
				gedcom.NewSimpleNode("CHAR", "UTF-8", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD\n1 SOUR Ancestry.com Family Trees": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode("HEAD", "", "", []gedcom.Node{
				gedcom.NewSimpleNode("SOUR", "Ancestry.com Family Trees", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD\n1 BIRT": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode("HEAD", "", "", []gedcom.Node{
				gedcom.NewSimpleNode("BIRT", "", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD\n1 GEDC\n2 VERS (2010.3)": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode("HEAD", "", "", []gedcom.Node{
				gedcom.NewSimpleNode("GEDC", "", "", []gedcom.Node{
					gedcom.NewSimpleNode("VERS", "(2010.3)", "", []gedcom.Node{}),
				}),
			}),
		},
	},
	"0 HEAD\n1 GEDC\n2 VERS 5.5": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode("HEAD", "", "", []gedcom.Node{
				gedcom.NewSimpleNode("GEDC", "", "", []gedcom.Node{
					gedcom.NewSimpleNode("VERS", "5.5", "", []gedcom.Node{}),
				}),
			}),
		},
	},
	"0 HEAD\n1 GEDC\n2 FORM LINEAGE-LINKED": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode("HEAD", "", "", []gedcom.Node{
				gedcom.NewSimpleNode("GEDC", "", "", []gedcom.Node{
					gedcom.NewSimpleNode("FORM", "LINEAGE-LINKED", "", []gedcom.Node{}),
				}),
			}),
		},
	},
	"0 HEAD\n1 BIRT\n2 PLAC Camperdown, Nsw, Australia": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode("HEAD", "", "", []gedcom.Node{
				gedcom.NewSimpleNode("BIRT", "", "", []gedcom.Node{
					gedcom.NewSimpleNode("PLAC", "Camperdown, Nsw, Australia", "", []gedcom.Node{}),
				}),
			}),
		},
	},
	"0 HEAD\n1 NAME Elliot Rupert de Peyster /Chance/": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode("HEAD", "", "", []gedcom.Node{
				gedcom.NewNameNode("Elliot Rupert de Peyster /Chance/", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD\n0 @P1@ INDI": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode("HEAD", "", "", []gedcom.Node{}),
			gedcom.NewSimpleNode("INDI", "", "P1", []gedcom.Node{}),
		},
	},
	"0 HEAD\n1 SEX M\n0 @P1@ INDI": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode("HEAD", "", "", []gedcom.Node{
				gedcom.NewSimpleNode("SEX", "M", "", []gedcom.Node{}),
			}),
			gedcom.NewSimpleNode("INDI", "", "P1", []gedcom.Node{}),
		},
	},
	"0 HEAD\n1 SEX M": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode("HEAD", "", "", []gedcom.Node{
				gedcom.NewSimpleNode("SEX", "M", "", []gedcom.Node{}),
			}),
		},
	},
	"0 HEAD\n1 BIRT\n2 PLAC Camperdown, Nsw, Australia\n1 SEX M": {
		Nodes: []gedcom.Node{
			gedcom.NewSimpleNode("HEAD", "", "", []gedcom.Node{
				gedcom.NewSimpleNode("BIRT", "", "", []gedcom.Node{
					gedcom.NewSimpleNode("PLAC", "Camperdown, Nsw, Australia", "", []gedcom.Node{}),
				}),
				gedcom.NewSimpleNode("SEX", "M", "", []gedcom.Node{}),
			}),
		},
	},
}

func TestDecoder_Decode(t *testing.T) {
	for ged, expected := range tests {
		t.Run("", func(t *testing.T) {
			decoder := gedcom.NewDecoder(strings.NewReader(ged))
			actual, err := decoder.Decode()

			assert.NoError(t, err, ged)
			assert.Equal(t, expected, actual, ged)
		})
	}
}

func trimSpaces(s string) string {
	s = strings.TrimSpace(s)
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
