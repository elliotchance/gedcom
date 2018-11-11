package main

import (
	"testing"

	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
	"errors"
)

func TestTokenizer_TokenizeString(t *testing.T) {
	TokenizeString := tf.Function(t, (*Tokenizer).TokenizeString)
	tz := NewTokenizer()

	TokenizeString(tz, "").Returns(&Tokens{Tokens: []Token{}})
	TokenizeString(tz, "|").Returns(&Tokens{Tokens: []Token{
		{TokenPipe, "|"},
	}})
	TokenizeString(tz, "||").Returns(&Tokens{Tokens: []Token{
		{TokenPipe, "|"},
		{TokenPipe, "|"},
	}})
	TokenizeString(tz, ";|").Returns(&Tokens{Tokens: []Token{
		{TokenSemiColon, ";"},
		{TokenPipe, "|"},
	}})
	TokenizeString(tz, ".Foo|").Returns(&Tokens{Tokens: []Token{
		{TokenAccessor, ".Foo"},
		{TokenPipe, "|"},
	}})
	TokenizeString(tz, ".Foo | Bar").Returns(&Tokens{Tokens: []Token{
		{TokenAccessor, ".Foo"},
		{TokenPipe, "|"},
		{TokenWord, "Bar"},
	}})
	TokenizeString(tz, "Foo is .Individuals").Returns(&Tokens{Tokens: []Token{
		{TokenWord, "Foo"},
		{TokenIs, "is"},
		{TokenAccessor, ".Individuals"},
	}})
	TokenizeString(tz, "Foo are .Individuals").Returns(&Tokens{Tokens: []Token{
		{TokenWord, "Foo"},
		{TokenAre, "are"},
		{TokenAccessor, ".Individuals"},
	}})
}

func TestTokens_Consume(t *testing.T) {
	for _, test := range []struct {
		s        string
		expected []TokenKind
		result   []Token
		err      error
	}{
		{
			".Foo | Bar",
			[]TokenKind{},
			[]Token(nil),
			nil,
		},
		{
			".Foo | Bar",
			[]TokenKind{TokenAccessor},
			[]Token{{TokenAccessor, ".Foo"}},
			nil,
		},
		{
			".Foo | Bar",
			[]TokenKind{TokenAccessor, TokenPipe},
			[]Token{{TokenAccessor, ".Foo"}, {TokenPipe, "|"}},
			nil,
		},
		{
			".Foo | Bar",
			[]TokenKind{TokenAccessor, TokenAccessor},
			[]Token{{TokenAccessor, ".Foo"}},
			errors.New("expected Accessor but found Pipe"),
		},
		{
			"Bar; Baz",
			[]TokenKind{TokenWord, TokenSemiColon, TokenWord},
			[]Token{
				{TokenWord, "Bar"},
				{TokenSemiColon, ";"},
				{TokenWord, "Baz"},
			},
			nil,
		},
	} {
		t.Run("", func(t *testing.T) {
			tokens := NewTokenizer().TokenizeString(test.s)
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
