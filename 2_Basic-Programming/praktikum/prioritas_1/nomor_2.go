package main

import "fmt"

func main() {
	var angka int

	fmt.Print("Masukkan jumlah angka : ")
	fmt.Scan(&angka)

	fmt.Println("-------------------------------------------------")

	for i := 1; i <= angka; i++ {
		if i%2 == 0 {
			fmt.Println(i, " - ", "angka genap")
		} else {
			fmt.Println(i, " - ", "angka ganjil")
		}
	}

}
