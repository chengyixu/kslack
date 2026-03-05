package auth

import (
	"fmt"
	"os"
)

func GetToken() (string, error) {
	token := os.Getenv("SLACK_BOT_TOKEN")
	if token == "" {
		return "", fmt.Errorf("SLACK_BOT_TOKEN environment variable is required")
	}
	return token, nil
}

func GetTeamID() string {
	return os.Getenv("SLACK_TEAM_ID")
}
