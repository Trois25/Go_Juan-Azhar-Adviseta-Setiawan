package main

import (
    "fmt"
    "sync"
    "time"
)

var mu sync.Mutex

func kelipatan(x int) {
    for i := 1; ; i++ {
        if i%x == 0 {
            mu.Lock() 
            fmt.Print(i, " ")
            mu.Unlock() 
            time.Sleep(3 * time.Second)
        }
    }
}

func main() {
    var angka int
    fmt.Print("Masukkan Angka : ")
    fmt.Scan(&angka)
    fmt.Print("\nKelipatan dari ", angka, ": ")
    go kelipatan(angka)

    select {}
}

/*
Note :
mutex merupakan cara untuk mengatasi race condition Sehingga jika kita memiliki goroutine melakukan lock maka goroutine 
yang lain tidak akan bisa melakukan lock sampai goroutine tersebut melakukan unlock, maka baru goroutine yang lain bisa 
melakukan lock. 
*/