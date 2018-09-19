package gedcom

import (
	"github.com/elliotchance/tf"
	"testing"
)

func TestSex_String(t *testing.T) {
	String := tf.Function(t, Sex.String)

	String("").Returns("Unknown")
	String(SexUnknown).Returns("Unknown")
	String(SexMale).Returns("Male")
	String(SexFemale).Returns("Female")
}
