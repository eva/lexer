package php

import (
	"../../parser/ast"
)

const (
	TokenAdd ast.TokenIdentity = iota + 1
)

var root = ast.NewNamespaceRoot()

var a ast.TokenSet = ast.TokenSet{
	ast.NewTokenLiteral(TokenAdd, root, "+"),
}
