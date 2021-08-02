package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Name struct {
    fname string
    lname string
}

func main() {
    var filePath string

    // 1. Read user input

    fmt.Println("Enter input file path:")
    _, err := fmt.Scanln(&filePath)
    if err != nil {
        fmt.Println(err)
        return
    }

    // 2. Read input file

    // 2.1. Define slice for names

    names := make([]Name, 1)

    // 2.2. Open file

    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println(err)
        return
    }

    // 2.3. Read lines, add entries to `names`

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lineStr := scanner.Text()
        splitLine := strings.Split(lineStr, " ")
        if len(splitLine) != 2 {
            fmt.Println("Invalid line: " + lineStr)
        }
        names = append(names, Name{fname: splitLine[0], lname: splitLine[1]})
    }

    // 2.4. Handle reading errors

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
        err := file.Close()
        if err != nil {
            fmt.Println(err)
        }
        return
    }

    // 2.5. Close file

    err = file.Close()
    if err != nil {
        fmt.Println(err)
    }

    // 3. Iterate through `names`

    for _, name := range names {
        fmt.Println(name.fname, name.lname)
    }
}
