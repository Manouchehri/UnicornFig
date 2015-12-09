package interpreter

import (
	"errors"
)

var (
	zeroi  = IntegerLiteral{0}
	zerof  = FloatLiteral{0.0}
	emptys = StringLiteral{""}
	falseb = BooleanLiteral{false}
)

/**
 * Shorthand functions to easily create instances of supported types.
 */

func NewString(str string) Value {
	emptyl := List{[]Value{}}
	emptym := Mapping{map[string]Value{}}
	return Value{StringT, StringLiteral{str}, zeroi, zerof, Name{}, falseb, Function{}, emptyl, emptym, false}
}

func NewInteger(n int64) Value {
	emptyl := List{[]Value{}}
	emptym := Mapping{map[string]Value{}}
	return Value{IntegerT, emptys, IntegerLiteral{n}, zerof, Name{}, falseb, Function{}, emptyl, emptym, false}
}

func NewFloat(n float64) Value {
	emptyl := List{[]Value{}}
	emptym := Mapping{map[string]Value{}}
	return Value{FloatT, emptys, zeroi, FloatLiteral{n}, Name{}, falseb, Function{}, emptyl, emptym, false}
}

func NewName(identifier string) Value {
	emptyl := List{[]Value{}}
	emptym := Mapping{map[string]Value{}}
	return Value{NameT, emptys, zeroi, zerof, Name{identifier}, falseb, Function{}, emptyl, emptym, false}
}

func NewBoolean(value bool) Value {
	emptyl := List{[]Value{}}
	emptym := Mapping{map[string]Value{}}
	return Value{BooleanT, emptys, zeroi, zerof, Name{}, BooleanLiteral{value}, Function{}, emptyl, emptym, false}
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
	emptyl := List{[]Value{}}
	emptym := Mapping{map[string]Value{}}
	return Value{FunctionT, emptys, zeroi, zerof, Name{}, falseb, Function{Name{name}, names, SExpression{}, true, Environment{}, fn}, emptyl, emptym, false}
}

func NewFunction(name string, argNames []string, body interface{}) Value {
	names := make([]Name, len(argNames))
	for i, arg := range argNames {
		names[i] = Name{arg}
	}
	emptyl := List{[]Value{}}
	emptym := Mapping{map[string]Value{}}
	return Value{FunctionT, emptys, zeroi, zerof, Name{}, falseb, Function{Name{name}, names, body, false, Environment{}, nil}, emptyl, emptym, false}
}

func NewList() Value {
	emptyl := List{[]Value{}}
	emptym := Mapping{map[string]Value{}}
	return Value{ListT, emptys, zeroi, zerof, Name{}, falseb, Function{}, emptyl, emptym, false}
}

func NewMap() Value {
	emptyl := List{[]Value{}}
	emptym := Mapping{map[string]Value{}}
	return Value{MapT, emptys, zeroi, zerof, Name{}, falseb, Function{}, emptyl, emptym, false}
}

/**
 * Unwrap a value to make the contents (stuff Go can compute with) available to a function.
 */
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

/**
 * Wrap a value back into one of Unicorn's Value instances.
 */
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
	case []interface{}:
		list := NewList()
		thingList := thing.([]interface{})
		for _, v := range thingList {
			wrapped, err := Wrap(v)
			if err != nil {
				return list, err
			}
			list.List.Data = append(list.List.Data, wrapped)
		}
		return list, nil
	case map[string]interface{}:
		mapping := NewMap()
		thingMap := thing.(map[string]interface{})
		for k, v := range thingMap {
			wrapped, err := Wrap(v)
			if err != nil {
				return mapping, err
			}
			mapping.Map.Data[k] = wrapped
		}
		return mapping, nil
	}
	return Value{}, errors.New("Cannot wrap values of the type of the argument provided.")
}
