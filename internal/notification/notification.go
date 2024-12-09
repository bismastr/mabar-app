package notification

import (
	"context"
	"fmt"

	"firebase.google.com/go/messaging"
)

type NotificationService struct {
	messaging *messaging.Client
}

func NewNotificationClient(messaging *messaging.Client) *NotificationService {
	return &NotificationService{
		messaging: messaging,
	}
}

func (n *NotificationService) SendMessage(ctx context.Context, req *SendNotificationWithTopicRequest) error {
	_, err := n.messaging.Send(ctx, &messaging.Message{
		Notification: &messaging.Notification{
			Title:    req.Title,
			Body:     req.Body,
			ImageURL: req.ImageURL,
		},
		Topic: req.Topic,
	})

	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}
