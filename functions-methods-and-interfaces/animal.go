package main

import (
    "bufio"
    "errors"
    "fmt"
    "os"
    "strings"
)

type Animal struct {
    food string
    locomotion string
    noise string
}

func (a Animal) Eat() string {
    return a.food
}

func (a Animal) Move() string {
    return a.locomotion
}

func (a Animal) Speak() string {
    return a.noise
}

var animals = map[string]Animal{
    "cow":   {food: "grass", locomotion: "walk",    noise: "moo"},
    "bird":  {food: "worms", locomotion: "fly",     noise: "peep"},
    "snake": {food: "mice",  locomotion: "slither", noise: "hsss"},
}

var actions = map[string] func (Animal) string {
    "eat":   Animal.Eat,
    "move":  Animal.Move,
    "speak": Animal.Speak,
}

func ProcessRequest(request string) (string, error) {
    requestComponents := strings.Split(request, " ")

    if len(requestComponents) != 2 {
        return "", errors.New(fmt.Sprintf("expected exactly 2 components, got %d", len(requestComponents)))
    }

    animalKind := requestComponents[0]
    actionType := requestComponents[1]

    if animalInstance, ok := animals[animalKind]; !ok {
        return "", errors.New(fmt.Sprintf("unexpected animal kind: %s", animalKind))
    } else {
        if actionFunction, ok := actions[actionType]; !ok {
            return "", errors.New(fmt.Sprintf("unexpected action type: %s", actionType))
        } else {
            return actionFunction(animalInstance), nil
        }
    }
}

func main() {
    s := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print(">")
        s.Scan()
        result, err := ProcessRequest(s.Text())

        if err != nil {
            fmt.Println("Invalid request: ", err)
        } else {
            fmt.Println(result)
        }
    }
}
