package main

import (
    "fmt"
    "sync"
)

type Chopstick struct {
    sync.Mutex
}

type Philosopher struct {
    leftChopstick *Chopstick
    rightChopstick *Chopstick
    id int
}

// Philosophers synchronization channel
var channel chan int

func (p Philosopher) eat(wg *sync.WaitGroup) {
    const numEats = 3

    for i := 0; i < numEats; i++ {
        <- channel

        p.leftChopstick.Lock()
        p.rightChopstick.Lock()

        fmt.Printf("starting to eat %d\n", p.id)

        p.leftChopstick.Unlock()
        p.rightChopstick.Unlock()

        fmt.Printf("finished eating %d\n", p.id)

        channel <- 1
    }

    wg.Done()
}

func main() {
    const numPhilosophers = 5
    const maxNumPhilosophersEating = 2

    // 1. Create chopsticks

    chopsticks := make([]*Chopstick, numPhilosophers)
    for i := 0; i < numPhilosophers; i++ {
        chopsticks[i] = new(Chopstick)
    }

    // 2. Create philosophers

    philosophers := make([]*Philosopher, numPhilosophers)
    for i := 0; i < numPhilosophers; i++ {
        philosophers[i] = &Philosopher{ chopsticks[i], chopsticks[(i + 1) % numPhilosophers], i + 1 }
    }

    // 3. Create WaitGroup

    var wg sync.WaitGroup
    wg.Add(numPhilosophers)

    // 4. Init synchronization channel

    channel = make(chan int, maxNumPhilosophersEating)
    for i := 0; i < maxNumPhilosophersEating; i++ {
        channel <- 1
    }

    // 5. Run goroutines

    for i := 0; i < numPhilosophers; i++ {
        go philosophers[i].eat(&wg)
    }

    // 6. Wait for completion

    wg.Wait()
}
