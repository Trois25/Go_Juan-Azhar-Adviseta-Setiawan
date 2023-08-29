package main

import "fmt"

func primeX(number int) int {

	// your code here
	var primeNumber []int
	for i := 2; i <= 100; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
			}
		}
		if isPrime == true {
			primeNumber = append(primeNumber, i)
		}
	}

	return primeNumber[number-1]
	
}

func main() {

	fmt.Println(primeX(1)) // 2

	fmt.Println(primeX(5)) // 11

	fmt.Println(primeX(8)) // 19

	fmt.Println(primeX(9)) // 23

	fmt.Println(primeX(10)) // 29

}
