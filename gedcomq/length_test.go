package main

import (
	"github.com/elliotchance/tf"
	"testing"
)

func TestLength_Evaluate(t *testing.T) {
	Evaluate := tf.Function(t, (*Length).Evaluate)

	Evaluate(&Length{}, nil).Returns(1, nil)
	Evaluate(&Length{}, []int{1, 2, 3}).Returns(3, nil)
	Evaluate(&Length{}, "foo bar").Returns(1, nil)
}
