package main

import "fmt"

func main() {
	var nilai float32
	fmt.Print("Masukkan nilai : ")
	fmt.Scan(&nilai)
	fmt.Println("-------------------------------------------------")
	if nilai >= 80 && nilai <=100{
		fmt.Println("Nilai : A")
	}else if nilai >= 65 && nilai <=79{
		fmt.Println("Nilai : B")
	}else if nilai >= 50 && nilai <=64{
		fmt.Println("Nilai : C")	
	}else if nilai >= 35 && nilai <=49{
		fmt.Println("Nilai : D")
	}else if nilai >= 0 && nilai <=34{
		fmt.Println("Nilai : E")
	}else{
		fmt.Println("Nilai Invalid")
	}
}