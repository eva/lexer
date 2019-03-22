package parser

import (
	"regexp"
	"testing"

	"./ast"
)

func TestTokenise(test *testing.T) {
	grammar := ast.Grammar{
		Namespaces: ast.NamespaceCollection{
			ast.Namespace{
				Identity: "root",
				Tokens: ast.TokenCollection{
					ast.TokenLiteral{Token: ast.Token{Identity: 1}, Literal: "foo"},
					ast.TokenLiteral{Token: ast.Token{Identity: 2}, Literal: "thing"},
					ast.TokenLiteral{Token: ast.Token{Identity: 3}, Literal: "some"},
					ast.TokenExpression{Token: ast.Token{Identity: 4}, Expression: regexp.MustCompile(`b(ox|atter)`)},
				},
			},
		},
	}

	dataset := []struct {
		input   string
		suceed  bool
		lexemes int
		index   int
	}{
		{"fo", false, 0, 0},
		{"foo", true, 1, 3},
		{"foobar", false, 1, 3},
		{"thing", true, 1, 5},
		{"some", true, 1, 4},
		{"somes", false, 1, 4},
		{"something", true, 2, 9},
		{"another-foo", false, 0, 0},
		{"somesomethingfoosome", true, 5, 20},

		{"bax", false, 0, 0},
		{"boxatter", false, 1, 3},
		{"battery", false, 1, 6},
		{"batterfooboxsome", true, 4, 16},
	}

	for i, data := range dataset {
		i = i + 1

		sequence, index, err := Tokenise(grammar, data.input)

		if data.suceed == true && err != nil {
			test.Errorf(`[%d] Expected to suceed but failed with error %v`, i, err)
			return
		}

		if data.suceed == false && err == nil {
			test.Errorf(`[%d] Expected to fail but no error was returned: %v`, i, err)
			return
		}

		if index != data.index {
			test.Errorf(`[%d] Index got to %v but expected to be %v`, i, index, data.index)
			return
		}

		if len(sequence) != data.lexemes {
			test.Errorf(`[%d] LexemeSequence length %v is expected to be %v`, i, len(sequence), data.lexemes)
			return
		}
	}
}

func TestTokeniseTraverseNamespace_StringQuoteExample(test *testing.T) {
	grammar := ast.Grammar{
		Namespaces: ast.NamespaceCollection{
			ast.Namespace{
				Identity: "root",
				Tokens: ast.TokenCollection{
					ast.TokenLiteral{Token: ast.Token{Identity: 1, TransitionTo: "string-double-quoted"}, Literal: `"`},
				},
			},
			ast.Namespace{
				Identity: "string-double-quoted",
				Tokens: ast.TokenCollection{
					ast.TokenLiteral{Token: ast.Token{Identity: 100, TransitionTo: ast.NamespaceIdentityShift}, Literal: `"`},
					ast.TokenExpression{Token: ast.Token{Identity: 2}, Expression: regexp.MustCompile(`(\\\"|[^"])+`)},
				},
			},
		},
	}

	dataset := []struct {
		input    string
		matched  bool
		haserror bool
		tokens   []ast.TokenIdentity
	}{
		{``, false, false, []ast.TokenIdentity{}},
		{`"`, true, true, []ast.TokenIdentity{1}},
		{`""`, true, false, []ast.TokenIdentity{1, 100}},
		{`"a`, true, true, []ast.TokenIdentity{1, 2}},
		{`"a"`, true, false, []ast.TokenIdentity{1, 2, 100}},
		{`"a\""`, true, false, []ast.TokenIdentity{1, 2, 100}},
	}

	for i, data := range dataset {
		i = i + 1

		sequence, _, err := Tokenise(grammar, data.input)

		if data.haserror == true && err == ErrTokeniserFinishedNotRoot {
			return
		}

		if data.matched != false && err != nil {
			test.Errorf(`[%d] Was not expecting to have an error: %v`, i, err)
			return
		}

		if len(data.tokens) != len(sequence) {
			test.Errorf(`[%d] Expecting that length of sequence is %d but got %d`, i, len(data.tokens), len(sequence))
			return
		}

		for position, lexeme := range sequence {
			token := data.tokens[position]

			if lexeme.GetTokenIdentity() != token {
				test.Errorf(`[%d] @%d Token %v was expected but got %v`, i, position, token, lexeme.GetTokenIdentity())
				return
			}
		}
	}
}

func TestTokeniseFirstToken(test *testing.T) {
	tokens := ast.TokenCollection{
		ast.TokenLiteral{Token: ast.Token{Identity: 1}, Literal: "foo"},
		ast.TokenLiteral{Token: ast.Token{Identity: 2}, Literal: "thing"},
		ast.TokenLiteral{Token: ast.Token{Identity: 3}, Literal: "some"},
	}

	dataset := []struct {
		input   string
		matched bool
		tokenid ast.TokenIdentity
		value   string
	}{
		{"foo", true, 1, "foo"},
		{"foobar", true, 1, "foo"},
		{"thing", true, 2, "thing"},
		{"some", true, 3, "some"},
		{"somes", true, 3, "some"},
		{"something", true, 3, "some"},
		{"another-foo", false, 0, ""},
	}

	for i, data := range dataset {
		i = i + 1

		namespace := ast.Namespace{Tokens: tokens}
		matched, lexeme := TokeniseFirstLexeme(0, data.input, namespace)

		if matched != data.matched {
			test.Errorf(`[%d] Matched %v is expected to be %v`, i, matched, data.matched)
			return
		}

		if matched == false {
			continue
		}

		identity := lexeme.GetTokenIdentity()
		value := lexeme.GetValue()

		if identity != data.tokenid {
			test.Errorf(`[%d] Lexeme token identity %v is expected to be %v`, i, identity, data.tokenid)
			return
		}

		if value != data.value {
			test.Errorf(`[%d] Lexeme value "%v" is expected to be "%v"`, i, value, data.value)
			return
		}
	}
}
