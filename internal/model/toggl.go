package model

type TogglManualTimeRequest struct {
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	/**
	 * Duration in seconds.
	 */
	Duration    int    `json:"duration"`
	Start       string `json:"start"`
	WorkspaceId int    `json:"workspace_id"`
	CreatedWith string `json:"created_with"`
}
