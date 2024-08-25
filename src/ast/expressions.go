package ast 
import "github.com/go-parser/src/lexer"

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
type BinaryExpr struct{
	Left Expr 
	operator lexer.Token 
	Right Expr 
}
func (n BinaryExpr) expr() {}