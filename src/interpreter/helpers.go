package interpreter

var (
	zeroi  = IntegerLiteral{0}
	zerof  = FloatLiteral{0.0}
	emptys = StringLiteral{""}
	falseb = BooleanLiteral{false}
	emptyl = ListNode{nil, nil}
	emptym = Mapping{map[string]interface{}{}}
)

func NewString(str string) Value {
	return Value{StringT, StringLiteral{str}, zeroi, zerof, Name{}, falseb, Function{}, &emptyl, emptym}
}

func NewInteger(n int64) Value {
	return Value{IntegerT, emptys, IntegerLiteral{n}, zerof, Name{}, falseb, Function{}, &emptyl, emptym}
}

func NewFloat(n float64) Value {
	return Value{FloatT, emptys, zeroi, FloatLiteral{n}, Name{}, falseb, Function{}, &emptyl, emptym}
}

func NewName(identifier string) Value {
	return Value{NameT, emptys, zeroi, zerof, Name{identifier}, falseb, Function{}, &emptyl, emptym}
}

func NewBoolean(value bool) Value {
	return Value{BooleanT, emptys, zeroi, zerof, Name{}, BooleanLiteral{value}, Function{}, &emptyl, emptym}
}

func NewSExpression(formName string, values ...interface{}) SExpression {
	emptyArray := make([]interface{}, 0)
	sexp := SExpression{Name{formName}, SExpressionT, emptyArray}
	for _, value := range values {
		sexp.Values = append(sexp.Values, value)
	}
	return sexp
}

func NewCallableFunction(name string, argNames []string, fn Builtin) Value {
	names := make([]Name, len(argNames))
	for i, arg := range argNames {
		names[i] = Name{arg}
	}
	return Value{FunctionT, emptys, zeroi, zerof, Name{}, falseb, Function{Name{name}, names, SExpression{}, true, fn}, &emptyl, emptym}
}

func NewFunction(name string, argNames []string, body interface{}) Value {
	names := make([]Name, len(argNames))
	for i, arg := range argNames {
		names[i] = Name{arg}
	}
	return Value{FunctionT, emptys, zeroi, zerof, Name{}, falseb, Function{Name{name}, names, body, false, nil}, &emptyl, emptym}
}

func NewList() Value {
	return Value{ListT, emptys, zeroi, zerof, Name{}, falseb, Function{}, &emptyl, emptym}
}

func NewMap() Value {
	return Value{MapT, emptys, zeroi, zerof, Name{}, falseb, Function{}, &emptyl, emptym}
}
