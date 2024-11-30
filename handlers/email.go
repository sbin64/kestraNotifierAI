package handlers

import (
	"fmt"
	"gofr-bot/utils"

	"gofr.dev/pkg/gofr"
)

// EmailOutreachHandler drafts and sends outreach emails
func EmailOutreachHandler(c *gofr.Context) (interface{}, error) {
	recipients := utils.FetchTargetRecipients(c)   // Fetch target emails
	kestrapost, err := utils.GetDraftFromRedis(c, "social_post")
	if err != nil {
		return nil, err
	}

	if err != nil {
		return "", fmt.Errorf("failed to fetch updates: %w", err)
	}

	for _, recipient := range recipients {
		email := utils.GenerateEmail(recipient, kestrapost)
		var subject = "Streamline Your Microservices with Kestra - IaC with plugins"
		err := utils.SendEmail(c, subject, recipient, email)
		if err != nil {
			return nil, err
		}
	}

	return map[string]string{"status": "Emails sent successfully"}, nil
}
