package main

import (
    "fmt"
    "time"
)

func kelipatan(x int, result_chan chan int) {
    for i := 1; ; i++ {
        if i%x == 0 {
            result_chan <- i
            time.Sleep(3 * time.Second)
        }
    }
}

func main() {
    var angka int
    result_chan := make(chan int)


    fmt.Print("Masukkan Angka : ")
    fmt.Scan(&angka)
    fmt.Printf("\nKelipatan dari %d: ", angka)
    go kelipatan(angka, result_chan)

    for {
        select {
        case result := <-result_chan:
            fmt.Print(result, " ")
        case <-time.After(10 * time.Second):
            fmt.Println("\nWaktu habis.")
            return
        }
    }
}

/*
Note : 
select pada code akan digunaakn seperti switch case dimana apabila data dari channel result_chan
masih masuk kedalam result maka program akan terus berjalan namun apabila telah tidak ada nilai
lagi maka program akan menunggu selama 10 detik sebelum menutup program
*/
