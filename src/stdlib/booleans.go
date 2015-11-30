package stdlib

import (
	uni "../interpreter"
	"errors"
)

func ToBoolKeyword(value bool) uni.Value {
	if value {
		return uni.NewBoolean(true)
	} else {
		return uni.NewBoolean(false)
	}
}

func SLIB_Negate(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 1 {
		return uni.Value{}, errors.New("Negation function expects exactly one argument.")
	}
	value := arguments[0].(bool)
	return ToBoolKeyword(!value), nil
}

func SLIB_IsZero(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 1 {
		return uni.Value{}, errors.New("Zero predicate function expects exactly one argument.")
	}
	isZero := false
	switch arguments[0].(type) {
	case int64:
		isZero = arguments[0].(int64) == int64(0)
	case float64:
		isZero = arguments[0].(float64) == float64(0.0)
	default:
		return uni.Value{}, errors.New("Zero predicate function expects a numeric argument.")
	}
	return ToBoolKeyword(isZero), nil
}

func SLIB_And(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) < 2 {
		return uni.Value{}, errors.New("And function expects two or more arguments.")
	}
	result := arguments[0].(bool)
	// Compound the values provided to the function.
	// We can short circuit as soon as `false` is encountered.
	for i := 1; i < len(arguments) && result; i++ {
		result = result && arguments[i].(bool)
	}
	return ToBoolKeyword(result), nil
}

func SLIB_Or(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) < 2 {
		return uni.Value{}, errors.New("And function expects two or more arguments.")
	}
	result := arguments[0].(bool)
	// We can short circuit as soon as `true` is encountered.
	for i := 1; i < len(arguments); i++ {
		result = result || arguments[i].(bool)
		if result {
			break
		}
	}
	return ToBoolKeyword(result), nil
}

func SLIB_Equal(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) < 2 {
		return uni.Value{}, errors.New("Equal function expects two or more arguments.")
	}
	result := true
	switch arguments[0].(type) {
	case int64:
		value := arguments[0].(int64)
		for i := 1; i < len(arguments) && result; i++ {
			result = value == arguments[i].(int64)
		}
	case float64:
		value := arguments[0].(float64)
		for i := 1; i < len(arguments) && result; i++ {
			result = value == arguments[i].(float64)
		}
	case string:
		value := arguments[0].(string)
		for i := 1; i < len(arguments) && result; i++ {
			result = value == arguments[i].(string)
		}
	case bool:
		value := arguments[0].(bool)
		for i := 1; i < len(arguments) && result; i++ {
			result = value == arguments[i].(bool)
		}
	}
	return ToBoolKeyword(result), nil
}
