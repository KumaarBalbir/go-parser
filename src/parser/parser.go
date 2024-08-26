package parser

import (
	"fmt"
	"github.com/go-parser/src/ast"
	"github.com/go-parser/src/lexer"
)

/*
This class definition defines a struct called `parser` in Go. Here's a succinct explanation of what each field does:

* `tokens []lexer.Token`: This field stores a list of tokens found in the source code being parsed.
* `pos int`: This field stores the current position in the token list.

Note that this struct does not have any methods, it only has fields to store the tokens and the current position.
*/
type parser struct {
	tokens []lexer.Token
	pos    int
}

/*
This is a Go function named `createParser` that creates and returns a new instance of
the `parser` struct. The function takes a slice of `lexer.Token` objects as input and
initializes the `parser` struct with this token list and a position of 0. Before creating
the `parser` instance, it calls two other functions: `createTokenLookups`
and `createTokenTypeLookups`, likely to set up some internal state or mappings used by the parser.
*/
func createParser(tokens []lexer.Token) *parser {
	createTokenLookups()
	createTokenTypeLookups()
	return &parser{
		tokens: tokens, pos: 0,
	}
}

/*
This Go function, `Parse`, takes a source string as input and returns a parsed abstract syntax tree (AST) as an `ast.BlockStmt`. Here's a succinct explanation of what the function does:

1. It tokenizes the input source string into a list of tokens using the `lexer.Tokenize` function.
2. It creates a new parser instance with the tokenized input using the `createParser` function.
3. It initializes an empty list to store parsed statements.
4. It enters a loop that continues as long as there are tokens left to parse.
5. Inside the loop, it parses a single statement using the `parse_stmt` function and appends it to the list of parsed statements.
6. Once all tokens have been parsed, it returns a new `ast.BlockStmt` instance with the list of parsed statements as its body.

In essence, this function is the entry point for parsing source code into an abstract syntax tree.
*/
func Parse(source string) ast.BlockStmt {
	tokens := lexer.Tokenize(source)
	p := createParser(tokens)
	body := make([]ast.Stmt, 0)
	// while we have tokens, continue to parse
	for p.hasTokens() {
		body = append(body, parse_stmt(p))
	}

	return ast.BlockStmt{
		Body: body,
	}
}

// HELPER FUNCTIONS

/*
This Go function, `currentToken`, returns the current token being processed
by the parser. It does this by accessing the `tokens` slice at the current position `p.pos`.
*/
func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

/*
This code snippet defines a method `currentTokenKind` on the `parser` struct in Go. The method returns the kind of the current token being processed by the parser, which is obtained by accessing the `Kind` field of the token returned by the `currentToken` method.

In other words, it gets the type of the current token (e.g., identifier, number, string, etc.) from the parser's current token.
*/
func (p *parser) currentTokenKind() lexer.TokenKind {
	return p.currentToken().Kind
}

/*
This Go function, `advance`, advances the parser's position to the next
token in the token list and returns the current token that was just advanced past.
*/
func (p *parser) advance() lexer.Token {

	tk := p.currentToken()
	p.pos++
	return tk
}

/*
This Go function, `hasTokens`, checks if the parser has more tokens to
process. It returns `true` if the parser's current position (`p.pos`) is
within the bounds of the token list (`p.tokens`) and the current token is
not an end-of-file (`EOF`) token.
*/
func (p *parser) hasTokens() bool {
	return p.pos < len(p.tokens) && p.currentTokenKind() != lexer.EOF
}

/*
This code snippet is a method called `expectError` defined on the `parser` struct in Go.
It takes two parameters: `expectedKind` of type `lexer.TokenKind` and `err` of type `any`.

The purpose of this method is to check if the current token in the parser matches the
expected token kind. If the token kind does not match, it checks if an error message is
provided (`err != nil`). If an error message is provided, it formats and panics with
the error message. If no error message is provided, it panics with a default error message.

The method returns the token that was advanced past in the parser.
*/
func (p *parser) expectError(expectedKind lexer.TokenKind, err any) lexer.Token {
	token := p.currentToken()
	kind := token.Kind

	if kind != expectedKind {
		if err == nil {
			err = fmt.Sprintf("Expected %s, but received %s instead\n", lexer.TokenKindString(expectedKind), lexer.TokenKindString(kind))

		}
		panic(err)
	}

	return p.advance()

}

/*
This code snippet defines a method `expect` on the `parser` struct in Go. It takes
a `lexer.TokenKind` parameter `expectedKind` and returns a `lexer.Token`.

The method calls the `expectError` method with `expectedKind` and `nil` as parameters.

In essence, this method checks if the current token in the parser matches the `expectedKind`.
If it does not match, it panics with a default error message. If it matches, it returns
the token that was advanced past in the parser.

This method is a convenience wrapper around `expectError` that provides a default
error message when the token kind does not match.
*/
func (p *parser) expect(expectedKind lexer.TokenKind) lexer.Token {
	return p.expectError(expectedKind, nil)
}
