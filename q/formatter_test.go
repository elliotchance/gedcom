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
	result       interface{}
	asJSON       []byte
	asPrettyJSON []byte
	asCSV        []byte
	csvHeader    []string
	csvError     error
}{
	{
		result:       nil,
		asJSON:       []byte("null"),
		asPrettyJSON: []byte("null"),
		asCSV:        []byte(nil),
		csvHeader:    nil,
		csvError:     errors.New("not a slice"),
	},
	{
		result: []*gedcom.NameNode{
			gedcom.NewNameNode(nil, "Elliot /Chance/", "", nil),
			gedcom.NewNameNode(nil, "Dina /Wyche/", "", nil),
		},
		asJSON: []byte(`[{"Tag":"NAME","Value":"Elliot /Chance/"},{"Tag":"NAME","Value":"Dina /Wyche/"}]`),
		asPrettyJSON: []byte(`[
  {
    "Tag": "NAME",
    "Value": "Elliot /Chance/"
  },
  {
    "Tag": "NAME",
    "Value": "Dina /Wyche/"
  }
]`),
		asCSV: []byte(`Tag,Value
NAME,Elliot /Chance/
NAME,Dina /Wyche/
`),
		csvHeader: []string{"Tag", "Value"},
		csvError:  nil,
	},
	{
		result: []map[string]interface{}{
			{"foo": "bar,", "baz": 123},
			{"foo": 4.56, "baz": `q"ux`},
		},
		asJSON: []byte(`[{"baz":123,"foo":"bar,"},{"baz":"q\"ux","foo":4.56}]`),
		asPrettyJSON: []byte(`[
  {
    "baz": 123,
    "foo": "bar,"
  },
  {
    "baz": "q\"ux",
    "foo": 4.56
  }
]`),
		asCSV: []byte(`baz,foo
123,"bar,"
"q""ux",4.56
`),
		csvHeader: []string{"baz", "foo"},
		csvError:  nil,
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
