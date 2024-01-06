package util

import (
	"errors"
	"os"
	"time"
)

// GetCurrentTimmyFile returns the timmy file for the current day.
func GetCurrentTimmyFile() (*os.File, error) {
	// Get directory from path variable TIMMY_PATH.
	directory := os.Getenv("TIMMY_PATH")
	if directory == "" {
		directory = os.Getenv("HOME") + "/.timmy-entries"
	}

	// Create directory if it doesn't exist.
	if _, err := os.Stat(directory); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(directory, 0755)
		if err != nil {
			return nil, err
		}
	}

	filename := directory + "/" + time.Now().Local().Format("2006-01-02") + ".timmy"

	// Create file if it doesn't exist. This file is for reading and writing and append.
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}
