package interpreter

// Literals

type Literal interface {
	IsLiteral() bool
}

type StringLiteral string
type IntegerLiteral int64
type FloatLiteral float64

func (s StringLiteral) IsLiteral() bool {
	return true
}

func (i IntegerLiteral) IsLiteral() bool {
	return true
}

func (f FloatLiteral) IsLiteral() bool {
	return true
}

// Names

type Name string

// Lists

// An OR type. Either has a literal in it or a name
type List struct {
	LiteralPtr *Literal
	NamePtr    *Name
	Next       *List
}

// S-Expressions

type SExpression struct {
	FormName Name
	Values   []interface{} // Values or S-Expressions
}

// Functions

type Function struct {
	FunctionName  Name
	ArgumentNames []Name
	Body          SExpression
}

// Values

// Another OR type. Either a literal, a name, a function, or a list
type Value struct {
	LiteralPtr  *Literal
	NamePtr     *Name
	FunctionPtr *Function
	ListPtr     *List
}

type Token string

const (
	NO_TOKEN      Token = ""
	START_LIST    Token = "[START_LIST]"
	START_SEXP    Token = "[START_SEXP]"
	START_STRING  Token = "[START_STRING]"
	START_COMMENT Token = "[START_COMMENT]"
	START_NUMBER  Token = "[START_NUMBER]"
	START_NAME    Token = "[START_NAME]"
	END_SEXP      Token = "[END_SEXP]"
	END_STRING    Token = "[END_STRING]"
	END_COMMENT   Token = "[END_COMMENT]"
	END_NUMBER    Token = "[END_NUMBER]"
	END_NAME      Token = "[END_NAME]"
)
