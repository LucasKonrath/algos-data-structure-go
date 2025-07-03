package twopointer

import "sort"

func threeSums(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // Skip duplicates for the first element
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++ // Skip duplicates for the second element
				}
				for left < right && nums[right] == nums[right-1] {
					right-- // Skip duplicates for the third element
				}
				left++  // Move left pointer to continue searching
				right-- // Move right pointer to continue searching
			} else if sum < 0 {
				left++ // Move left pointer to increase sum
			} else {
				right-- // Move right pointer to decrease sum
			}
		}
	}
	return result
}
