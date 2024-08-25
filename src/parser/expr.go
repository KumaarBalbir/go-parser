package parser 

import (
	"github.com/go-parser/src/ast"
	"fmt"
	"strconv"
	"github.com/go-parser/src/lexer"
)


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

		if !exists{
			panic(fmt.Sprintf("LED handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
		}

		left = led_fn(p, left, bp_lu[p.currentTokenKind()])
	}

	return left

	// While we have a LED and the current bp is < bp of current token
	// continue parsing left hand side  
}


func parse_primary_expr(p *parser) ast.Expr {
	switch p.currentTokenKind(){
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

func parse_binary_expr(p *parser, left ast.Expr, bp binding_power) ast.Expr{
	operatorToken := p.advance()
	right := parse_expr(p,bp)

	return ast.BinaryExpr{
		Left: left,
		Operator: operatorToken,
		Right: right,
	}
}