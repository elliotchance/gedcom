package gedcom_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"testing"
)

type fakeQ struct {
	value interface{}
}

func (f fakeQ) MarshalQ() interface{} {
	return f.value
}

type fakeQ2 struct {
	Value interface{}
}

func TestMarshalQ(t *testing.T) {
	MarshalQ := tf.Function(t, gedcom.MarshalQ)

	MarshalQ(nil).Returns(nil)

	MarshalQ("foo").Returns("foo")

	MarshalQ(map[string]string{
		"a": "b",
	}).Returns(map[string]string{
		"a": "b",
	})

	MarshalQ(map[string]interface{}{
		"a": fakeQ{123},
	}).Returns(map[string]interface{}{
		"a": fakeQ{123},
	})

	MarshalQ(fakeQ{123}).Returns(123)

	MarshalQ(fakeQ{fakeQ{456}}).Returns(fakeQ{456})
}
