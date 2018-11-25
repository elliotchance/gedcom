package q_test

import (
	"testing"

	"errors"
	"github.com/elliotchance/gedcom/q"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestTokenizer_TokenizeString(t *testing.T) {
	TokenizeString := tf.Function(t, (*q.Tokenizer).TokenizeString)
	tz := q.NewTokenizer()

	TokenizeString(tz, "").Returns(&q.Tokens{Tokens: []q.Token{}})
	TokenizeString(tz, "|").Returns(&q.Tokens{Tokens: []q.Token{
		{q.TokenPipe, "|"},
	}})
	TokenizeString(tz, "||").Returns(&q.Tokens{Tokens: []q.Token{
		{q.TokenPipe, "|"},
		{q.TokenPipe, "|"},
	}})
	TokenizeString(tz, ";|").Returns(&q.Tokens{Tokens: []q.Token{
		{q.TokenSemiColon, ";"},
		{q.TokenPipe, "|"},
	}})
	TokenizeString(tz, ".Foo|").Returns(&q.Tokens{Tokens: []q.Token{
		{q.TokenAccessor, ".Foo"},
		{q.TokenPipe, "|"},
	}})
	TokenizeString(tz, ".Foo | Bar").Returns(&q.Tokens{Tokens: []q.Token{
		{q.TokenAccessor, ".Foo"},
		{q.TokenPipe, "|"},
		{q.TokenWord, "Bar"},
	}})
	TokenizeString(tz, "Foo is .Individuals").Returns(&q.Tokens{Tokens: []q.Token{
		{q.TokenWord, "Foo"},
		{q.TokenWord, "is"},
		{q.TokenAccessor, ".Individuals"},
	}})
	TokenizeString(tz, "Foo are .Individuals").Returns(&q.Tokens{Tokens: []q.Token{
		{q.TokenWord, "Foo"},
		{q.TokenWord, "are"},
		{q.TokenAccessor, ".Individuals"},
	}})
	TokenizeString(tz, "Foo ?").Returns(&q.Tokens{Tokens: []q.Token{
		{q.TokenWord, "Foo"},
		{q.TokenQuestionMark, "?"},
	}})

	for _, operator := range q.Operators {
		expectedTokens := []q.Token{}
		for _, token := range operator.Tokens {
			expectedTokens = append(expectedTokens, q.Token{token, string(token)})
		}

		TokenizeString(tz, operator.Name).
			Returns(&q.Tokens{Tokens: expectedTokens})
	}
}

func TestTokens_Consume(t *testing.T) {
	for _, test := range []struct {
		s        string
		expected []q.TokenKind
		result   []q.Token
		err      error
	}{
		{
			".Foo | Bar",
			[]q.TokenKind{},
			[]q.Token(nil),
			nil,
		},
		{
			".Foo | Bar",
			[]q.TokenKind{q.TokenAccessor},
			[]q.Token{{q.TokenAccessor, ".Foo"}},
			nil,
		},
		{
			".Foo | Bar",
			[]q.TokenKind{q.TokenAccessor, q.TokenPipe},
			[]q.Token{{q.TokenAccessor, ".Foo"}, {q.TokenPipe, "|"}},
			nil,
		},
		{
			".Foo | Bar",
			[]q.TokenKind{q.TokenAccessor, q.TokenAccessor},
			[]q.Token{{q.TokenAccessor, ".Foo"}},
			errors.New("expected accessor but found |"),
		},
		{
			"Bar; Baz",
			[]q.TokenKind{q.TokenWord, q.TokenSemiColon, q.TokenWord},
			[]q.Token{
				{q.TokenWord, "Bar"},
				{q.TokenSemiColon, ";"},
				{q.TokenWord, "Baz"},
			},
			nil,
		},
		{
			".Foo | ?",
			[]q.TokenKind{q.TokenAccessor, q.TokenPipe, q.TokenQuestionMark},
			[]q.Token{
				{q.TokenAccessor, ".Foo"},
				{q.TokenPipe, "|"},
				{q.TokenQuestionMark, "?"},
			},
			nil,
		},
		{
			"First(13)",
			[]q.TokenKind{
				q.TokenWord,
				q.TokenOpenBracket,
				q.TokenNumber,
				q.TokenCloseBracket,
			},
			[]q.Token{
				{q.TokenWord, "First"},
				{q.TokenOpenBracket, "("},
				{q.TokenNumber, "13"},
				{q.TokenCloseBracket, ")"},
			},
			nil,
		},
		{
			"",
			[]q.TokenKind{q.TokenEOF},
			[]q.Token{{q.TokenEOF, ""}},
			nil,
		},
		{
			"{foo:bar}",
			[]q.TokenKind{
				q.TokenOpenCurly,
				q.TokenWord,
				q.TokenColon,
				q.TokenWord,
				q.TokenCloseCurly,
			},
			[]q.Token{
				{q.TokenOpenCurly, "{"},
				{q.TokenWord, "foo"},
				{q.TokenColon, ":"},
				{q.TokenWord, "bar"},
				{q.TokenCloseCurly, "}"},
			},
			nil,
		},
		{
			"foo, bar, baz",
			[]q.TokenKind{
				q.TokenWord,
				q.TokenComma,
				q.TokenWord,
				q.TokenComma,
				q.TokenWord,
			},
			[]q.Token{
				{q.TokenWord, "foo"},
				{q.TokenComma, ","},
				{q.TokenWord, "bar"},
				{q.TokenComma, ","},
				{q.TokenWord, "baz"},
			},
			nil,
		},
	} {
		t.Run("", func(t *testing.T) {
			tokens := q.NewTokenizer().TokenizeString(test.s)
			result, err := tokens.Consume(test.expected...)
			assert.Equal(t, test.result, result)

			if test.err == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, test.err.Error())
			}
		})
	}
}
