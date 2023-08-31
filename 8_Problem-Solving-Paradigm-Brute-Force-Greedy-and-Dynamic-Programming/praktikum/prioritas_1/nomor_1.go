package main

import (
	"fmt"
	"strconv"
)

func arrayn(n int) []string{
	ans := make([]string, n+1)
	for i := 0; i <= n; i++ {
		biner := strconv.FormatInt(int64(i), 2)
		ans[i] = biner
	}
	return ans
}



func main() {
	var n int
	fmt.Print("Input : ")
	fmt.Scan(&n)

	fmt.Print("Output : " , arrayn(n))
}