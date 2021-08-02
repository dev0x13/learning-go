package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "strings"
)

type Animal interface {
    Eat()
    Move()
    Speak()
}

type Cow struct {}
type Bird struct {}
type Snake struct {}

func (c Cow) Eat() {
    fmt.Println("grass")
}

func (c Cow) Move() {
    fmt.Println("walk")
}

func (c Cow) Speak() {
    fmt.Println("moo")
}

func (b Bird) Eat() {
    fmt.Println("worms")
}

func (b Bird) Move() {
    fmt.Println("fly")
}

func (b Bird) Speak() {
    fmt.Println("peep")
}

func (s Snake) Eat() {
    fmt.Println("mice")
}

func (s Snake) Move() {
    fmt.Println("slither")
}

func (s Snake) Speak() {
    fmt.Println("hsss")
}

var animals = map[string] Animal {}

var actions = map[string] func (Animal) {
    "eat":   Animal.Eat,
    "move":  Animal.Move,
    "speak": Animal.Speak,
}

func ProcessRequest(request string) error {
    requestComponents := strings.Split(request, " ")

    if len(requestComponents) != 3 {
        return errors.New(fmt.Sprintf("expected exactly 3 components, got %d", len(requestComponents)))
    }

    switch requestComponents[0] {
    case "newanimal":
        switch requestComponents[2] {
        case "cow":
            animals[requestComponents[1]] = Cow{}
        case "bird":
            animals[requestComponents[1]] = Bird{}
        case "snake":
            animals[requestComponents[1]] = Snake{}
        default:
            return errors.New(fmt.Sprintf("unexpected animal kind: %s", requestComponents[2]))
        }

        fmt.Println("Created it!")

        return nil
    case "query":
        animal, found := animals[requestComponents[1]]

        if !found {
            return errors.New(fmt.Sprintf("no animal found using the given name: %s", requestComponents[1]))
        }

        if actionFunction, ok := actions[requestComponents[2]]; !ok {
            return errors.New(fmt.Sprintf("unexpected action type: %s", requestComponents[2]))
        } else {
            actionFunction(animal)
        }

        return nil
    default:
        return errors.New(fmt.Sprintf("unexpected action type: %s", requestComponents[0]))
    }
}

func main() {
    s := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print(">")
        s.Scan()
        err := ProcessRequest(s.Text())

        if err != nil {
            fmt.Println("Invalid request:", err)
        }
    }
}