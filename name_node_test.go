package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/stretchr/testify/assert"
	"testing"
)

var nameTests = []struct {
	node          *gedcom.NameNode
	title         string
	prefix        string
	givenName     string
	surnamePrefix string
	surname       string
	suffix        string
	str           string
}{
	{
		node:          gedcom.NewNameNode(nil, "", "", nil),
		title:         "",
		prefix:        "",
		givenName:     "",
		surnamePrefix: "",
		surname:       "",
		suffix:        "",
		str:           "",
	},
	{
		node:          gedcom.NewNameNode(nil, "/Double  Last/", "", nil),
		title:         "",
		prefix:        "",
		givenName:     "",
		surnamePrefix: "",
		surname:       "Double Last",
		suffix:        "",
		str:           "Double Last",
	},
	{
		node:          gedcom.NewNameNode(nil, "//", "", nil),
		title:         "",
		prefix:        "",
		givenName:     "",
		surnamePrefix: "",
		surname:       "",
		suffix:        "",
		str:           "",
	},
	{
		// This is an invalid case. I don't mind that the data returned seems
		// garbled. It's better than nothing.
		node:          gedcom.NewNameNode(nil, "a / b", "", nil),
		title:         "",
		prefix:        "",
		givenName:     "a",
		surnamePrefix: "",
		surname:       "",
		suffix:        "/ b",
		str:           "a / b",
	},
	{
		node:          gedcom.NewNameNode(nil, "Double First", "", nil),
		title:         "",
		prefix:        "",
		givenName:     "Double First",
		surnamePrefix: "",
		surname:       "",
		suffix:        "",
		str:           "Double First",
	},
	{
		node:          gedcom.NewNameNode(nil, "First /Last/", "", nil),
		title:         "",
		prefix:        "",
		givenName:     "First",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "",
		str:           "First Last",
	},
	{
		node:          gedcom.NewNameNode(nil, "First   Middle /Last/", "", nil),
		title:         "",
		prefix:        "",
		givenName:     "First Middle",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "",
		str:           "First Middle Last",
	},
	{
		node:          gedcom.NewNameNode(nil, "First /Last/  Suffix ", "", nil),
		title:         "",
		prefix:        "",
		givenName:     "First",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "Suffix",
		str:           "First Last Suffix",
	},
	{
		node:          gedcom.NewNameNode(nil, "   /Last/ Suffix", "", nil),
		title:         "",
		prefix:        "",
		givenName:     "",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "Suffix",
		str:           "Last Suffix",
	},
	{
		// The GivenName overrides the givenName name if provided. When multiple
		// GivenNames are provided then it will always use the first one.
		node: gedcom.NewNameNode(nil, "First /Last/ II", "", []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagGivenName, " Other  Name ", "", nil),
			gedcom.NewSimpleNode(nil, gedcom.TagGivenName, "Uh-oh", "", nil),
		}),
		title:         "",
		prefix:        "",
		givenName:     "Other Name",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "II",
		str:           "Other Name Last II",
	},
	{
		// The Surname overrides the surname name if provided. When multiple
		// Surnames are provided then it will always use the first one.
		node: gedcom.NewNameNode(nil, "First /Last/ II", "", []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagSurname, " Other  name ", "", nil),
			gedcom.NewSimpleNode(nil, gedcom.TagSurname, "uh-oh", "", nil),
		}),
		title:         "",
		prefix:        "",
		givenName:     "First",
		surnamePrefix: "",
		surname:       "Other name",
		suffix:        "II",
		str:           "First Other name II",
	},
	{
		node: gedcom.NewNameNode(nil, "First /Last/ Esq.", "", []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagNamePrefix, " Mr ", "", nil),
			gedcom.NewSimpleNode(nil, gedcom.TagNamePrefix, "Dr", "", nil),
		}),
		title:         "",
		prefix:        "Mr",
		givenName:     "First",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "Esq.",
		str:           "Mr First Last Esq.",
	},
	{
		// The NameSuffix overrides the suffix in the name if provided.
		// When multiple name suffixes are provided then it will always use the
		// first one.
		node: gedcom.NewNameNode(nil, "First /Last/ Suffix", "", []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagNameSuffix, " Esq. ", "", nil),
			gedcom.NewSimpleNode(nil, gedcom.TagNameSuffix, "Dr", "", nil),
			gedcom.NewSimpleNode(nil, gedcom.TagNamePrefix, "Sir", "", nil),
		}),
		title:         "",
		prefix:        "Sir",
		givenName:     "First",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "Esq.",
		str:           "Sir First Last Esq.",
	},
	{
		node: gedcom.NewNameNode(nil, "First /Last/ Esq.", "", []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagSurnamePrefix, " Foo ", "", nil),
			gedcom.NewSimpleNode(nil, gedcom.TagSurnamePrefix, "Bar", "", nil),
		}),
		title:         "",
		prefix:        "",
		givenName:     "First",
		surnamePrefix: "Foo",
		surname:       "Last",
		suffix:        "Esq.",
		str:           "First Foo Last Esq.",
	},
	{
		node: gedcom.NewNameNode(nil, "First /Last/ Esq.", "", []gedcom.Node{
			gedcom.NewSimpleNode(nil, gedcom.TagTitle, " Grand  Duke ", "", nil),
			gedcom.NewSimpleNode(nil, gedcom.TagTitle, "Nobody", "", nil),
		}),
		title:         "Grand Duke",
		prefix:        "",
		givenName:     "First",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "Esq.",
		str:           "Grand Duke First Last Esq.",
	},
}

func TestNameNode_GivenName(t *testing.T) {
	for _, test := range nameTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.GivenName(), test.givenName)
		})
	}
}

func TestNameNode_Surname(t *testing.T) {
	for _, test := range nameTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Surname(), test.surname)
		})
	}
}

func TestNameNode_SurnamePrefix(t *testing.T) {
	for _, test := range nameTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.SurnamePrefix(), test.surnamePrefix)
		})
	}
}

func TestNameNode_Prefix(t *testing.T) {
	for _, test := range nameTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Prefix(), test.prefix)
		})
	}
}

func TestNameNode_Suffix(t *testing.T) {
	for _, test := range nameTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Suffix(), test.suffix)
		})
	}
}

func TestNameNode_Title(t *testing.T) {
	for _, test := range nameTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.Title(), test.title)
		})
	}
}

func TestNameNode_String(t *testing.T) {
	for _, test := range nameTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.String(), test.str)
		})
	}
}
