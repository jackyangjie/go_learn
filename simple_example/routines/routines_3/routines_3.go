package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {
	wg.Add(2)
	//go incCounter(1)
	//go incCounter(2)

	//atomicCounter(1)
	//atomicCounter(2)
	lockCounter(1)
	lockCounter(2)
	wg.Wait()
	fmt.Println("Finished", counter)
}

func incCounter(i int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter
		runtime.Gosched()
		value++
		counter = value

	}
}

//cas
func atomicCounter(i int) {
	defer wg.Done()

	for c := 0; c < 2; c++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}

//lock
func lockCounter(i int) {
	defer wg.Done()
	for c := 0; c < 2; c++ {
		mutex.Lock()
		{
			counter++
		}
		mutex.Unlock()
	}

}
