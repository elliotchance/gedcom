package q_test

import (
	"bytes"
	"errors"
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/gedcom/q"
	"github.com/stretchr/testify/assert"
	"testing"
)

var formatterTests = []struct {
	result interface{}

	// json
	asJSON []byte

	// pretty-json
	asPrettyJSON []byte

	// csv
	asCSV     []byte
	csvHeader []string
	csvError  error

	// gedcom
	asGEDCOM    []byte
	gedcomError error
}{
	{
		result:       nil,
		asJSON:       []byte("null\n"),
		asPrettyJSON: []byte("null\n"),
		asCSV:        []byte(nil),
		csvHeader:    nil,
		csvError:     errors.New("not a slice"),
	},
	{
		result: []*gedcom.NameNode{
			gedcom.NewNameNode("Elliot /Chance/"),
			gedcom.NewNameNode("Dina /Wyche/"),
		},
		asJSON: []byte(`[{"Tag":"NAME","Value":"Elliot /Chance/"},{"Tag":"NAME","Value":"Dina /Wyche/"}]
`),
		asPrettyJSON: []byte(`[
  {
    "Tag": "NAME",
    "Value": "Elliot /Chance/"
  },
  {
    "Tag": "NAME",
    "Value": "Dina /Wyche/"
  }
]
`),
		asCSV: []byte(`Tag,Value
NAME,Elliot /Chance/
NAME,Dina /Wyche/
`),
		csvHeader: []string{"Tag", "Value"},
		csvError:  nil,
		asGEDCOM:  []byte("0 NAME Elliot /Chance/\n0 NAME Dina /Wyche/\n"),
	},
	{
		result: gedcom.NewDocument().AddIndividual("P1",
			gedcom.NewNameNode("Elliot /Chance/"),
			gedcom.NewNameNode("Dina /Wyche/"),
		),
		asJSON: []byte(`{"Nodes":[{"Tag":"NAME","Value":"Elliot /Chance/"},{"Tag":"NAME","Value":"Dina /Wyche/"}],"Pointer":"P1","Tag":"INDI"}
`),
		asPrettyJSON: []byte(`{
  "Nodes": [
    {
      "Tag": "NAME",
      "Value": "Elliot /Chance/"
    },
    {
      "Tag": "NAME",
      "Value": "Dina /Wyche/"
    }
  ],
  "Pointer": "P1",
  "Tag": "INDI"
}
`),
		csvError: errors.New("not a slice"),
		asGEDCOM: []byte("0 @P1@ INDI\n1 NAME Elliot /Chance/\n1 NAME Dina /Wyche/\n"),
	},
	{
		result: []map[string]interface{}{
			{"foo": "bar,", "baz": 123},
			{"foo": 4.56, "baz": `q"ux`},
		},
		asJSON: []byte(`[{"baz":123,"foo":"bar,"},{"baz":"q\"ux","foo":4.56}]
`),
		asPrettyJSON: []byte(`[
  {
    "baz": 123,
    "foo": "bar,"
  },
  {
    "baz": "q\"ux",
    "foo": 4.56
  }
]
`),
		asCSV: []byte(`baz,foo
123,"bar,"
"q""ux",4.56
`),
		csvHeader:   []string{"baz", "foo"},
		csvError:    nil,
		gedcomError: errors.New("map[string]interface {} does not implement gedcom.GEDCOMStringer"),
	},
}

func TestJSONFormatter_Write(t *testing.T) {
	for _, test := range formatterTests {
		t.Run("", func(t *testing.T) {
			buffer := bytes.Buffer{}
			formatter := &q.JSONFormatter{&buffer}
			err := formatter.Write(test.result)
			assert.NoError(t, err)
			assert.Equal(t, test.asJSON, buffer.Bytes())
		})
	}
}

func TestPrettyJSONFormatter_Write(t *testing.T) {
	for _, test := range formatterTests {
		t.Run("", func(t *testing.T) {
			buffer := bytes.Buffer{}
			formatter := &q.PrettyJSONFormatter{&buffer}
			err := formatter.Write(test.result)
			assert.NoError(t, err)
			assert.Equal(t, test.asPrettyJSON, buffer.Bytes())
		})
	}
}

func TestCSVFormatter_Write(t *testing.T) {
	for _, test := range formatterTests {
		t.Run("", func(t *testing.T) {
			buffer := bytes.Buffer{}
			formatter := &q.CSVFormatter{&buffer}
			err := formatter.Write(test.result)

			if test.csvError == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err, test.csvError)
			}

			assert.Equal(t, test.asCSV, buffer.Bytes())
		})
	}
}

func TestCSVFormatter_Header(t *testing.T) {
	for _, test := range formatterTests {
		t.Run("", func(t *testing.T) {
			formatter := &q.CSVFormatter{}
			header, err := formatter.Header(test.result)

			if test.csvError == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err, test.csvError)
			}

			assert.Equal(t, test.csvHeader, header)
		})
	}
}

func TestGEDCOMFormatter_Write(t *testing.T) {
	for _, test := range formatterTests {
		t.Run("", func(t *testing.T) {
			buffer := bytes.Buffer{}
			formatter := &q.GEDCOMFormatter{&buffer}
			err := formatter.Write(test.result)

			if test.gedcomError == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err, test.gedcomError)
			}

			assert.Equal(t, test.asGEDCOM, buffer.Bytes())
		})
	}
}
