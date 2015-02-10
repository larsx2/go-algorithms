package main

import "fmt"

// Returns the array index where target is located
func LinearSearch(array []int, target int) int {

	for idx, el := range array {
		if el == target {
			return idx
		}
	}

	return -1
}

func main() {
	array := []int{2, 3, 1, 8, 9, 10, 23, 21, 65, 19}
	target := 23

	fmt.Println("Array is", array, "target is", target)
	result := LinearSearch(array, target)

	if result == -1 {
		fmt.Println("Not Found")
	} else {
		fmt.Println("Found at index:", result)
	}
}
