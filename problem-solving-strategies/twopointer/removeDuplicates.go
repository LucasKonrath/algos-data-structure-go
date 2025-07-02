package twopointer

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	tortoise := 0
	for hare := 1; hare < len(nums); hare++ {
		if nums[hare] != nums[tortoise] {
			tortoise++
			nums[tortoise] = nums[hare]
		}
	}
	return tortoise + 1
}
