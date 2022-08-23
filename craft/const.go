package craft

// 状态集合
var (
	// 初始状态
	Init = "Init"
	// int关键词状态集合
	Int1 = "int1"
	Int2 = "int2"
	Int3 = "int3"
	// 运算符状态集合
	Plus  = "plus"  // 加法
	Minus = "minus" // 减法
	Star  = "star"  // 乘法
	Slash = "slash" // 除法
	// 比较符状态集合
	Gt = "gt" // 大于
	Ge = "ge" // 大于等于
	// 数字状态
	IntLiteral = "IntLiteral"
	// 标志符状态
	Assignment = "Assignment"
)
