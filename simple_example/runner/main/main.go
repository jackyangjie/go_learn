package main

import (
	"log"
	"os"
	"simple_example/runner"
	"time"
)

const timeout = 3 * time.Second

func main() {
	log.Println("Starting work.")
	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask())
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeOut:
			log.Println("Terminating due to timeout")
			os.Exit(1)

		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(1)
		}
	}
	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(i int) {
		log.Printf("Processor - Task #%d.", i)
		time.Sleep(time.Duration(i) * time.Second)
	}
}
