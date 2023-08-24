package main


import (

"fmt"

)


func munculSekali(angka string) []int {

// your code here
count := make(map[int32]int)
var newArray []int

for _, baris := range angka {
	count[baris]++
	fmt.Println(count)
}

for _, baris := range angka {
	if count[baris] == 1 {
		newArray = append(newArray, int(baris-'0'))
	}
}

return newArray

}


func main() {

// Test cases

fmt.Println(munculSekali("1234123")) // [4]

fmt.Println(munculSekali("76523752")) // [6 3]

fmt.Println(munculSekali("12345")) // [1 2 3 4 5]

fmt.Println(munculSekali("1122334455")) // []

fmt.Println(munculSekali("0872504")) // [8 7 2 5 4]

}