package workers

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/gafonsouDeV/LearningGo/projects/fakeImageProcessing/utils"
)

func Worker(
	id int,
	jobs <-chan string,
	results chan<- string,
	errors chan<- error,
	wg *sync.WaitGroup,
	logMutex *sync.Mutex,
	logs *[]string,
) {
	for img := range jobs {
		processImage(id, img, results, errors, wg, logMutex, logs)
	}
}

func processImage(
	id int,
	img string,
	results chan<- string,
	errors chan<- error,
	wg *sync.WaitGroup,
	logMutex *sync.Mutex,
	logs *[]string,
) {
	fmt.Printf("Worker %d → Downloading %s...\n", id, img)
	time.Sleep(utils.RandDelay())

	if rand.Intn(10) < 2 {
		err := fmt.Errorf("Worker %d failed processing %s", id, img)
		errors <- err

		logMutex.Lock()

		*logs = append(*logs, err.Error())
		logMutex.Unlock()

		wg.Done()
		return
	}

	fmt.Printf("Worker %d → Processing %s...\n", id, img)
	time.Sleep(utils.RandDelay())

	result := fmt.Sprintf("Worker %d finished %s", id, img)
	results <- result

	logMutex.Lock()
	*logs = append(*logs, result)
	logMutex.Unlock()
	wg.Done()
}
