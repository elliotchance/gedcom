package q

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// BinaryExpr evaluates a binary operator expression.
type BinaryExpr struct {
	Left, Right Expression
	Operator    string
}

// Operators contains the tokens and functions for all operators.
//
// It is important that the operators are ordered so that the operators with
// most tokens are read first. This prevents it from consuming operators that
// are subsets of others.
var Operators = []struct {
	Name     string
	Tokens   []TokenKind
	Function func(left, right interface{}) (bool, error)
}{
	{"!=", []TokenKind{TokenNot, TokenEqual}, notEqual},
	{">=", []TokenKind{TokenGreaterThan, TokenEqual}, greaterThanEqual},
	{"<=", []TokenKind{TokenLessThan, TokenEqual}, lessThanEqual},
	{"=", []TokenKind{TokenEqual}, equal},
	{">", []TokenKind{TokenGreaterThan}, greaterThan},
	{"<", []TokenKind{TokenLessThan}, lessThan},
}

func (e *BinaryExpr) Evaluate(engine *Engine, input interface{}, args []*Statement) (interface{}, error) {
	in := reflect.ValueOf(input)

	// If it is a slice we need to Evaluate each one.
	if in.Kind() == reflect.Slice {
		mutResults := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(true)), 0, in.Len())

		for i := 0; i < in.Len(); i++ {
			result, err := e.Evaluate(engine, in.Index(i).Interface(), args)
			if err != nil {
				return nil, err
			}

			mutResults = reflect.Append(mutResults, reflect.ValueOf(result))
		}

		return mutResults.Interface(), nil
	}

	left, err := e.Left.Evaluate(engine, input, args)
	if err != nil {
		return nil, err
	}

	right, err := e.Right.Evaluate(engine, input, args)
	if err != nil {
		return nil, err
	}

	for _, operator := range Operators {
		if operator.Name == e.Operator {
			return operator.Function(left, right)
		}
	}

	return nil, fmt.Errorf("no such operator: %s", e.Operator)
}

func binaryFloats(left, right string) (float64, float64, bool) {
	floatLeft, errLeft := strconv.ParseFloat(left, 64)
	floatRight, errRight := strconv.ParseFloat(right, 64)

	// Compare as numbers.
	if errLeft == nil && errRight == nil {
		return floatLeft, floatRight, true
	}

	return 0, 0, false
}

func binaryStrings(left, right interface{}) (sLeft string, sRight string) {
	if s, ok := left.(string); ok {
		sLeft = s
	} else {
		sLeft = fmt.Sprintf("%v", left)
	}

	if s, ok := right.(string); ok {
		sRight = s
	} else {
		sRight = fmt.Sprintf("%v", right)
	}

	return
}

func compareStrings(s, t string, op func(string, string) bool) bool {
	// This is not the most ideal way to do comparisons because it doesn't
	// handle more complex characters as nicely as strings.EqualFold.

	s = strings.TrimSpace(strings.ToLower(s))
	t = strings.TrimSpace(strings.ToLower(t))

	return op(s, t)
}

func equal(left, right interface{}) (bool, error) {
	sLeft, sRight := binaryStrings(left, right)

	// Compare as numbers.
	if floatLeft, floatRight, ok := binaryFloats(sLeft, sRight); ok {
		return floatLeft == floatRight, nil
	}

	compare := func(s, t string) bool {
		return s == t
	}

	// At least one side could not be converted to a float, so compare them as
	// strings.
	return compareStrings(sLeft, sRight, compare), nil
}

func notEqual(left, right interface{}) (bool, error) {
	result, err := equal(left, right)
	if err != nil {
		return false, err
	}

	return !result, nil
}

func greaterThan(left, right interface{}) (bool, error) {
	sLeft, sRight := binaryStrings(left, right)

	if floatLeft, floatRight, ok := binaryFloats(sLeft, sRight); ok {
		return floatLeft > floatRight, nil
	}

	compare := func(s, t string) bool {
		return s > t
	}

	return compareStrings(sLeft, sRight, compare), nil
}

func greaterThanEqual(left, right interface{}) (bool, error) {
	sLeft, sRight := binaryStrings(left, right)

	if floatLeft, floatRight, ok := binaryFloats(sLeft, sRight); ok {
		return floatLeft >= floatRight, nil
	}

	compare := func(s, t string) bool {
		return s >= t
	}

	return compareStrings(sLeft, sRight, compare), nil
}

func lessThan(left, right interface{}) (bool, error) {
	sLeft, sRight := binaryStrings(left, right)

	if floatLeft, floatRight, ok := binaryFloats(sLeft, sRight); ok {
		return floatLeft < floatRight, nil
	}

	compare := func(s, t string) bool {
		return s < t
	}

	return compareStrings(sLeft, sRight, compare), nil
}

func lessThanEqual(left, right interface{}) (bool, error) {
	sLeft, sRight := binaryStrings(left, right)

	if floatLeft, floatRight, ok := binaryFloats(sLeft, sRight); ok {
		return floatLeft <= floatRight, nil
	}

	compare := func(s, t string) bool {
		return s <= t
	}

	return compareStrings(sLeft, sRight, compare), nil
}
