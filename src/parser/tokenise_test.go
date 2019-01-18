package parser

import (
	"testing"

	"./ast"
)

func TestTokenise(test *testing.T) {
	tokenset := ast.TokenSet{
		ast.TokenLiteral{Token: ast.Token{Identity: 1}, Literal: "foo"},
		ast.TokenLiteral{Token: ast.Token{Identity: 2}, Literal: "thing"},
		ast.TokenLiteral{Token: ast.Token{Identity: 3}, Literal: "some"},
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
	}

	for i, data := range dataset {
		i = i + 1

		lexemesequence, index, err := Tokenise(data.input, tokenset)

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

		if len(lexemesequence) != data.lexemes {
			test.Errorf(`[%d] LexemeSequence length %v is expected to be %v`, i, len(lexemesequence), data.lexemes)
			return
		}
	}
}

func TestTokeniseFirstToken(test *testing.T) {
	tokenset := ast.TokenSet{
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

		matched, lexeme := TokeniseFirstLexeme(data.input, tokenset)

		if matched != data.matched {
			test.Errorf(`[%d] Matched %v is expected to be %v`, i, matched, data.matched)
			return
		}

		if matched == false {
			continue
		}

		identity := lexeme.Token.GetIdentity()
		value := lexeme.Value

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
