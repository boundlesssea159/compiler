package craft

// Lexer 词法解析器
type Lexer struct {
	s      string  // 脚本语句
	tokens []Token // 解析出来的所有token
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
func (lexer *Lexer) tokenize() {
	token := Token{Tp: Init}
	start := 0
	for start < len(lexer.s) {
		s := lexer.s[start]
		if token.Tp == Init {
			token.Tp = lexer.initState(s)
			token.Text = append(token.Text, s)
			start++
		} else if token.Tp == Int1 {
			if s == 'n' {
				token.Tp = Int2
				token.Text = append(token.Text, s)
				start++
			} else if isDigital(s) || isAlpha(s) {
				token.Tp = Id
				token.Text = append(token.Text, s)
				start++
			} else {
				lexer.tokens = append(lexer.tokens, token)
				token = Token{Tp: Init}
			}
		} else if token.Tp == Int2 {
			if s == 't' {
				token.Tp = Int
				token.Text = append(token.Text, s)
				start++
			} else if isDigital(s) || isAlpha(s) {
				token.Tp = Id
				token.Text = append(token.Text, s)
				start++
			} else { // "in>" 怎么办？还是会解析出in和>两个token吗？词法解析不考虑合法性？
				lexer.tokens = append(lexer.tokens, token)
				token = Token{Tp: Init}
			}
		} else if token.Tp == Int {
			if isBlank(s) {
				lexer.tokens = append(lexer.tokens, token)
				token = Token{Tp: Init}
				start++
			}else {
				token.Tp = Id
				token.Text = append(token.Text, s)
				start++
			}
		}
	}
}

// 第一步：状态机定义（状态定义，流转分析）

// 第二步：词法分析器实现
