package main

import (
	"fmt"
	"strings"
)


func Compare(a, b string) string {

// your code here
var check = strings.Contains(a,b)
var result string
if check {
	result = b
}else{
	result = a
}
return result
}


func main() {

fmt.Println(Compare("AKA", "AKASHI")) //AKA

fmt.Println(Compare("KANGOORO", "KANG")) //KANG

fmt.Println(Compare("KI", "KIJANG")) //KI

fmt.Println(Compare("KUPU-KUPU", "KUPU")) //KUPU

fmt.Println(Compare("ILALANG", "ILA")) //ILA

}

