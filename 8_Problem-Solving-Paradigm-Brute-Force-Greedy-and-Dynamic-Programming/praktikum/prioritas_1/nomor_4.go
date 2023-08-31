package main

import (
	"fmt"
)

func SimpleEquations(a, b, c int) {
	check := false
	max := a
	min := a
	if max < b || min > b {
		max = b
		min = b
	}else{
		max = c
		min = c
	}

	if max >= 100000 || min <= 1{
		fmt.Println("Number out of range")
	}else{
		for x := 1; x <= max+1 && !check; x++ {
			for y := 1; y <= max+1 && !check; y++ {
				for z := 1; z <= max+1 && !check; z++ {
					if x != y && x != z && y != z && x+y+z == a && x*y*z == b && x*x+y*y+z*z == c {
						fmt.Print(x, y, z)
						check = true
					}
				}
			}
		}
	}

	if !check {
		fmt.Println("no solution")
	}
}

func main() {
	SimpleEquations(1, 2, 3) //no solution
	SimpleEquations(6, 6, 14)  // 1 2 3
}
