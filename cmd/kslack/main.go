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
	Long:  "Command-line interface for Slack Web API. Auth via SLACK_TOKEN env var.",
}

func getClient() *api.Client {
	token, err := auth.GetToken()
	if err != nil {
		output.Error(err.Error())
		os.Exit(1)
	}
	return api.NewClient(token)
}

// ============================================================================
// channel commands
// ============================================================================

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

var channelMembersCmd = &cobra.Command{
	Use:   "members",
	Short: "List members of a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		limit, _ := cmd.Flags().GetInt("limit")

		params := url.Values{}
		params.Set("channel", channel)
		params.Set("limit", strconv.Itoa(limit))

		result, err := client.Get("conversations.members", params)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var channelCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		name, _ := cmd.Flags().GetString("name")

		payload := map[string]string{"name": name}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("conversations.create", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var channelArchiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Archive a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")

		payload := map[string]string{"channel": channel}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("conversations.archive", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var channelUnarchiveCmd = &cobra.Command{
	Use:   "unarchive",
	Short: "Unarchive a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")

		payload := map[string]string{"channel": channel}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("conversations.unarchive", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var channelInviteCmd = &cobra.Command{
	Use:   "invite",
	Short: "Invite users to a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		users, _ := cmd.Flags().GetString("users")

		payload := map[string]string{"channel": channel, "users": users}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("conversations.invite", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var channelKickCmd = &cobra.Command{
	Use:   "kick",
	Short: "Remove a user from a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		user, _ := cmd.Flags().GetString("user")

		payload := map[string]string{"channel": channel, "user": user}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("conversations.kick", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var channelJoinCmd = &cobra.Command{
	Use:   "join",
	Short: "Join a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")

		payload := map[string]string{"channel": channel}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("conversations.join", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var channelLeaveCmd = &cobra.Command{
	Use:   "leave",
	Short: "Leave a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")

		payload := map[string]string{"channel": channel}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("conversations.leave", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var channelSetTopicCmd = &cobra.Command{
	Use:   "set-topic",
	Short: "Set channel topic",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		topic, _ := cmd.Flags().GetString("topic")

		payload := map[string]string{"channel": channel, "topic": topic}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("conversations.setTopic", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var channelSetPurposeCmd = &cobra.Command{
	Use:   "set-purpose",
	Short: "Set channel purpose/description",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		purpose, _ := cmd.Flags().GetString("purpose")

		payload := map[string]string{"channel": channel, "purpose": purpose}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("conversations.setPurpose", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var channelRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		name, _ := cmd.Flags().GetString("name")

		payload := map[string]string{"channel": channel, "name": name}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("conversations.rename", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var channelMarkCmd = &cobra.Command{
	Use:   "mark",
	Short: "Set read cursor in a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		ts, _ := cmd.Flags().GetString("timestamp")

		payload := map[string]string{"channel": channel, "ts": ts}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("conversations.mark", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

// ============================================================================
// message commands
// ============================================================================

var messageCmd = &cobra.Command{
	Use:   "message",
	Short: "Send, read, and manage messages",
}

var messageSendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a message to a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		text, _ := cmd.Flags().GetString("text")

		payload := map[string]string{"channel": channel, "text": text}
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

var messageThreadCmd = &cobra.Command{
	Use:   "thread",
	Short: "List replies in a thread",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		ts, _ := cmd.Flags().GetString("timestamp")
		limit, _ := cmd.Flags().GetInt("limit")

		params := url.Values{}
		params.Set("channel", channel)
		params.Set("ts", ts)
		params.Set("limit", strconv.Itoa(limit))

		result, err := client.Get("conversations.replies", params)
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

		payload := map[string]string{"channel": channel, "ts": ts, "text": text}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("chat.update", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var messageDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a message",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		ts, _ := cmd.Flags().GetString("timestamp")

		payload := map[string]string{"channel": channel, "ts": ts}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("chat.delete", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var messageScheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Schedule a message",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		text, _ := cmd.Flags().GetString("text")
		postAt, _ := cmd.Flags().GetString("post-at")

		payload := map[string]string{"channel": channel, "text": text, "post_at": postAt}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("chat.scheduleMessage", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

// ============================================================================
// user commands
// ============================================================================

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

var userInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get user info",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		user, _ := cmd.Flags().GetString("user")

		params := url.Values{}
		params.Set("user", user)

		result, err := client.Get("users.info", params)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var userProfileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Get user profile",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		user, _ := cmd.Flags().GetString("user")

		params := url.Values{}
		if user != "" {
			params.Set("user", user)
		}

		result, err := client.Get("users.profile.get", params)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

// ============================================================================
// reaction commands
// ============================================================================

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

		payload := map[string]string{
			"channel":   channel,
			"timestamp": timestamp,
			"name":      emoji,
		}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("reactions.add", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var reactionRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a reaction from a message",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		timestamp, _ := cmd.Flags().GetString("timestamp")
		emoji, _ := cmd.Flags().GetString("emoji")

		payload := map[string]string{
			"channel":   channel,
			"timestamp": timestamp,
			"name":      emoji,
		}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("reactions.remove", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var reactionGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get reactions for a message",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		timestamp, _ := cmd.Flags().GetString("timestamp")

		params := url.Values{}
		params.Set("channel", channel)
		params.Set("timestamp", timestamp)
		params.Set("full", "true")

		result, err := client.Get("reactions.get", params)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

// ============================================================================
// pin commands
// ============================================================================

var pinCmd = &cobra.Command{
	Use:   "pin",
	Short: "Manage pinned messages",
}

var pinListCmd = &cobra.Command{
	Use:   "list",
	Short: "List pinned items in a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")

		params := url.Values{}
		params.Set("channel", channel)

		result, err := client.Get("pins.list", params)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var pinAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Pin a message to a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		timestamp, _ := cmd.Flags().GetString("timestamp")

		payload := map[string]string{"channel": channel, "timestamp": timestamp}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("pins.add", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var pinRemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Unpin a message from a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		timestamp, _ := cmd.Flags().GetString("timestamp")

		payload := map[string]string{"channel": channel, "timestamp": timestamp}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("pins.remove", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

// ============================================================================
// file commands
// ============================================================================

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Manage files",
}

var fileListCmd = &cobra.Command{
	Use:   "list",
	Short: "List files",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		limit, _ := cmd.Flags().GetInt("limit")

		params := url.Values{}
		params.Set("count", strconv.Itoa(limit))
		if channel != "" {
			params.Set("channel", channel)
		}

		result, err := client.Get("files.list", params)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var fileInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get file info",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		file, _ := cmd.Flags().GetString("file")

		params := url.Values{}
		params.Set("file", file)

		result, err := client.Get("files.info", params)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var fileDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a file",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		file, _ := cmd.Flags().GetString("file")

		payload := map[string]string{"file": file}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("files.delete", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

// ============================================================================
// search commands
// ============================================================================

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search messages and files",
}

var searchMessagesCmd = &cobra.Command{
	Use:   "messages",
	Short: "Search messages",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		query, _ := cmd.Flags().GetString("query")
		count, _ := cmd.Flags().GetInt("count")
		sort, _ := cmd.Flags().GetString("sort")

		params := url.Values{}
		params.Set("query", query)
		params.Set("count", strconv.Itoa(count))
		if sort != "" {
			params.Set("sort", sort)
		}

		result, err := client.Get("search.messages", params)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var searchFilesCmd = &cobra.Command{
	Use:   "files",
	Short: "Search files",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		query, _ := cmd.Flags().GetString("query")
		count, _ := cmd.Flags().GetInt("count")

		params := url.Values{}
		params.Set("query", query)
		params.Set("count", strconv.Itoa(count))

		result, err := client.Get("search.files", params)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

// ============================================================================
// reminder commands
// ============================================================================

var reminderCmd = &cobra.Command{
	Use:   "reminder",
	Short: "Manage reminders",
}

var reminderAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a reminder",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		text, _ := cmd.Flags().GetString("text")
		time, _ := cmd.Flags().GetString("time")

		payload := map[string]string{"text": text, "time": time}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("reminders.add", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var reminderListCmd = &cobra.Command{
	Use:   "list",
	Short: "List reminders",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()

		result, err := client.Get("reminders.list", url.Values{})
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var reminderDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a reminder",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		reminder, _ := cmd.Flags().GetString("reminder")

		payload := map[string]string{"reminder": reminder}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("reminders.delete", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var reminderCompleteCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a reminder as complete",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		reminder, _ := cmd.Flags().GetString("reminder")

		payload := map[string]string{"reminder": reminder}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("reminders.complete", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

// ============================================================================
// bookmark commands
// ============================================================================

var bookmarkCmd = &cobra.Command{
	Use:   "bookmark",
	Short: "Manage channel bookmarks",
}

var bookmarkListCmd = &cobra.Command{
	Use:   "list",
	Short: "List bookmarks in a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")

		params := url.Values{}
		params.Set("channel_id", channel)

		result, err := client.Get("bookmarks.list", params)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var bookmarkAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a bookmark to a channel",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		channel, _ := cmd.Flags().GetString("channel")
		title, _ := cmd.Flags().GetString("title")
		link, _ := cmd.Flags().GetString("link")

		payload := map[string]string{
			"channel_id": channel,
			"title":      title,
			"type":       "link",
			"link":       link,
		}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("bookmarks.add", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

// ============================================================================
// dm commands (direct messages)
// ============================================================================

var dmCmd = &cobra.Command{
	Use:   "dm",
	Short: "Direct messages",
}

var dmOpenCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a DM with user(s)",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		users, _ := cmd.Flags().GetString("users")

		payload := map[string]string{"users": users}
		jsonBody, _ := json.Marshal(payload)

		result, err := client.PostJSON("conversations.open", jsonBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

var dmSendCmd = &cobra.Command{
	Use:   "send",
	Short: "Open DM and send a message",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := getClient()
		users, _ := cmd.Flags().GetString("users")
		text, _ := cmd.Flags().GetString("text")

		// First open the conversation
		openPayload := map[string]string{"users": users}
		openBody, _ := json.Marshal(openPayload)

		openResult, err := client.PostJSON("conversations.open", openBody)
		if err != nil {
			return fmt.Errorf("opening DM: %w", err)
		}

		// Parse channel ID from response
		var openResp struct {
			OK      bool `json:"ok"`
			Channel struct {
				ID string `json:"id"`
			} `json:"channel"`
		}
		if err := json.Unmarshal(openResult, &openResp); err != nil {
			return fmt.Errorf("parsing DM open response: %w", err)
		}
		if !openResp.OK || openResp.Channel.ID == "" {
			return fmt.Errorf("failed to open DM conversation")
		}

		// Send message
		msgPayload := map[string]string{"channel": openResp.Channel.ID, "text": text}
		msgBody, _ := json.Marshal(msgPayload)

		result, err := client.PostJSON("chat.postMessage", msgBody)
		if err != nil {
			return err
		}
		output.RawJSON(result)
		return nil
	},
}

// ============================================================================
// init + main
// ============================================================================

func init() {
	// channel
	channelListCmd.Flags().Int("limit", 100, "Max channels to return")
	channelListCmd.Flags().String("types", "public_channel", "Channel types (public_channel, private_channel, mpim, im)")
	channelInfoCmd.Flags().String("channel", "", "Channel ID")
	channelInfoCmd.MarkFlagRequired("channel")
	channelMembersCmd.Flags().String("channel", "", "Channel ID")
	channelMembersCmd.Flags().Int("limit", 100, "Max members")
	channelMembersCmd.MarkFlagRequired("channel")
	channelCreateCmd.Flags().String("name", "", "Channel name")
	channelCreateCmd.MarkFlagRequired("name")
	channelArchiveCmd.Flags().String("channel", "", "Channel ID")
	channelArchiveCmd.MarkFlagRequired("channel")
	channelUnarchiveCmd.Flags().String("channel", "", "Channel ID")
	channelUnarchiveCmd.MarkFlagRequired("channel")
	channelInviteCmd.Flags().String("channel", "", "Channel ID")
	channelInviteCmd.Flags().String("users", "", "Comma-separated user IDs")
	channelInviteCmd.MarkFlagRequired("channel")
	channelInviteCmd.MarkFlagRequired("users")
	channelKickCmd.Flags().String("channel", "", "Channel ID")
	channelKickCmd.Flags().String("user", "", "User ID")
	channelKickCmd.MarkFlagRequired("channel")
	channelKickCmd.MarkFlagRequired("user")
	channelJoinCmd.Flags().String("channel", "", "Channel ID")
	channelJoinCmd.MarkFlagRequired("channel")
	channelLeaveCmd.Flags().String("channel", "", "Channel ID")
	channelLeaveCmd.MarkFlagRequired("channel")
	channelSetTopicCmd.Flags().String("channel", "", "Channel ID")
	channelSetTopicCmd.Flags().String("topic", "", "Topic text")
	channelSetTopicCmd.MarkFlagRequired("channel")
	channelSetTopicCmd.MarkFlagRequired("topic")
	channelSetPurposeCmd.Flags().String("channel", "", "Channel ID")
	channelSetPurposeCmd.Flags().String("purpose", "", "Purpose text")
	channelSetPurposeCmd.MarkFlagRequired("channel")
	channelSetPurposeCmd.MarkFlagRequired("purpose")
	channelRenameCmd.Flags().String("channel", "", "Channel ID")
	channelRenameCmd.Flags().String("name", "", "New channel name")
	channelRenameCmd.MarkFlagRequired("channel")
	channelRenameCmd.MarkFlagRequired("name")
	channelMarkCmd.Flags().String("channel", "", "Channel ID")
	channelMarkCmd.Flags().String("timestamp", "", "Message timestamp")
	channelMarkCmd.MarkFlagRequired("channel")
	channelMarkCmd.MarkFlagRequired("timestamp")
	channelCmd.AddCommand(channelListCmd, channelInfoCmd, channelMembersCmd,
		channelCreateCmd, channelArchiveCmd, channelUnarchiveCmd,
		channelInviteCmd, channelKickCmd, channelJoinCmd, channelLeaveCmd,
		channelSetTopicCmd, channelSetPurposeCmd, channelRenameCmd, channelMarkCmd)

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
	messageThreadCmd.Flags().String("channel", "", "Channel ID")
	messageThreadCmd.Flags().String("timestamp", "", "Parent message timestamp")
	messageThreadCmd.Flags().Int("limit", 20, "Max replies")
	messageThreadCmd.MarkFlagRequired("channel")
	messageThreadCmd.MarkFlagRequired("timestamp")
	messageUpdateCmd.Flags().String("channel", "", "Channel ID")
	messageUpdateCmd.Flags().String("timestamp", "", "Message timestamp")
	messageUpdateCmd.Flags().String("text", "", "New message text")
	messageUpdateCmd.MarkFlagRequired("channel")
	messageUpdateCmd.MarkFlagRequired("timestamp")
	messageUpdateCmd.MarkFlagRequired("text")
	messageDeleteCmd.Flags().String("channel", "", "Channel ID")
	messageDeleteCmd.Flags().String("timestamp", "", "Message timestamp")
	messageDeleteCmd.MarkFlagRequired("channel")
	messageDeleteCmd.MarkFlagRequired("timestamp")
	messageScheduleCmd.Flags().String("channel", "", "Channel ID")
	messageScheduleCmd.Flags().String("text", "", "Message text")
	messageScheduleCmd.Flags().String("post-at", "", "Unix timestamp for when to send")
	messageScheduleCmd.MarkFlagRequired("channel")
	messageScheduleCmd.MarkFlagRequired("text")
	messageScheduleCmd.MarkFlagRequired("post-at")
	messageCmd.AddCommand(messageSendCmd, messageListCmd, messageReplyCmd, messageThreadCmd,
		messageUpdateCmd, messageDeleteCmd, messageScheduleCmd)

	// user
	userListCmd.Flags().Int("limit", 100, "Max users")
	userInfoCmd.Flags().String("user", "", "User ID")
	userInfoCmd.MarkFlagRequired("user")
	userProfileCmd.Flags().String("user", "", "User ID (omit for self)")
	userCmd.AddCommand(userListCmd, userInfoCmd, userProfileCmd)

	// reaction
	reactionAddCmd.Flags().String("channel", "", "Channel ID")
	reactionAddCmd.Flags().String("timestamp", "", "Message timestamp")
	reactionAddCmd.Flags().String("emoji", "", "Emoji name (without colons)")
	reactionAddCmd.MarkFlagRequired("channel")
	reactionAddCmd.MarkFlagRequired("timestamp")
	reactionAddCmd.MarkFlagRequired("emoji")
	reactionRemoveCmd.Flags().String("channel", "", "Channel ID")
	reactionRemoveCmd.Flags().String("timestamp", "", "Message timestamp")
	reactionRemoveCmd.Flags().String("emoji", "", "Emoji name (without colons)")
	reactionRemoveCmd.MarkFlagRequired("channel")
	reactionRemoveCmd.MarkFlagRequired("timestamp")
	reactionRemoveCmd.MarkFlagRequired("emoji")
	reactionGetCmd.Flags().String("channel", "", "Channel ID")
	reactionGetCmd.Flags().String("timestamp", "", "Message timestamp")
	reactionGetCmd.MarkFlagRequired("channel")
	reactionGetCmd.MarkFlagRequired("timestamp")
	reactionCmd.AddCommand(reactionAddCmd, reactionRemoveCmd, reactionGetCmd)

	// pin
	pinListCmd.Flags().String("channel", "", "Channel ID")
	pinListCmd.MarkFlagRequired("channel")
	pinAddCmd.Flags().String("channel", "", "Channel ID")
	pinAddCmd.Flags().String("timestamp", "", "Message timestamp")
	pinAddCmd.MarkFlagRequired("channel")
	pinAddCmd.MarkFlagRequired("timestamp")
	pinRemoveCmd.Flags().String("channel", "", "Channel ID")
	pinRemoveCmd.Flags().String("timestamp", "", "Message timestamp")
	pinRemoveCmd.MarkFlagRequired("channel")
	pinRemoveCmd.MarkFlagRequired("timestamp")
	pinCmd.AddCommand(pinListCmd, pinAddCmd, pinRemoveCmd)

	// file
	fileListCmd.Flags().String("channel", "", "Filter by channel ID")
	fileListCmd.Flags().Int("limit", 20, "Max files")
	fileInfoCmd.Flags().String("file", "", "File ID")
	fileInfoCmd.MarkFlagRequired("file")
	fileDeleteCmd.Flags().String("file", "", "File ID")
	fileDeleteCmd.MarkFlagRequired("file")
	fileCmd.AddCommand(fileListCmd, fileInfoCmd, fileDeleteCmd)

	// search
	searchMessagesCmd.Flags().String("query", "", "Search query")
	searchMessagesCmd.Flags().Int("count", 20, "Max results")
	searchMessagesCmd.Flags().String("sort", "", "Sort by: score or timestamp")
	searchMessagesCmd.MarkFlagRequired("query")
	searchFilesCmd.Flags().String("query", "", "Search query")
	searchFilesCmd.Flags().Int("count", 20, "Max results")
	searchFilesCmd.MarkFlagRequired("query")
	searchCmd.AddCommand(searchMessagesCmd, searchFilesCmd)

	// reminder
	reminderAddCmd.Flags().String("text", "", "Reminder text")
	reminderAddCmd.Flags().String("time", "", "When to remind (Unix timestamp or natural language)")
	reminderAddCmd.MarkFlagRequired("text")
	reminderAddCmd.MarkFlagRequired("time")
	reminderDeleteCmd.Flags().String("reminder", "", "Reminder ID")
	reminderDeleteCmd.MarkFlagRequired("reminder")
	reminderCompleteCmd.Flags().String("reminder", "", "Reminder ID")
	reminderCompleteCmd.MarkFlagRequired("reminder")
	reminderCmd.AddCommand(reminderAddCmd, reminderListCmd, reminderDeleteCmd, reminderCompleteCmd)

	// bookmark
	bookmarkListCmd.Flags().String("channel", "", "Channel ID")
	bookmarkListCmd.MarkFlagRequired("channel")
	bookmarkAddCmd.Flags().String("channel", "", "Channel ID")
	bookmarkAddCmd.Flags().String("title", "", "Bookmark title")
	bookmarkAddCmd.Flags().String("link", "", "Bookmark URL")
	bookmarkAddCmd.MarkFlagRequired("channel")
	bookmarkAddCmd.MarkFlagRequired("title")
	bookmarkAddCmd.MarkFlagRequired("link")
	bookmarkCmd.AddCommand(bookmarkListCmd, bookmarkAddCmd)

	// dm
	dmOpenCmd.Flags().String("users", "", "Comma-separated user IDs")
	dmOpenCmd.MarkFlagRequired("users")
	dmSendCmd.Flags().String("users", "", "Comma-separated user IDs")
	dmSendCmd.Flags().String("text", "", "Message text")
	dmSendCmd.MarkFlagRequired("users")
	dmSendCmd.MarkFlagRequired("text")
	dmCmd.AddCommand(dmOpenCmd, dmSendCmd)

	rootCmd.AddCommand(channelCmd, messageCmd, userCmd, reactionCmd,
		pinCmd, fileCmd, searchCmd, reminderCmd, bookmarkCmd, dmCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
