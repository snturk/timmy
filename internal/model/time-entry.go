package model

import (
	"bytes"
	"github.com/snturk/timmy/internal/constants"
	"strconv"
	"time"
)

// TimeEntry represents a single time entry.
type TimeEntry struct {
	// The start date and time of the entry.
	Start time.Time `json:"start"`
	// The end date and time of the entry, if it has ended.
	End time.Time `json:"end",omitempty`
	// The task that the entry is for.
	Task string `json:"task"`
	// Is the entry currently running?
	Running bool `json:"running"`
}

// String returns a string representation of the time entry.
func (entry TimeEntry) String() string {
	// 'Start' - 'End' - 'Task' - 'Running'
	var entryString bytes.Buffer
	entryString.WriteString(entry.Start.Format(constants.DateTimeFormat))
	entryString.WriteString(constants.FileDataDivider)
	entryString.WriteString(entry.End.Format(constants.DateTimeFormat))
	entryString.WriteString(constants.FileDataDivider)
	entryString.WriteString(entry.Task)
	entryString.WriteString(constants.FileDataDivider)
	entryString.WriteString(strconv.FormatBool(entry.Running))

	return entryString.String()
}

// ParseTimeEntry parses a string representation of a time entry.
func ParseTimeEntry(entryString string) (TimeEntry, error) {
	var entry TimeEntry
	var err error

	// 'Start' - 'End' - 'Task' - 'Running'
	entryData := bytes.Split([]byte(entryString), []byte(constants.FileDataDivider))

	// Parse the start time.
	entry.Start, err = time.Parse(constants.DateTimeFormat, string(entryData[0]))
	if err != nil {
		return entry, err
	}

	// Parse the end time.
	entry.End, err = time.Parse(constants.DateTimeFormat, string(entryData[1]))
	if err != nil {
		return entry, err
	}

	// Parse the task.
	entry.Task = string(entryData[2])

	// Parse the running status.
	entry.Running, err = strconv.ParseBool(string(entryData[3]))
	if err != nil {
		return entry, err
	}

	return entry, nil
}
