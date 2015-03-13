package main

import "fmt"

func BinarySearch(items []int, target int) int {
	first, last := 0, len(items)-1

	for first <= last {
		mid := (first + last) / 2
		val := items[mid]

		switch {
		case target == val:
			return mid

		case target < val:
			last = mid - 1

		case target > val:
			first = mid + 1
		}
	}

	return -1
}

func main() {
	array := []int{1, 3, 5, 7, 12, 15, 24, 50}
	target := 5

	fmt.Println(">> Looking for", target, "in", array)
	result := BinarySearch(array, target)

	if result == -1 {
		fmt.Println("[-] Not Found")
	} else {
		fmt.Println("[+] Found at index:", result)
	}

}
