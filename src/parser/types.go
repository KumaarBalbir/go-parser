package parser

import (
	"fmt"
	"github.com/go-parser/src/ast"
	"github.com/go-parser/src/lexer"
)

/*
defines a type alias named type_nud_handler. This type alias represents a function
that takes a pointer to a parser struct as a parameter and returns an ast.Type. In the
context of the provided Go code, type_nud_handler is used as a type for a map that
associates token kinds from the lexer package with their corresponding nud (null denotation)
handlers for parsing types.

The type_nud_handler function type is used in the declaration of type_nud_lookup,
which is a map that maps token kinds to their respective nud handlers. This map is
used in the parse_type function to determine which nud handler to call for the current token kind.
*/
type type_nud_handler func(p *parser) ast.Type

/*
defines a type alias named type_led_handler. This type alias represents a
function that takes three parameters: a pointer to a parser struct, an ast.Type
representing the left-hand side of the expression, and a binding_power value. The
function returns an ast.Type.

In the context of the provided Go code, type_led_handler is used as a type for a map
that associates token kinds from the lexer package with their corresponding
left-denotation (infix) handlers for parsing types. The type_led_handler function
type is used in the declaration of type_led_lookup, which is a map that maps token
kinds to their respective left-denotation handlers. This map is used in the
parse_type function to determine which left-denotation handler to call
for the current token kind.

The type_led function sets up a left-denotation operator for a given token kind
in the parser. It takes three parameters: the token kind, the binding power of
the token, and the function to be called when the token is encountered. The
function updates two lookup tables: type_bp_lu and type_led_lu, which store the
binding power and the handling function for the token, respectively.
*/
type type_led_handler func(p *parser, left ast.Type, bp binding_power) ast.Type

type type_nud_lookup map[lexer.TokenKind]type_nud_handler
type type_led_lookup map[lexer.TokenKind]type_led_handler
type type_bp_lookup map[lexer.TokenKind]binding_power

var type_bp_lu = bp_lookup{}
var type_nud_lu = type_nud_lookup{}
var type_led_lu = type_led_lookup{}

/*
This function, `type_led`, sets up a left-denotation (infix) operator for a given token kind in the parser. It takes three parameters:

* `kind`: the type of token to be handled
* `bp`: the binding power of the token
* `led_fn`: the function to be called when the token is encountered

The function updates two lookup tables: `type_bp_lu` and `type_led_lu`, which store the binding power and the handling function for the token, respectively.
*/
func type_led(kind lexer.TokenKind, bp binding_power, led_fn type_led_handler) {
	type_bp_lu[kind] = bp
	type_led_lu[kind] = led_fn
}

/*
This function, `type_nud`, sets up a null-denotation (prefix) operator for a given token kind in the parser. It takes two parameters:

* `kind`: the type of token to be handled
* `nud_fn`: the function to be called when the token is encountered

The function updates the `type_nud_lu` lookup table, which stores the handling function for the token.
*/
func type_nud(kind lexer.TokenKind, nud_fn type_nud_handler) {
	type_nud_lu[kind] = nud_fn
}

/*
This function, `createTokenTypeLookups`, sets up token type lookups for the parser. It defines two null-denotation (prefix) operators:

*   When the parser encounters an `IDENTIFIER` token, it will call the `parse_symbol_type` function to parse the identifier as a symbol type.
*   When the parser encounters an `OPEN_BRACKET` token, it will call the `parse_array_type` function to parse the array type.
*/
func createTokenTypeLookups() {
	type_nud(lexer.IDENTIFIER, parse_symbol_type)
	type_nud(lexer.OPEN_BRACKET, parse_array_type)
}

/*
This function, `parse_symbol_type`, parses a symbol type from the current token in
the parser. It expects the current token to be an `IDENTIFIER` and returns
an `ast.SymbolType` object with the identifier's value as its name.
*/
func parse_symbol_type(p *parser) ast.Type {
	return ast.SymbolType{
		Name: p.expect(lexer.IDENTIFIER).Value}
}

/*
This function, `parse_array_type`, parses an array type from the current token in the parser. It does the following:

1. Advances past the current token (expected to be `OPEN_BRACKET`).
2. Expects a `CLOSE_BRACKET` token to follow.
3. Parses the underlying type of the array using the `parse_type` function with the default binding power (`default_bp`).
4. Returns an `ast.ArrayType` object with the parsed underlying type.

In essence, this function constructs an array type by parsing the underlying type and ensuring it's enclosed in square brackets.
*/
func parse_array_type(p *parser) ast.Type {
	p.advance()
	p.expect(lexer.CLOSE_BRACKET)
	var underlyingType = parse_type(p, default_bp)
	return ast.ArrayType{
		Underlying: underlyingType,
	}
}

/*
This Go function, `parse_type`, recursively parses a type expression in the abstract syntax tree (AST) using a parser object `p` and a binding power `bp`. It handles two types of type expressions:

*   Null-denotation (prefix) expressions: It calls the handling function stored in the `type_nud_lu` lookup table based on the current token kind.
*   Left-denotation (infix) expressions: It calls the handling function stored in the `type_led_lu` lookup table based on the current token kind, as long as the binding power of the current token is greater than the input `bp`.

The function returns the fully parsed type expression.
*/
func parse_type(p *parser, bp binding_power) ast.Type {
	// First parse the NUD
	tokenKind := p.currentTokenKind()
	nud_fn, exists := type_nud_lu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("TYPE_NUD handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
	}

	left := nud_fn(p)

	for type_bp_lu[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		led_fn, exists := type_led_lu[tokenKind]

		if !exists {
			panic(fmt.Sprintf("TYPE_LED handler expected for token %s\n", lexer.TokenKindString(tokenKind)))
		}

		left = led_fn(p, left, type_bp_lu[p.currentTokenKind()])
	}

	// While we have a LED and the current bp is < bp of current token
	// continue parsing left hand side

	// Move the return statement outside of the loop
	return left
}
