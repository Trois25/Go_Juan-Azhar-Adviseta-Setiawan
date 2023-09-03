package main

import (
    "fmt"
    "time"
)

func kelipatan (x int){
    for i := 1; ; i++ {
        if i%x == 0 {
            fmt.Print(i, " ")
            time.Sleep(3 * time.Second)
        }
        
    }
}

func main() {
    var angka int
    fmt.Print("Masukkan Angka : ")
    fmt.Scan(&angka)
    fmt.Print("\nKelipatan dari ", angka,": ")
    go kelipatan(angka)

    select {}
}
/* 
note:
select digunakan untuk memanage beberapa goroutine dan menunggu program untuk dimatikan
*/
