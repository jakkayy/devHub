package domain

import "context"

// DiscordField represents a field inside a Discord Rich Embed
type DiscordField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

// DiscordEmbed represents a Rich Embed payload for Discord Webhook
type DiscordEmbed struct {
	Title       string         `json:"title"`
	Description string         `json:"description,omitempty"`
	URL         string         `json:"url,omitempty"`
	Color       int            `json:"color,omitempty"` // Integer color code (e.g. 0x00FF00 for green)
	Fields      []DiscordField `json:"fields,omitempty"`
	Timestamp   string         `json:"timestamp,omitempty"`
}

// DiscordMessageDTO represents the full payload sent to Discord Webhook
type DiscordMessageDTO struct {
	Username  string         `json:"username,omitempty"`
	AvatarURL string         `json:"avatar_url,omitempty"`
	Content   string         `json:"content,omitempty"`
	Embeds    []DiscordEmbed `json:"embeds,omitempty"`
}

// NotificationEventDTO represents event payload triggered internally or via API
type NotificationEventDTO struct {
	EventType string `json:"event_type" binding:"required"` // e.g. "task_update", "pr_review", "api_contract"
	Title     string `json:"title" binding:"required"`
	Message   string `json:"message" binding:"required"`
	TaskID    string `json:"task_id,omitempty"`
	URL       string `json:"url,omitempty"`
	VSCodeURI string `json:"vscode_uri,omitempty"`
}

// DiscordClientInterface defines contract for sending webhook messages
type DiscordClientInterface interface {
	SendWebhook(ctx context.Context, webhookURL string, msg DiscordMessageDTO) error
}
