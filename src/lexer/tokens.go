package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota
	NULL
	TRUE
	FALSE
	NUMBER
	STRING
	IDENTIFIER

	// Grouping & Braces
	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_CURLY
	CLOSE_CURLY
	OPEN_PAREN
	CLOSE_PAREN

	// Equivilance
	ASSIGNMENT
	EQUALS
	NOT_EQUALS
	NOT

	// Conditional
	LESS
	LESS_EQUALS
	GREATER
	GREATER_EQUALS

	// Logical
	OR
	AND

	// Symbols
	DOT
	DOT_DOT
	SEMI_COLON
	COLON
	QUESTION
	COMMA

	// Shorthand
	PLUS_PLUS
	MINUS_MINUS
	PLUS_EQUALS
	MINUS_EQUALS
	NULLISH_ASSIGNMENT // ??=

	//Maths
	PLUS
	DASH
	SLASH
	STAR
	PERCENT

	// Reserved Keywords
	LET
	CONST
	CLASS
	NEW
	IMPORT
	FROM
	FN
	IF
	ELSE
	FOREACH
	WHILE
	FOR
	EXPORT
	TYPEOF
	IN

	// Misc
	NUM_TOKENS
)

var isReservedKeyword = map[string]TokenKind{
	"let":     LET,
	"const":   CONST,
	"fn":      FN,
	"if":      IF,
	"else":    ELSE,
	"for":     FOR,
	"while":   WHILE,
	"new":     NEW,
	"import":  IMPORT,
	"from":    FROM,
	"class":   CLASS,
	"true":    TRUE,
	"false":   FALSE,
	"foreach": FOREACH,
	"export":  EXPORT,
	"typeof":  TYPEOF,
	"in":      IN,
}

/*
The `Token` struct in Go represents a token in the lexer. It has two fields:

- `Kind`: stores the kind of the token, which is an enum value from `TokenKind`.
- `Value`: stores the literal value of the token, which is a string.

The `Token` struct does not have any methods.
*/
type Token struct {
	Kind  TokenKind
	Value string
}

// isOneOfMany Checks if the token kind is one of the expected tokens.
//
// expectedTokens - A variable number of TokenKind values to check against.
// Return type: bool
func (tk Token) isOneOfMany(expectedTokens ...TokenKind) bool {
	for _, expected := range expectedTokens {
		if expected == tk.Kind {
			return true
		}
	}

	return false
}

// Debug Prints a debug representation of the token.
//
// It prints the token kind and its literal value if it's an identifier, number, or string.
// Otherwise, it only prints the token kind.
// Return type: No return value.
func (token Token) Debug() {
	if token.isOneOfMany(IDENTIFIER, NUMBER, STRING) {
		fmt.Printf("%s (%s)\n", TokenKindString(token.Kind), token.Value)
	} else {
		fmt.Printf("%s ()\n", TokenKindString(token.Kind))
	}
}

// NewToken Creates a new token with the specified kind and value.
//
// kind - The type of token to be created.
// value - The value of the token to be created.
// Return type: Token

func NewToken(kind TokenKind, value string) Token {
	return Token{
		kind, value,
	}
	// TokenKindString Returns a string representation of a TokenKind.
	//
	// kind - The TokenKind to be converted to a string.
	// Return type: string
}
func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "eof"
	case NULL:
		return "null"
	case NUMBER:
		return "number"
	case STRING:
		return "string"
	case TRUE:
		return "true"
	case FALSE:
		return "false"
	case IDENTIFIER:
		return "identifier"
	case OPEN_BRACKET:
		return "open_bracket"
	case CLOSE_BRACKET:
		return "close_bracket"
	case OPEN_CURLY:
		return "open_curly"
	case CLOSE_CURLY:
		return "close_curly"
	case OPEN_PAREN:
		return "open_paren"
	case CLOSE_PAREN:
		return "close_paren"
	case ASSIGNMENT:
		return "assignment"
	case EQUALS:
		return "equals"
	case NOT_EQUALS:
		return "not_equals"
	case NOT:
		return "not"
	case LESS:
		return "less"
	case LESS_EQUALS:
		return "less_equals"
	case GREATER:
		return "greater"
	case GREATER_EQUALS:
		return "greater_equals"
	case OR:
		return "or"
	case AND:
		return "and"
	case DOT:
		return "dot"
	case DOT_DOT:
		return "dot_dot"
	case SEMI_COLON:
		return "semi_colon"
	case COLON:
		return "colon"
	case QUESTION:
		return "question"
	case COMMA:
		return "comma"
	case PLUS_PLUS:
		return "plus_plus"
	case MINUS_MINUS:
		return "minus_minus"
	case PLUS_EQUALS:
		return "plus_equals"
	case MINUS_EQUALS:
		return "minus_equals"
	case NULLISH_ASSIGNMENT:
		return "nullish_assignment"
	case PLUS:
		return "plus"
	case DASH:
		return "dash"
	case SLASH:
		return "slash"
	case STAR:
		return "star"
	case PERCENT:
		return "percent"
	case LET:
		return "let"
	case CONST:
		return "const"
	case CLASS:
		return "class"
	case NEW:
		return "new"
	case IMPORT:
		return "import"
	case FROM:
		return "from"
	case FN:
		return "fn"
	case IF:
		return "if"
	case ELSE:
		return "else"
	case FOREACH:
		return "foreach"
	case FOR:
		return "for"
	case WHILE:
		return "while"
	case EXPORT:
		return "export"
	case IN:
		return "in"
	default:
		return fmt.Sprintf("unknown(%d)", kind)
	}
}
