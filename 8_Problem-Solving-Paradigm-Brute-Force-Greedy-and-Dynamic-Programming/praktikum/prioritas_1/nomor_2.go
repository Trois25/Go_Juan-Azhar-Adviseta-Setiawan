package main

import "fmt"

func segitigaPascal(rows int) [][]int {
    if rows <= 0 {
        return nil
    }	

    array := make([][]int, rows)
    for i := 0; i < rows; i++ {
        array[i] = make([]int, i+1)
        array[i][0] = 1
        array[i][i] = 1

        for j := 1; j < i; j++ {
            array[i][j] = array[i-1][j-1] + array[i-1][j]
        }
    }

    return array
}

func main() {
	var result [][]int
    var many int

	fmt.Print("NumRows = ")
	fmt.Scan(&many)
    pascal := segitigaPascal(many)


    for _, row := range pascal {
		result = append(result, row)
    }

	fmt.Print(result)
}
