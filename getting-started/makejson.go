package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    var name string
    var address string

    fmt.Println("Enter your name:")
    _, err := fmt.Scanln(&name)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Enter your address:")
    _, err = fmt.Scanln(&address)
    if err != nil {
        fmt.Println(err)
        return
    }

    dataMap := map[string]string{
        "name": name,
        "address": address,
    }

    // Here we don't handle the error since it could not actually happen
    // in this particular case
    dataJson, _ := json.Marshal(dataMap)

    fmt.Println(string(dataJson))
}
