package main

import "fmt"
import "strconv"
import "sort"

func main() {
    slice := make([]int, 0, 3)
    var input string

    for {
        _, err := fmt.Scan(&input)

        if err == nil {
            if input == "X" {
                break;
            }

            i, err := strconv.Atoi(input)

            if err == nil {
                slice = append(slice, i)
                sort.Ints(slice)
                fmt.Println(slice)
            } else {
                fmt.Println(err)
            }
        } else {
            fmt.Println(err)
        }
    }
}
