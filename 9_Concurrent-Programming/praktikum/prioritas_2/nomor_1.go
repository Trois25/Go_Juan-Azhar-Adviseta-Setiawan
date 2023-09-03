package main

import (
	"fmt"
	"sync"
)

func main() {
	var text string
	fmt.Print("Masukkan text : ")
	fmt.Scan(&text)
	banyakHuruf := make(map[rune]int)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, char := range text {
		wg.Add(1)
		go func(c rune) {
			defer wg.Done()
			mu.Lock()
			banyakHuruf[c]++
			mu.Unlock()
		}(char)
	}

	wg.Wait()

	fmt.Println("Frekuensi Huruf:")
	for huruf, banyak := range banyakHuruf {
		fmt.Printf("%c: %d\n", huruf, banyak)
	}
}
/*
Note
digunakan waitgroup untuk menunggu koordinasi goroutine dengan menambahkan 1 wg untuk menghitung goroutine
digunakan mutex untuk menghindari race condition pada saat mengakses banyakHuruf
wg.Wait digunakan untuk memastikan segala yang terdapat di wg dieksekusi sebelum mematikan program

*/