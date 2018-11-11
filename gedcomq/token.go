package main

import (
	"regexp"
	"unicode"
	"fmt"
)

type TokenKind string

const (
	TokenEOF       = TokenKind("EOF")
	TokenAccessor  = TokenKind("Accessor")
	TokenPipe      = TokenKind("Pipe")
	TokenWord      = TokenKind("Word")
	TokenIs        = TokenKind("Is")
	TokenAre       = TokenKind("Are")
	TokenSemiColon = TokenKind("SemiColon")
)

var TokenRegexp = []struct {
	re   *regexp.Regexp
	kind TokenKind
}{
	{regexp.MustCompile(`^\|$`), TokenPipe},
	{regexp.MustCompile(`^;$`), TokenSemiColon},
	{regexp.MustCompile(`^is$`), TokenIs},
	{regexp.MustCompile(`^are$`), TokenAre},
	{regexp.MustCompile(`^\.[A-Z][a-zA-Z0-9_]*$`), TokenAccessor},
	{regexp.MustCompile(`^[A-Z][a-zA-Z0-9_]*$`), TokenWord},
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
		// Ignore whitespace.
		if unicode.IsSpace(rune(s[i])) {
			continue
		}

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

				token := Token{
					Kind:  test.kind,
					Value: string(buf),
				}
				tokens = append(tokens, token)
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
		p := Token{Kind: TokenEOF}
		if t.Position < len(t.Tokens) {
			p = t.Tokens[t.Position]
		}

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

func (t *Tokens) Peek(expected ...TokenKind) (tokens []Token, err error) {
	originalPosition := t.Position

	defer func() {
		t.Position = originalPosition
	}()

	return t.Consume(expected...)
}
