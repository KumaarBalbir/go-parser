package parser

import (
	"fmt"
	"strconv"

	"github.com/go-parser/src/ast"
	"github.com/go-parser/src/lexer"
)

// parse_expr Parses an expression based on the given parser and binding power.
//
// p - The parser instance used to parse the expression.
// bp - The binding power of the expression.
// Return type: ast.Expr

/*
This is a recursive descent parser function in Go, specifically designed to parse expressions. Here's a succinct explanation:

* It takes a `parser` instance (`p`) and a `binding_power` (`bp`) as inputs.
* It first parses the "NUD" (Null denotation, or prefix operator) of the current token.
* If a NUD handler is found, it calls the handler to parse the expression.
* Then, it enters a loop where it checks if the current token's binding power is greater than the input `bp`.
* If it is, it parses the "LED" (Left denotation, or infix operator) of the current token.
* If an LED handler is found, it calls the handler to parse the expression, passing the current parser, the left-hand side of the expression, and the binding power of the current token.
* The loop continues until the binding power of the current token is no longer greater than the input `bp`.
* Finally, the function returns the fully parsed expression.

In essence, this function is responsible for parsing expressions with prefix and infix operators, using a recursive descent approach.
*/
func parse_expr(p *parser, bp binding_power) ast.Expr {
	// First parse the NUD
	tokenKind := p.currentTokenKind()
	nud_fn, exists := nud_lu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("NUD handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
	}

	left := nud_fn(p)

	for bp_lu[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		led_fn, exists := led_lu[tokenKind]

		if !exists {
			panic(fmt.Sprintf("LED handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
		}

		left = led_fn(p, left, bp_lu[p.currentTokenKind()])
	}

	return left
}

// parse_primary_expr Parses a primary expression from the current token.
//
// p - The parser instance used to parse the expression.
// Return type: ast.Expr

/*
This function, `parse_primary_expr`, parses a primary expression from the current token in the parser. It handles three types of primary expressions:

*   `lexer.NUMBER`: Parses a number literal and returns an `ast.NumberExpr` with the parsed float value.
*   `lexer.STRING`: Parses a string literal and returns an `ast.StringExpr` with the string value.
*   `lexer.IDENTIFIER`: Parses an identifier and returns an `ast.SymbolExpr` with the identifier's value.

If the current token is none of the above, it panics with an error message indicating an unexpected token.
*/
func parse_primary_expr(p *parser) ast.Expr {
	switch p.currentTokenKind() {
	case lexer.NUMBER:
		number, _ := strconv.ParseFloat(p.advance().Value, 64)
		return ast.NumberExpr{
			Value: number,
		}

	case lexer.STRING:
		return ast.StringExpr{
			Value: p.advance().Value,
		}

	case lexer.IDENTIFIER:
		return ast.SymbolExpr{
			Value: p.advance().Value,
		}
	default:
		panic(fmt.Sprintf("Parser:Error -> Unexpected token %s\n", lexer.TokenKindString(p.currentTokenKind())))
	}
}

// parse_binary_expr Parses a binary expression from the current token.
//
// p - The parser instance used to parse the expression.
// left - The left operand of the binary expression.
// bp - The binding power of the expression.
// Return type: ast.Expr

/*
This function parses a binary expression from the current token in the parser.
It takes the left operand of the binary expression, advances the parser to the operator token,
and then recursively parses the right operand. The function returns a BinaryExpr struct containing
the left operand, operator token, and right operand.
*/
func parse_binary_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	operatorToken := p.advance()
	right := parse_expr(p, bp)

	return ast.BinaryExpr{
		Left:     left,
		Operator: operatorToken,
		Right:    right,
	}
}

// parse_prefix_expr Parses a prefix expression from the current token.
//
// p - The parser instance used to parse the expression.
// Return type: ast.Expr

/*
This function parses a prefix expression from the current token in the parser.
It advances the parser to the operator token, recursively parses the right-hand side expression,
and returns a `PrefixExpr` struct containing the operator token and the right-hand side expression.
*/
func parse_prefix_expr(p *parser) ast.Expr {
	operatorToken := p.advance()
	rhs := parse_expr(p, default_bp)
	return ast.PrefixExpr{
		Operator:  operatorToken,
		RightExpr: rhs,
	}
}

/*
This code snippet defines a function called `parse_assignment_expr` in Go.
It takes three parameters: a pointer to a `parser` object, an `ast.Expr` object called `left`,
and a `binding_power` object.

Inside the function, it advances the parser to the next token using the `advance()` method of
the `parser` object. It then calls another function called `parse_expr` to parse the right-hand
side expression.

The function returns an `ast.AssignmentExpr` object, which is a struct that contains the
operator token, the right-hand side expression, and the left-hand side expression.

The purpose of this function is to parse and construct an assignment expression in the
abstract syntax tree (AST) of a programming language.
*/
func parse_assignment_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr {
	operatorToken := p.advance()
	rhs := parse_expr(p, bp)
	return ast.AssignmentExpr{
		Operator: operatorToken,
		Value:    rhs,
		Assigne:  left,
	}
}

/*
This code snippet defines a function called `parse_grouping_expr` in Go. The function takes a
pointer to a `parser` object as a parameter.

Inside the function, it advances the parser to the next token using the `advance()` method of
the `parser` object.

Then, it calls another function called `parse_expr` to parse the expression inside the grouping.
The result is stored in the `expression` variable.

Finally, it expects the closing parenthesis token using the `expect()` method of the `parser`
object, and returns the parsed expression.

In summary, this function parses an expression inside grouping parentheses in an abstract
syntax tree (AST) using a parser object.
*/
func parse_grouping_expr(p *parser) ast.Expr {
	p.advance() // advance past grouping start
	expression := parse_expr(p, default_bp)
	p.expect(lexer.CLOSE_PAREN) // advance past close
	return expression
}
