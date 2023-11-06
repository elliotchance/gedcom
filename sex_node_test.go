package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom/v39"
	"github.com/elliotchance/tf"
)

func TestSexNode_String(t *testing.T) {
	String := tf.Function(t, (*gedcom.SexNode).String)

	String(gedcom.NewSexNode("")).Returns("Unknown")
	String(gedcom.NewSexNode(gedcom.SexUnknown)).Returns("Unknown")
	String(gedcom.NewSexNode(gedcom.SexMale)).Returns("Male")
	String(gedcom.NewSexNode(gedcom.SexFemale)).Returns("Female")
}

func TestSexNode_OwnershipWord(t *testing.T) {
	OwnershipWord := tf.Function(t, (*gedcom.SexNode).OwnershipWord)

	OwnershipWord(gedcom.NewSexNode("")).Returns("their")
	OwnershipWord(gedcom.NewSexNode(gedcom.SexUnknown)).Returns("their")
	OwnershipWord(gedcom.NewSexNode(gedcom.SexMale)).Returns("his")
	OwnershipWord(gedcom.NewSexNode(gedcom.SexFemale)).Returns("her")
}
