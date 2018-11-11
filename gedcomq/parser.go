package main

import (
	"errors"
)

// Parser converts the query string into an Engine that can be evaluated.
type Parser struct {
	tokens *Tokens
}

// NewParser creates a new parser.
func NewParser() *Parser {
	return &Parser{}
}

// ParseString returns a new Engine by parsing the query string.
func (p *Parser) ParseString(q string) (*Engine, error) {
	engine := NewEngine()
	p.tokens = NewTokenizer().TokenizeString(q)

	variable, err := p.consumeVariable()
	if err != nil {
		return nil, err
	}

	engine.Variables = append(engine.Variables, variable)

	for {
		if _, err := p.tokens.Consume(TokenEOF); err == nil {
			break
		}

		_, err := p.tokens.Consume(TokenSemiColon)
		if err != nil {
			return nil, err
		}

		variable, err := p.consumeVariable()
		if err != nil {
			return nil, err
		}

		engine.Variables = append(engine.Variables, variable)
	}

	return engine, nil
}

func (p *Parser) consumeVariable() (v *Variable, err error) {
	if v, err = p.consumeNamedVariable(); err == nil {
		return
	}

	if v, err = p.consumeUnnamedVariable(); err == nil {
		return
	}

	return nil, errors.New("expected variable name or expressions")
}

func (p *Parser) consumeNamedVariable() (variable *Variable, err error) {
	variable = &Variable{}
	var tokens []Token

	tokens, err = p.tokens.Consume(TokenWord, TokenIs)
	if err == nil {
		variable.Name = tokens[0].Value
		variable.Expressions, err = p.consumeExpressions()

		return
	}

	tokens, err = p.tokens.Consume(TokenWord, TokenAre)
	if err == nil {
		variable.Name = tokens[0].Value
		variable.Expressions, err = p.consumeExpressions()

		return
	}

	return nil, errors.New("expected Variable")
}

func (p *Parser) consumeUnnamedVariable() (variable *Variable, err error) {
	variable = &Variable{}
	variable.Expressions, err = p.consumeExpressions()

	return
}

func (p *Parser) consumeExpressions() ([]Expression, error) {
	expressions := []Expression{}

	v, err := p.consumeExpression()
	if err != nil {
		return nil, err
	}

	expressions = append(expressions, v)

	for {
		if _, err := p.tokens.Consume(TokenEOF); err == nil {
			break
		}

		if _, err := p.tokens.Peek(TokenSemiColon); err == nil {
			break
		}

		_, err := p.tokens.Consume(TokenPipe)
		if err != nil {
			return nil, err
		}

		v, err := p.consumeExpression()
		if err != nil {
			return nil, err
		}

		expressions = append(expressions, v)
	}

	return expressions, nil
}

func (p *Parser) consumeExpression() (Expression, error) {
	if v, err := p.consumeAccessor(); err == nil {
		return v, nil
	}

	if v, err := p.consumeWord(); err == nil {
		return v, nil
	}

	return nil, errors.New("expected accessor, function or variable")
}

func (p *Parser) consumeAccessor() (*AccessorExpr, error) {
	t, err := p.tokens.Consume(TokenAccessor)
	if err != nil {
		return nil, err
	}

	return &AccessorExpr{
		Query: t[0].Value,
	}, nil
}

func (p *Parser) consumeWord() (Expression, error) {
	t, err := p.tokens.Consume(TokenWord)
	if err != nil {
		return nil, err
	}

	// Function
	if v, ok := Functions[t[0].Value]; ok {
		return v, nil
	}

	// Variable
	return &VariableExpr{VariableName: t[0].Value}, nil
}
