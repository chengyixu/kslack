package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/minervacap2022/klik-slack-cli/internal/api"
	"github.com/minervacap2022/klik-slack-cli/internal/auth"
	"github.com/minervacap2022/klik-slack-cli/internal/output"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kslack",
	Short: "Slack CLI for KLIK platform",
	Long:  "Command-line interface for Slack Web API. Auth via SLACK_BOT_TOKEN env var.",
}

func getClient() *api.Client {
	token, err := auth.GetToken()
	if err != nil {
		output.Error(err.Error())
		os.Exit(1)
	}
	return api.NewClient(token)
}

// --- channel commands ---

var channelCmd = &cobra.Command{
	Use:   "channel",
	Short: "Manage Slack channels",
}

var channelListCmd = &cobra.Command{
	Use:   "list",
	Short: "List channels",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		limit, _ := cmd.Flags().GetInt("limit")
		types, _ := cmd.Flags().GetString("types")

		params := url.Values{}
		params.Set("limit", strconv.Itoa(limit))
		params.Set("types", types)
		params.Set("exclude_archived", "true")

		result, err := client.Get("conversations.list", params)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var channelInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get channel info",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")

		params := url.Values{}
		params.Set("channel", channel)

		result, err := client.Get("conversations.info", params)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

// --- message commands ---

var messageCmd = &cobra.Command{
	Use:   "message",
	Short: "Send and read messages",
}

var messageSendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message to a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		text, _ := cmd.Flags().GetString("text")

		payload := map[string]string{
			"channel": channel,
			"text":    text,
		}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("chat.postMessage", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var messageListCmd = &cobra.Command{
	Use:   "list",
	Short: "List messages in a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		limit, _ := cmd.Flags().GetInt("limit")

		params := url.Values{}
		params.Set("channel", channel)
		params.Set("limit", strconv.Itoa(limit))

		result, err := client.Get("conversations.history", params)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var messageReplyCmd = &cobra.Command{
	Use:   "reply",
	Short: "Reply to a thread",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		thread, _ := cmd.Flags().GetString("thread")
		text, _ := cmd.Flags().GetString("text")

		payload := map[string]string{
			"channel":   channel,
			"text":      text,
			"thread_ts": thread,
		}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("chat.postMessage", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var messageUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a message",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		ts, _ := cmd.Flags().GetString("timestamp")
		text, _ := cmd.Flags().GetString("text")

		payload := map[string]string{
			"channel": channel,
			"ts":      ts,
			"text":    text,
		}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("chat.update", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

// --- user commands ---

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users",
}

var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "List workspace users",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		limit, _ := cmd.Flags().GetInt("limit")

		params := url.Values{}
		params.Set("limit", strconv.Itoa(limit))

		result, err := client.Get("users.list", params)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

// --- reaction commands ---

var reactionCmd = &cobra.Command{
	Use:   "reaction",
	Short: "Manage reactions",
}

var reactionAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a reaction to a message",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		timestamp, _ := cmd.Flags().GetString("timestamp")
		emoji, _ := cmd.Flags().GetString("emoji")

		params := url.Values{}
		params.Set("channel", channel)
		params.Set("timestamp", timestamp)
		params.Set("name", emoji)

		result, err := client.Post("reactions.add", params)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

func init() {
	// channel
	channelListCmd.Flags().Int("limit", 100, "Max channels to return")
	channelListCmd.Flags().String("types", "public_channel,private_channel", "Channel types")
	channelInfoCmd.Flags().String("channel", "", "Channel ID")
	channelInfoCmd.MarkFlagRequired("channel")
	channelCmd.AddCommand(channelListCmd, channelInfoCmd)

	// message
	messageSendCmd.Flags().String("channel", "", "Channel ID or name")
	messageSendCmd.Flags().String("text", "", "Message text")
	messageSendCmd.MarkFlagRequired("channel")
	messageSendCmd.MarkFlagRequired("text")

	messageListCmd.Flags().String("channel", "", "Channel ID")
	messageListCmd.Flags().Int("limit", 20, "Max messages")
	messageListCmd.MarkFlagRequired("channel")

	messageReplyCmd.Flags().String("channel", "", "Channel ID")
	messageReplyCmd.Flags().String("thread", "", "Thread timestamp")
	messageReplyCmd.Flags().String("text", "", "Reply text")
	messageReplyCmd.MarkFlagRequired("channel")
	messageReplyCmd.MarkFlagRequired("thread")
	messageReplyCmd.MarkFlagRequired("text")

	messageUpdateCmd.Flags().String("channel", "", "Channel ID")
	messageUpdateCmd.Flags().String("timestamp", "", "Message timestamp")
	messageUpdateCmd.Flags().String("text", "", "New message text")
	messageUpdateCmd.MarkFlagRequired("channel")
	messageUpdateCmd.MarkFlagRequired("timestamp")
	messageUpdateCmd.MarkFlagRequired("text")

	messageCmd.AddCommand(messageSendCmd, messageListCmd, messageReplyCmd, messageUpdateCmd)

	// user
	userListCmd.Flags().Int("limit", 100, "Max users")
	userCmd.AddCommand(userListCmd)

	// reaction
	reactionAddCmd.Flags().String("channel", "", "Channel ID")
	reactionAddCmd.Flags().String("timestamp", "", "Message timestamp")
	reactionAddCmd.Flags().String("emoji", "", "Emoji name (without colons)")
	reactionAddCmd.MarkFlagRequired("channel")
	reactionAddCmd.MarkFlagRequired("timestamp")
	reactionAddCmd.MarkFlagRequired("emoji")
	reactionCmd.AddCommand(reactionAddCmd)

	rootCmd.AddCommand(channelCmd, messageCmd, userCmd, reactionCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
