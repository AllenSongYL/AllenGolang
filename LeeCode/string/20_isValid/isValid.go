package main

// 给定一个只包括 "("，")"，"{"，"}"，"["，"]" 的字符串 s ，判断字符串是否有效。\
// copilot
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]string, 0)
	for _, i := range s {
		v := string(i)
		if v == "(" || v == "[" || v == "{" {
			stack = append(stack, string(v))
		} else {
			if len(stack) == 0 {
				return false
			}
			if v == ")" && stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
			} else if v == "]" && stack[len(stack)-1] == "[" {
				stack = stack[:len(stack)-1]
			} else if v == "}" && stack[len(stack)-1] == "{" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func main() {

}
