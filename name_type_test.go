package gedcom

import (
	"github.com/elliotchance/tf"
	"testing"
)

func TestNameType_String(t *testing.T) {
	String := tf.Function(t, NameType.String)

	String("").Returns("Normal")
	String(NameTypeNormal).Returns("Normal")
	String(NameTypeMarriedName).Returns("Married Name")
	String(NameTypeAlsoKnownAs).Returns("Also Known As")
	String(NameTypeMaidenName).Returns("Maiden Name")
	String(NameTypeNickname).Returns("Nickname")
	String("Foo bar").Returns("Foo bar")
}
