package stdlib

import (
	uni "../interpreter"
	"errors"
)

func toBoolKeyword(value bool) uni.Value {
	if value {
		return uni.NewName("true")
	} else {
		return uni.NewName("false")
	}
}

func SLIB_Negate(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	if len(arguments) != 1 {
		return errors.New("Negation function expects exactly one argument."), uni.Value{}, env
	}
	value := arguments[0].(bool)
	return nil, toBoolKeyword(!value), env
}

func SLIB_IsZero(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	if len(arguments) != 1 {
		return errors.New("Zero predicate function expects exactly one argument."), uni.Value{}, env
	}
	value := arguments[0].(int64)
	isZero := value == int64(0)
	return nil, toBoolKeyword(isZero), env
}

func SLIB_And(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	if len(arguments) < 2 {
		return errors.New("And function expects two or more arguments."), uni.Value{}, env
	}
	result := arguments[0].(bool)
	// Compound the values provided to the function.
	// We can short circuit as soon as `false` is encountered.
	for i := 1; i < len(arguments) && result; i++ {
		result = result && arguments[i].(bool)
	}
	return nil, toBoolKeyword(result), env
}

func SLIB_Or(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	if len(arguments) < 2 {
		return errors.New("And function expects two or more arguments."), uni.Value{}, env
	}
	result := arguments[0].(bool)
	// We can short circuit as soon as `true` is encountered.
	for i := 1; i < len(arguments); i++ {
		result = result || arguments[i].(bool)
		if result {
			break
		}
	}
	return nil, toBoolKeyword(result), env
}

func SLIB_Equal(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	if len(arguments) < 2 {
		return errors.New("Equal function expects two or more arguments."), uni.Value{}, env
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
	return nil, toBoolKeyword(result), env
}
