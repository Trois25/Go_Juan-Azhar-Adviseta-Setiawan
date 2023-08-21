package main

import "fmt"

func pow (x,n int) int {
	if n == 0 {
		return 1
	}

	temp := pow(x, n/2)
	if n%2 == 0 {
		return temp * temp
	} else {
		if n > 0 {
			return x * temp * temp
		} else {
			return (temp * temp) / x
		}
	}
}

func main() {

	var x,n int

	fmt.Print("nilai x: ")
	fmt.Scan(&x)

	fmt.Print("nilai n: ")
	fmt.Scan(&n)

	fmt.Print("Output : " , pow(x, n))
}

