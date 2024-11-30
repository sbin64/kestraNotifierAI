package handlers

import (
	"fmt"
	"gofr-bot/utils"

	"gofr.dev/pkg/gofr"
)

// GeneratePostHandler drafts a social media post
func GeneratePostHandler(c *gofr.Context) (interface{}, error) {
	kestraUpdates, err := utils.FetchkestraUpdates(c)  // Pull kestra updates

	if err != nil {
		return "", fmt.Errorf("failed to fetch updates: %w", err)
	}

	post := utils.GeneratePost(c, kestraUpdates) // AI-generated post

	// Save post draft to Redis for review
	utils.SaveDraftToRedis(c, "social_post", post)
	return map[string]string{"draft": post}, nil
}

// ApprovePostHandler publishes an approved post
func ApprovePostHandler(c *gofr.Context) (interface{}, error) {
	post, err := utils.GetDraftFromRedis(c, "social_post")
	if err != nil {
		return nil, err
	}

	err = utils.PublishToSocialMedia(c, post) // Publish to LinkedIn/Twitter
	if err != nil {
		return nil, err
	}

	return map[string]string{"status": "Post published successfully"}, nil
}
