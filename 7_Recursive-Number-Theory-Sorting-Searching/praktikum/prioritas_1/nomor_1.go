package main


import "fmt"


func fibonacci(number int) int {
	
	var count int
	awal := 0
	akhir :=1
	if number == 0 {
		return awal
	}
	for i := 0; i <= number; i++ {
		if(i == 1 ){
			count = awal
			continue
		}
		if (i == 2){
			count = akhir
			continue
		}
		count = awal + akhir
		awal = akhir
		akhir = count
		
	}
	return count
}



func main() {

fmt.Println(fibonacci(0)) // 0

fmt.Println(fibonacci(2)) // 1

fmt.Println(fibonacci(9)) // 34

fmt.Println(fibonacci(10)) // 55

fmt.Println(fibonacci(12)) // 144

}