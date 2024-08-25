package parser 
import (
	"github.com/go-parser/src/ast"
	"github.com/go-parser/src/lexer"

)

func parse_stmt (p *parser) ast.Stmt {
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]

	if exists{
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