package stdlib

import (
  uni "../interpreter"
  "fmt"
)

func SLIB_Print(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment) {
  fmt.Println(arguments...)
  return nil, uni.Value{}, env
}
