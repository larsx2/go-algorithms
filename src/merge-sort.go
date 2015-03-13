package main

import "fmt"

func MergeSort(l []int) []int {
	if len(l) <= 1 {
		return l
	}

	mid := len(l) / 2
	a := MergeSort(l[:mid])
	b := MergeSort(l[mid:])

	return Merge(a, b)
}

func Merge(a []int, b []int) []int {
	c := []int{}

	i, j := 0, 0
	for i <= len(a) && j <= len(b) {
		if i == len(a) {
			return append(c, b[j:]...)
		}

		if j == len(b) {
			return append(c, a[i:]...)
		}

		if a[i] < b[j] {
			c = append(c, a[i])
			i += 1
		} else {
			c = append(c, b[j])
			j += 1
		}
	}

	return c
}

func main() {
	unorderedArray := []int{2, 5, 0, 3, 1, 12, 123, 13, 12312312, -1}
	fmt.Println(MergeSort(unorderedArray))
}
