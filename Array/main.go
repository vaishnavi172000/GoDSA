package main

import (
	"fmt"
	"math"
)

func FindSecondlargest(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return -1
	}
	max := math.MinInt
	secondMax := math.MinInt
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			secondMax = max
			max = arr[i]
		}
	}

	return secondMax
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
// Given an integer array nums sorted in non-decreasing order,
// remove the duplicates in-place such that each unique element appears only once.
// The relative order of the elements should be kept the same. Then return the number of unique elements in nums.

// Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:

// Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially.
// The remaining elements of nums are not important as well as the size of nums.
// Return k.

func removeDuplicates(arr []int) int {
	n := len(arr)
	cnt := 0
	j := 0
	for i := 0; i < n; i++ {
		arr[j] = arr[i]
		for i != n-1 && arr[i] == arr[i+1] {
			i++
		}
		j++
		cnt++
	}
	return cnt

}

func reverse(arr []int) {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}

}

// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]
func rotateRight(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])

}

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and
// choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
// Example 1:
// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
// Example 2:

// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transactions are done and the max profit = 0.
func maxProfit(arr []int) int {
	l := len(arr)
	if l <= 1 {
		return 0
	}
	min := arr[0]
	profit := 0
	for i := range arr {
		if profit < arr[i]-min {
			profit = arr[i] - min
		}
		if min > arr[i] {
			min = arr[i]
		}
	}
	return profit
}

// https://leetcode.com/problems/two-sum/description/
// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution,
// and you may not use the same element twice.
// You can return the answer in any order.
// Example 1:

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:

// Input: nums = [3,3], target = 6
// Output: [0,1]
func twoSum(nums []int, target int) []int {
	var result []int
	mp := make(map[int]int)
	for i, num := range nums {
		if val, ok := mp[target-num]; !ok {
			mp[num] = i
		} else {
			result = append(result, []int{i, val}...)
			return result
		}

	}
	return result

}

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/submissions/1617122398/
// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order,
// find two numbers such that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where 1 <= index1 < index2 <= numbers.length.
// Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
// The tests are generated such that there is exactly one solution. You may not use the same element twice.
// Your solution must use only constant extra space.
// Example 1:

// Input: numbers = [2,7,11,15], target = 9
// Output: [1,2]
// Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].
// Example 2:

// Input: numbers = [2,3,4], target = 6
// Output: [1,3]
// Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].
// Example 3:

// Input: numbers = [-1,0], target = -1
// Output: [1,2]
// Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].

// 2 sum return the elements
func twoSumElements(nums []int, target int) []int {
	var result []int
	//slices.Sort(nums)
	i := 0
	j := len(nums) - 1
	for i < j {
		sum := nums[i] + nums[j]
		if sum == target {
			result = append(result, []int{nums[i], nums[j]}...)
			return result
		}
		if sum > target {
			j--
		} else {
			i++
		}
	}
	return result

}

// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
// Given a roman numeral, convert it to an integer.

// Example 1:

// Input: s = "III"
// Output: 3
// Explanation: III = 3.
// Example 2:

// Input: s = "LVIII"
// Output: 58
// Explanation: L = 50, V= 5, III = 3.
// Example 3:

// Input: s = "MCMXCIV"
// Output: 1994
// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

func romanToInt(s string) int {
	m1 := make(map[string]int)
	m1["I"] = 1
	m1["V"] = 5
	m1["X"] = 10
	m1["L"] = 50
	m1["C"] = 100
	m1["D"] = 500
	m1["M"] = 1000
	l := len(s)
	sum := 0
	for i, v := range s {
		if i != l-1 && (m1[string(v)] < m1[string(s[i+1])]) {
			sum = sum - m1[string(v)]
		} else {
			sum += m1[string(v)]
		}
	}
	return sum

}

// 8
// https://leetcode.com/problems/move-zeroes/
// Given an integer array nums, move all 0's to the end of it
//  while maintaining the relative order of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.

// Example 1:

// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
// Example 2:

// Input: nums = [0]
// Output: [0]

func moveZeroes(nums []int) {
	i := 0
	n := len(nums)
	for j := 0; j < n; j++ {
		if nums[j] != 0 {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			i++
		}

	}

}

// 9
// rerturn the length of longest subraarya whose sum is k
func subarraySumlenLongest(nums []int, k int) int {
	preSumMap := make(map[int]int)
	n := len(nums)
	sum := nums[0]
	preSumMap[sum] = 0
	l := 0
	if sum == k {
		l = 1
	}

	for i := 1; i < n; i++ {
		sum += nums[i]
		if sum == k {
			l = max(l, i+1)
		}
		if sum > k {
			if v, ok := preSumMap[sum-k]; ok {
				if l < (i - v) {
					l = i - v
				}
			}
		}
		preSumMap[sum] = i + 1
	}
	return l

}

// 10 kanades algo
// [] array is considered that is max sum can be 0
func maxSubArray(nums []int) int {
	sum := 0
	maxE := math.MinInt
	for _, v := range nums {
		sum += v
		if sum < 0 {
			sum = 0
			continue
		}
		if sum > maxE {
			maxE = sum
		}

	}
	return maxE
}

// 11
// https://leetcode.com/problems/rotate-image/description/?envType=problem-list-v2&envId=array
// You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
// You have to rotate the image in-place, which means you have to modify the input 2D matrix directly.
// DO NOT allocate another 2D matrix and do the rotation.
// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [[7,4,1],[8,5,2],[9,6,3]]
func rotate90(matrix [][]int) {
	n := len(matrix)
	// transpose
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// reverse the rows
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}

	}

}

// 12
// https://leetcode.com/problems/longest-repeating-character-replacement/description/
// You are given a string s and an integer k. You can choose any character of the string and change it to any other uppercase English character.
//  You can perform this operation at most k times.

// Return the length of the longest substring containing the same letter you can get after performing the above operations.

// Example 1:

// Input: s = "ABAB", k = 2
// Output: 4
// Explanation: Replace the two 'A's with two 'B's or vice versa.
// Example 2:

// Input: s = "AABABBA", k = 1
// Output: 4
// Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
// The substring "BBBB" has the longest repeating letters, which is 4.
// There may exists other ways to achieve this answer too.
func characterReplacement(s string, k int) int {
	n := len(s)
	l := 0
	r := 0
	maxlen := 0
	maxf := 0
	mp := make(map[rune]int)
	for r < n {
		mp[rune(s[r])]++
		maxf = max(maxf, mp[rune(s[r])])
		if r-l+1-maxf > k {
			mp[rune(s[l])]--
			if mp[rune(s[l])] == 0 {
				delete(mp, rune(s[l]))
			}
			l++
		}
		maxlen = max(maxlen, r-l+1)
		r++
	}
	return maxlen

}

// 13
// https://leetcode.com/problems/count-number-of-nice-subarrays/description/
// Given a binary array nums and an integer goal, return the number of non-empty subarrays with a sum goal.

// A subarray is a contiguous part of the array.

// Example 1:

// Input: nums = [1,0,1,0,1], goal = 2
// Output: 4
// Explanation: The 4 subarrays are bolded and underlined below:
// [1,0,1,0,1]
// [1,0,1,0,1]
// [1,0,1,0,1]
// [1,0,1,0,1]
// Example 2:

// Input: nums = [0,0,0,0,0], goal = 0
// Output: 15
func numSubarraysWithSum(nums []int, goal int) int {
	return numSubArrSumLessThanEqual(nums, goal) - numSubArrSumLessThanEqual(nums, goal-1)

}

func numSubArrSumLessThanEqual(nums []int, goal int) int {
	if goal < 0 {
		return 0
	}
	cnt := 0
	l := 0
	r := 0
	n := len(nums)
	sum := 0

	for r < n {
		if nums[r] == 1 {
			sum++
		}
		for sum > goal {
			sum -= nums[l]
			l++
		}
		if sum <= goal {
			cnt += r - l + 1
		}
		//fmt.Println(cnt, l, r);
		r++

	}
	fmt.Println(cnt)
	return cnt
}

// 14
// https://leetcode.com/problems/binary-subarrays-with-sum/description/
// Given an array of integers nums and an integer k. A continuous subarray is called nice if there are k odd numbers on it.

// Return the number of nice sub-arrays.
// Example 1:

// Input: nums = [1,1,2,1,1], k = 3
// Output: 2
// Explanation: The only sub-arrays with 3 odd numbers are [1,1,2,1] and [1,2,1,1].
// Example 2:

// Input: nums = [2,4,6], k = 1
// Output: 0
// Explanation: There are no odd numbers in the array.
// Example 3:

// Input: nums = [2,2,2,1,2,2,1,2,2,2], k = 2
// Output: 16

func numberOfSubarrays1(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	n := len(nums)
	oddCnt := 0
	cnt := 0
	r := 0
	l := 0
	for r < n {
		if nums[r]%2 == 1 {
			cnt++
		}
		if cnt > k {
			for cnt > k && l < r {
				if nums[l]%2 == 1 {
					cnt--
				}
				l++
			}
		}
		if cnt <= k {
			oddCnt += r - l + 1
		}
		r++
	}
	return oddCnt

}

func numberOfSubarrays(nums []int, k int) int {
	return numberOfSubarrays1(nums, k) - numberOfSubarrays1(nums, k-1)
}

// 15
// https://leetcode.com/problems/rearrange-array-elements-by-sign/description/
// You are given a 0-indexed integer array nums of even length consisting of an equal number of positive and negative integers.

// You should return the array of nums such that the the array follows the given conditions:

// Every consecutive pair of integers have opposite signs.
// For all integers with the same sign, the order in which they were present in nums is preserved.
// The rearranged array begins with a positive integer.
// Return the modified array after rearranging the elements to satisfy the aforementioned conditions.

// Example 1:

// Input: nums = [3,1,-2,-5,2,-4]
// Output: [3,-2,1,-5,2,-4]
// Explanation:
// The positive integers in nums are [3,1,2]. The negative integers are [-2,-5,-4].
// The only possible way to rearrange them such that they satisfy all conditions is [3,-2,1,-5,2,-4].
// Other ways such as [1,-2,2,-5,3,-4], [3,1,2,-2,-5,-4], [-2,3,-5,1,-4,2] are incorrect because they do not satisfy one or more conditions.
// Example 2:

// Input: nums = [-1,1]
// Output: [1,-1]
// Explanation:
// 1 is the only positive integer and -1 the only negative integer in nums.
// So nums is rearranged to [1,-1].

func rearrangeArray(nums []int) []int {
	postIndex := 0
	negIndex := 1
	n := len(nums)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		if nums[i] < 0 {
			result[negIndex] = nums[i]
			negIndex += 2
		} else {
			result[postIndex] = nums[i]
			postIndex += 2
		}

	}
	return result

}

// 16
// https://leetcode.com/problems/minimum-pair-removal-to-sort-array-i/description
// Given an array nums, you can perform the following operation any number of times:

// Select the adjacent pair with the minimum sum in nums.
// If multiple such pairs exist, choose the leftmost one.
// Replace the pair with their sum.
// Return the minimum number of operations needed to make the array non-decreasing.

// An array is said to be non-decreasing
// if each element is greater than or equal to its previous element (if it exists).
// Example 1:

// Input: nums = [5,2,3,1]

// Output: 2

// Explanation:

// The pair (3,1) has the minimum sum of 4. After replacement, nums = [5,2,4].
// The pair (2,4) has the minimum sum of 6. After replacement, nums = [5,6].
// The array nums became non-decreasing in two operations.

// Example 2:

// Input: nums = [1,2,2]

// Output: 0

// Explanation:

// The array nums is already sorted.
func mergeArray(nums []int) []int {
	sum := math.MaxInt
	idx := 0
	n := len(nums)
	for i := 0; i < n-1; i++ {
		sum1 := nums[i] + nums[i+1]
		if sum1 < sum {
			sum = sum1
			idx = i
		}
	}
	nums[idx] = sum
	for i := idx + 1; i < n-1; i++ {
		nums[i] = nums[i+1]
	}
	return nums[:n-1]
}
func isArraySorted(nums []int) bool {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}
func minimumPairRemoval(nums []int) int {
	cnt := 0
	for !isArraySorted(nums) {
		cnt++
		nums = mergeArray(nums)
	}
	return cnt

}

// Given an m x n integer matrix matrix,
// if an element is 0, set its entire row and column to 0's.

// You must do it in place.

// Example 1:

// Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
// Output: [[1,0,1],[0,0,0],[1,0,1]]
// Example 2:

// Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
func setZeroes(matrix [][]int) {
	colZ := 1
	n := len(matrix)
	m := len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				if j != 0 {
					matrix[0][j] = 0
				} else {
					colZ = 0
				}
			}

		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for j := 1; j < m; j++ {
			matrix[0][j] = 0
		}
	}

	if colZ == 0 {
		for i := 0; i < n; i++ {
			matrix[i][0] = 0
		}
	}

}

func main() {
	var n int
	fmt.Println("Enter the size of the array")
	fmt.Scan(&n)
	var arr []int
	for i := 0; i < n; i++ {
		var num int
		fmt.Println("Enter element at index ", i)
		fmt.Scan(&num)
		arr = append(arr, num)
	}
	fmt.Println("Second largest element is")
	fmt.Println(FindSecondlargest(arr))

	fmt.Println("Array after reversing")
	reverse(arr)
	fmt.Println(arr)

	fmt.Println("Array after rotatng right by 2")
	rotateRight(arr, 2)
	fmt.Println(arr)

}
