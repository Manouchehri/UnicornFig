package stdlib

import (
	uni "../interpreter"
)

var StandardLibrary uni.Environment = uni.Environment{
	"true":     uni.NewBoolean(true),
	"false":    uni.NewBoolean(false),
	"pi":       uni.NewFloat(3.141592653589793),
	"mul":      uni.NewCallableFunction("mul", []string{"a", "b"}, SLIB_Multiply),
	"div":      uni.NewCallableFunction("div", []string{"a", "b"}, SLIB_Divide),
	"add":      uni.NewCallableFunction("add", []string{"a", "b"}, SLIB_Add),
	"sub":      uni.NewCallableFunction("sub", []string{"a", "b"}, SLIB_Subtract),
	"concat":   uni.NewCallableFunction("concat", []string{"s1", "s2"}, SLIB_Concatenate),
	"substr":   uni.NewCallableFunction("substr", []string{"str", "start", "end"}, SLIB_Substring),
	"index":    uni.NewCallableFunction("index", []string{"s1", "s2"}, SLIB_Index),
	"length":   uni.NewCallableFunction("length", []string{"str"}, SLIB_Length),
	"upcase":   uni.NewCallableFunction("upcase", []string{"str"}, SLIB_Upcase),
	"downcase": uni.NewCallableFunction("downcase", []string{"str"}, SLIB_Downcase),
	"not":      uni.NewCallableFunction("not", []string{"value"}, SLIB_Negate),
	"zero":     uni.NewCallableFunction("zero", []string{"n"}, SLIB_IsZero),
	"and":      uni.NewCallableFunction("and", []string{"b1", "b2"}, SLIB_And),
	"or":       uni.NewCallableFunction("or", []string{"b1", "b2"}, SLIB_Or),
	"equal":    uni.NewCallableFunction("equal", []string{"a", "b"}, SLIB_Equal),
	"print":    uni.NewCallableFunction("print", []string{"msg"}, SLIB_Print),
}
