package slidingwindow

func longestUniquesSubstring(s string) string {
	charIndex := make(map[rune]int)
	start := 0
	maxLength := 0
	maxStart := 0

	for i, char := range s {
		// Check if the character is already in the map and its index is within the current window
		if lastIndex, found := charIndex[char]; found && lastIndex >= start {
			start = lastIndex + 1 // Move the start to the right of the last occurrence
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1 // Update maxLength if the current window is larger
			maxStart = start          // Update the starting index of the longest substring
		}
		charIndex[char] = i // Update the last index of the character
	}
	return s[maxStart : maxStart+maxLength] // Return the longest unique substring
}
