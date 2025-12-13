package workers

import (
	"fmt"
	"sync"

	"github.com/gafonsouDeV/LearningGo/projects/fileProcessor/models"
	"github.com/gafonsouDeV/LearningGo/projects/fileProcessor/utils"
)

func Worker(
	id int,
	jobs <-chan models.Job,
	results chan<- string,
	errors chan<- error,
	wg *sync.WaitGroup,
	logger *utils.Logger,
	outputDir string,
) {
	logger.Info(fmt.Sprintf("Worker %d started", id))
	for job := range jobs {
		logger.Info(fmt.Sprintf(
			"Worker %d processing Job %d (%s)",
			id,
			job.ID,
			job.Filename,
		))

		stats, err := utils.ReadFileStats(job.FilePath)
		if err != nil {
			errors <- fmt.Errorf(
				"Worker %d failed writing output for %s: %w",
				id,
				job.Filename,
				err,
			)
			wg.Done()
			continue
		}

		if err := utils.WriteResult(outputDir, job.Filename, stats); err != nil {
			errors <- fmt.Errorf(
				"Worker %d failed writing output for %s: %w",
				id,
				job.Filename,
				err,
			)
			wg.Done()
			continue
		}

		result := fmt.Sprintf(
			"Worker %d processed %s (lines=%d words=%d bytes=%d)",
			id,
			job.Filename,
			stats.Lines,
			stats.Words,
			stats.Bytes64,
		)

		results <- result
		wg.Done()
	}

	logger.Info(fmt.Sprintf("Worker %d stopped", id))
}
