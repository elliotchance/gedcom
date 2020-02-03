package gedcom_test

import (
	"github.com/elliotchance/gedcom/tag"
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

var nameTests = []struct {
	node          *gedcom.NameNode
	title         string // Title()
	prefix        string // Prefix()
	givenName     string // GivenName()
	surnamePrefix string // SurnamePrefix()
	surname       string // Surname()
	suffix        string // Suffix()
	str           string // String()
	gedcomName    string // GedcomName()
}{
	{
		node:          nil,
		title:         "",
		prefix:        "",
		givenName:     "",
		surnamePrefix: "",
		surname:       "",
		suffix:        "",
		str:           "",
		gedcomName:    "",
	},
	{
		node:          gedcom.NewNameNode(""),
		title:         "",
		prefix:        "",
		givenName:     "",
		surnamePrefix: "",
		surname:       "",
		suffix:        "",
		str:           "",
		gedcomName:    "",
	},
	{
		node:          gedcom.NewNameNode("/Double  Last/"),
		title:         "",
		prefix:        "",
		givenName:     "",
		surnamePrefix: "",
		surname:       "Double Last",
		suffix:        "",
		str:           "Double Last",
		gedcomName:    "/Double Last/",
	},
	{
		node:          gedcom.NewNameNode("//"),
		title:         "",
		prefix:        "",
		givenName:     "",
		surnamePrefix: "",
		surname:       "",
		suffix:        "",
		str:           "",
		gedcomName:    "",
	},
	{
		// This is an invalid case. I don't mind that the data returned seems
		// garbled. It's better than nothing.
		node:          gedcom.NewNameNode("a / b"),
		title:         "",
		prefix:        "",
		givenName:     "a",
		surnamePrefix: "",
		surname:       "",
		suffix:        "/ b",
		str:           "a / b",
		gedcomName:    "a / b",
	},
	{
		node:          gedcom.NewNameNode("Double First"),
		title:         "",
		prefix:        "",
		givenName:     "Double First",
		surnamePrefix: "",
		surname:       "",
		suffix:        "",
		str:           "Double First",
		gedcomName:    "Double First",
	},
	{
		node:          gedcom.NewNameNode("First /Last/"),
		title:         "",
		prefix:        "",
		givenName:     "First",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "",
		str:           "First Last",
		gedcomName:    "First /Last/",
	},
	{
		node:          gedcom.NewNameNode("First   Middle /Last/"),
		title:         "",
		prefix:        "",
		givenName:     "First Middle",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "",
		str:           "First Middle Last",
		gedcomName:    "First Middle /Last/",
	},
	{
		node:          gedcom.NewNameNode("First /Last/  Suffix "),
		title:         "",
		prefix:        "",
		givenName:     "First",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "Suffix",
		str:           "First Last Suffix",
		gedcomName:    "First /Last/ Suffix",
	},
	{
		node:          gedcom.NewNameNode("   /Last/ Suffix"),
		title:         "",
		prefix:        "",
		givenName:     "",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "Suffix",
		str:           "Last Suffix",
		gedcomName:    "/Last/ Suffix",
	},
	{
		// The GivenName overrides the givenName name if provided. When multiple
		// GivenNames are provided then it will always use the first one.
		node: gedcom.NewNameNode("First /Last/ II",
			gedcom.NewNode(tag.TagGivenName, " Other  Name ", ""),
			gedcom.NewNode(tag.TagGivenName, "Uh-oh", ""),
		),
		title:         "",
		prefix:        "",
		givenName:     "Other Name",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "II",
		str:           "Other Name Last II",
		gedcomName:    "Other Name /Last/ II",
	},
	{
		// The Surname overrides the surname name if provided. When multiple
		// Surnames are provided then it will always use the first one.
		node: gedcom.NewNameNode("First /Last/ II",
			gedcom.NewNode(tag.TagSurname, " Other  name ", ""),
			gedcom.NewNode(tag.TagSurname, "uh-oh", ""),
		),
		title:         "",
		prefix:        "",
		givenName:     "First",
		surnamePrefix: "",
		surname:       "Other name",
		suffix:        "II",
		str:           "First Other name II",
		gedcomName:    "First /Other name/ II",
	},
	{
		node: gedcom.NewNameNode("First /Last/ Esq.",
			gedcom.NewNode(tag.TagNamePrefix, " Mr ", ""),
			gedcom.NewNode(tag.TagNamePrefix, "Dr", ""),
		),
		title:         "",
		prefix:        "Mr",
		givenName:     "First",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "Esq.",
		str:           "Mr First Last Esq.",
		gedcomName:    "Mr First /Last/ Esq.",
	},
	{
		// The NameSuffix overrides the suffix in the name if provided.
		// When multiple name suffixes are provided then it will always use the
		// first one.
		node: gedcom.NewNameNode("First /Last/ Suffix",
			gedcom.NewNode(tag.TagNameSuffix, " Esq. ", ""),
			gedcom.NewNode(tag.TagNameSuffix, "Dr", ""),
			gedcom.NewNode(tag.TagNamePrefix, "Sir", ""),
		),
		title:         "",
		prefix:        "Sir",
		givenName:     "First",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "Esq.",
		str:           "Sir First Last Esq.",
		gedcomName:    "Sir First /Last/ Esq.",
	},
	{
		node: gedcom.NewNameNode("First /Last/ Esq.",
			gedcom.NewNode(tag.TagSurnamePrefix, " Foo ", ""),
			gedcom.NewNode(tag.TagSurnamePrefix, "Bar", ""),
		),
		title:         "",
		prefix:        "",
		givenName:     "First",
		surnamePrefix: "Foo",
		surname:       "Last",
		suffix:        "Esq.",
		str:           "First Foo Last Esq.",
		gedcomName:    "First Foo /Last/ Esq.",
	},
	{
		node: gedcom.NewNameNode("First /Last/ Esq.",
			gedcom.NewNode(tag.TagTitle, " Grand  Duke ", ""),
			gedcom.NewNode(tag.TagTitle, "Nobody", ""),
		),
		title:         "Grand Duke",
		prefix:        "",
		givenName:     "First",
		surnamePrefix: "",
		surname:       "Last",
		suffix:        "Esq.",
		str:           "Grand Duke First Last Esq.",
		gedcomName:    "Grand Duke First /Last/ Esq.",
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

func TestNameNode_GedcomName(t *testing.T) {
	for _, test := range nameTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.node.GedcomName(), test.gedcomName)
		})
	}
}

func TestNameNode_Type(t *testing.T) {
	Type := tf.Function(t, (*gedcom.NameNode).Type)

	Type((*gedcom.NameNode)(nil)).Returns(gedcom.NameTypeNormal)
}

func TestNameNode_Format(t *testing.T) {
	Format := tf.Function(t, (*gedcom.NameNode).Format)

	Format(nil, "").Returns("")
	Format(nil, "%f %l").Returns("")

	name := gedcom.NewNameNode("",
		gedcom.NewNode(tag.TagGivenName, "Given", ""),
		gedcom.NewNode(tag.TagSurname, "Surname", ""),
		gedcom.NewNode(tag.TagNamePrefix, "Prefix", ""),
		gedcom.NewNode(tag.TagNameSuffix, "Suffix", ""),
		gedcom.NewNode(tag.TagSurnamePrefix, "SurnamePrefix", ""),
		gedcom.NewNode(tag.TagTitle, "Title", ""),
	)

	Format(name, "").Returns("")
	Format(name, "%").Returns("%")
	Format(name, "%a").Returns("%a")
	Format(name, "%A").Returns("%A")
	Format(name, "%%").Returns("%")

	Format(name, "%f").Returns("Given")
	Format(name, "%l").Returns("Surname")
	Format(name, "%m").Returns("SurnamePrefix")
	Format(name, "%p").Returns("Prefix")
	Format(name, "%s").Returns("Suffix")
	Format(name, "%t").Returns("Title")

	Format(name, "%F").Returns("GIVEN")
	Format(name, "%L").Returns("SURNAME")
	Format(name, "%M").Returns("SURNAMEPREFIX")
	Format(name, "%P").Returns("PREFIX")
	Format(name, "%S").Returns("SUFFIX")
	Format(name, "%T").Returns("TITLE")

	Format(name, "HI %t").Returns("HI Title")
	Format(name, "HI %t bar").Returns("HI Title bar")
	Format(name, "%l, %f").Returns("Surname, Given")

	Format(name, gedcom.NameFormatWritten).Returns("Title Prefix Given SurnamePrefix Surname Suffix")
	Format(name, gedcom.NameFormatGEDCOM).Returns("Title Prefix Given SurnamePrefix /Surname/ Suffix")
	Format(name, gedcom.NameFormatIndex).Returns("SurnamePrefix Surname, Title Prefix Given Suffix")

	name = gedcom.NewNameNode("Bob /Smith/")

	Format(name, "%f %L").Returns("Bob SMITH")
	Format(name, "%f%L").Returns("BobSMITH")
	Format(name, "%f %m (%l)").Returns("Bob (Smith)")
}
