package main

import (
	"log"
	"main/utils"
)

/*
	给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

	有效字符串需满足：

	左括号必须用相同类型的右括号闭合。
	左括号必须以正确的顺序闭合。
	注意空字符串可被认为是有效字符串。

	示例 1:

	输入: "()"
	输出: true
	示例 2:

	输入: "()[]{}"
	输出: true
	示例 3:

	输入: "(]"
	输出: false
	示例 4:

	输入: "([)]"
	输出: false
	示例 5:

	输入: "{[]}"
	输出: true

	来源：力扣（LeetCode）
	链接：https://leetcode-cn.com/problems/valid-parentheses
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


// 判断括号是否有效
func isValid(s string) bool {
	lenS := len(s)

	if lenS == 0 {
		return true
	}

	if lenS % 2 != 0 {
		return false
	}

	stack := utils.NewStack()
	parenthesesMap := map[byte]byte {'}': '{', ']': '[', ')': '('}

	// 当出现右括号的时候，就不断得从栈中弹出数据匹配，判断是否是左括号
	for idx, llen := 0, len(s); idx < llen; idx = idx +1 {
		if s[idx] == ' ' {
			continue
		}

		if s[idx] == ']' || s[idx] == '}' || s[idx] == ')' {
			for ; ;  {
				if stack.Empty().(bool) {
					return false
				}

				peek := stack.Pop().(byte)
				value := parenthesesMap[s[idx]]
				if peek == value {
					break
				}
			}

		} else {
			stack.Push(s[idx])
		}
	}

	return stack.Empty().(bool)
}


func main() {
	result := isValid("")
	if !result {
		log.Printf("error: %v", result)
		return
	}

	result = isValid("}")
	if result {
		log.Printf("error: %v", result)
		return
	}

	result = isValid("()[]{}")
	if !result {
		log.Printf("error: %v", result)
		return
	}

	result = isValid("([)]")
	if result {
		log.Printf("error: %v", result)
		return
	}

	result = isValid("([]")
	if result {
		log.Printf("error: %v", result)
		return
	}

	result = isValid("{[]}")
	if !result {
		log.Printf("error: %v", result)
		return
	}

	result = isValid("}}")
	if result {
		log.Printf("error: %v", result)
		return
	}

	log.Printf("success")
}
