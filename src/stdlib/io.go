package stdlib

import (
	uni "../interpreter"
	"fmt"
)

func SLIB_Print(arguments ...interface{}) (uni.Value, error) {
	fmt.Println(arguments...)
	return uni.Value{}, nil
}
