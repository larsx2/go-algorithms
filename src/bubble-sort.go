package main

import "fmt"

// using a 'while true' approach
func BubbleSort(items []int) []int {
	swapped := true

	for swapped {
		swapped = false

		for idx := 0; idx < len(items)-1; idx++ {
			if items[idx] > items[idx+1] {
				items[idx], items[idx+1] = items[idx+1], items[idx]
				swapped = true
			}
		}
	}

	return items
}

// cleaner version
func BubbleSort2(items []int) []int {

	for i, j := 0, 1; i < len(items)-1; i, j = i+1, j+1 {
		if items[i] > items[j] {
			items[i], items[j] = items[j], items[i]
		}
	}

	return items
}

func main() {
	unorderedArray := []int{2, 5, 0, 3, 1, 12, 123, 13, 12312312, -1}
	fmt.Println(BubbleSort(unorderedArray))
	fmt.Println(BubbleSort2(unorderedArray))
}
