package user

type User struct {
	Name       string `json:"name"`
	AvatarURL  string `json:"avatar_url"`
	ID         int64  `json:"id"`
	DiscordUID int64  `json:"discord_uid"`
}
