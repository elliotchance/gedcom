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
func (p *Parser) ParseString(q string) (engine *Engine, err error) {
	engine = &Engine{}
	p.tokens = NewTokenizer().TokenizeString(q)

	engine.Statements, err = p.consumeStatements()
	if err != nil {
		return nil, err
	}

	if _, err := p.tokens.Consume(TokenEOF); err != nil {
		return nil, err
	}

	return engine, nil
}

//   Statements := Statement NextStatement
//               | Statement
func (p *Parser) consumeStatements() (statements []*Statement, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	statement, err := p.consumeStatement()
	if err != nil {
		return nil, err
	}

	statements = append(statements, statement)

	for {
		statement, err := p.consumeNextStatement()
		if err != nil {
			break
		}

		statements = append(statements, statement)
	}

	return
}

//   NextStatement := ";" Statement
func (p *Parser) consumeNextStatement() (_ *Statement, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	_, err = p.tokens.Consume(TokenSemiColon)
	if err != nil {
		return nil, err
	}

	return p.consumeStatement()
}

//   Statement := word [ are | is ] Expressions
//              | Expressions
func (p *Parser) consumeStatement() (statement *Statement, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	statement = &Statement{}

	if t, err := p.tokens.Consume(TokenWord, TokenAre); err == nil {
		statement.VariableName = t[0].Value
	}

	if t, err := p.tokens.Consume(TokenWord, TokenIs); err == nil {
		statement.VariableName = t[0].Value
	}

	statement.Expressions, err = p.consumeExpressions()

	return
}

//   NextExpression := "|" Expression
func (p *Parser) consumeNextExpression() (_ Expression, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	_, err = p.tokens.Consume(TokenPipe)
	if err != nil {
		return nil, err
	}

	return p.consumeExpression()
}

//   Expressions := Expression NextExpression*
func (p *Parser) consumeExpressions() (expressions []Expression, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	v, err := p.consumeExpression()
	if err != nil {
		return nil, err
	}

	expressions = append(expressions, v)

	for {
		v, err := p.consumeNextExpression()
		if err != nil {
			break
		}

		expressions = append(expressions, v)
	}

	return
}

//   Expression := Accessor | Word | QuestionMark
func (p *Parser) consumeExpression() (expression Expression, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	if expression, err = p.consumeAccessor(); err == nil {
		return expression, nil
	}

	if expression, err = p.consumeWord(); err == nil {
		return expression, nil
	}

	if expression, err = p.consumeQuestionMark(); err == nil {
		return expression, nil
	}

	return nil, errors.New("expected expression")
}

//   Accessor := accessor
func (p *Parser) consumeAccessor() (expr *AccessorExpr, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	var t []Token
	t, err = p.tokens.Consume(TokenAccessor)
	if err != nil {
		return nil, err
	}

	return &AccessorExpr{
		Query: t[0].Value,
	}, nil
}

//   Word := word
func (p *Parser) consumeWord() (expr Expression, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	var t []Token
	t, err = p.tokens.Consume(TokenWord)
	if err != nil {
		return nil, err
	}

	// Function
	if v, ok := Functions[t[0].Value]; ok {
		return v, nil
	}

	// Variable
	return &VariableExpr{Name: t[0].Value}, nil
}

//   QuestionMark := "?"
func (p *Parser) consumeQuestionMark() (expr *QuestionMarkExpr, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	_, err = p.tokens.Consume(TokenQuestionMark)
	if err != nil {
		return nil, err
	}

	return &QuestionMarkExpr{}, nil
}
