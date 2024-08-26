package ast

import (
	"github.com/go-parser/src/lexer"
)

// Literal expressions
/*
The NumberExpr class represents a numeric expression in the abstract syntax tree (AST). It has a single field Value of type float64, which stores the numerical value of the expression.
*/
type NumberExpr struct {
	Value float64
}

// expr is a method that serves as a marker interface for all expression types in the abstract syntax tree (AST).
// It does not contain any methods or fields, but it allows for type assertions and polymorphism in the AST.
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

/*
This class definition defines a BinaryExpr struct in Go, which represents a binary expression in an abstract syntax tree (AST). Here's a succinct explanation of what each field does:

Left Expr: This field stores the left operand of the binary expression.
Operator lexer.Token: This field stores the operator of the binary expression, such as +, -, *, etc.
Right Expr: This field stores the right operand of the binary expression.
Note that this struct does not have any methods, it only has fields to store the operands and operator of a binary expression.
*/
// 14 + 3 * 4
type BinaryExpr struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
}

/*
This code defines a method expr on the BinaryExpr struct in Go, which satisfies the expr interface. The method is empty and serves only to indicate that BinaryExpr implements the expr interface, allowing for polymorphism in the abstract syntax tree (AST).
*/
func (n BinaryExpr) expr() {}

/*
This class definition defines a PrefixExpr struct in Go, which represents a prefix expression in an abstract syntax tree (AST). Here's a succinct explanation of what each field does:

Operator lexer.Token: This field stores the operator of the prefix expression, such as +, -, !, etc.
RightExpr Expr: This field stores the right operand of the prefix expression.
The PrefixExpr struct has a method expr() which is not shown in this snippet, but it is likely an empty method that serves to indicate that PrefixExpr implements the expr interface, allowing for polymorphism in the abstract syntax tree (AST).
*/
type PrefixExpr struct {
	Operator  lexer.Token
	RightExpr Expr
}

func (n PrefixExpr) expr() {}

// a = a + 1
// a +=5
/*
This class definition defines an AssignmentExpr struct in Go, which represents an assignment expression in an abstract syntax tree (AST). Here's a succinct explanation of what each field does:

Assigne Expr: This field stores the left-hand side of the assignment, which is the variable or expression being assigned to.
Operator lexer.Token: This field stores the operator of the assignment, such as =, +=, -=, etc.
Value Expr: This field stores the right-hand side of the assignment, which is the value being assigned to the left-hand side.
Note that this struct does not have any methods, it only has fields to store the components of an assignment expression.
*/
type AssignmentExpr struct {
	Assigne  Expr
	Operator lexer.Token
	Value    Expr
}

func (n AssignmentExpr) expr() {}
