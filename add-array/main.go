package main

import (
	"fmt"
)

func Sum(n int) int {
	total := 0
	//From 1 to N, 1 + 2 + 3 + 4 + 5 +.. + n
	for i := 1; i <= n; i++ {
		total = total + i
		//log.Printf(" Index %d:: Total  %d ", i, total)
	}
	return total
}

func main() {
	fmt.Println(Sum(100))
	// go sumR(100)
	// time.Sleep(1 * time.Minute)
}

func SumR(n int) {
	total := 0
	for i := 1; i <= n; i++ {
		total = total + i
	}
}

func Add(n int) {
	go SumR(n)
}
