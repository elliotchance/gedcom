package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"testing"
)

func TestDateNodes_Minimum(t *testing.T) {
	Minimum := tf.Function(t, gedcom.DateNodes.Minimum)

	at3Sep1923 := gedcom.NewDateNode(nil, "3 Sep 1923", "", nil)
	at4Mar1923 := gedcom.NewDateNode(nil, "4 Mar 1923", "", nil)
	at5Mar1923 := gedcom.NewDateNode(nil, "5 Mar 1923", "", nil)

	// Nils
	Minimum(([]*gedcom.DateNode)(nil)).Returns((*gedcom.DateNode)(nil))
	Minimum([]*gedcom.DateNode{}).Returns((*gedcom.DateNode)(nil))

	// Values
	Minimum([]*gedcom.DateNode{
		at3Sep1923,
	}).Returns(at3Sep1923)

	Minimum([]*gedcom.DateNode{
		at3Sep1923,
		at4Mar1923,
	}).Returns(at4Mar1923)

	Minimum([]*gedcom.DateNode{
		at3Sep1923,
		at4Mar1923,
		at5Mar1923,
	}).Returns(at4Mar1923)
}

func TestDateNodes_Maximum(t *testing.T) {
	Maximum := tf.Function(t, gedcom.DateNodes.Maximum)

	at3Sep1923 := gedcom.NewDateNode(nil, "3 Sep 1923", "", nil)
	at4Mar1923 := gedcom.NewDateNode(nil, "4 Mar 1923", "", nil)
	at5Mar1923 := gedcom.NewDateNode(nil, "5 Mar 1923", "", nil)

	// Nils
	Maximum(([]*gedcom.DateNode)(nil)).Returns((*gedcom.DateNode)(nil))
	Maximum([]*gedcom.DateNode{}).Returns((*gedcom.DateNode)(nil))

	// Values
	Maximum([]*gedcom.DateNode{
		at3Sep1923,
	}).Returns(at3Sep1923)

	Maximum([]*gedcom.DateNode{
		at3Sep1923,
		at4Mar1923,
	}).Returns(at3Sep1923)

	Maximum([]*gedcom.DateNode{
		at4Mar1923,
		at3Sep1923,
		at5Mar1923,
	}).Returns(at3Sep1923)
}
