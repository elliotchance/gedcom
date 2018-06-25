package gedcom_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/elliotchance/gedcom"
	"strings"
)

func TestDecoder_Decode(t *testing.T) {
	tests := map[string]*gedcom.DocumentNode{
		"": {
			Nodes: []gedcom.Node{},
		},
		"\n\n": {
			Nodes: []gedcom.Node{},
		},
		"0 HEAD": {
			Nodes: []gedcom.Node{
				gedcom.NewSimpleNode(0, "HEAD", "", "", []gedcom.Node{}),
			},
		},
		"0 HEAD\n1 CHAR UTF-8": {
			Nodes: []gedcom.Node{
				gedcom.NewSimpleNode(0, "HEAD", "", "", []gedcom.Node{
					gedcom.NewSimpleNode(1, "CHAR", "UTF-8", "", []gedcom.Node{}),
				}),
			},
		},
		"0 HEAD\n\n1 CHAR UTF-8\n": {
			Nodes: []gedcom.Node{
				gedcom.NewSimpleNode(0, "HEAD", "", "", []gedcom.Node{
					gedcom.NewSimpleNode(1, "CHAR", "UTF-8", "", []gedcom.Node{}),
				}),
			},
		},
		"0 HEAD\n1 CHAR UTF-8\n1 SOUR Ancestry.com Family Trees": {
			Nodes: []gedcom.Node{
				gedcom.NewSimpleNode(0, "HEAD", "", "", []gedcom.Node{
					gedcom.NewSimpleNode(1, "CHAR", "UTF-8", "", []gedcom.Node{}),
					gedcom.NewSimpleNode(1, "SOUR", "Ancestry.com Family Trees", "", []gedcom.Node{}),
				}),
			},
		},
		"0 HEAD\n1 CHAR UTF-8\n1 CHAR UTF-8": {
			Nodes: []gedcom.Node{
				gedcom.NewSimpleNode(0, "HEAD", "", "", []gedcom.Node{
					gedcom.NewSimpleNode(1, "CHAR", "UTF-8", "", []gedcom.Node{}),
					gedcom.NewSimpleNode(1, "CHAR", "UTF-8", "", []gedcom.Node{}),
				}),
			},
		},
		"0 HEAD\n1 SOUR Ancestry.com Family Trees": {
			Nodes: []gedcom.Node{
				gedcom.NewSimpleNode(0, "HEAD", "", "", []gedcom.Node{
					gedcom.NewSimpleNode(1, "SOUR", "Ancestry.com Family Trees", "", []gedcom.Node{}),
				}),
			},
		},
		"0 HEAD\n1 BIRT": {
			Nodes: []gedcom.Node{
				gedcom.NewSimpleNode(0, "HEAD", "", "", []gedcom.Node{
					gedcom.NewSimpleNode(1, "BIRT", "", "", []gedcom.Node{}),
				}),
			},
		},
		"0 HEAD\n1 GEDC\n2 VERS (2010.3)": {
			Nodes: []gedcom.Node{
				gedcom.NewSimpleNode(0, "HEAD", "", "", []gedcom.Node{
					gedcom.NewSimpleNode(1, "GEDC", "", "", []gedcom.Node{
						gedcom.NewSimpleNode(2, "VERS", "(2010.3)", "", []gedcom.Node{}),
					}),
				}),
			},
		},
		"0 HEAD\n1 GEDC\n2 VERS 5.5": {
			Nodes: []gedcom.Node{
				gedcom.NewSimpleNode(0, "HEAD", "", "", []gedcom.Node{
					gedcom.NewSimpleNode(1, "GEDC", "", "", []gedcom.Node{
						gedcom.NewSimpleNode(2, "VERS", "5.5", "", []gedcom.Node{}),
					}),
				}),
			},
		},
		"0 HEAD\n1 GEDC\n2 FORM LINEAGE-LINKED": {
			Nodes: []gedcom.Node{
				gedcom.NewSimpleNode(0, "HEAD", "", "", []gedcom.Node{
					gedcom.NewSimpleNode(1, "GEDC", "", "", []gedcom.Node{
						gedcom.NewSimpleNode(2, "FORM", "LINEAGE-LINKED", "", []gedcom.Node{}),
					}),
				}),
			},
		},
		"0 HEAD\n1 BIRT\n2 PLAC Camperdown, Nsw, Australia": {
			Nodes: []gedcom.Node{
				gedcom.NewSimpleNode(0, "HEAD", "", "", []gedcom.Node{
					gedcom.NewSimpleNode(1, "BIRT", "", "", []gedcom.Node{
						gedcom.NewSimpleNode(2, "PLAC", "Camperdown, Nsw, Australia", "", []gedcom.Node{}),
					}),
				}),
			},
		},
		"0 HEAD\n1 NAME Elliot Rupert de Peyster /Chance/": {
			Nodes: []gedcom.Node{
				gedcom.NewSimpleNode(0, "HEAD", "", "", []gedcom.Node{
					gedcom.NewSimpleNode(1, "NAME", "Elliot Rupert de Peyster /Chance/", "", []gedcom.Node{}),
				}),
			},
		},
		"0 HEAD\n0 @P1@ INDI": {
			Nodes: []gedcom.Node{
				gedcom.NewSimpleNode(0, "HEAD", "", "", []gedcom.Node{}),
				gedcom.NewSimpleNode(0, "INDI", "", "P1", []gedcom.Node{}),
			},
		},
		"0 HEAD\n1 SEX M\n0 @P1@ INDI": {
			Nodes: []gedcom.Node{
				gedcom.NewSimpleNode(0, "HEAD", "", "", []gedcom.Node{
					gedcom.NewSimpleNode(1, "SEX", "M", "", []gedcom.Node{}),
				}),
				gedcom.NewSimpleNode(0, "INDI", "", "P1", []gedcom.Node{}),
			},
		},
		"0 HEAD\n1 SEX M": {
			Nodes: []gedcom.Node{
				gedcom.NewSimpleNode(0, "HEAD", "", "", []gedcom.Node{
					gedcom.NewSimpleNode(1, "SEX", "M", "", []gedcom.Node{}),
				}),
			},
		},
		"0 HEAD\n1 BIRT\n2 PLAC Camperdown, Nsw, Australia\n1 SEX M": {
			Nodes: []gedcom.Node{
				gedcom.NewSimpleNode(0, "HEAD", "", "", []gedcom.Node{
					gedcom.NewSimpleNode(1, "BIRT", "", "", []gedcom.Node{
						gedcom.NewSimpleNode(2, "PLAC", "Camperdown, Nsw, Australia", "", []gedcom.Node{}),
					}),
					gedcom.NewSimpleNode(1, "SEX", "M", "", []gedcom.Node{}),
				}),
			},
		},
	}

	for ged, expected := range tests {
		t.Run("", func(t *testing.T) {
			decoder := gedcom.NewDecoder(strings.NewReader(ged))
			actual, err := decoder.Decode()

			assert.NoError(t, err, ged)
			assert.Equal(t, expected, actual, ged)
		})
	}
}
