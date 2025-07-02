package stack

func ReverseString(s string) string {
	stack := &Stack{}
	for _, char := range s {
		stack.Push(int(char))
	}
	var reversed []rune

	for r, ok := stack.Pop(); ok; r, ok = stack.Pop() {
		reversed = append(reversed, rune(r))
	}
	return string(reversed)
}
