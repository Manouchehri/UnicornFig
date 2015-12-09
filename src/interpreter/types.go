package interpreter

import (
	"errors"
)

// Types

type ValueType int

const (
	UnassignedT  ValueType = iota
	LiteralT     ValueType = iota
	StringT      ValueType = iota
	IntegerT     ValueType = iota
	FloatT       ValueType = iota
	NameT        ValueType = iota
	BooleanT     ValueType = iota
	FunctionT    ValueType = iota
	SExpressionT ValueType = iota
	SpecialFormT ValueType = iota
	ListT        ValueType = iota
	MapT         ValueType = iota
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

type BooleanLiteral struct {
	Contained bool
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

func (b BooleanLiteral) Type() ValueType {
	return BooleanT
}

// S-Expressions

type SExpression struct {
	FormName Name
	Type     ValueType
	Values   []interface{} // Values or S-Expressions
}

// Functions

type Builtin func(...interface{}) (Value, error)

/**
 * Represents both user-defined functions, which are built on top of builtins,
 * as well as builtin functions.  In the case of user-defined fucntions, a body S-Expression
 * is provided to be evaluated until a builtin is reached that can be executed as Go code.
 * The IsCallable and Callable fields handle the latter case.
 */
type Function struct {
	FunctionName  Name
	ArgumentNames []Name
	Body          interface{} // Can be a Value or an S-Expression
	IsCallable    bool
	Scope         Environment
	Callable      Builtin
}

func (fn Function) Call(unwrapped ...interface{}) (Value, error) {
	if !fn.IsCallable {
		return Value{}, errors.New("Not a callable function")
	} else {
		return (fn.Callable)(unwrapped...)
	}
}

// Another OR type. Either a literal, a name, a function, or a list
type Value struct {
	Type     ValueType
	String   StringLiteral
	Integer  IntegerLiteral
	Float    FloatLiteral
	Name     Name
	Boolean  BooleanLiteral
	Function Function
	List     List
	Map      Mapping
	Ignored  bool // Should we ignore the value when producing an output config file?
}

// Lists

type List struct {
	Data []Value
}

// Maps

type Mapping struct {
	Data map[string]Value
}

/**
 * Token types for the parser
 */

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
