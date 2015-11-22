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

type Builtin func(Environment, ...interface{}) (error, Value, Environment)

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
	Callable      Builtin
}

func (fn Function) Call(env Environment, unwrapped ...interface{}) (error, Value, Environment) {
	if !fn.IsCallable {
		return errors.New("Not a callable function"), Value{}, Environment{}
	} else {
		return (fn.Callable)(env, unwrapped...)
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
}

/**
 * Special forms
 */

type SpecialForm interface {
	Form() SpecialFormType
}

type SpecialFormType int

const (
	DefinitionFormT SpecialFormType = iota
	FunctionFormT   SpecialFormType = iota
	ConditionFormT  SpecialFormType = iota
)

type DefinitionForm struct {
	Definitions []SExpression
}

type FunctionForm struct {
	Name      Name
	Arguments SExpression
	Body      SExpression
}

type ConditionForm struct {
	Condition SExpression
	Branch1   SExpression
	Branch2   SExpression
}

func (df DefinitionForm) Form() SpecialFormType {
	return DefinitionFormT
}

func (ff FunctionForm) Form() SpecialFormType {
	return FunctionFormT
}

func (cf ConditionForm) Form() SpecialFormType {
	return ConditionFormT
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
