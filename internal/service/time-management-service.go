package service

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/snturk/timmy/internal/model"
	"github.com/snturk/timmy/internal/util"
)

// StartTimeEntry starts a new time entry, and saves it to the file.
func StartTimeEntry(task string) error {
	entry := model.TimeEntry{
		Start:   time.Now().Local(),
		Running: true,
		Task:    task,
	}

	// Write that into file.

	file, err := util.GetCurrentTimmyFile()
	if err != nil {
		return err
	}
	defer file.Close()

	// Check if there is a running time entry.
	if util.CheckIfTimeEntryIsRunning() {
		return errors.New("there is a running time entry, please stop it first")
	}

	// Write entry to file.
	entryString := entry.String()
	_, err = file.WriteString(entryString + "\n")
	if err != nil {
		return err
	}

	return nil
}

// StopTimeEntry stops the currently running time entry, and saves it to the file.
func StopTimeEntry() error {
	file, err := util.GetCurrentTimmyFile()
	if err != nil {
		return err
	}
	defer file.Close()
	lineCount := 0

	// Check if there is a running time entry.
	if !util.CheckIfTimeEntryIsRunning() {
		return errors.New("there is no running time entry, please start one first")
	}

	// Read file line by line.
	scanner := bufio.NewScanner(file)
	var buffer bytes.Buffer
	for scanner.Scan() {
		lineCount++
		line := scanner.Text()

		// Parse line.
		entry, err := model.ParseTimeEntry(line)
		if err != nil {
			return err
		}

		// Check if entry is running.
		//If it is, replace it with the stopped entry. If it's not, write it to the buffer.
		if entry.Running {
			entry.Running = false
			entry.End = time.Now().Local()

			// Write entry to file.
			entryString := entry.String()
			buffer.WriteString(entryString + "\n")
		} else {
			buffer.WriteString(line + "\n")
		}
	}

	// Write the buffer to the file.
	err = file.Truncate(0)
	if err != nil {
		return err
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}
	_, err = file.WriteString(buffer.String())
	if err != nil {
		return err
	}
	fmt.Println("Stopped the running time entry.")

	return nil
}

// PrintTodayBrief prints a brief summary of today's time entries.
func PrintTodayBrief() error {
	hasEntryFlag := false
	briefText := bytes.Buffer{}
	file, err := util.GetCurrentTimmyFile()
	if err != nil {
		return err
	}
	defer file.Close()
	var durationMap = make(map[string]time.Duration)

	// Read file line by line.
	scanner := bufio.NewScanner(file)

	// Print header.
	briefText.WriteString("Your daily time entries:\n")

	for scanner.Scan() {
		hasEntryFlag = true
		entryName, duration, err := util.GetTaskDuration(scanner)
		if err != nil {
			return err
		}

		durationMap[entryName] += duration
	}

	if !hasEntryFlag {
		fmt.Println("You have no time entries for today.")
		return nil
	}

	totalDuration := time.Duration(0)
	for task, duration := range durationMap {
		// Print task and duration.
		briefText.WriteString("- " + task + ": " + duration.String() + "\n")

		// Add to total duration.
		totalDuration += duration
	}

	// Print total duration.
	briefText.WriteString("Total: " + totalDuration.String() + "\n")

	fmt.Println(briefText.String())

	return nil
}

// PrintCurrent prints the currently running time entry. If there is none, it prints a message.
func PrintCurrent() error {
	file, err := util.GetCurrentTimmyFile()
	if err != nil {
		return err
	}
	defer file.Close()

	// Read file line by line.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Parse line.
		entry, err := model.ParseTimeEntry(line)
		if err != nil {
			return err
		}

		// Check if entry is running.
		//If it is, print it. If it's not, do nothing.
		if entry.Running {
			fmt.Println("Your current time entry:")
			// Do not print after the . (dot).
			fmt.Println("- " + entry.Task + ": " + entry.GetDuration().Round(time.Second).String())
			return nil
		}
	}

	fmt.Println("You have no running time entries.")

	return nil
}

// IsAnyTimeEntryRunning returns true if there is any time entry running.
func IsAnyTimeEntryRunning() bool {
	file, err := util.GetCurrentTimmyFile()
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
		//If it is, return true. If it's not, do nothing.
		if entry.Running {
			return true
		}
	}

	return false
}

// HasAnyTimeEntry returns true if there is any time entry.
func HasAnyTimeEntry() bool {
	file, err := util.GetCurrentTimmyFile()
	if err != nil {
		return false
	}
	defer file.Close()

	// Read file line by line.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return true
	}

	return false
}
