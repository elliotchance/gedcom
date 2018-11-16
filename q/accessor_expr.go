package q

import (
	"fmt"
	"reflect"
)

// AccessorExpr is used to fetch the value of a property or to invoke a method.
//
// The simplest form is ".Foo" where Foo could be a property or method.
//
// When an accessor is used on a slice the accessor is performed on each
// element, generating a new slice of that returned type.
type AccessorExpr struct {
	Query string
}

// Evaluate  will automatically handle conversions between pointer and
// non-pointers to find the property or method and return the value. However, if
// it is a method it must not take any arguments.
//
// It will return an error if a property or method could not be found by that
// name.
func (e *AccessorExpr) Evaluate(engine *Engine, input interface{}) (interface{}, error) {
	in := reflect.ValueOf(input)
	accessor := e.Query[1:]

	// If it is a slice we need to Evaluate each one.
	if in.Kind() == reflect.Slice {
		t := TypeOfSliceElement(input)
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		returnType := e.getReturnType(accessor, reflect.New(t).Interface())

		results := reflect.MakeSlice(reflect.SliceOf(returnType), 0, 0)

		for i := 0; i < in.Len(); i++ {
			result, err := e.Evaluate(engine, in.Index(i).Interface())
			if err != nil {
				return nil, err
			}

			results = reflect.Append(results, reflect.ValueOf(result))
		}

		return results.Interface(), nil
	}

	var err error
	input, err = e.evaluateAccessor(accessor, input)

	if err != nil {
		return nil, err
	}

	return input, nil
}

func (e *AccessorExpr) evaluateAccessor(accessor string, input interface{}) (r interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("cannot use .%s on %s", accessor, getType(input))
		}
	}()

	method, field := e.getMethodOrField(accessor, input)

	switch {
	case method != nil:
		return callMethod(*method)

	case field != nil:
		return field.Interface(), nil
	}

	return nil, fmt.Errorf(`%s does not have a method or property named "%s"`,
		getType(input), accessor)
}

func (e *AccessorExpr) getReturnType(accessor string, input interface{}) reflect.Type {
	method, field := e.getMethodOrField(accessor, input)

	switch {
	case method != nil:
		return (*method).Type().Out(0)

	case field != nil:
		return field.Type()
	}

	return nil
}

func (e *AccessorExpr) getMethodOrField(accessor string, input interface{}) (*reflect.Value, *reflect.Value) {
	method := e.getMethod(accessor, input)

	if method != nil {
		return method, nil
	}

	field := e.getField(accessor, input)

	return nil, field
}

func (e *AccessorExpr) getMethod(accessor string, input interface{}) *reflect.Value {
	defer func() {
		// The nil return value will be handled higher up.
		recover()
	}()

	in := ValueToPointer(reflect.ValueOf(input))

	s := in.String()
	s += ""

	// Try the method on the pointer.
	methodByName := in.MethodByName(accessor)
	if methodByName.IsValid() {
		return &methodByName
	}

	// If that doesn't work, try calling it on the dereferenced value.
	methodByName = in.Elem().MethodByName(accessor)
	if methodByName.IsValid() {
		return &methodByName
	}

	return nil
}

func (e *AccessorExpr) getField(accessor string, input interface{}) *reflect.Value {
	defer func() {
		// The nil return value will be handled higher up.
		recover()
	}()

	in := ValueToPointer(reflect.ValueOf(input))

	// Try the property on the pointer.
	fieldByName := in.Elem().FieldByName(accessor)
	if fieldByName.IsValid() {
		return &fieldByName
	}

	// Try the property on the struct.
	if in.Kind() == reflect.Struct {
		fieldByName = in.FieldByName(accessor)
		if fieldByName.IsValid() {
			return &fieldByName
		}
	}

	return nil
}

func callMethod(methodByName reflect.Value) (interface{}, error) {
	result := methodByName.Call([]reflect.Value{})

	return result[0].Interface(), nil
}
