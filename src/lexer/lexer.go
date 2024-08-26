package lexer

import (
	"fmt"
	"regexp"
)

/*
This line of code defines a type alias in Go called `regexHandler`. It represents a function that takes two parameters: a pointer to a `lexer` struct (`lex *lexer`) and a pointer to a `regexp.Regexp` object (`regex *regexp.Regexp`).
*/
type regexHandler func(lex *lexer, regex *regexp.Regexp)

/*
This class definition defines a struct called `regexPattern` in Go. Here's a succinct explanation of what each field does:

* `regex *regexp.Regexp`: This field stores a compiled regular expression that can be used to match patterns in text.
* `handler regexHandler`: This field stores a function that takes a `lexer` and a `regexp.Regexp` as parameters, and is used to handle matches found by the regular expression.

Note that this struct does not have any methods, it only has fields to store a regular expression and a handler function.
*/
type regexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

/*
This class definition defines a struct called `lexer` in Go. Here's a succinct explanation of what each field does:

* `Tokens []Token`: stores a list of tokens found in the source code.
* `source string`: stores the source code being lexed.
* `pos int`: stores the current position in the source code.
* `patterns []regexPattern`: stores a list of regular expression patterns used to match tokens in the source code.
*/
type lexer struct {
	Tokens   []Token
	source   string
	pos      int
	patterns []regexPattern
}

// advanceN Advances the lexer's position by a specified number of characters.
//
// n - The number of characters to advance the lexer's position.
// No return value.
func (lex *lexer) advanceN(n int) {
	lex.pos += n
}

// push Appends a token to the lexer's token list.
//
// token - The token to be appended to the lexer's token list.
// No return value.
func (lex *lexer) push(token Token) {
	lex.Tokens = append(lex.Tokens, token)
}

// at Returns the character at the current position in the lexer's source code.
//
// lex - The lexer instance.
// Return type: byte
func (lex *lexer) at() byte {
	return lex.source[lex.pos]
}

// remainder Returns the remaining part of the source code being lexed.
//
// No parameters.
// Return type: string
func (lex *lexer) remainder() string {
	return lex.source[lex.pos:]
}

// at_eof checks if the lexer has reached the end of the source code.
//
// It returns a boolean value indicating whether the lexer is at the end of the source code.
func (lex *lexer) at_eof() bool {
	return lex.pos >= len(lex.source)
}

// Tokenize Tokenizes the source code into a list of tokens.
//
// source - The source code to be tokenized.
// Return type: []Token
func Tokenize(source string) []Token {
	lex := createLexer(source)
	// iterate while we still have tokens
	for !lex.at_eof() {
		matched := false
		for _, pattern := range lex.patterns {
			// match the pattern
			loc := pattern.regex.FindStringIndex(lex.remainder())
			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}

		}
		// you may improve this error message, like location of error
		if !matched {
			panic(fmt.Sprintf("Lexer:Error -> unrecognized token near %s\n", lex.remainder()))
		}

	}
	lex.push(NewToken(EOF, "eof"))
	return lex.Tokens
}

// defaultHandler Returns a regexHandler function that handles tokens of a specific kind.
//
// kind - The type of token to be handled.
// value - The value of the token to be handled.
// Return type: regexHandler
func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(lex *lexer, regex *regexp.Regexp) {
		// advance the lexer's position past the value we just reached
		lex.advanceN(len(value))
		lex.push(NewToken(kind, value))
	}
}

// createLexer Creates a new lexer instance from the given source code.
//
// source - The source code to be lexed.
// Return type: *lexer
func createLexer(source string) *lexer {
	return &lexer{
		pos:    0,
		source: source,
		Tokens: make([]Token, 0),
		patterns: []regexPattern{
			// define all of the regex patterns
			{regexp.MustCompile(`\s+`), skipHandler},
			{regexp.MustCompile(`\/\/.*`), commentHandler},
			{regexp.MustCompile(`"[^"]*"`), stringHandler},
			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), numberHandler},
			{regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`), symbolHandler},
			{regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			{regexp.MustCompile(`\]`), defaultHandler(CLOSE_BRACKET, "]")},
			{regexp.MustCompile(`\{`), defaultHandler(OPEN_CURLY, "{")},
			{regexp.MustCompile(`\}`), defaultHandler(CLOSE_CURLY, "}")},
			{regexp.MustCompile(`\(`), defaultHandler(OPEN_PAREN, "(")},
			{regexp.MustCompile(`\)`), defaultHandler(CLOSE_PAREN, ")")},
			{regexp.MustCompile(`==`), defaultHandler(EQUALS, "==")},
			{regexp.MustCompile(`!=`), defaultHandler(NOT_EQUALS, "!=")},
			{regexp.MustCompile(`=`), defaultHandler(ASSIGNMENT, "=")},
			{regexp.MustCompile(`!`), defaultHandler(NOT, "!")},
			{regexp.MustCompile(`<=`), defaultHandler(LESS_EQUALS, "<=")},
			{regexp.MustCompile(`<`), defaultHandler(LESS, "<")},
			{regexp.MustCompile(`>=`), defaultHandler(GREATER_EQUALS, ">=")},
			{regexp.MustCompile(`>`), defaultHandler(GREATER, ">")},
			{regexp.MustCompile(`\|\|`), defaultHandler(OR, "||")},
			{regexp.MustCompile(`&&`), defaultHandler(AND, "&&")},
			{regexp.MustCompile(`\.\.`), defaultHandler(DOT_DOT, "..")},
			{regexp.MustCompile(`\.`), defaultHandler(DOT, ".")},
			{regexp.MustCompile(`;`), defaultHandler(SEMI_COLON, ";")},
			{regexp.MustCompile(`:`), defaultHandler(COLON, ":")},
			{regexp.MustCompile(`\?\?=`), defaultHandler(NULLISH_ASSIGNMENT, "??=")},
			{regexp.MustCompile(`\?`), defaultHandler(QUESTION, "?")},
			{regexp.MustCompile(`,`), defaultHandler(COMMA, ",")},
			{regexp.MustCompile(`\+\+`), defaultHandler(PLUS_PLUS, "++")},
			{regexp.MustCompile(`--`), defaultHandler(MINUS_MINUS, "--")},
			{regexp.MustCompile(`\+=`), defaultHandler(PLUS_EQUALS, "+=")},
			{regexp.MustCompile(`-=`), defaultHandler(MINUS_EQUALS, "-=")},
			{regexp.MustCompile(`\+`), defaultHandler(PLUS, "+")},
			{regexp.MustCompile(`-`), defaultHandler(DASH, "-")},
			{regexp.MustCompile(`/`), defaultHandler(SLASH, "/")},
			{regexp.MustCompile(`\*`), defaultHandler(STAR, "*")},
			{regexp.MustCompile(`%`), defaultHandler(PERCENT, "%")},
			// order of patterns matter, if = is before == then we will get two assignments
			// operator instead of one equals
		},
	}
}

// numberHandler Handles a numeric token match in the source code.
//
// lex - The lexer instance used to process the source code.
// regex - The regular expression used to match the numeric token.
// Return type: No return value.
func numberHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	lex.push(NewToken(NUMBER, match))
	lex.advanceN(len(match))

}

// skipHandler Skips a token match in the source code.
//
// lex - The lexer instance used to process the source code.
// regex - The regular expression used to match the token.
// Return type: No return value.
func skipHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	lex.advanceN(match[1])

}

// stringHandler Handles a string token match in the source code.
//
// lex - The lexer instance used to process the source code.
// regex - The regular expression used to match the string token.
// Return type: No return value.
func stringHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	stringLiteral := lex.remainder()[match[0]+1 : match[1]-1]
	lex.push(NewToken(STRING, stringLiteral))
	lex.advanceN(len(stringLiteral) + 2)
}

// commentHandler Handles a comment token match in the source code.
//
// lex - The lexer instance used to process the source code.
// regex - The regular expression used to match the comment token.
// Return type: No return value.
func commentHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	lex.advanceN(match[1])

}

// symbolHandler Handles a symbol token match in the source code.
//
// lex - The lexer instance used to process the source code.
// regex - The regular expression used to match the symbol token.
// Return type: No return value.
func symbolHandler(lex *lexer, regex *regexp.Regexp) {
	value := regex.FindString(lex.remainder())

	if kind, exists := isReservedKeyword[value]; exists {
		lex.push(NewToken(kind, value))
	} else {
		lex.push(NewToken(IDENTIFIER, value))
	}
	lex.advanceN(len(value))

}
