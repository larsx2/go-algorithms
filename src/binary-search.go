package main

import (
	"fmt"
	"os"
)

func BinarySearch(items []int, target int) int {
	first, last := 0, len(items)-1

	for first <= last {
		middle := (first + last) / 2
		currValue := items[middle]

		switch {
		case target == currValue:
			return middle

		case target < currValue:
			last = middle - 1

		case target > currValue:
			first = middle + 1
		}
	}

	return -1
}

func main() {
	array := []int{1, 3, 5, 7, 12, 15, 24, 50}
	target := 5

	fmt.Println("Array is", array, "target is", target)
	result := BinarySearch(array, target)

	if result == -1 {
		fmt.Println("Not Found")
		os.Exit(1)
	}

	fmt.Println("Found at index:", result)
}
