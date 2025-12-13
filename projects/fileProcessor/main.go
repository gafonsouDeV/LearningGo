package main

import (
	"fmt"
	"log"
	"path/filepath"
	"sync"

	"github.com/gafonsouDeV/LearningGo/projects/fileProcessor/models"
	"github.com/gafonsouDeV/LearningGo/projects/fileProcessor/utils"
	"github.com/gafonsouDeV/LearningGo/projects/fileProcessor/workers"
)

const (
	inputDir      = "input"
	outputDir     = "output"
	logFile       = "processor.log"
	workersNumber = 3
)

func main() {
	fmt.Println("The file processing will start soon")

	logger, err := utils.NewLogger(logFile)
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Close()

	logger.Info("Starting Concurrent File Processor")

	files, err := utils.ListFiles(inputDir)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	if len(files) == 0 {
		logger.Info("No files to process")
		return
	}

	jobs := make(chan models.Job)
	results := make(chan string)
	errors := make(chan error)
	done := make(chan bool)

	var wg sync.WaitGroup

	for i := 1; i <= workersNumber; i++ {
		go workers.Worker(
			i,
			jobs,
			results,
			errors,
			&wg,
			logger,
			outputDir,
		)
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	go func() {
		for i, file := range files {
			wg.Add(1)
			job := models.Job{
				ID:       i + 1,
				FilePath: filepath.Join(inputDir, file),
				Filename: file,
			}

			jobs <- job
		}

		close(jobs)
	}()

	var processed int
	var failed int

	for {
		select {
		case res := <-results:
			processed++
			logger.Info(res)
		case err := <-errors:
			failed++
			logger.Error(err.Error())
		case <-done:
			logger.Info("All jobs finished")
			fmt.Printf("Files processed: %d\n", processed)
			fmt.Printf("Errors: %d\n", failed)
			fmt.Printf("Output directory: %s\n", outputDir)
			return
		}
	}
}
