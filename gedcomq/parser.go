package main

import "strings"

// Parser converts the query string into an Engine that can be evaluated.
type Parser struct{}

// NewParser creates a new parser.
func NewParser() *Parser {
	return &Parser{}
}

// ParseString returns a new Engine by parsing the query string.
func (p *Parser) ParseString(q string) *Engine {
	engine := &Engine{}

	if q == "" {
		return engine
	}

	for _, e := range strings.Split(q, "|") {
		expression := getExpression(e)

		engine.Expressions = append(engine.Expressions, expression)
	}

	return engine
}

func getExpression(e string) Expression {
	q := strings.TrimSpace(e)

	if q[0] == '.' {
		return &Accessor{
			Query: q,
		}
	}

	return Functions[q]
}
