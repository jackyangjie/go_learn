package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	wg.Add(2)
	court := make(chan int)
	go play("Nadal", court)
	go play("Djokovic", court)
	court <- 1
	wg.Wait()
}

func play(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("play %s,missed \n", name)
			return
		}
		r := rand.Intn(100)
		if r%13 == 0 {
			fmt.Printf("play %s,missed \n ", name)
			close(court)
			return
		}
		fmt.Printf("play %s ,Hit %d \n", name, ball)
		ball++
		court <- ball
	}
}
