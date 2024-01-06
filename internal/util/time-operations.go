package util

import (
	"bufio"
	"github.com/snturk/timmy/internal/model"
	"time"
)

// GetTaskDuration calculates the duration of a time entry, and returns it.
func GetTaskDuration(scanner *bufio.Scanner) (string, time.Duration, error) {
	line := scanner.Text()
	entry, err := model.ParseTimeEntry(line)
	if err != nil {
		return "", time.Duration(0), err
	}

	// If it's Running, don't print it.
	if entry.Running {
		return "", time.Duration(0), nil
	}

	// Calculate duration.
	return entry.Task, entry.End.Sub(entry.Start), nil
}

// CheckIfTimeEntryIsRunning checks if there is a running time entry.
func CheckIfTimeEntryIsRunning() bool {
	file, err := GetCurrentTimmyFile()
	if err != nil {
		return false
	}
	defer file.Close()

	// Read file line by line.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Parse line.
		entry, err := model.ParseTimeEntry(line)
		if err != nil {
			return false
		}

		// Check if entry is running.
		if entry.Running {
			return true
		}
	}

	return false
}
