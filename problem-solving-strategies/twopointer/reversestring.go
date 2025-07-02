package twopointer

func reverseString(s string) string {
	bytes := []byte(s)
	left, right := 0, len(bytes)-1

	for left < right {
		bytes[left], bytes[right] = bytes[right], bytes[left]
		left++
		right--
	}
	return string(bytes)
}
