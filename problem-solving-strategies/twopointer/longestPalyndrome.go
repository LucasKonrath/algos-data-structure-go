package twopointer

func longestPalyndrome(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1
	for i := 0; i < len(s); i++ { // Iterate through each character in the string
		if len(s)-i <= maxLen/2 {
			break // No longer possible to find a longer palindrome
		}
		left, right := i, i

		for right < len(s)-1 && s[right+1] == s[right] { // Skip duplicates
			right++ // Skip duplicates
		}
		i = right + 1
		for left > 0 && right < len(s)-1 && s[left-1] == s[right+1] { // Expand around the center
			left-- // Expand around the center
			right++
		}
		newLen := right - left + 1
		if newLen > maxLen { // Found a longer palindrome
			start = left
			maxLen = newLen
		}
	}
	return s[start : start+maxLen] // Return the longest palindromic substring
}
