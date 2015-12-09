package stdlib

import (
	uni "../interpreter"
	"errors"
	"fmt"
	"os"
)

func SLIB_Print(arguments ...interface{}) (uni.Value, error) {
	fmt.Println(arguments...)
	return uni.Value{}, nil
}

/**
 * Get the value of an environment variable
 */
func SLIB_Environment(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 1 {
		return uni.Value{}, errors.New("Environment function expects one argument.")
	}
	switch arguments[0].(type) {
	case string:
		break
	default:
		return uni.Value{}, errors.New("Environment function expects its argument to be a string.")
	}
	envVar := os.Getenv(arguments[0].(string))
	return uni.NewString(envVar), nil
}

func SLIB_Ignore(arguments ...interface{}) (uni.Value, error) {
	if len(arguments) != 1 {
		return uni.Value{}, errors.New("Ignore function expects only one argument.")
	}
	wrapped, err := uni.Wrap(arguments[0])
	wrapped.Ignored = true
	return wrapped, err
}
