package ast

import (
	"github.com/go-parser/src/lexer"
)

// Literal expressions
type NumberExpr struct {
	Value float64
}

func (n NumberExpr) expr() {}

type StringExpr struct {
	Value string
}

func (n StringExpr) expr() {}

type SymbolExpr struct {
	Value string
}

func (n SymbolExpr) expr() {}

// COMPLEX EXPRESSIONS

// 14 + 3 * 4
type BinaryExpr struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
}

func (n BinaryExpr) expr() {}

type PrefixExpr struct {
	Operator  lexer.Token
	RightExpr Expr
}

func (n PrefixExpr) expr() {}

// a = a + 1
// a +=5
type AssignmentExpr struct {
	Assigne  Expr
	Operator lexer.Token
	Value    Expr
}

func (n AssignmentExpr) expr() {}
