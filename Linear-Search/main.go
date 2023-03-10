package main

import "fmt"

func linearSort(arr []int, s int) int {
	for i, v := range arr {
		if s == v {
			return i
		}
	}

	return -1
}

func main() {
	var n = []int{9, 1, 33, 21, 78, 42, 4}

	fmt.Println("Index of :: ", linearSort(n, 78)) // 4
	index := linearSort(n, 78)
	fmt.Printf("Value of %d index is :: %d", index, n[index])
	fmt.Println()
}
