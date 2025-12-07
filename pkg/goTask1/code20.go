package goTask1

//20. 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。

// 示例 1：
// 输入：s = "()"
// 输出：true

// 示例 2：
// 输入：s = "()[]{}"
// 输出：true

// 示例 3：
// 输入：s = "(]"
// 输出：false

// 示例 4：
// 输入：s = "([])"
// 输出：true

// 示例 5：
// 输入：s = "([)]"
// 输出：false

func IsValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	runes1 := []rune(s)
	runes2 := make([]rune, 0, len(runes1)/2)

	for _, value := range runes1 {
		switch value {
		case '(':
			runes2 = append(runes2, ')')
		case '[':
			runes2 = append(runes2, ']')
		case '{':
			runes2 = append(runes2, '}')
		case ')', ']', '}':
			if len(runes2) == 0 {
				return false
			}
			if runes2[len(runes2)-1] != value {
				return false
			}
			runes2 = runes2[:len(runes2)-1]
		default:
			return false
		}
	}

	return len(runes2) == 0
}
