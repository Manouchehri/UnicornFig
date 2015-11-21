package stdlib

import (
  uni "../interpreter"
)

var StandardLibrary uni.Environment = uni.Environment {
  "mul": uni.NewCallableFunction("mul", []string{"a", "b"}, SLIB_Multiply),
  "div": uni.NewCallableFunction("div", []string{"a", "b"}, SLIB_Divide),
  "add": uni.NewCallableFunction("add", []string{"a", "b"}, SLIB_Add),
  "sub": uni.NewCallableFunction("sub", []string{"a", "b"}, SLIB_Subtract),
  "concat": uni.NewCallableFunction("concat", []string{"s1", "s2"}, SLIB_Concatenate),
  "substr": uni.NewCallableFunction("substr", []string{"str", "start", "end"}, SLIB_Substring),
  "index": uni.NewCallableFunction("index", []string{"s1", "s2"}, SLIB_Index),
  "length": uni.NewCallableFunction("length", []string{"str"}, SLIB_Length),
  "upcase": uni.NewCallableFunction("upcase", []string{"str"}, SLIB_Upcase),
  "downcase": uni.NewCallableFunction("downcase", []string{"str"}, SLIB_Downcase),
  "print": uni.NewCallableFunction("print", []string{"msg"}, SLIB_Print),
}
