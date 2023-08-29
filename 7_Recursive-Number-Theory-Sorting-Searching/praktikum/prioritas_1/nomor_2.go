package main

import (
	"fmt"
	"sort"
)


type pair struct {

name string

count int

}


func MostAppearItem(items []string) []pair {

// your code here
ascending := make(map[string]int)
var data []pair

for i := 0; i < len(items); i++ {
	ascending[items[i]]++
}

for name, count := range ascending {
		data = append(data, pair{name, count})
	}

sort.SliceStable(data, func(i, j int) bool {
	return data[i].count < data[j].count
})

return data

}


func main() {

pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}) // golang->1 ruby->2 js->4

for _, list := range pairs {

fmt.Print(list.name, " -> ", list.count, " ")

}

fmt.Println()


pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}) // C->1 D->2 B->3 A->4

for _, list := range pairs {

fmt.Print(list.name, " -> ", list.count, " ")

}

fmt.Println()


pairs = MostAppearItem([]string{"football", "basketball", "tenis"}) // football->1 basketball->1 tenis->1

for _, list := range pairs {

fmt.Print(list.name, " -> ", list.count, " ")

}

}