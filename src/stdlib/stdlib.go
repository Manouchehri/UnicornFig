package stdlib

import (
	uni "../interpreter"
)

var ConstantNames = []string{
	"true",
	"false",
	"pi",
}

var StandardLibrary uni.Environment = uni.Environment{
	"true":     uni.NewBoolean(true),
	"false":    uni.NewBoolean(false),
	"pi":       uni.NewFloat(3.141592653589793),
	"*":        uni.NewCallableFunction("*", []string{"a", "b"}, SLIB_Multiply),
	"/":        uni.NewCallableFunction("/", []string{"a", "b"}, SLIB_Divide),
	"+":        uni.NewCallableFunction("+", []string{"a", "b"}, SLIB_Add),
	"-":        uni.NewCallableFunction("-", []string{"a", "b"}, SLIB_Subtract),
	"%":        uni.NewCallableFunction("%", []string{"a", "b"}, SLIB_Modulo),
	"concat":   uni.NewCallableFunction("concat", []string{"s1", "s2"}, SLIB_Concatenate),
	"substr":   uni.NewCallableFunction("substr", []string{"str", "start", "end"}, SLIB_Substring),
	"index":    uni.NewCallableFunction("index", []string{"s1", "s2"}, SLIB_Index),
	"length":   uni.NewCallableFunction("length", []string{"str"}, SLIB_Length),
	"upcase":   uni.NewCallableFunction("upcase", []string{"str"}, SLIB_Upcase),
	"downcase": uni.NewCallableFunction("downcase", []string{"str"}, SLIB_Downcase),
	"split":    uni.NewCallableFunction("split", []string{"_str_", "_sep_"}, SLIB_Split),
	"at":       uni.NewCallableFunction("at", []string{"_str_", "_index_"}, SLIB_AtIndex),
	"not":      uni.NewCallableFunction("not", []string{"value"}, SLIB_Negate),
	"zero?":    uni.NewCallableFunction("zero?", []string{"n"}, SLIB_IsZero),
	"and":      uni.NewCallableFunction("and", []string{"b1", "b2"}, SLIB_And),
	"or":       uni.NewCallableFunction("or", []string{"b1", "b2"}, SLIB_Or),
	"=":        uni.NewCallableFunction("=", []string{"a", "b"}, SLIB_Equal),
	">":        uni.NewCallableFunction(">", []string{"a", "b"}, SLIB_GreaterThan),
	"<":        uni.NewCallableFunction("<", []string{"a", "b"}, SLIB_LessThan),
	">=":       uni.NewCallableFunction(">=", []string{"a", "b"}, SLIB_GreaterOrEqual),
	"<=":       uni.NewCallableFunction("<=", []string{"a", "b"}, SLIB_LessOrEqual),
	"list":     uni.NewCallableFunction("list", []string{}, SLIB_List),
	"first":    uni.NewCallableFunction("first", []string{"_list_"}, SLIB_First),
	"tail":     uni.NewCallableFunction("tail", []string{"_list_"}, SLIB_Tail),
	"append":   uni.NewCallableFunction("append", []string{"_list_", "_value_"}, SLIB_Append),
	"size":     uni.NewCallableFunction("size", []string{"_list_"}, SLIB_Size),
	"mapping":  uni.NewCallableFunction("mapping", []string{}, SLIB_Map),
	"assoc":    uni.NewCallableFunction("assoc", []string{"_map_", "_key1_", "_value1_"}, SLIB_Associate),
	"get":      uni.NewCallableFunction("get", []string{"_map_", "_key_"}, SLIB_GetMap),
	"keys":     uni.NewCallableFunction("keys", []string{"_map_"}, SLIB_Keys),
	"print":    uni.NewCallableFunction("print", []string{"msg"}, SLIB_Print),
	"env":      uni.NewCallableFunction("env", []string{"_envvar_"}, SLIB_Environment),
	"ignored":  uni.NewCallableFunction("ignored", []string{"_value_"}, SLIB_Ignore),
}
