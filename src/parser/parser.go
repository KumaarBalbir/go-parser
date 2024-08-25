package parser 
import (
	"github.com/go-parser/src/lexer"
	"github.com/go-parser/src/ast"
	"fmt"
)

type parser struct {
	
	tokens []lexer.Token 
	pos int 
}

func createParser(tokens []lexer.Token) *parser{
	createTokenLookups()
	return &parser{
		tokens: tokens, pos:0,
	}
}

func Parse(source string)ast.BlockStmt {
	tokens := lexer.Tokenize(source)
	p := createParser(tokens)
  body := make([]ast.Stmt,0)
	// while we have tokens, continue to parse
	for p.hasTokens() {
		body = append(body,parse_stmt(p))
	}

	return ast.BlockStmt{
		Body: body,
	}
}

// HELPER FUNCTIONS

func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

func (p *parser) currentTokenKind() lexer.TokenKind {
	return p.currentToken().Kind
}

func (p *parser) advance() lexer.Token {

	tk := p.currentToken()
	p.pos++
	return tk
}

func (p *parser) hasTokens() bool {
	return p.pos <len(p.tokens) && p.currentTokenKind() != lexer.EOF
}

func (p *parser) expectError (expectedKind lexer.TokenKind, err any) lexer.Token {
	token := p.currentToken()
	kind := token.Kind 

	if kind != expectedKind {
   if err == nil{
		err = fmt.Sprintf("Expected %s, but received %s instead\n", lexer.TokenKindString(expectedKind), lexer.TokenKindString(kind))

	 }
	 panic(err)
	}

	return p.advance()

}

func (p *parser) expect(expectedKind lexer.TokenKind) lexer.Token {
  return p.expectError(expectedKind, nil)
}