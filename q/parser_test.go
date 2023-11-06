package q_test

import (
	"testing"

	"errors"

	"github.com/elliotchance/gedcom/v39/q"
	"github.com/elliotchance/tf"
)

func TestNewParser(t *testing.T) {
	NewParser := tf.Function(t, q.NewParser)

	NewParser().Returns(&q.Parser{})
}

func TestParser_ParseString(t *testing.T) {
	ParseString := tf.Function(t, (*q.Parser).ParseString)
	parser := q.NewParser()

	ParseString(parser, "").Returns(nil, errors.New("expected expression"))

	ParseString(parser, ".Individuals").Returns(&q.Engine{
		Statements: []*q.Statement{
			{
				Expressions: []q.Expression{
					&q.AccessorExpr{Query: ".Individuals"},
				},
			},
		},
	}, nil)

	ParseString(parser, ".Individuals | .Name").Returns(&q.Engine{
		Statements: []*q.Statement{
			{
				Expressions: []q.Expression{
					&q.AccessorExpr{Query: ".Individuals"},
					&q.AccessorExpr{Query: ".Name"},
				},
			},
		},
	}, nil)

	ParseString(parser, "Foo is .Individuals | .Name").Returns(&q.Engine{
		Statements: []*q.Statement{
			{
				VariableName: "Foo",
				Expressions: []q.Expression{
					&q.AccessorExpr{Query: ".Individuals"},
					&q.AccessorExpr{Query: ".Name"},
				},
			},
		},
	}, nil)

	ParseString(parser, "Bar are .Individuals | .Name").Returns(&q.Engine{
		Statements: []*q.Statement{
			{
				VariableName: "Bar",
				Expressions: []q.Expression{
					&q.AccessorExpr{Query: ".Individuals"},
					&q.AccessorExpr{Query: ".Name"},
				},
			},
		},
	}, nil)

	ParseString(parser, "Foo Bar").Returns(nil,
		errors.New("expected EOF but found word"))

	ParseString(parser, "{}").Returns(&q.Engine{
		Statements: []*q.Statement{
			{
				VariableName: "",
				Expressions: []q.Expression{
					&q.ObjectExpr{},
				},
			},
		},
	}, nil)

	ParseString(parser, "{ foo: .OK }").Returns(&q.Engine{
		Statements: []*q.Statement{
			{
				Expressions: []q.Expression{
					&q.ObjectExpr{Data: map[string]*q.Statement{
						"foo": {
							Expressions: []q.Expression{
								&q.AccessorExpr{Query: ".OK"},
							},
						},
					}},
				},
			},
		},
	}, nil)

	ParseString(parser, "{ foo: .OK, bar: .Yes }").Returns(&q.Engine{
		Statements: []*q.Statement{
			{
				VariableName: "",
				Expressions: []q.Expression{
					&q.ObjectExpr{Data: map[string]*q.Statement{
						"foo": {
							Expressions: []q.Expression{
								&q.AccessorExpr{Query: ".OK"},
							},
						},
						"bar": {
							Expressions: []q.Expression{
								&q.AccessorExpr{Query: ".Yes"},
							},
						},
					}},
				},
			},
		},
	}, nil)

	ParseString(parser, ".Foo = .Bar").Returns(&q.Engine{
		Statements: []*q.Statement{
			{
				Expressions: []q.Expression{
					&q.BinaryExpr{
						Left:     &q.AccessorExpr{Query: ".Foo"},
						Operator: "=",
						Right:    &q.AccessorExpr{Query: ".Bar"},
					},
				},
			},
		},
	}, nil)

	ParseString(parser, ".Foo != 3").Returns(&q.Engine{
		Statements: []*q.Statement{
			{
				Expressions: []q.Expression{
					&q.BinaryExpr{
						Left:     &q.AccessorExpr{Query: ".Foo"},
						Operator: "!=",
						Right:    &q.ConstantExpr{Value: "3"},
					},
				},
			},
		},
	}, nil)

	for _, operator := range q.Operators {
		ParseString(parser, `"foo"`+operator.Name+`"3.12"`).Returns(&q.Engine{
			Statements: []*q.Statement{
				{
					Expressions: []q.Expression{
						&q.BinaryExpr{
							Left:     &q.ConstantExpr{Value: "foo"},
							Operator: operator.Name,
							Right:    &q.ConstantExpr{Value: "3.12"},
						},
					},
				},
			},
		}, nil)
	}

	ParseString(parser, `Only(.Foo = "bar")`).Returns(&q.Engine{
		Statements: []*q.Statement{
			{
				Expressions: []q.Expression{
					&q.CallExpr{
						Function: &q.OnlyExpr{},
						Args: []*q.Statement{
							{
								Expressions: []q.Expression{
									&q.BinaryExpr{
										Left:     &q.AccessorExpr{Query: ".Foo"},
										Operator: "=",
										Right:    &q.ConstantExpr{Value: "bar"},
									},
								},
							},
						},
					},
				},
			},
		},
	}, nil)

	ParseString(parser, `Combine(.Individuals, .Individuals)`).Returns(&q.Engine{
		Statements: []*q.Statement{
			{
				Expressions: []q.Expression{
					&q.CallExpr{
						Function: &q.CombineExpr{},
						Args: []*q.Statement{
							{
								Expressions: []q.Expression{
									&q.AccessorExpr{Query: ".Individuals"},
								},
							},
							{
								Expressions: []q.Expression{
									&q.AccessorExpr{Query: ".Individuals"},
								},
							},
						},
					},
				},
			},
		},
	}, nil)

	ParseString(parser, `Combine(.Individuals | .Names, .Individuals | .Names)`).Returns(&q.Engine{
		Statements: []*q.Statement{
			{
				Expressions: []q.Expression{
					&q.CallExpr{
						Function: &q.CombineExpr{},
						Args: []*q.Statement{
							{
								Expressions: []q.Expression{
									&q.AccessorExpr{Query: ".Individuals"},
									&q.AccessorExpr{Query: ".Names"},
								},
							},
							{
								Expressions: []q.Expression{
									&q.AccessorExpr{Query: ".Individuals"},
									&q.AccessorExpr{Query: ".Names"},
								},
							},
						},
					},
				},
			},
		},
	}, nil)
}
