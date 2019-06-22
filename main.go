package main

func main() {
	if true {
		s1 := "(()"
		s2 := ")()())"
		println(longestValidParentheses(s1))
		println(longestValidParentheses(s2))
	}
}
