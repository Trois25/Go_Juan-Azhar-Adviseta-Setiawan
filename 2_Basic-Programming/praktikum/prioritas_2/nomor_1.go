package main

import "fmt"

func main() {
	var tinggi int

	fmt.Print("Masukkan tinggi segitiga : ")
	fmt.Scan(&tinggi)

	for i := 0; i <= tinggi; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

}