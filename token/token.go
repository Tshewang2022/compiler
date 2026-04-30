package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// defining keywords for my language
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// checks if given identifier is a keyword
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//identifiers +literal
	IDENT = "IDENT"
	INT   = "INT" //1232434

	// operatiors
	ASSIGN = "="
	PLUS   = "+"

	//delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
