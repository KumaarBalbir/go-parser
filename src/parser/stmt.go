package parser

import (
	"github.com/go-parser/src/ast"
	"github.com/go-parser/src/lexer"
)

/*
This function, `parse_stmt`, attempts to parse a statement in the abstract syntax tree (AST) using a parser object `p`.

Here's a step-by-step explanation:

1. It checks if there's a statement handler function (`stmt_fn`) associated with the current token kind in the parser's `stmt_lu` lookup table.
2. If a handler function exists, it calls the handler function with the parser object `p` and returns the result.
3. If no handler function exists, it falls back to parsing an expression statement using the `parse_expression_stmt` function and returns the result.

In essence, this function dispatches the parsing of a statement to a specific handler function based on the current token kind, or defaults to parsing an expression statement if no specific handler is found.
*/
func parse_stmt(p *parser) ast.Stmt {
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]

	if exists {
		return stmt_fn(p)
	}

	return parse_expression_stmt(p)

}

/*
This function, `parse_expression_stmt`, parses an expression statement in the abstract syntax tree (AST) using a parser object `p`. It does the following:

1. Parses an expression using the `parse_expr` function with the default binding power (`default_bp`).
2. Expects a semicolon (`lexer.SEMI_COLON`) token to follow the expression.
3. Returns an `ast.ExpressionStmt` object containing the parsed expression.

In essence, this function constructs an expression statement by parsing an expression and ensuring it's terminated with a semicolon.
*/
func parse_expression_stmt(p *parser) ast.ExpressionStmt {
	expression := parse_expr(p, default_bp)
	p.expect(lexer.SEMI_COLON)

	return ast.ExpressionStmt{
		Expression: expression,
	}
}

/*
This code snippet defines a function called `parse_var_decl_stmt` in Go, which parses a
variable declaration statement in the abstract syntax tree (AST) using a parser object `p`.
It handles the parsing of variable declarations with optional explicit types and assignment
values, and returns an `ast.VarDeclStmt` object containing the parsed information.
*/
func parse_var_decl_stmt(p *parser) ast.Stmt {
	var explicitType ast.Type
	startToken := p.advance().Kind
	isConstant := startToken == lexer.CONST
	varName := p.expectError(lexer.IDENTIFIER, "Inside variable declaration expected to find variable name").Value
	// Explicit type could be present
	if p.currentTokenKind() == lexer.COLON {
		p.expect(lexer.COLON)
		explicitType = parse_type(p, default_bp)
	}
	var assignmentValue ast.Expr
	if p.currentTokenKind() != lexer.SEMI_COLON {
		p.expect(lexer.ASSIGNMENT)
		assignmentValue = parse_expr(p, assignment)
	} else if explicitType == nil {
		panic("Missing either right hand side in var declaration or explicit type.")
	}

	p.expect(lexer.SEMI_COLON)
	if isConstant && assignmentValue == nil {
		panic("Cannot define constant without providing value")
	}

	return ast.VarDeclStmt{
		ExplicitType:  explicitType,
		IsConstant:    isConstant,
		VariableName:  varName,
		AssignedValue: assignmentValue,
	}
}
