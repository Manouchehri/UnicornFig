package interpreter

// Types

type ValueType int

const (
	UnassignedT  ValueType = iota
	LiteralT     ValueType = iota
	StringT      ValueType = iota
	IntegerT     ValueType = iota
	FloatT       ValueType = iota
	NameT        ValueType = iota
	FunctionT    ValueType = iota
	SExpressionT ValueType = iota
	ValueT       ValueType = iota
)

// Literals

type Literal interface {
	Type() ValueType
}

type StringLiteral struct {
	Contained string
}

type IntegerLiteral struct {
	Contained int64
}

type FloatLiteral struct {
	Contained float64
}

type Name struct {
	Contained string
}

func (s StringLiteral) Type() ValueType {
	return StringT
}

func (i IntegerLiteral) Type() ValueType {
	return IntegerT
}

func (f FloatLiteral) Type() ValueType {
	return FloatT
}

func (n Name) Type() ValueType {
	return NameT
}

// S-Expressions

type SExpression struct {
	FormName Name
	Type     ValueType
	Values   []interface{} // Values or S-Expressions
}

// Functions

type Function struct {
	FunctionName  Name
	ArgumentNames []Name
	Body          SExpression
}

// Another OR type. Either a literal, a name, a function, or a list
type Value struct {
	Type     ValueType
	String   StringLiteral
	Integer  IntegerLiteral
	Float    FloatLiteral
	Name     Name
	Function Function
}

type Token string

const (
	NO_TOKEN      Token = ""
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
