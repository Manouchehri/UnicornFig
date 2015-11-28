package interpreter

import (
	"errors"
)

var (
	zeroi  = IntegerLiteral{0}
	zerof  = FloatLiteral{0.0}
	emptys = StringLiteral{""}
	falseb = BooleanLiteral{false}
	emptyl = List{[]Value{}}
	emptym = Mapping{map[string]Value{}}
)

func NewString(str string) Value {
	return Value{StringT, StringLiteral{str}, zeroi, zerof, Name{}, falseb, Function{}, emptyl, emptym}
}

func NewInteger(n int64) Value {
	return Value{IntegerT, emptys, IntegerLiteral{n}, zerof, Name{}, falseb, Function{}, emptyl, emptym}
}

func NewFloat(n float64) Value {
	return Value{FloatT, emptys, zeroi, FloatLiteral{n}, Name{}, falseb, Function{}, emptyl, emptym}
}

func NewName(identifier string) Value {
	return Value{NameT, emptys, zeroi, zerof, Name{identifier}, falseb, Function{}, emptyl, emptym}
}

func NewBoolean(value bool) Value {
	return Value{BooleanT, emptys, zeroi, zerof, Name{}, BooleanLiteral{value}, Function{}, emptyl, emptym}
}

func NewSExpression(formName string, values ...interface{}) SExpression {
	emptyArray := make([]interface{}, 0)
	sexp := SExpression{Name{formName}, SExpressionT, emptyArray}
	for _, value := range values {
		sexp.Values = append(sexp.Values, value)
	}
	return sexp
}

func NewCallableFunction(name string, argNames []string, fn Builtin) Value {
	names := make([]Name, len(argNames))
	for i, arg := range argNames {
		names[i] = Name{arg}
	}
	return Value{FunctionT, emptys, zeroi, zerof, Name{}, falseb, Function{Name{name}, names, SExpression{}, true, fn}, emptyl, emptym}
}

func NewFunction(name string, argNames []string, body interface{}) Value {
	names := make([]Name, len(argNames))
	for i, arg := range argNames {
		names[i] = Name{arg}
	}
	return Value{FunctionT, emptys, zeroi, zerof, Name{}, falseb, Function{Name{name}, names, body, false, nil}, emptyl, emptym}
}

func NewList() Value {
	return Value{ListT, emptys, zeroi, zerof, Name{}, falseb, Function{}, emptyl, emptym}
}

func NewMap() Value {
	return Value{MapT, emptys, zeroi, zerof, Name{}, falseb, Function{}, emptyl, emptym}
}

func Unwrap(value Value) interface{} {
	switch value.Type {
	case StringT:
		return value.String.Contained
	case IntegerT:
		return value.Integer.Contained
	case FloatT:
		return value.Float.Contained
	case NameT:
		return value.Name.Contained
	case BooleanT:
		return value.Boolean.Contained
	case FunctionT:
		// TODO - This is definitely NOT going to work in most places
		return value.Function.Callable
	case ListT:
		values := make([]interface{}, len(value.List.Data))
		for i, val := range value.List.Data {
			values[i] = Unwrap(val)
		}
		return values
	case MapT:
		unwrapped := make(map[string]interface{})
		for key, val := range value.Map.Data {
			unwrapped[key] = Unwrap(val)
		}
		return unwrapped
	}
	return nil
}

func Wrap(thing interface{}) (Value, error) {
	switch thing.(type) {
	case int64:
		return NewInteger(thing.(int64)), nil
	case float64:
		return NewFloat(thing.(float64)), nil
	case string:
		return NewString(thing.(string)), nil
	case bool:
		value := thing.(bool)
		if value {
			return NewName("true"), nil
		}
		return NewName("false"), nil
	}
	return Value{}, errors.New("Cannot wrap values of the type of the argument provided.")
}
