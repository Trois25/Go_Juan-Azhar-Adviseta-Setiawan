package main

import (
    "fmt"
    "time"
)

func kelipatan(x int, bil_chan chan int) {
    defer close(bil_chan) 
    for i := 1; ; i++ {
        if i%x == 0 {
            bil_chan <- i
            time.Sleep(3 * time.Second)
        }
    }
}

func main() {
    var angka int
	result_chan := make(chan int, 5)
    fmt.Print("Masukkan Angka : ")
    fmt.Scan(&angka) 

    fmt.Print("\nKelipatan dari ", angka ," : ")
    go kelipatan(angka, result_chan)

    for result := range result_chan {
        fmt.Print(result, " ")
    }
}

/*
Note : 
result_chan := make(chan int, 5)
pada channel diatasa diterapkan buffered channel untuk goroutine yakni dimana
pada code diatas hanya diizinkan 5 goroutine secara bersamaan kedalam variabel result_chan
dan selama buffer tidak penuh maka akan dijalankan secara synchronous
*/
