package q

import (
	"testing"

	"errors"
	"github.com/elliotchance/tf"
)

func TestNewParser(t *testing.T) {
	NewParser := tf.Function(t, NewParser)

	NewParser().Returns(&Parser{})
}

func TestParser_ParseString(t *testing.T) {
	ParseString := tf.Function(t, (*Parser).ParseString)
	parser := NewParser()

	ParseString(parser, "").Returns(nil, errors.New("expected expression"))

	ParseString(parser, ".Individuals").Returns(&Engine{
		Statements: []*Statement{
			{
				Expressions: []Expression{
					&AccessorExpr{Query: ".Individuals"},
				},
			},
		},
	}, nil)

	ParseString(parser, ".Individuals | .Name").Returns(&Engine{
		Statements: []*Statement{
			{
				Expressions: []Expression{
					&AccessorExpr{Query: ".Individuals"},
					&AccessorExpr{Query: ".Name"},
				},
			},
		},
	}, nil)

	ParseString(parser, "Foo is .Individuals | .Name").Returns(&Engine{
		Statements: []*Statement{
			{
				VariableName: "Foo",
				Expressions: []Expression{
					&AccessorExpr{Query: ".Individuals"},
					&AccessorExpr{Query: ".Name"},
				},
			},
		},
	}, nil)

	ParseString(parser, "Bar are .Individuals | .Name").Returns(&Engine{
		Statements: []*Statement{
			{
				VariableName: "Bar",
				Expressions: []Expression{
					&AccessorExpr{Query: ".Individuals"},
					&AccessorExpr{Query: ".Name"},
				},
			},
		},
	}, nil)

	ParseString(parser, "Foo Bar").Returns(nil,
		errors.New("expected EOF but found word"))

	ParseString(parser, "{}").Returns(&Engine{
		Statements: []*Statement{
			{
				VariableName: "",
				Expressions: []Expression{
					&ObjectExpr{},
				},
			},
		},
	}, nil)

	ParseString(parser, "{ foo: .OK }").Returns(&Engine{
		Statements: []*Statement{
			{
				Expressions: []Expression{
					&ObjectExpr{Data: map[string]*Statement{
						"foo": {
							Expressions: []Expression{
								&AccessorExpr{Query: ".OK"},
							},
						},
					}},
				},
			},
		},
	}, nil)

	ParseString(parser, "{ foo: .OK, bar: .Yes }").Returns(&Engine{
		Statements: []*Statement{
			{
				VariableName: "",
				Expressions: []Expression{
					&ObjectExpr{Data: map[string]*Statement{
						"foo": {
							Expressions: []Expression{
								&AccessorExpr{Query: ".OK"},
							},
						},
						"bar": {
							Expressions: []Expression{
								&AccessorExpr{Query: ".Yes"},
							},
						},
					}},
				},
			},
		},
	}, nil)
}
