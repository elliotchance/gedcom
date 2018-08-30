package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"testing"
)

func TestGedcomLine(t *testing.T) {
	GedcomLine := tf.Function(t, gedcom.GedcomLine)

	GedcomLine(0, (*gedcom.SimpleNode)(nil)).Returns("")

	GedcomLine(0, gedcom.NewBirthNode(nil, "foo", "72", nil)).
		Returns("0 @72@ BIRT foo")

	GedcomLine(3, gedcom.NewSimpleNode(nil, gedcom.TagDeath, "bar", "baz", nil)).
		Returns("3 @baz@ DEAT bar")

	GedcomLine(2, gedcom.NewDateNode(nil, "3 SEP 1945", "", nil)).
		Returns("2 DATE 3 SEP 1945")

	GedcomLine(-1, gedcom.NewBirthNode(nil, "foo", "72", nil)).
		Returns("@72@ BIRT foo")
}
