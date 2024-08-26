package parser

import (
	"github.com/go-parser/src/ast"
	"github.com/go-parser/src/lexer"
)

func parse_stmt(p *parser) ast.Stmt {
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]

	if exists {
		return stmt_fn(p)
	}

	return parse_expression_stmt(p)

}

func parse_expression_stmt(p *parser) ast.ExpressionStmt {
	expression := parse_expr(p, default_bp)
	p.expect(lexer.SEMI_COLON)

	return ast.ExpressionStmt{
		Expression: expression,
	}
}

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
