package main

import (
	"github.com/elliotchance/tf"
	"testing"
)

func TestLengthExpr_Evaluate(t *testing.T) {
	Evaluate := tf.Function(t, (*LengthExpr).Evaluate)
	engine := &Engine{}

	Evaluate(&LengthExpr{}, engine, nil).Returns(1, nil)
	Evaluate(&LengthExpr{}, engine, []int{1, 2, 3}).Returns(3, nil)
	Evaluate(&LengthExpr{}, engine, "foo bar").Returns(1, nil)
}
