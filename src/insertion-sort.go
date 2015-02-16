package main

import "fmt"

func InsertionSort(items []int) []int {
	fmt.Println(items)
	for i := 1; i < len(items); i++ {
		j := i

		for j > 0 && items[j-1] > items[j] {
			items[j], items[j-1] = items[j-1], items[j]
			j = j - 1
		}
	}

	return items
}

func main() {
	array := []int{5, 4, 3, 2, 1, -4}
	fmt.Println(InsertionSort(array))
}
