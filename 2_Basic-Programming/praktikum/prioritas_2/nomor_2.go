package main

import "fmt"

func main() {

	var angka int

	fmt.Print("Masukkan angka : ")
	fmt.Scan(&angka)

	fmt.Println("Faktor dari ", angka, " adalah: ")
	for i := 1; i <= angka; i++ {
		if angka%i == 0 {
			fmt.Println(i)
		}
	}
}