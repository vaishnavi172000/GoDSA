package main

import "fmt"

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := range n - 1 {
		isSwapped := false
		for j := range n - 1 - i {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSwapped = true
			}
		}
		if !isSwapped {
			break
		}
	}
	return arr
}

func insertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j = i - 1; j >= 0 && arr[j] > key; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = key

	}
	return arr

}

func RecursiveBubbleSort(arr []int, n int) {
	if n > 0 {
		for i := range n - 1 {
			if arr[i] > arr[i+1] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
			}
		}
		RecursiveBubbleSort(arr, n-1)
	}

}

// RecursiveInsertionSort sorts the array using insertion sort recursively
// It takes two parameters array and n (length of the array)
// It sorts the array in ascending order
func RecursiveInsertionSort(arr []int, n int) {
	if n > 2 {
		RecursiveInsertionSort(arr, n-1)

	}
	key := arr[n-1]
	i := n - 2
	for i = n - 2; i >= 0 && arr[i] > key; i-- {
		fmt.Println(i)
		arr[i+1] = arr[i]
	}
	arr[i+1] = key

}

func main() {
	var n int
	var e int
	fmt.Println("Enter the number of elements")
	fmt.Scan(&n)

	mySlice := make([]int, n)

	fmt.Println("Enter the elements")
	for i := range n {
		fmt.Println("Enter ", i+1, "Element value")
		fmt.Scan(&mySlice[i])
	}

	fmt.Println("Enter your choice: \n1. Bubble Sort \n2. Insertion Sort \n3. Recursive Bubble Sort\n4. Recursive Insertion Sort")
	fmt.Scan(&e)
	switch e {
	case 1:
		bubbleSort(mySlice)
		fmt.Println("Sorted Using Bubble sort")
		fmt.Println(bubbleSort(mySlice))
	case 2:
		insertionSort(mySlice)
		fmt.Println("Sorted Using Insertion Sort")
		fmt.Println(insertionSort(mySlice))
	case 3:
		RecursiveBubbleSort(mySlice, n)
		fmt.Println("Sorted Using Recursive Bubble Sort")
		RecursiveBubbleSort(mySlice, n)
		fmt.Println(mySlice)
	case 4:
		RecursiveInsertionSort(mySlice, n)
		fmt.Println("Sorted Using Recursive Insertion Sort")
		RecursiveInsertionSort(mySlice, n)
		fmt.Println(mySlice)
	default:
		fmt.Println("Invalid Choice")

	}

}
