package q_test

import (
	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"github.com/elliotchance/gedcom/q"
)

func TestTypeOfSliceElement(t *testing.T) {
	TypeOfSliceElement := tf.Function(t, q.TypeOfSliceElement)

	TypeOfSliceElement([]string{"foo", "bar"}).Returns(reflect.TypeOf(""))
	TypeOfSliceElement([]int{1, 2, 3}).Returns(reflect.TypeOf(0))
	TypeOfSliceElement([]int{}).Returns(reflect.TypeOf(0))
	TypeOfSliceElement([]interface{}{12.3, "bar"}).Returns(reflect.TypeOf(nil))
	TypeOfSliceElement(123).Returns(nil)
	TypeOfSliceElement(gedcom.IndividualNode{}).Returns(nil)
	TypeOfSliceElement(&gedcom.IndividualNode{}).Returns(nil)
	TypeOfSliceElement(gedcom.IndividualNodes{}).Returns(reflect.TypeOf(&gedcom.IndividualNode{}))
}

func TestValueToPointer(t *testing.T) {
	actual := q.ValueToPointer(reflect.ValueOf(3.5))
	assert.Equal(t, *actual.Interface().(*float64), 3.5)

	actual = q.ValueToPointer(reflect.ValueOf(123))
	assert.Equal(t, *actual.Interface().(*int), 123)

	a := "foo"
	actual = q.ValueToPointer(reflect.ValueOf(&a))
	assert.Equal(t, *actual.Interface().(*string), "foo")
}
