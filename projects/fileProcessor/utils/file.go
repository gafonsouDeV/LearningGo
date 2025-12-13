package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileStats struct {
	Lines   int
	Words   int
	Bytes64 int64
}

func ListFiles(filepath string) ([]string, error) {
	entries, err := os.ReadDir(filepath)

	if err != nil {
		return nil, err
	}

	var files []string

	for _, entrie := range entries {
		if !entrie.IsDir() {
			files = append(files, entrie.Name())
		}
	}

	return files, nil
}

func ReadFileStats(path string) (FileStats, error) {
	file, err := os.Open(path)

	if err != nil {
		return FileStats{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines, words int

	for scanner.Scan() {
		lines++
		words += len(strings.Fields(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return FileStats{}, err
	}

	info, err := file.Stat()
	if err != nil {
		return FileStats{}, err
	}

	return FileStats{
		Lines:   lines,
		Words:   words,
		Bytes64: info.Size(),
	}, nil
}

func WriteResult(outputDir, filename string, stats FileStats) error {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	outhPath := filepath.Join(outputDir, filename+time.Now().Format("20060102150405")+".out")

	content := fmt.Sprintf(
		"Lines: %d\nWords: %d\nnBytes: %d\n",
		stats.Lines,
		stats.Words,
		stats.Bytes64,
	)

	return os.WriteFile(outhPath, []byte(content), 0644)
}
