package craft

// Lexer 词法解析器
type Lexer struct {
}

// 根据字符判断当前应该处于什么状态
func (lexer *Lexer) initState(s byte) string {
	if isDigital(s) {
		return IntLiteral
	} else if isAlpha(s) {
		if s == 'i' {
			return Int1
		} else {
			return Id
		}
	} else if s == '=' {
		return Assignment
	} else if s == '>' {
		return Gt
	} else if s == '<' {
		return Lt
	} else if s == '+' {
		return Plus
	} else if s == '-' {
		return Minus
	} else if s == '*' {
		return Star
	} else if s == '/' {
		return Slash
	} else if s == '(' {
		return LeftParen
	} else if s == ')' {
		return RightParen
	} else if s == ';' {
		return SemiColon
	}

	return ""
}

// 基于当前状态来解析出token
func (lexer *Lexer) tokenize(state string, s string) {
}

// 第一步：状态机定义（状态定义，流转分析）

// 第二步：词法分析器实现
