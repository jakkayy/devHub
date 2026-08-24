package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jakkayy/devHub/services/backend-api/internal/domain"
)

type DiscordClient struct {
	httpClient *http.Client
}

func NewDiscordClient() *DiscordClient {
	return &DiscordClient{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// SendWebhook posts a Rich Embed message payload to a Discord Webhook URL
func (c *DiscordClient) SendWebhook(ctx context.Context, webhookURL string, msg domain.DiscordMessageDTO) error {
	if webhookURL == "" {
		return fmt.Errorf("discord webhook URL is empty")
	}

	payload, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal discord message payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, webhookURL, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create http request for discord webhook: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send discord webhook request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("discord webhook API error: status code %d", resp.StatusCode)
	}

	return nil
}
