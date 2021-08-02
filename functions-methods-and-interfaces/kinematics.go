package main

import (
    "fmt"
    "math"
)

type DisplaceFnType = func(time float64) float64

func GenDisplaceFn(acceleration float64, initialVelocity float64, initialDisplacement float64) DisplaceFnType {
    return func(time float64) float64 {
        return 0.5 * acceleration * math.Pow(time, 2) + initialVelocity * time + initialDisplacement
    }
}

func Prompt(promptMessage string, value *float64) bool {
    fmt.Println(promptMessage)
    _, err := fmt.Scanf("%f", value)

    if err != nil {
        fmt.Println(err)
        return false
    }

    return true
}

func main() {
    var acceleration float64
    var initialVelocity float64
    var initialDisplacement float64
    var time float64

    if !Prompt("Acceleration: ", &acceleration) {
        return
    }

    if !Prompt("Initial velocity: ", &initialVelocity) {
        return
    }

    if !Prompt("Initial displacement: ", &initialDisplacement) {
        return
    }

    if !Prompt("Time: ", &time) {
        return
    }

    fmt.Println("Calculated displacement: ",
        GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)(time))
}
