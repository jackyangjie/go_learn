package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

const (
	numberGoroutines = 4
	taskLoad         = 10
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks := make(chan string, taskLoad)
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		//tasks <- fmt.Sprintf("Task:%d",post)
		tasks <- "Task:" + strconv.Itoa(post)
	}

	close(tasks)
	wg.Wait()
}

func worker(tasks chan string, gr int) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker %d :shutting down \n", gr)
			return
		}
		fmt.Printf("Worker %d : Started %s \n", gr, task)
		n := rand.Int63n(100)
		time.Sleep(time.Duration(n) * time.Millisecond)
		fmt.Printf("Worker: %d ：Complated : %s \n", gr, task)

	}
}
