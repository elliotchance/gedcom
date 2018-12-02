package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"testing"
)

func TestGEDCOMLine(t *testing.T) {
	GEDCOMLine := tf.Function(t, gedcom.GEDCOMLine)
	nameNode := gedcom.NewNameNode(nil, "Joe /Bloggs/", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "3 Sep 1943", "", nil),
	})

	GEDCOMLine(nil, 0).Returns("")

	GEDCOMLine(nameNode, gedcom.NoIndent).Returns("NAME Joe /Bloggs/")

	GEDCOMLine(nameNode, 0).Returns("0 NAME Joe /Bloggs/")

	GEDCOMLine(nameNode, 5).Returns("5 NAME Joe /Bloggs/")
}

func TestGEDCOMString(t *testing.T) {
	GEDCOMString := tf.Function(t, gedcom.GEDCOMString)
	nameNode := gedcom.NewNameNode(nil, "Joe /Bloggs/", "", []gedcom.Node{
		gedcom.NewDateNode(nil, "3 Sep 1943", "", nil),
	})

	GEDCOMString(nil, 0).Returns("")

	GEDCOMString(nameNode, gedcom.NoIndent).
		Returns("NAME Joe /Bloggs/\nDATE 3 Sep 1943\n")

	GEDCOMString(nameNode, 0).
		Returns("0 NAME Joe /Bloggs/\n1 DATE 3 Sep 1943\n")

	GEDCOMString(nameNode, 5).
		Returns("5 NAME Joe /Bloggs/\n6 DATE 3 Sep 1943\n")
}
