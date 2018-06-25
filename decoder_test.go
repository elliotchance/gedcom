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
			Nodes: []*gedcom.SimpleNode{},
		},
		"\n\n": {
			Nodes: []*gedcom.SimpleNode{},
		},
		"0 HEAD": {
			Nodes: []*gedcom.SimpleNode{
				{
					Indent:   0,
					Tag:      "HEAD",
					Value:    "",
					Pointer:  "",
					Children: []*gedcom.SimpleNode{},
				},
			},
		},
		"0 HEAD\n1 CHAR UTF-8": {
			Nodes: []*gedcom.SimpleNode{
				{
					Indent:  0,
					Tag:     "HEAD",
					Value:   "",
					Pointer: "",
					Children: []*gedcom.SimpleNode{
						{
							Indent:   1,
							Tag:      "CHAR",
							Value:    "UTF-8",
							Pointer:  "",
							Children: []*gedcom.SimpleNode{},
						},
					},
				},
			},
		},
		"0 HEAD\n\n1 CHAR UTF-8\n": {
			Nodes: []*gedcom.SimpleNode{
				{
					Indent:  0,
					Tag:     "HEAD",
					Value:   "",
					Pointer: "",
					Children: []*gedcom.SimpleNode{
						{
							Indent:   1,
							Tag:      "CHAR",
							Value:    "UTF-8",
							Pointer:  "",
							Children: []*gedcom.SimpleNode{},
						},
					},
				},
			},
		},
		"0 HEAD\n1 CHAR UTF-8\n1 SOUR Ancestry.com Family Trees": {
			Nodes: []*gedcom.SimpleNode{
				{
					Indent:  0,
					Tag:     "HEAD",
					Value:   "",
					Pointer: "",
					Children: []*gedcom.SimpleNode{
						{
							Indent:   1,
							Tag:      "CHAR",
							Value:    "UTF-8",
							Pointer:  "",
							Children: []*gedcom.SimpleNode{},
						},
						{
							Indent:   1,
							Tag:      "SOUR",
							Value:    "Ancestry.com Family Trees",
							Pointer:  "",
							Children: []*gedcom.SimpleNode{},
						},
					},
				},
			},
		},
		"0 HEAD\n1 CHAR UTF-8\n1 CHAR UTF-8": {
			Nodes: []*gedcom.SimpleNode{
				{
					Indent:  0,
					Tag:     "HEAD",
					Value:   "",
					Pointer: "",
					Children: []*gedcom.SimpleNode{
						{
							Indent:   1,
							Tag:      "CHAR",
							Value:    "UTF-8",
							Pointer:  "",
							Children: []*gedcom.SimpleNode{},
						},
						{
							Indent:   1,
							Tag:      "CHAR",
							Value:    "UTF-8",
							Pointer:  "",
							Children: []*gedcom.SimpleNode{},
						},
					},
				},
			},
		},
		"0 HEAD\n1 SOUR Ancestry.com Family Trees": {
			Nodes: []*gedcom.SimpleNode{
				{
					Indent:  0,
					Tag:     "HEAD",
					Value:   "",
					Pointer: "",
					Children: []*gedcom.SimpleNode{
						{
							Indent:   1,
							Tag:      "SOUR",
							Value:    "Ancestry.com Family Trees",
							Pointer:  "",
							Children: []*gedcom.SimpleNode{},
						},
					},
				},
			},
		},
		"0 HEAD\n1 BIRT": {
			Nodes: []*gedcom.SimpleNode{
				{
					Indent:  0,
					Tag:     "HEAD",
					Value:   "",
					Pointer: "",
					Children: []*gedcom.SimpleNode{
						{
							Indent:   1,
							Tag:      "BIRT",
							Value:    "",
							Pointer:  "",
							Children: []*gedcom.SimpleNode{},
						},
					},
				},
			},
		},
		"0 HEAD\n1 GEDC\n2 VERS (2010.3)": {
			Nodes: []*gedcom.SimpleNode{
				{
					Indent:  0,
					Tag:     "HEAD",
					Value:   "",
					Pointer: "",
					Children: []*gedcom.SimpleNode{
						{
							Indent:  1,
							Tag:     "GEDC",
							Value:   "",
							Pointer: "",
							Children: []*gedcom.SimpleNode{
								{
									Indent:   2,
									Tag:      "VERS",
									Value:    "(2010.3)",
									Pointer:  "",
									Children: []*gedcom.SimpleNode{},
								},
							},
						},
					},
				},
			},
		},
		"0 HEAD\n1 GEDC\n2 VERS 5.5": {
			Nodes: []*gedcom.SimpleNode{
				{
					Indent:  0,
					Tag:     "HEAD",
					Value:   "",
					Pointer: "",
					Children: []*gedcom.SimpleNode{
						{
							Indent:  1,
							Tag:     "GEDC",
							Value:   "",
							Pointer: "",
							Children: []*gedcom.SimpleNode{
								{
									Indent:   2,
									Tag:      "VERS",
									Value:    "5.5",
									Pointer:  "",
									Children: []*gedcom.SimpleNode{},
								},
							},
						},
					},
				},
			},
		},
		"0 HEAD\n1 GEDC\n2 FORM LINEAGE-LINKED": {
			Nodes: []*gedcom.SimpleNode{
				{
					Indent:  0,
					Tag:     "HEAD",
					Value:   "",
					Pointer: "",
					Children: []*gedcom.SimpleNode{
						{
							Indent:  1,
							Tag:     "GEDC",
							Value:   "",
							Pointer: "",
							Children: []*gedcom.SimpleNode{
								{
									Indent:   2,
									Tag:      "FORM",
									Value:    "LINEAGE-LINKED",
									Pointer:  "",
									Children: []*gedcom.SimpleNode{},
								},
							},
						},
					},
				},
			},
		},
		"0 HEAD\n1 BIRT\n2 PLAC Camperdown, Nsw, Australia": {
			Nodes: []*gedcom.SimpleNode{
				{
					Indent:  0,
					Tag:     "HEAD",
					Value:   "",
					Pointer: "",
					Children: []*gedcom.SimpleNode{
						{
							Indent:  1,
							Tag:     "BIRT",
							Value:   "",
							Pointer: "",
							Children: []*gedcom.SimpleNode{
								{
									Indent:   2,
									Tag:      "PLAC",
									Value:    "Camperdown, Nsw, Australia",
									Pointer:  "",
									Children: []*gedcom.SimpleNode{},
								},
							},
						},
					},
				},
			},
		},
		"0 HEAD\n1 NAME Elliot Rupert de Peyster /Chance/": {
			Nodes: []*gedcom.SimpleNode{
				{
					Indent:  0,
					Tag:     "HEAD",
					Value:   "",
					Pointer: "",
					Children: []*gedcom.SimpleNode{
						{
							Indent:   1,
							Tag:      "NAME",
							Value:    "Elliot Rupert de Peyster /Chance/",
							Pointer:  "",
							Children: []*gedcom.SimpleNode{},
						},
					},
				},
			},
		},
		"0 HEAD\n0 @P1@ INDI": {
			Nodes: []*gedcom.SimpleNode{
				{
					Indent:   0,
					Tag:      "HEAD",
					Value:    "",
					Pointer:  "",
					Children: []*gedcom.SimpleNode{},
				},
				{
					Indent:   0,
					Tag:      "INDI",
					Value:    "",
					Pointer:  "P1",
					Children: []*gedcom.SimpleNode{},
				},
			},
		},
		"0 HEAD\n1 SEX M\n0 @P1@ INDI": {
			Nodes: []*gedcom.SimpleNode{
				{
					Indent:  0,
					Tag:     "HEAD",
					Value:   "",
					Pointer: "",
					Children: []*gedcom.SimpleNode{
						{
							Indent:   1,
							Tag:      "SEX",
							Value:    "M",
							Pointer:  "",
							Children: []*gedcom.SimpleNode{},
						},
					},
				},
				{
					Indent:   0,
					Tag:      "INDI",
					Value:    "",
					Pointer:  "P1",
					Children: []*gedcom.SimpleNode{},
				},
			},
		},
		"0 HEAD\n1 SEX M": {
			Nodes: []*gedcom.SimpleNode{
				{
					Indent:  0,
					Tag:     "HEAD",
					Value:   "",
					Pointer: "",
					Children: []*gedcom.SimpleNode{
						{
							Indent:   1,
							Tag:      "SEX",
							Value:    "M",
							Pointer:  "",
							Children: []*gedcom.SimpleNode{},
						},
					},
				},
			},
		},
		"0 HEAD\n1 BIRT\n2 PLAC Camperdown, Nsw, Australia\n1 SEX M": {
			Nodes: []*gedcom.SimpleNode{
				{
					Indent:  0,
					Tag:     "HEAD",
					Value:   "",
					Pointer: "",
					Children: []*gedcom.SimpleNode{
						{
							Indent:  1,
							Tag:     "BIRT",
							Value:   "",
							Pointer: "",
							Children: []*gedcom.SimpleNode{
								{
									Indent:   2,
									Tag:      "PLAC",
									Value:    "Camperdown, Nsw, Australia",
									Pointer:  "",
									Children: []*gedcom.SimpleNode{},
								},
							},
						},
						{
							Indent:   1,
							Tag:      "SEX",
							Value:    "M",
							Pointer:  "",
							Children: []*gedcom.SimpleNode{},
						},
					},
				},
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
