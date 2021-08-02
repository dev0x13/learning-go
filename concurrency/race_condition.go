package main

import (
    "fmt"
    "sync"
)

var x = 1

func worker1(wg *sync.WaitGroup) {
    defer wg.Done()

    x *= 2
}

func worker2(wg *sync.WaitGroup) {
    defer wg.Done()

    x += 1
}

func main() {
    var wg sync.WaitGroup

    wg.Add(1)
    go worker1(&wg)

    wg.Add(1)
    go worker2(&wg)

    wg.Wait()

    // Race condition is a situation when the result of the parallelized
    // program/function/routine etc. is non deterministic and defined by
    // the parallel execution order, which usually differs between program
    // runs du to the system scheduler works. Usually this situation happens
    // because of the concurrent read and write operations of the same data.
    // In this program result can be either 4 or 3, because two goroutines
    // modify 'x' variable concurrently, and the result depends on which 
    // goroutine will be executed first.
    // kek
    fmt.Println(x)
}
