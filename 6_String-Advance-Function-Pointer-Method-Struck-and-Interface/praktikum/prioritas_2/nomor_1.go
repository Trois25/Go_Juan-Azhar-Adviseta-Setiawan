package main


import (
	"fmt"
)


func caesar(offset int, input string) string {

// your code here
var result string
for _, char := range input {
	if char >= 'a' && char <= 'z' {
		char = 'a' + (char-'a'+rune(offset))%26
	}
	result += string(char)
}

return result

}


func main() {

fmt.Println(caesar(3, "abc")) // def

fmt.Println(caesar(2, "alta")) // def

fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi

fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza

fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl

}