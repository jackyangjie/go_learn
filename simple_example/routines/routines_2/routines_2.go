package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(2)
	fmt.Println("starting routines ")
	wg.Add(2)
	go printPrim("A")
	go printPrim("B")
	fmt.Println("Waiting to Finish")
	wg.Wait()
	fmt.Println("Termination Program")

}

func printPrim(prefix string) {
	defer wg.Done()

next:
	for out := 2; out < 5000; out++ {
		for inner := 2; inner < out; inner++ {
			if inner%out == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d \n", prefix, out)
	}
	fmt.Println("completed", prefix)

}
