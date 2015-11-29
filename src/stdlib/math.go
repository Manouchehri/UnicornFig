package stdlib

import (
	uni "../interpreter"
	"errors"
)

func SLIB_Multiply(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) < 2 {
		return uni.Value{}, errors.New("Multiply function expects two or more arguments.")
	}
	product := arguments[0].(int64)
	for i := 1; i < len(arguments); i++ {
		product *= arguments[i].(int64)
	}
	return uni.NewInteger(product), nil
}

func SLIB_Divide(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) < 2 {
		return uni.Value{}, errors.New("Divide function expects two or more arguments.")
	}
	result := arguments[0].(int64)
	for i := 1; i < len(arguments); i++ {
		next := arguments[i].(int64)
		if next == 0 {
			return uni.Value{}, errors.New("Cannot divide by zero.")
		}
		result /= next
	}
	return uni.NewInteger(result), nil
}

func SLIB_Add(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) < 2 {
		return uni.Value{}, errors.New("Add function expects two or more arguments.")
	}
	sum := arguments[0].(int64)
	for i := 1; i < len(arguments); i++ {
		sum += arguments[i].(int64)
	}
	return uni.NewInteger(sum), nil
}

func SLIB_Subtract(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) < 2 {
		return uni.Value{}, errors.New("Subtract function expects two or more arguments.")
	}
	result := arguments[0].(int64)
	for i := 1; i < len(arguments); i++ {
		result -= arguments[i].(int64)
	}
	return uni.NewInteger(result), nil
}

func SLIB_GreaterThan(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 2 {
		return uni.Value{}, errors.New("Greater-Than function expects exactly two arguments.")
	}
	result := arguments[0].(int64) > arguments[1].(int64)
	return ToBoolKeyword(result), nil
}

func SLIB_LessThan(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 2 {
		return uni.Value{}, errors.New("Less-Than function expects exactly two arguments.")
	}
	result := arguments[0].(int64) < arguments[1].(int64)
	return ToBoolKeyword(result), nil
}

func SLIB_GreaterOrEqual(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 2 {
		return uni.Value{}, errors.New("Greater-Than-Or-Equal function expects exactly two arguments.")
	}
	result := arguments[0].(int64) >= arguments[1].(int64)
	return ToBoolKeyword(result), nil
}

func SLIB_LessOrEqual(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 2 {
		return uni.Value{}, errors.New("Less-Than-Or-Equal function expects exactly two arguments.")
	}
	result := arguments[0].(int64) >= arguments[1].(int64)
	return ToBoolKeyword(result), nil
}

func SLIB_Modulo(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) < 2 {
		return uni.Value{}, errors.New("Modulo function expects two or more arguments.")
	}
	result := arguments[0].(int64)
	for i := 1; i < len(arguments); i++ {
		result = result % arguments[i].(int64)
	}
	return uni.NewInteger(result), nil
}
