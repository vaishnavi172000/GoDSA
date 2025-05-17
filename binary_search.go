package godsa

import (
	"fmt"
	"slices"
)

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

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/?envType=problem-list-v2&envId=array
// Given an array of integers nums sorted in non-decreasing order,
// find the starting and ending position of a given target value.
// If target is not found in the array, return [-1, -1].
// You must write an algorithm with O(log n) runtime complexity.
// Example 1:

// Input: nums = [5,7,7,8,8,10], target = 8
// Output: [3,4]
// Example 2:

// Input: nums = [5,7,7,8,8,10], target = 6
// Output: [-1,-1]
// Example 3:

// Input: nums = [], target = 0
// Output: [-1,-1]
func searchRange(nums []int, target int) []int {
	fo := findFirstOccurance(nums, target)
	if fo == -1 {
		return []int{-1, -1}
	}
	lo := findLastOccurance(nums, target)
	return []int{fo, lo}

}

func findFirstOccurance(nums []int, target int) int {
	ans := -1
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			ans = mid
			high = mid - 1
			continue
		}

		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}
	return ans
}

func findLastOccurance(nums []int, target int) int {
	ans := -1
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			ans = mid
			low = mid + 1
			continue
		}

		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}
	return ans
}

// Given two integer arrays nums1 and nums2, return an array of their intersection.
// Each element in the result must be unique and you may return the result in any order.

// Example 1:

// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2]
// Example 2:

// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [9,4]
// Explanation: [4,9] is also accepted.

func intersection(nums1 []int, nums2 []int) []int {
	n1 := len(nums1)
	n2 := len(nums2) // smaller array
	var result []int
	slices.Sort(nums1)
	slices.Sort(nums2)
	i := 0
	j := 0
	for i < n1 && j < n2 {
		key1 := nums1[i]
		key2 := nums2[j]
		if nums1[i] == nums2[j] {
			result = append(result, nums2[j])
			for i != n1-1 && nums1[i+1] == key1 {
				i++
			}
			i++
			for j != n2-1 && nums2[j+1] == key2 {
				j++
			}
			j++

		} else {
			if nums1[i] > nums2[j] {
				for j != n2-1 && nums2[j+1] == key2 {
					j++
				}
				j++
			} else {
				for i != n1-1 && nums1[i+1] == key1 {
					i++
				}
				i++

			}
		}

	}
	return result

}

// https://leetcode.com/problems/valid-perfect-square/description/?envType=problem-list-v2&envId=binary-search
// Given a positive integer num, return true if num is a perfect square or false otherwise.

// A perfect square is an integer that is the square of an integer. In other words, it is the product of some integer with itself.

// You must not use any built-in library function, such as sqrt.

// Example 1:

// Input: num = 16
// Output: true
// Explanation: We return true because 4 * 4 = 16 and 4 is an integer.
// Example 2:

// Input: num = 14
// Output: false
// Explanation: We return false because 3.742 * 3.742 = 14 and 3.742 is not an integer.
func isPerfectSquare(num int) bool {
	low := 0
	high := num
	for low <= high {
		mid := low + (high-low)/2
		if mid*mid == num {
			return true
		}
		if mid*mid > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false

}

func findMax(mat [][]int, n int, m int, mid int) int {
	maxI := mat[0][mid]
	inx := 0
	for i := 1; i < n; i++ {
		if maxI < mat[i][mid] {
			maxI = mat[i][mid]
			inx = i
		}
	}
	return inx

}

// https://leetcode.com/problems/find-a-peak-element-ii/description/
// A peak element in a 2D grid is an element that is strictly
// greater than all of its adjacent neighbors to the left, right, top, and bottom.

// Given a 0-indexed m x n matrix mat where no two adjacent cells are equal,
//  find any peak element mat[i][j] and return the length 2 array [i,j].

// You may assume that the entire matrix is surrounded by an outer perimeter with the value -1 in each cell.

// You must write an algorithm that runs in O(m log(n)) or O(n log(m)) time.
// Input: mat = [[1,4],[3,2]]
// Output: [0,1]
// Explanation: Both 3 and 4 are peak elements so [1,0] and [0,1] are both acceptable answers.

func findPeakGrid(mat [][]int) []int {
	n := len(mat)
	m := len(mat[0])
	low := 0
	high := m - 1
	for low <= high {
		mid := low + (high-low)/2
		maxRowIndex := findMax(mat, n, m, mid)
		left := -1
		right := -1
		if mid > 0 {
			left = mat[maxRowIndex][mid-1]
		}

		if mid < m-1 {
			right = mat[maxRowIndex][mid+1]
		}

		if mat[maxRowIndex][mid] > left && mat[maxRowIndex][mid] > right {
			return []int{maxRowIndex, mid}
		}
		if mat[maxRowIndex][mid] < left {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	return []int{-1, -1}

}

// https://leetcode.com/problems/arranging-coins/description/?envType=problem-list-v2&envId=binary-search
// You have n coins and you want to build a staircase with these coins. The staircase consists of k rows where the ith row has exactly i coins. The last row of the staircase may be incomplete.

// Given the integer n, return the number of complete rows of the staircase you will build.

// Example 1:

// Input: n = 5
// Output: 2
// Explanation: Because the 3rd row is incomplete, we return 2.
// Example 2:

// Input: n = 8
// Output: 3
// Explanation: Because the 4th row is incomplete, we return 3.
func arrangeCoins(n int) int {
	low := int64(0)
	high := int64(n)
	for low <= high {
		mid := low + (high-low)/2
		n1 := int(mid*(mid+1)) / 2
		if n1 == n {
			return int(mid)
		}
		if n1 > n {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return int(high)

}

// https://leetcode.com/problems/search-in-rotated-sorted-array/?envType=problem-list-v2&envId=binary-search
// There is an integer array nums sorted in ascending order (with distinct values).
// Prior to being passed to your function,
// nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length)
// such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
//
//	For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].
//
// Given the array nums after the possible rotation and an integer target,
// return the index of target if it is in nums, or -1 if it is not in nums.
// You must write an algorithm with O(log n) runtime complexity.
func searchInRotatedArray(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[low] {
			if nums[mid] >= target && nums[low] <= target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] <= target && nums[high] >= target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1

}

//There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).

// Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length)
// such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).
// For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].

// Given the array nums after the rotation and an integer target,
//  return true if target is in nums, or false if it is not in nums.

// You must decrease the overall operation steps as much as possible.

// Example 1:

// Input: nums = [2,5,6,0,0,1,2], target = 0
// Output: true
// Example 2:

// Input: nums = [2,5,6,0,0,1,2], target = 3
// Output: false

func searchInRotatedArrayWithDuplicates(arr []int, target int) bool {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return true
		}
		if arr[low] == arr[mid] && arr[mid] == arr[high] {
			low++
			high--
			continue
		}
		if arr[low] <= arr[mid] {
			if arr[low] <= target && arr[mid] >= target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if arr[high] >= target && arr[mid] <= target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return false

}
