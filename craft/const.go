package craft

// 状态集合
var (
	// 初始状态
	Init = "Init"
	// int关键词状态集合
	Int1 = "int1"
	Int2 = "int2"
	Int  = "int"
	// 运算符状态集合
	Plus  = "plus"  // +
	Minus = "minus" // -
	Star  = "star"  // *
	Slash = "slash" // /
	// 比较符状态集合
	Gt = "gt" // >
	Ge = "ge" // >=
	Eq = "eq" // ==
	Le = "le" // <=
	Lt = "lt" // <
	// id状态
	Id = "id"

	SemiColon  = "semiColon"  // ;
	LeftParen  = "leftParen"  // (
	RightParen = "rightParen" // )
	Assignment = "Assignment" // =

	If   = "if"
	Else = "else"

	IntLiteral    = "IntLiteral"    // 数字字面量
	StringLiteral = "stringLiteral" // 字符串字面量
)
