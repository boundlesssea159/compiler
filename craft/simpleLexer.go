package craft

// Lexer 词法解析器
type Lexer struct {
}

// 根据字符判断当前应该处于什么状态
func (lexer *Lexer) initState(s byte) string {
	// 进入int关键词状态
	if s == 'i' {
		return Int1
	}
	// 进入数字状态
	if isDigital(s) {
		return IntLiteral
	}

	return ""
}

// 基于当前状态来解析出token
func (lexer *Lexer) tokenize(state string, s string) {

}

// 第一步：状态机定义（状态定义，流转分析）

// 第二步：词法分析器实现
