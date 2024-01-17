package config

import (
	"os"
	"strconv"
)

// Model represents the configurations of the application.
type Model struct {
	// Toggl API token.
	TogglApiKey string `json:"toggl_api_key"`
	// Toggl workspace ID.
	TogglWorkspaceId int64 `json:"toggl_workspace_id"`
}

// GetConfig returns the configuration of the application.
func GetConfig() (Model, error) {
	var config Model
	workspaceId, err := strconv.ParseInt(os.Getenv("TOGGL_WORKSPACE_ID"), 10, 64)
	if err != nil {
		return config, err
	}

	togglApiToken := os.Getenv("TOGGL_API_TOKEN")

	config = Model{
		TogglApiKey:      togglApiToken,
		TogglWorkspaceId: workspaceId,
	}

	return config, nil
}
