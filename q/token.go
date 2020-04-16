package q

import (
	"fmt"
	"regexp"
)

type TokenKind string

const (
	// Special
	TokenEOF = TokenKind("EOF")

	// Ignored
	TokenWhitespace = TokenKind("whitespace")

	// Words
	TokenAccessor = TokenKind("accessor")
	TokenWord     = TokenKind("word")
	TokenNumber   = TokenKind("number")
	TokenString   = TokenKind("string")

	// Operators
	TokenPipe         = TokenKind("|")
	TokenSemiColon    = TokenKind(";")
	TokenQuestionMark = TokenKind("?")
	TokenOpenBracket  = TokenKind("(")
	TokenCloseBracket = TokenKind(")")
	TokenOpenCurly    = TokenKind("{")
	TokenCloseCurly   = TokenKind("}")
	TokenColon        = TokenKind(":")
	TokenComma        = TokenKind(",")
	TokenEqual        = TokenKind("=")
	TokenNot          = TokenKind("!")
	TokenGreaterThan  = TokenKind(">")
	TokenLessThan     = TokenKind("<")
)

var TokenRegexp = []struct {
	re   *regexp.Regexp
	kind TokenKind
}{
	{regexp.MustCompile(`^\s+$`), TokenWhitespace},
	{regexp.MustCompile(`^\|$`), TokenPipe},
	{regexp.MustCompile(`^;$`), TokenSemiColon},
	{regexp.MustCompile(`^\?$`), TokenQuestionMark},
	{regexp.MustCompile(`^\($`), TokenOpenBracket},
	{regexp.MustCompile(`^\)$`), TokenCloseBracket},
	{regexp.MustCompile(`^\{$`), TokenOpenCurly},
	{regexp.MustCompile(`^\}$`), TokenCloseCurly},
	{regexp.MustCompile(`^:$`), TokenColon},
	{regexp.MustCompile(`^,$`), TokenComma},
	{regexp.MustCompile(`^!$`), TokenNot},
	{regexp.MustCompile(`^=$`), TokenEqual},
	{regexp.MustCompile(`^>$`), TokenGreaterThan},
	{regexp.MustCompile(`^<$`), TokenLessThan},
	{regexp.MustCompile(`^".*"$`), TokenString},
	{regexp.MustCompile(`^\.[a-zA-Z0-9_]*$`), TokenAccessor},
	{regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`), TokenWord},
	{regexp.MustCompile(`^[0-9]+$`), TokenNumber},
}

type Token struct {
	Kind  TokenKind
	Value string
}

type Tokenizer struct{}

func NewTokenizer() *Tokenizer {
	return &Tokenizer{}
}

func (t *Tokenizer) TokenizeString(s string) *Tokens {
	tokens := []Token{}
	buf := []byte{}

Begin:
	for i := 0; i < len(s); i++ {
		buf = append(buf, s[i])

		// Try to match a token. At this point it may be possible to match
		// multiple tokens which is why its important that we check them in
		// order. The first match always wins.
		for _, test := range TokenRegexp {
			if test.re.Match(buf) {
				// Now attempt to consume as many characters as we can that
				// still match the regexp.
				for ; i+1 < len(s) && test.re.Match(append(buf, s[i+1])); i++ {
					buf = append(buf, s[i+1])
				}

				if test.kind != TokenWhitespace {
					token := Token{
						Kind:  test.kind,
						Value: string(buf),
					}

					tokens = append(tokens, token)
				}

				buf = nil
				continue Begin
			}
		}
	}

	return &Tokens{tokens, 0}
}

type Tokens struct {
	Tokens   []Token
	Position int
}

func (t *Tokens) Consume(expected ...TokenKind) (tokens []Token, err error) {
	// Attempt to consume the tokens. If something goes wrong the Position is
	// not moved forward and an error is returned.
	originalPosition := t.Position

	for _, kind := range expected {
		p := t.tokenOrEOF()

		if p.Kind == kind {
			tokens = append(tokens, p)
			t.Position++
		} else {
			t.Position = originalPosition
			err = fmt.Errorf("expected %s but found %s", kind, p.Kind)
		}
	}

	return
}

func (t *Tokens) tokenOrEOF() Token {
	if t.Position < len(t.Tokens) {
		return t.Tokens[t.Position]
	}

	return Token{Kind: TokenEOF}
}

func (t *Tokens) Rollback(position int, err *error) {
	if *err != nil {
		t.Position = position
	}
}
