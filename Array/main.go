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

}
