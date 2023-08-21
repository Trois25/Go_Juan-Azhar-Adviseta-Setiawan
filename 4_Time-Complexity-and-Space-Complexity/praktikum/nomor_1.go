package main

import "fmt"

func primeNumber(number int) bool{

	if number <= 1{
		return false
	}
	if number <= 3{
		return true
	}
	if number%2 ==0 || number%3 ==0{
		return false
	}
	for i := 5; i*i <= number; i+=6 {
		if number%i == 0 || number%(i+2)==0{
			return false
		}
	}

	return true
}

func main() {
	var angka int

	fmt.Print("Masukkan angka : ")
	fmt.Scan(&angka)

	primeNumber(angka)

	if primeNumber(angka) {
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
}