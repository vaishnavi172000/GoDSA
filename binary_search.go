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

// You are given an integer array bloomDay, an integer m and an integer k.
// You want to make m bouquets. To make a bouquet, you need to use k adjacent flowers from the garden.
// The garden consists of n flowers, the ith flower will bloom in the bloomDay[i] and then can be used in exactly one bouquet.
// Return the minimum number of days you need to wait to be able to make m bouquets from the garden.
// If it is impossible to make m bouquets return -1.

// Example 1:

// Input: bloomDay = [1,10,3,10,2], m = 3, k = 1
// Output: 3
// Explanation: Let us see what happened in the first three days. x means flower bloomed and _ means flower did not bloom in the garden.
// We need 3 bouquets each should contain 1 flower.
// After day 1: [x, _, _, _, _]   // we can only make one bouquet.
// After day 2: [x, _, _, _, x]   // we can only make two bouquets.
// After day 3: [x, _, x, _, x]   // we can make 3 bouquets. The answer is 3.
func minDays(bloomDay []int, m int, k int) int {
	n := len(bloomDay)
	max := bloomDay[0]
	min := bloomDay[0]
	for i := 1; i < n; i++ {
		if max < bloomDay[i] {
			max = bloomDay[i]
		}
		if min > bloomDay[i] {
			min = bloomDay[i]
		}
	}

	low := min
	high := max
	ans := -1

	for low <= high {
		mid := (low + high) / 2
		nob := maxbouq(bloomDay, mid, k)
		fmt.Println("mid:", mid, "bouquets:", nob, "needed:", m)

		if nob >= m {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return ans
}

func maxbouq(bloom []int, day int, k int) int {
	n := len(bloom)
	bouquets := 0
	flowers := 0

	for i := 0; i < n; i++ {
		if bloom[i] <= day {
			flowers++
			if flowers == k {
				bouquets++
				flowers = 0
			}
		} else {
			flowers = 0
		}
	}
	return bouquets
}
