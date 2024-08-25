package lexer 

import (
	"fmt"
	"regexp"
)

type regexHandler func (lex *lexer, regex *regexp.Regexp)

type regexPattern struct {
	regex *regexp.Regexp 
	handler regexHandler
}

type lexer struct {
	Tokens []Token
	source string
	pos int
	patterns []regexPattern
}

func (lex *lexer) advanceN (n int){
	lex.pos+=n
}

func (lex *lexer) push (token Token){
	lex.Tokens = append(lex.Tokens, token)
}

func (lex *lexer) at () byte {
	return lex.source[lex.pos]
}

func (lex *lexer) remainder () string {
	return lex.source[lex.pos:]
}

func (lex *lexer) at_eof () bool {
	return lex.pos >= len(lex.source)
}

func Tokenize (source string) []Token {
	lex := createLexer(source)
  // iterate while we still have tokens 
	for !lex.at_eof(){
		matched := false
		for _, pattern := range lex.patterns {
			// match the pattern
			loc := pattern.regex.FindStringIndex(lex.remainder())
			if loc !=nil && loc[0]==0{
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}

	}
	// you may improve this error message, like location of error
	if !matched{
	  panic(fmt.Sprintf("Lexer:Error -> unrecognized token near %s\n", lex.remainder()))
	}
	
}
lex.push(NewToken(EOF, "eof"))
return lex.Tokens
}

func defaultHandler(kind TokenKind, value string) regexHandler{
	return func(lex *lexer, regex *regexp.Regexp) {
		// advance the lexer's position past the value we just reached
		lex.advanceN(len(value))
		lex.push(NewToken(kind, value))
	}
}

func createLexer(source string) *lexer {
	return &lexer{
		pos: 0,
		source: source,
		Tokens: make([]Token,0),
		patterns: []regexPattern{
			// define all of the regex patterns
			{regexp.MustCompile(`\s+`), skipHandler},
			// {regexp.MustCompile(`\/\/.*`), commentHandler},
			{regexp.MustCompile(`"[^"]*"`), stringHandler},
			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), numberHandler},
			// {regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`), symbolHandler},
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

func numberHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	lex.push(NewToken(NUMBER, match))
	lex.advanceN(len(match))
	 
}

func skipHandler(lex *lexer, regex *regexp.Regexp) {
match := regex.FindStringIndex(lex.remainder())
lex.advanceN(match[1])

	 
}

func stringHandler(lex *lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder()) 
	stringLiteral := lex.remainder()[match[0]:match[1]]
	lex.push(NewToken(STRING, stringLiteral))
	lex.advanceN(len(stringLiteral))	 
	}

// func commentHandler(lex *lexer, regex *regexp.Regexp) {
// 	match := regex.FindString(lex.remainder())
// 	lex.push(NewToken(NUMBER, match))
// 	lex.advanceN(len(match))
	 
// }