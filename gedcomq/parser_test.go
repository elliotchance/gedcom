package main

import (
	"testing"

	"github.com/elliotchance/tf"
	"errors"
)

func TestNewParser(t *testing.T) {
	NewParser := tf.Function(t, NewParser)

	NewParser().Returns(&Parser{})
}

func TestParser_ParseString(t *testing.T) {
	ParseString := tf.Function(t, (*Parser).ParseString)
	parser := NewParser()

	ParseString(parser, "").Returns(nil, errors.New("expected variable name or expressions"))

	ParseString(parser, ".Individuals").Returns(&Engine{
		Variables: []*Variable{
			{
				Expressions: []Expression{
					&AccessorExpr{Query: ".Individuals"},
				},
			},
		},
	}, nil)

	ParseString(parser, ".Individuals | .Name").Returns(&Engine{
		Variables: []*Variable{
			{
				Expressions: []Expression{
					&AccessorExpr{Query: ".Individuals"},
					&AccessorExpr{Query: ".Name"},
				},
			},
		},
	}, nil)

	ParseString(parser, "Foo is .Individuals | .Name").Returns(&Engine{
		Variables: []*Variable{
			{
				Name: "Foo",
				Expressions: []Expression{
					&AccessorExpr{Query: ".Individuals"},
					&AccessorExpr{Query: ".Name"},
				},
			},
		},
	}, nil)

	ParseString(parser, "Bar are .Individuals | .Name").Returns(&Engine{
		Variables: []*Variable{
			{
				Name: "Bar",
				Expressions: []Expression{
					&AccessorExpr{Query: ".Individuals"},
					&AccessorExpr{Query: ".Name"},
				},
			},
		},
	}, nil)
}
