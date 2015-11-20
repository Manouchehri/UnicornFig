package stdlib

import (
  uni "../interpreter"
)

var StandardLibrary uni.Environment = uni.Environment {
  "mult": uni.NewCallableFunction("mult", []string{"a", "b"}, SLIB_Multiply),
  "print": uni.NewCallableFunction("print", []string{"msg"}, SLIB_Print),
}
