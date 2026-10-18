package main

import "fmt"

//func isValid(s string) bool {
//	n := len(s)
//	if n%2 == 1 {
//		return false
//	}
//	pairs := map[byte]byte{
//		')': '(',
//		']': '[',
//		'}': '{',
//	}
//	//stack := []byte{}
//	var stack []byte
//	for i := 0; i < n; i++ {
//		if pairs[s[i]] > 0 { // 判断是不是右括号
//			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
//				return false
//			}
//			stack = stack[:len(stack)-1]
//		} else {
//			stack = append(stack, s[i])
//		}
//	}
//	return len(stack) == 0
//}

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	var stack []byte
	for i := 0; i < n; i++ {
		// 左括号，入栈
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			// 右括号，出栈比较
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (s[i] == ')' && top != '(') ||
				(s[i] == ']' && top != '[') ||
				(s[i] == '}' && top != '{') {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
	m := map[string]byte{"a": '{', "b": '}', "c": 'c'}
	fmt.Println(m["a"], m["b"], m["qqq"])
	m2 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(m2["a"], m2["b"], m2["ccc"])
}

/*
判断逻辑：
当 s[i] 是一个右括号（)、] 或 }）时，pairs[s[i]] 会返回对应的左括号（(、[ 或 {），这些字符的ASCII值都大于0。
当 s[i] 不是右括号时（即为左括号或其它字符），由于map中没有该键，Go会返回零值（空字节 \x00），其值为0。
因此：

如果 pairs[s[i]] > 0 为真，则说明 s[i] 是一个右括号
如果 pairs[s[i]] > 0 为假，则说明 s[i] 不是右括号（是左括号）
这样通过 pairs[s[i]] > 0 就可以区分当前字符是左括号还是右括号了。

这种设计巧妙地利用了Go语言map的特性：当访问不存在的键时返回零值，从而实现了一种简洁的左右括号判断方式。
*/
