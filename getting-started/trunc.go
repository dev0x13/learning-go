package main

import "fmt"

func main() {
    var value float32

    for {
        _, err := fmt.Scan(&value)

        if err == nil {
            fmt.Println(int(value))
        } else {
            fmt.Println(err)
        }
    }
}
