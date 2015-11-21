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
