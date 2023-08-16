package main

import "fmt"

func main(){

	var tinggi float32 
	var sisi_a float32
	var sisi_b float32

	fmt.Print("Masukkan tinggi : ")
	fmt.Scan(&tinggi)
	fmt.Print("Masukkan sisi a : ")
	fmt.Scan(&sisi_a)
	fmt.Print("Masukkan sisi b : ")
	fmt.Scan(&sisi_b)

	fmt.Println("-------------------------------------------------")
	
	var luas float32 = 0.5 * (sisi_a + sisi_b) * tinggi
	fmt.Println("Luas trapesium : " , " " ,luas)
}