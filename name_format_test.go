package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/tf"
)

func TestNewNameFormatByName(t *testing.T) {
	NewNameFormatByName := tf.Function(t, gedcom.NewNameFormatByName)

	NewNameFormatByName("written").Returns(gedcom.NameFormatWritten, true)
	NewNameFormatByName("gedcom").Returns(gedcom.NameFormatGEDCOM, true)
	NewNameFormatByName("index").Returns(gedcom.NameFormatIndex, true)

	NewNameFormatByName("%L %f").Returns("%L %f", false)
	NewNameFormatByName("").Returns("", false)
}
