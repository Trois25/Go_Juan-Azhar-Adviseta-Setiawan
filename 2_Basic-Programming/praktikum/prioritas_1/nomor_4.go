package main

import "fmt"

func main() {
	total := 100

	for i := 1; i<=total; i++{
		if i%3 == 0{
			fmt.Println("Fizz")
		}else if i%5 == 0 {
			fmt.Println("Buzz")
		}else{
			fmt.Println(i)
		}
	}
}