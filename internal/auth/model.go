package auth

type User struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	UserID    string `json:"user_id"`
}
