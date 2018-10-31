package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
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

	// When comparing date ranges we must look at the specific bounds, rather
	// than just the average.
	btw1923And1943 := gedcom.NewDateNode(nil, "Between 1923 and 1943", "", nil)
	btw1924And1934 := gedcom.NewDateNode(nil, "Between 1924 and 1934", "", nil)
	Minimum([]*gedcom.DateNode{
		btw1923And1943, // avg = 1933, value = 1923
		btw1924And1934, // avg = 1929, value = 1924
	}).Returns(btw1923And1943)
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

	// When comparing date ranges we must look at the specific bounds, rather
	// than just the average.
	btw1903And1924 := gedcom.NewDateNode(nil, "Between 1904 and 1924", "", nil)
	btw1913And1923 := gedcom.NewDateNode(nil, "Between 1913 and 1923", "", nil)
	Maximum([]*gedcom.DateNode{
		btw1903And1924, // avg = 1914, value = 1924
		btw1913And1923, // avg = 1919, value = 1923
	}).Returns(btw1903And1924)
}

func TestDateNodes_StripZero(t *testing.T) {
	StripZero := tf.Function(t, gedcom.DateNodes.StripZero)

	at3Sep1923 := gedcom.NewDateNode(nil, "3 Sep 1923", "", nil)
	at5Mar1923 := gedcom.NewDateNode(nil, "5 Mar 1923", "", nil)
	zeroDate := gedcom.NewDateNode(nil, "foo bar", "", nil)

	// Nils.
	StripZero(nil).Returns(nil)
	StripZero(gedcom.DateNodes{}).Returns(nil)

	// Valid cases.
	StripZero(gedcom.DateNodes{at3Sep1923}).
		Returns(gedcom.DateNodes{at3Sep1923})
	StripZero(gedcom.DateNodes{at3Sep1923, at3Sep1923}).
		Returns(gedcom.DateNodes{at3Sep1923, at3Sep1923})
	StripZero(gedcom.DateNodes{at3Sep1923, at5Mar1923}).
		Returns(gedcom.DateNodes{at3Sep1923, at5Mar1923})

	// With zero dates.
	StripZero(gedcom.DateNodes{zeroDate}).
		Returns(nil)
	StripZero(gedcom.DateNodes{at3Sep1923, zeroDate}).
		Returns(gedcom.DateNodes{at3Sep1923})
	StripZero(gedcom.DateNodes{zeroDate, at5Mar1923}).
		Returns(gedcom.DateNodes{at5Mar1923})
	StripZero(gedcom.DateNodes{zeroDate, at3Sep1923, zeroDate, zeroDate, at5Mar1923}).
		Returns(gedcom.DateNodes{at3Sep1923, at5Mar1923})
}
