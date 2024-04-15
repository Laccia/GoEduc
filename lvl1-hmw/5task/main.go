package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(id int, done chan struct{}) {
	for i := 0; i < 1000; i++ {
		fmt.Println("Worker", id, "processing:", i)
	}
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})

	workers := 5

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGTERM)

	go func() {
		<-signalChan
		close(done)
	}()

	for i := 0; i < workers; i++ {
		go worker(i, done)
	}

	for _ = range done {
		fmt.Println("All workers finished.")
		break
	}

	time.Sleep(1 * time.Second)
	fmt.Println("Program stopped.")
}
