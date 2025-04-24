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
