package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"testing"
)

func TestNewNameFormatByName(t *testing.T) {
	NewNameFormatByName := tf.Function(t, gedcom.NewNameFormatByName)

	NewNameFormatByName("written").Returns(gedcom.NameFormatWritten, true)
	NewNameFormatByName("gedcom").Returns(gedcom.NameFormatGEDCOM, true)
	NewNameFormatByName("index").Returns(gedcom.NameFormatIndex, true)

	NewNameFormatByName("%L %f").Returns("%L %f", false)
	NewNameFormatByName("").Returns("", false)
}
