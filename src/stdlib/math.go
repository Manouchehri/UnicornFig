package stdlib

import (
  uni "../interpreter"
  "errors"
)

func SLIB_Multiply(env uni.Environment, arguments ...interface{}) (error, uni.Value, uni.Environment){
  if len(arguments) < 2 {
    return errors.New("Multiply function expects two or more arguments."), uni.Value{}, env
  }
  product := arguments[0].(int64)
  for i := 1; i < len(arguments); i++ {
    product *= arguments[i].(int64)
  }
  return nil, uni.NewInteger(product), env
}
