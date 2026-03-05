package auth

import (
	"fmt"
	"os"
)

func GetToken() (string, error) {
	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		return "", fmt.Errorf("SLACK_TOKEN environment variable is required (user or bot token)")
	}
	return token, nil
}

func GetTeamID() string {
	return os.Getenv("SLACK_TEAM_ID")
}
