package util

import (
	"bufio"
	"github.com/snturk/timmy/internal/model"
)

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
