package stack

func isBalanced(s string) bool {
	stack := &Stack{}
	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack.Push(int(char))
		} else if char == ')' || char == '}' || char == ']' {
			if stack.isEmpty() {
				return false
			}
			top, _ := stack.Pop()
			if !isMatchingPair(top, int(char)) {
				return false
			}
		}
	}
	return stack.isEmpty()
}

func isMatchingPair(top int, i int) bool {
	if (top == '(' && i == ')') || (top == '{' && i == '}') || (top == '[' && i == ']') {
		return true
	}
	return false
}
