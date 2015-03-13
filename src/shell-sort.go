package main

import "fmt"

func ShellSort(l []int) []int {

	gaps := []int{701, 301, 132, 57, 23, 10, 4, 1}

	for _, gap := range gaps {
		for i := gap; i < len(l); i++ {
			tmp := l[i]

			j := i
			for ; j >= gap && l[j-gap] > tmp; j -= gap {
				l[j] = l[j-gap]
			}
			l[j] = tmp
		}
	}

	return l
}

func main() {
	unorderedArray := []int{2, 5, 0, 3, 1, 12, 123, 13, 12312312, -1}
	fmt.Println(ShellSort(unorderedArray))
}
