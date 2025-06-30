package maxsubarray

func maxSubArray(nums []int) int {
	return maxSubArrayRecur(nums, 0, len(nums)-1)
}

func maxSubArrayRecur(nums []int, left, right int) int {
	if left == right {
		return nums[left]
	}

	mid := (left + right) / 2

	leftMax := maxSubArrayRecur(nums, left, mid)
	rightMax := maxSubArrayRecur(nums, mid+1, right)

	crossMax := maxCrossingSum(nums, left, mid, right)

	return max(leftMax, max(rightMax, crossMax))
}

func maxCrossingSum(nums []int, left, mid, right int) int {
	leftSum := -1 << 31 // Minimum possible value
	sum := 0
	for i := mid; i >= left; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	rightSum := -1 << 31 // Minimum possible value
	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	return leftSum + rightSum
}
