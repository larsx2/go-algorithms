package main

import "fmt"

func SelectionSort(items []int) []int {

	for i := 0; i < len(items); i++ {
		min := i

		for j := i + 1; j < len(items); j++ {
			if items[j] < items[min] {
				min = j
			}
		}

		items[i], items[min] = items[min], items[i]
	}

	return items
}

func main() {
	unorderedArray := []int{2, 5, 0, 3, 1, 12, 123, 13, 12312312, -1}
	fmt.Println("Sorted Array:  ", SelectionSort(unorderedArray))
}
