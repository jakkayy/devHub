package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/jakkayy/devHub/services/backend-api/internal/config"
	"github.com/jakkayy/devHub/services/backend-api/internal/domain"
)

type DiscordUsecase struct {
	discordClient domain.DiscordClientInterface
	cfg           *config.Config
}

func NewDiscordUsecase(discordClient domain.DiscordClientInterface, cfg *config.Config) *DiscordUsecase {
	return &DiscordUsecase{
		discordClient: discordClient,
		cfg:           cfg,
	}
}

// SendEventNotification processes events and dispatches Rich Embed to Discord Webhook
func (u *DiscordUsecase) SendEventNotification(ctx context.Context, event domain.NotificationEventDTO) error {
	webhookURL := u.cfg.DiscordWebhookURL
	if webhookURL == "" {
		// Mock fallback logging if webhook URL is not set in environment
		fmt.Printf("[DiscordUsecase] Webhook URL not set. Mock Notification Triggered: %s - %s\n", event.Title, event.Message)
		return nil
	}

	var embedColor int
	switch event.EventType {
	case "task_update":
		embedColor = 0x10B981 // Emerald Green
	case "pr_review":
		embedColor = 0xA855F7 // Purple
	case "api_contract":
		embedColor = 0x0EA5E9 // Sky Blue
	default:
		embedColor = 0x6366F1 // Indigo
	}

	fields := []domain.DiscordField{
		{Name: "Event Type", Value: event.EventType, Inline: true},
	}

	if event.TaskID != "" {
		fields = append(fields, domain.DiscordField{Name: "Task ID", Value: event.TaskID, Inline: true})
	}

	if event.VSCodeURI != "" {
		fields = append(fields, domain.DiscordField{Name: "IDE Link", Value: fmt.Sprintf("[Open in VS Code](%s)", event.VSCodeURI), Inline: false})
	}

	embed := domain.DiscordEmbed{
		Title:       fmt.Sprintf("🔔 %s", event.Title),
		Description: event.Message,
		URL:         event.URL,
		Color:       embedColor,
		Fields:      fields,
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	msg := domain.DiscordMessageDTO{
		Username:  "devHub Bot",
		AvatarURL: "https://raw.githubusercontent.com/github/explore/main/topics/terminal/terminal.png",
		Embeds:    []domain.DiscordEmbed{embed},
	}

	return u.discordClient.SendWebhook(ctx, webhookURL, msg)
}
