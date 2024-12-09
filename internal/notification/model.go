package notification

type SendNotificationWithTopicRequest struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	ImageURL string `json:"image_url"`
	Topic    string `json:"topic"`
}
