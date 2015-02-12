package main

import "fmt"

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

func main() {
	unorderedArray := []int{2, 5, 0, 3, 1, 12, 123, 13, 12312312, -1}
	fmt.Println(BubbleSort(unorderedArray))
}
