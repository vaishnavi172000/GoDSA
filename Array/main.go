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

func rotateRight(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])

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
