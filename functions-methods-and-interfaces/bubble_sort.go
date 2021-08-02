package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func Swap(slice []int, idx int) {
    slice[idx], slice[idx + 1] = slice[idx + 1], slice[idx]
}

func BubbleSort(slice []int) {
    for i := 0; i < len(slice) - 1; i++ {
        for j := 0; j < len(slice) - i - 1; j++ {
            if slice[j] > slice[j + 1] {
                Swap(slice, j)
            }
        }
    }
}

func main() {
    var maxIntegers int = 10

    fmt.Printf("Please type a sequence of up to %d integers separated by space\n", maxIntegers)

    in := bufio.NewReader(os.Stdin)
    stringInput, err := in.ReadString('\n')

    if err != nil {
        fmt.Println("Error while consuming input: " + err.Error())
        return
    }

    stringInput = strings.TrimSuffix(stringInput, "\n")

    tmp := strings.Split(stringInput, " ")

    if len(tmp) > maxIntegers {
        fmt.Printf("Expected up to %d integers, got %d\n", maxIntegers, len(tmp))
        return
    }

    slice := make([]int, len(tmp))

    for i, raw := range tmp {
        slice[i], err = strconv.Atoi(raw)

        if err != nil {
            fmt.Println("Error while parsing input: " + err.Error())
            return
        }
    }

    BubbleSort(slice)

    fmt.Print("Sorted sequence: ")
    fmt.Println(slice)
}
