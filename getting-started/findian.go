package main

import "fmt"
import "strings"

func main() {
    var input string

    _, err := fmt.Scan(&input)

    if err == nil {
        input = strings.ToLower(input)
        if strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Contains(input, "a") {
            fmt.Println("Found!")
        } else {
            fmt.Println("Not Found!")
        }
    } else {
        fmt.Println(err)
    }
}
