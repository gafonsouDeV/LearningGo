package main

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/gafonsouDeV/LearningGo/projects/fakeImageProcessing/workers"
)

func main() {
	var images [6]string
	for index := range len(images) {
		images[index] = "image" + strconv.Itoa(index+1) + ".png"
	}

	const workerCount = 3

	jobs := make(chan string)
	results := make(chan string)
	errors := make(chan error)
	done := make(chan bool)

	var wg sync.WaitGroup
	var logMutex sync.Mutex
	var logs []string

	for i := 1; i < workerCount; i++ {
		go workers.Worker(i, jobs, results, errors, &wg, &logMutex, &logs)
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	go func() {
		for _, img := range images {
			wg.Add(1)
			jobs <- img
		}
		close(jobs)
	}()

	for {
		select {
		case res := <-results:
			fmt.Println("[RESULT]", res)
		case err := <-errors:
			fmt.Println("[ERROR]", err)
		case <-done:
			fmt.Println("\n--- All images processed! ---")
			fmt.Printf("Total logs stored: %d\n", len(logs))
			fmt.Println("\nLog summary:")
			for _, entry := range logs {
				fmt.Println(" -", entry)
			}
			return
		}
	}

}
