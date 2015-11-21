package stdlib

import (
	uni "../interpreter"
	"errors"
)

func SLIB_Multiply(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	if len(arguments) < 2 {
		return errors.New("Multiply function expects two or more arguments."), uni.Value{}, env
	}
	product := arguments[0].(int64)
	for i := 1; i < len(arguments); i++ {
		product *= arguments[i].(int64)
	}
	return nil, uni.NewInteger(product), env
}

func SLIB_Divide(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	if len(arguments) < 2 {
		return errors.New("Divide function expects two or more arguments."), uni.Value{}, env
	}
	result := arguments[0].(int64)
	for i := 1; i < len(arguments); i++ {
		next := arguments[i].(int64)
		if next == 0 {
			return errors.New("Cannot divide by zero."), uni.Value{}, env
		}
		result /= next
	}
	return nil, uni.NewInteger(result), env
}

func SLIB_Add(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	if len(arguments) < 2 {
		return errors.New("Add function expects two or more arguments."), uni.Value{}, env
	}
	sum := arguments[0].(int64)
	for i := 1; i < len(arguments); i++ {
		sum += arguments[i].(int64)
	}
	return nil, uni.NewInteger(sum), env
}

func SLIB_Subtract(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
	if len(arguments) < 2 {
		return errors.New("Subtract function expects two or more arguments."), uni.Value{}, env
	}
	result := arguments[0].(int64)
	for i := 1; i < len(arguments); i++ {
		result -= arguments[i].(int64)
	}
	return nil, uni.NewInteger(result), env
}
