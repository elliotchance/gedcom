package main

import (
	"testing"

	"github.com/elliotchance/tf"
)

func TestNewParser(t *testing.T) {
	NewParser := tf.Function(t, NewParser)

	NewParser().Returns(&Parser{})
}

func TestParser_ParseString(t *testing.T) {
	ParseString := tf.Function(t, (*Parser).ParseString)
	parser := NewParser()

	ParseString(parser, "").Returns(&Engine{})

	ParseString(parser, ".Individuals").Returns(&Engine{
		Expressions: []Expression{
			&Accessor{Query: ".Individuals"},
		},
	})

	ParseString(parser, ".Individuals | .Name").Returns(&Engine{
		Expressions: []Expression{
			&Accessor{Query: ".Individuals"},
			&Accessor{Query: ".Name"},
		},
	})
}
