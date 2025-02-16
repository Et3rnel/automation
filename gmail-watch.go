package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func callGmailWatch(config *Config) error {
	ctx := context.Background()
	// Ensure that GCP authentication is properly configured (via GOOGLE_APPLICATION_CREDENTIALS or gcloud auth)
	srv, err := gmail.NewService(ctx, option.WithScopes(gmail.MailGoogleComScope))
	if err != nil {
		return fmt.Errorf("failed to create Gmail client: %v", err)
	}

	watchReq := &gmail.WatchRequest{
		LabelIds:  []string{"INBOX"},
		TopicName: fmt.Sprintf("projects/%s/topics/%s", config.GCPProjectId, TopicName),
	}

	res, err := srv.Users.Watch("me", watchReq).Do()
	if err != nil {
		return fmt.Errorf("error calling watch: %v", err)
	}

	log.Printf("Watch successfully activated, historyId: %v, expiration: %v", res.HistoryId, res.Expiration)
	return nil
}
