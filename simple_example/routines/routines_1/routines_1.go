package main

import (
	"fmt"
	. "runtime"
	"sync"
)

func main() {
	GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("start routines")

	go func() {
		defer wg.Done()

		for c := 0; c < 3; c++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for c := 0; c < 3; c++ {
			for a := 'A'; a < 'A'+26; a++ {
				fmt.Printf("%c ", a)
			}
		}

	}()
	fmt.Println("Waiting to finish ")
	wg.Wait()
	fmt.Println("Termination program")

}
