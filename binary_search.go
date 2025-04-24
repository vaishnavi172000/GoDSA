package godsa

import "fmt"

// BinarySearch uses binary search algo to find the element in the array
// Complexity o(logn)
// Space Complexity o(1)
func BinarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		fmt.Println(mid)
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1

}
