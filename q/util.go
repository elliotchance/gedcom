package q

import (
	"reflect"
)

func getType(v interface{}) string {
	t := reflect.TypeOf(v)

	if t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	}

	return t.Name()
}

// TypeOfSliceElement returns the type of element from a slice. The input should
// not be a reflect.Value, but an actual value.
//
// If v is not a slice then nil is returned.
func TypeOfSliceElement(v interface{}) reflect.Type {
	vType := reflect.TypeOf(v)

	if vType.Kind() != reflect.Slice {
		return nil
	}

	e := vType.Elem()
	if e.String() == "interface {}" {
		return nil
	}

	return e
}

// ValueToPointer converts a value to a pointer. If the value is already a
// pointer then it is passed through.
func ValueToPointer(v reflect.Value) reflect.Value {
	if v.Kind() != reflect.Ptr {
		vp := reflect.New(v.Type())
		vp.Elem().Set(v)

		return vp
	}

	return v
}
