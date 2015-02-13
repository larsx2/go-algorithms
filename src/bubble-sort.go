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
	for i := 0; i < len(items)-1; i++ {
		for j := 0; j < len(items)-1; j++ {
			if items[j+1] < items[j] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}

	return items
}

func main() {
	unorderedArray := []int{2, 5, 0, 3, 1, 12, 123, 13, 12312312, -1}
	unorderedArray2 := []int{3, 1, 4, 2, 12, 32, 56, 100}
	fmt.Println(BubbleSort(unorderedArray))
	fmt.Println(BubbleSort2(unorderedArray2))
}
