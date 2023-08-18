package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan kata : ")
	kalimat,_:=reader.ReadString('\n')

	kalimat = strings.TrimSpace(kalimat)

	palindorm := true
	
	panjang_kalimat := len(kalimat)

	for i := 0; i < panjang_kalimat/2; i++ {
		if kalimat[i] != kalimat[panjang_kalimat-i-1]{
			palindorm = false
			break
		}
	}

	fmt.Println("Captured: " + kalimat )
	if palindorm == true {
		fmt.Println("Palindorm")
	}else{
		fmt.Println("Bukan Palindorm")
	}
	
}