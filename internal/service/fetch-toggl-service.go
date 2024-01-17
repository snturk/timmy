package service

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/snturk/timmy/internal/config"
	"github.com/snturk/timmy/internal/constants"
	"github.com/snturk/timmy/internal/model"
	"github.com/snturk/timmy/internal/util"
	"io"
	"net/http"
	"strconv"
)

// FetchTodayToToggl fetches today's time entries from the file and saves them to Toggl.
func FetchTodayToToggl() error {

	// Check if there is any time entry.
	if !HasAnyTimeEntry() {
		return fmt.Errorf("there is no time entry, please start one first")
	}
	// Check if there is a running time entry.
	if IsAnyTimeEntryRunning() {
		return fmt.Errorf("there is a running time entry, please stop it first")
	}

	// Get today's time entries from the file.
	file, err := util.GetCurrentTimmyFile()
	if err != nil {
		return err
	}
	defer file.Close()

	// Read file line by line.
	scanner := bufio.NewScanner(file)
	var buffer bytes.Buffer
	for scanner.Scan() {
		entry, err := model.ParseTimeEntry(scanner.Text())
		if err != nil {
			return err
		}

		// If the entry is fetched or running, write it to the buffer.
		if entry.Fetched || entry.Running {
			buffer.WriteString(entry.String() + "\n")
			continue
		}

		// If the entry is not fetched, fetch it.
		updatedEntry, err := fetchEntryToToggl(entry)
		if err != nil {
			return err
		}

		// After fetching the entry, write it to the file.
		buffer.WriteString(updatedEntry.String() + "\n")
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

	return nil
}

func fetchEntryToToggl(entry model.TimeEntry) (model.TimeEntry, error) {
	configModel, err := config.GetConfig()
	if err != nil {
		return entry, err
	}
	workspaceIdString := strconv.FormatInt(configModel.TogglWorkspaceId, 10)
	togglApiUrl := "https://api.track.toggl.com/api/v9/workspaces/" + workspaceIdString + "/time_entries"

	// Create the request body TogglManualTimeRequest
	requestBody := model.TogglManualTimeRequest{
		Description: entry.Task,
		Tags:        []string{},
		Duration:    int(entry.GetDuration().Seconds()),
		Start:       entry.Start.Format(constants.DateTimeFormat),
		CreatedWith: "timmy-cli",
		WorkspaceId: int(configModel.TogglWorkspaceId),
	}

	// Send the request to Toggl.
	jsonStr, err := json.Marshal(requestBody)
	if err != nil {
		return entry, err
	}

	// Send the request to Toggl.
	req, err := http.NewRequest("POST", togglApiUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		return entry, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(configModel.TogglApiKey, "api_token")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.TimeEntry{}, err
	}
	if err != nil {
		return entry, err
	}
	if resp.StatusCode != 200 {
		all, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(all))
		return entry, fmt.Errorf("error while fetching entry to Toggl: %v", resp.Status)
	}
	defer resp.Body.Close()

	// Mark the entry as fetched.
	entry.Fetched = true
	return entry, nil
}
