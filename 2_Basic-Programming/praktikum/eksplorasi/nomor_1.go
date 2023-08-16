package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan kata : ")
	kalimat,_:=reader.ReadString('\n')

	palindorm := true
	
	panjang_kalimat := len(kalimat)

	for i := 0; i < panjang_kalimat; i++ {
		if kalimat[i] != kalimat[panjang_kalimat-i-1]{
			palindorm = true
		}else{		
			palindorm = false
		}
	}

	fmt.Print("Captured: " + kalimat )
	if palindorm==true {
		fmt.Println("Palindorm")
	}else{
		fmt.Println("Bukan Palindorm")
	}
	
}