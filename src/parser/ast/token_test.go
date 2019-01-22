package ast

import "testing"

func TestToken_IsTokenKind(test *testing.T) {
	var token interface{} = Token{}
	_, instanceof := token.(TokenKind)

	if instanceof == false {
		test.Error(`Expected Token to be of TokenKind`)
		return
	}
}

func TestToken_CannotTransition(test *testing.T) {
	token := Token{}
	should, namespace := token.Transition()

	if should != false {
		test.Error(`Expected with no TransitionTo the transition should be false`)
		return
	}

	if namespace != NamespaceIdentityNone {
		test.Error(`Expected when the transition is false that an empty NamespaceIdentity is returned`)
		return
	}
}

func TestToken_CanTransition(test *testing.T) {
	token := Token{
		TransitionTo: "new-transition",
	}

	should, namespace := token.Transition()

	if should != true {
		test.Error(`Expected with TransitionTo the transition should be true`)
		return
	}

	if namespace != "new-transition" {
		test.Error(`Expected when the transition is true that NamespaceIdentity is returned`)
		return
	}
}

func TestToken_MatchAlwaysFails(test *testing.T) {
	token := Token{}
	matched, offset := token.Match("foo")

	if matched != false {
		test.Error(`Token should always fail match call`)
		return
	}

	if offset != NoTokenOffset {
		test.Error(`Token should always return NoTokenOffset`)
		return
	}
}
