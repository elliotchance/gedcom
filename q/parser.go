package q

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

//   AreOrIs := "are" | "is"
func (p *Parser) consumeAreOrIs() (err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	t, err := p.tokens.Consume(TokenWord)
	if err == nil && (t[0].Value == "are" || t[0].Value == "is") {
		return
	}

	return err
}

//   Statement := NamedStatement | UnnamedStatement
func (p *Parser) consumeStatement() (statement *Statement, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	if s, err := p.consumeNamedStatement(); err == nil {
		return s, nil
	}

	return p.consumeUnnamedStatement()
}

//   NamedStatement := word AreOrIs Expressions
func (p *Parser) consumeNamedStatement() (statement *Statement, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	t, err := p.tokens.Consume(TokenWord)
	if err != nil {
		return nil, err
	}

	err = p.consumeAreOrIs()
	if err != nil {
		return nil, err
	}

	exprs, err := p.consumeExpressions()
	if err != nil {
		return nil, err
	}

	return &Statement{VariableName: t[0].Value, Expressions: exprs}, nil
}

//   UnnamedStatement := Expressions
func (p *Parser) consumeUnnamedStatement() (statement *Statement, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	exprs, err := p.consumeExpressions()
	if err != nil {
		return nil, err
	}

	return &Statement{Expressions: exprs}, nil
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

//   Expressions := Expression NextExpression...
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

//   Expression := Accessor | Word | QuestionMark | BinaryExpression
func (p *Parser) consumeExpression() (expression Expression, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	if expression, err = p.consumeConstant(); err == nil {
		goto end
	}

	if expression, err = p.consumeAccessor(); err == nil {
		goto end
	}

	if expression, err = p.consumeVariableOrFunction(); err == nil {
		goto end
	}

	if expression, err = p.consumeQuestionMark(); err == nil {
		goto end
	}

	if expression, err = p.consumeObject(); err == nil {
		goto end
	}

	return nil, errors.New("expected expression")

end:
	if op, err := p.consumeOperator(); err == nil {
		if right, err := p.consumeExpression(); err == nil {
			return &BinaryExpr{
				Left:     expression,
				Operator: op,
				Right:    right,
			}, nil
		}
	}

	return expression, nil
}

//   Constant := number | string
func (p *Parser) consumeConstant() (_ *ConstantExpr, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	t, err := p.tokens.Consume(TokenNumber)
	if err == nil {
		return &ConstantExpr{Value: t[0].Value}, err
	}

	t, err = p.tokens.Consume(TokenString)
	if err == nil {
		// Trim off "".
		return &ConstantExpr{Value: t[0].Value[1 : len(t[0].Value)-1]}, err
	}

	return nil, errors.New("no constant found")
}

//   Operator := "=" | "!=" | ">" | "<" | ">=" | "<="
func (p *Parser) consumeOperator() (_ string, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	for _, operator := range Operators {
		originalPosition := p.tokens.Position
		_, err := p.tokens.Consume(operator.Tokens...)
		if err == nil {
			return operator.Name, nil
		} else {
			p.tokens.Position = originalPosition
		}
	}

	return "", errors.New("operator expected")
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

//   VariableOrFunction := word [ "(" number ")" ]
func (p *Parser) consumeVariableOrFunction() (expr Expression, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	var t []Token
	t, err = p.tokens.Consume(TokenWord)
	if err != nil {
		return nil, err
	}

	// Ignore error because function args are optional.
	args, _ := p.consumeFunctionArgs()

	// Function
	if v, ok := Functions[t[0].Value]; ok {
		return &CallExpr{v, args}, nil
	}

	// Variable
	return &VariableExpr{Name: t[0].Value}, nil
}

//   FunctionArgs := "(" Statement ")"
func (p *Parser) consumeFunctionArgs() (args []*Statement, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	_, err = p.tokens.Consume(TokenOpenBracket)
	if err != nil {
		return nil, err
	}

	statement, err := p.consumeStatement()
	if err != nil {
		return nil, err
	}

	_, err = p.tokens.Consume(TokenCloseBracket)
	if err != nil {
		return nil, err
	}

	return []*Statement{statement}, nil
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

//   Object := ObjectWithoutKeys | ObjectWithKeys
func (p *Parser) consumeObject() (expr Expression, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	if e, err := p.consumeObjectWithoutKeys(); err == nil {
		return e, nil
	}

	return p.consumeObjectWithKeys()
}

//   KeyValues := KeyValue
func (p *Parser) consumeKeyValue() (key string, value *Statement, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	t, err := p.tokens.Consume(TokenWord, TokenColon)
	if err != nil {
		return "", nil, err
	}

	value, err = p.consumeStatement()
	if err != nil {
		return "", nil, err
	}

	return t[0].Value, value, nil
}

//   ObjectWithoutKeys := "{" "}"
func (p *Parser) consumeObjectWithoutKeys() (expr Expression, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	_, err = p.tokens.Consume(TokenOpenCurly, TokenCloseCurly)

	return &ObjectExpr{}, err
}

//   ObjectWithKeys := "{" KeyValues "}"
func (p *Parser) consumeObjectWithKeys() (expr Expression, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	_, err = p.tokens.Consume(TokenOpenCurly)
	if err != nil {
		return nil, err
	}

	data, err := p.consumeKeyValues()
	if err != nil {
		return nil, err
	}

	if _, err = p.tokens.Consume(TokenCloseCurly); err != nil {
		return nil, err
	}

	return &ObjectExpr{Data: data}, nil
}

//   KeyValues := KeyValue NextKeyValue...
func (p *Parser) consumeKeyValues() (data map[string]*Statement, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	key, value, err := p.consumeKeyValue()
	if err != nil {
		return nil, err
	}

	data = map[string]*Statement{key: value}

	for {
		key, value, err := p.consumeNextKeyValue()
		if err != nil {
			break
		}

		data[key] = value
	}

	return data, nil
}

//   NextKeyValue := "," KeyValue
func (p *Parser) consumeNextKeyValue() (key string, value *Statement, err error) {
	defer p.tokens.Rollback(p.tokens.Position, &err)

	if _, err = p.tokens.Consume(TokenComma); err != nil {
		return "", nil, err
	}

	return p.consumeKeyValue()
}
