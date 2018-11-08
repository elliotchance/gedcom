package main

import (
	"github.com/elliotchance/tf"
	"testing"
)

type MyStruct struct {
	Property int
}

func (ms *MyStruct) Foo() string {
	return "bar"
}

func (ms MyStruct) Baz() []string {
	return []string{"qux", "quux"}
}

func TestAccessor_Evaluate(t *testing.T) {
	Evaluate := tf.NamedFunction(t, "Accessor_Evaluate", (*Accessor).Evaluate)

	ms1 := &MyStruct{Property: 123}
	ms2 := MyStruct{Property: 456}

	Evaluate(&Accessor{Query: ".Foo"}, ms1).Returns("bar", nil)
	Evaluate(&Accessor{Query: ".Foo"}, ms2).Returns("bar", nil)

	Evaluate(&Accessor{Query: ".Baz"}, ms1).Returns([]string{"qux", "quux"}, nil)
	Evaluate(&Accessor{Query: ".Baz"}, ms2).Returns([]string{"qux", "quux"}, nil)

	Evaluate(&Accessor{Query: ".Property"}, ms1).Returns(123, nil)
	Evaluate(&Accessor{Query: ".Property"}, ms2).Returns(456, nil)

	Evaluate(&Accessor{Query: ".Missing"}, ms1).
		Errors(`MyStruct does not have a method or property named "Missing"`)
	Evaluate(&Accessor{Query: ".Missing"}, ms2).
		Errors(`MyStruct does not have a method or property named "Missing"`)
}
