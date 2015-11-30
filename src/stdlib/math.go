package stdlib

import (
	uni "../interpreter"
	"errors"
)

func SLIB_Multiply(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) < 2 {
		return uni.Value{}, errors.New("Multiply function expects two or more arguments.")
	}
	anyFloats := false
	var product float64 = float64(1.0)
	for i := 0; i < len(arguments); i++ {
		switch arguments[i].(type) {
		case int64:
			product *= float64(arguments[i].(int64))
		case float64:
			product *= arguments[i].(float64)
			anyFloats = true
		default:
			return uni.Value{}, errors.New("Multiply expects all arguments to be numbers.")
		}
	}
	if anyFloats {
		return uni.NewFloat(product), nil
	} else {
		return uni.NewInteger(int64(product)), nil
	}
}

func SLIB_Divide(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) < 2 {
		return uni.Value{}, errors.New("Divide function expects two or more arguments.")
	}
	anyFloats := false
	var result float64
	switch arguments[0].(type) {
	case int64:
		result = float64(arguments[0].(int64))
	case float64:
		result = arguments[0].(float64)
		anyFloats = true
	default:
		return uni.Value{}, errors.New("Divide expects all arguments to be numbers.")
	}
	for i := 1; i < len(arguments); i++ {
		switch arguments[i].(type) {
		case int64:
			next := arguments[i].(int64)
			if next == int64(0) {
				return uni.Value{}, errors.New("Cannot divide by zero.")
			}
			result /= float64(next)
		case float64:
			next := arguments[i].(float64)
			anyFloats = true
			if next == float64(0.0) {
				return uni.Value{}, errors.New("Cannot divide by zero.")
			}
			result /= next
		default:
			return uni.Value{}, errors.New("Divide expects all arguments to be numbers.")
		}
	}
	if anyFloats {
		return uni.NewFloat(result), nil
	} else {
		return uni.NewInteger(int64(result)), nil
	}
}

func SLIB_Add(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) < 2 {
		return uni.Value{}, errors.New("Add function expects two or more arguments.")
	}
	anyFloats := false
	var sum float64 = float64(0.0)
	for i := 0; i < len(arguments); i++ {
		switch arguments[i].(type) {
		case int64:
			sum += float64(arguments[i].(int64))
		case float64:
			sum += arguments[i].(float64)
			anyFloats = true
		default:
			return uni.Value{}, errors.New("Add expects all arguments to be numbers.")
		}
	}
	if anyFloats {
		return uni.NewFloat(sum), nil
	} else {
		return uni.NewInteger(int64(sum)), nil
	}
}

func SLIB_Subtract(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) < 2 {
		return uni.Value{}, errors.New("Subtract function expects two or more arguments.")
	}
	anyFloats := false
	var result float64
	switch arguments[0].(type) {
	case int64:
		result = float64(arguments[0].(int64))
	case float64:
		result = arguments[0].(float64)
		anyFloats = true
	default:
		return uni.Value{}, errors.New("Subtract expects all arguments to be numbers.")
	}
	for i := 1; i < len(arguments); i++ {
		switch arguments[i].(type) {
		case int64:
			result -= float64(arguments[i].(int64))
		case float64:
			result -= arguments[i].(float64)
		default:
			return uni.Value{}, errors.New("Subtract expects all arguments to be numbers.")
		}
	}
	if anyFloats {
		return uni.NewFloat(result), nil
	} else {
		return uni.NewInteger(int64(result)), nil
	}
}

func SLIB_GreaterThan(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 2 {
		return uni.Value{}, errors.New("Greater-Than function expects exactly two arguments.")
	}
	var first, second float64
	switch arguments[0].(type) {
	case int64:
		first = float64(arguments[0].(int64))
	case float64:
		first = arguments[0].(float64)
	default:
		return uni.Value{}, errors.New("Greater-Than function expects two numbers.")
	}
	switch arguments[1].(type) {
	case int64:
		first = float64(arguments[1].(int64))
	case float64:
		first = arguments[1].(float64)
	default:
		return uni.Value{}, errors.New("Greater-Than function expects two numbers.")
	}
	result := first > second
	return ToBoolKeyword(result), nil
}

func SLIB_LessThan(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 2 {
		return uni.Value{}, errors.New("Less-Than function expects exactly two arguments.")
	}
	var first, second float64
	switch arguments[0].(type) {
	case int64:
		first = float64(arguments[0].(int64))
	case float64:
		first = arguments[0].(float64)
	default:
		return uni.Value{}, errors.New("Less-Than function expects two numbers.")
	}
	switch arguments[1].(type) {
	case int64:
		first = float64(arguments[1].(int64))
	case float64:
		first = arguments[1].(float64)
	default:
		return uni.Value{}, errors.New("Less-Than function expects two numbers.")
	}
	result := first > second
	return ToBoolKeyword(result), nil
}

func SLIB_GreaterOrEqual(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 2 {
		return uni.Value{}, errors.New("Greater-Than-Or-Equal function expects exactly two arguments.")
	}
	var first, second float64
	switch arguments[0].(type) {
	case int64:
		first = float64(arguments[0].(int64))
	case float64:
		first = arguments[0].(float64)
	default:
		return uni.Value{}, errors.New("Greater-Than-Or-Equal function expects two numbers.")
	}
	switch arguments[1].(type) {
	case int64:
		first = float64(arguments[1].(int64))
	case float64:
		first = arguments[1].(float64)
	default:
		return uni.Value{}, errors.New("Greater-Than-Or-Equal function expects two numbers.")
	}
	result := first > second
	return ToBoolKeyword(result), nil
}

func SLIB_LessOrEqual(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 2 {
		return uni.Value{}, errors.New("Less-Than-Or-Equal function expects exactly two arguments.")
	}
	var first, second float64
	switch arguments[0].(type) {
	case int64:
		first = float64(arguments[0].(int64))
	case float64:
		first = arguments[0].(float64)
	default:
		return uni.Value{}, errors.New("Less-Than-Or-Equal function expects two numbers.")
	}
	switch arguments[1].(type) {
	case int64:
		first = float64(arguments[1].(int64))
	case float64:
		first = arguments[1].(float64)
	default:
		return uni.Value{}, errors.New("Less-Than-Or-Equal function expects two numbers.")
	}
	result := first > second
	return ToBoolKeyword(result), nil
}

func SLIB_Modulo(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) < 2 {
		return uni.Value{}, errors.New("Modulo function expects exactly two arguments.")
	}
	var first, second int64
	switch arguments[0].(type) {
	case int64:
		first = arguments[0].(int64)
	default:
		return uni.Value{}, errors.New("Modulo function expects both arguments to be integers.")
	}
	switch arguments[1].(type) {
	case int64:
		second = arguments[1].(int64)
	default:
		return uni.Value{}, errors.New("Modulo function expects both arguments to be integers.")
	}
	var result int64 = first % second
	return uni.NewInteger(result), nil
}
