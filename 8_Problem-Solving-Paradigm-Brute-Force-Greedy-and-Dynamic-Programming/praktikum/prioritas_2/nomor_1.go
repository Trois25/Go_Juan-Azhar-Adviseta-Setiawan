package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	n := len(jumps)
	if n <= 2 {
		return 0
	}

	mincost := make([]int, n)
	mincost[0] = 0
	mincost[1] = int(math.Abs(float64(jumps[1] - jumps[0])))

	for i := 2; i < n; i++ {
		nilai1 := mincost[i-1] + int(math.Abs(float64(jumps[i] - jumps[i-1])))
		nilai2 := mincost[i-2] + int(math.Abs(float64(jumps[i] - jumps[i-2])))
		mincost[i] = int(math.Min(float64(nilai1), float64(nilai2)))
	}

	return mincost[n-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20})) // Output: 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // Output: 40
}
