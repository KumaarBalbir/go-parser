package parser

import (
	"github.com/go-parser/src/ast"
	"github.com/go-parser/src/lexer"
)

type binding_power int

const (
	default_bp binding_power = iota
	comma
	assignment
	logical
	relational
	additive
	multiplicative
	unary
	call
	member
	primary
)

type stmt_handler func(p *parser) ast.Stmt

type nud_handler func(p *parser) ast.Expr
type led_handler func(p *parser, left ast.Expr, bp binding_power) ast.Expr

type stmt_lookup map[lexer.TokenKind]stmt_handler
type nud_lookup map[lexer.TokenKind]nud_handler
type led_lookup map[lexer.TokenKind]led_handler
type bp_lookup map[lexer.TokenKind]binding_power

var bp_lu = bp_lookup{}
var nud_lu = nud_lookup{}
var led_lu = led_lookup{}
var stmt_lu = stmt_lookup{}

func led(kind lexer.TokenKind, bp binding_power, led_fn led_handler) {
	bp_lu[kind] = bp
	led_lu[kind] = led_fn
}

func nud(kind lexer.TokenKind, nud_fn nud_handler) {
	nud_lu[kind] = nud_fn
}

func stmt(kind lexer.TokenKind, stmt_fn stmt_handler) {
	bp_lu[kind] = default_bp
	stmt_lu[kind] = stmt_fn
}

// array[index] // computed expression // LED
// const foo = [1, 2, 3]; // Array/Slice literal // NUD
// let foo; []number; // TYPE_NUD

/*
This Go function, `createTokenLookups`, sets up token lookups for a parser. It defines the binding power and parsing functions for various token kinds, including:

* Assignment operators (`=`)
* Logical operators (`&&`, `||`, `..`)
* Relational operators (`<`, `>`, `==`, `!=`)
* Additive operators (`+`, `-`)
* Multiplicative operators (`*`, `/`, `%`)
* Literals and symbols (`number`, `string`, `identifier`, `(`, `-`)
* Statements (`const`, `let`)

The `led` function sets up left-denotation (infix) operators, while the `nud` function sets up null-denotation (prefix) operators. The `stmt` function sets up statement handlers.

In essence, this function tells the parser how to handle different tokens and what parsing functions to call when encountering them.
*/
func createTokenLookups() {
	led(lexer.ASSIGNMENT, assignment, parse_assignment_expr)
	led(lexer.PLUS_EQUALS, assignment, parse_assignment_expr)
	led(lexer.MINUS_EQUALS, assignment, parse_assignment_expr)

	// Logical
	led(lexer.AND, logical, parse_binary_expr)
	led(lexer.OR, logical, parse_binary_expr)
	led(lexer.DOT_DOT, logical, parse_binary_expr) // 10..math.random()

	// Relational
	led(lexer.LESS_EQUALS, relational, parse_binary_expr)
	led(lexer.LESS, relational, parse_binary_expr)
	led(lexer.GREATER_EQUALS, relational, parse_binary_expr)
	led(lexer.GREATER, relational, parse_binary_expr)
	led(lexer.NOT_EQUALS, relational, parse_binary_expr)
	led(lexer.EQUALS, relational, parse_binary_expr)

	// Additive
	led(lexer.PLUS, additive, parse_binary_expr)
	led(lexer.DASH, additive, parse_binary_expr)

	// Multiplicative
	led(lexer.STAR, multiplicative, parse_binary_expr)
	led(lexer.SLASH, multiplicative, parse_binary_expr)
	led(lexer.PERCENT, multiplicative, parse_binary_expr)

	// Literals & Symbols
	nud(lexer.NUMBER, parse_primary_expr)

	nud(lexer.STRING, parse_primary_expr)

	nud(lexer.IDENTIFIER, parse_primary_expr)
	nud(lexer.OPEN_PAREN, parse_grouping_expr)
	nud(lexer.DASH, parse_prefix_expr)

	// statements
	stmt(lexer.CONST, parse_var_decl_stmt)
	stmt(lexer.LET, parse_var_decl_stmt)

}
