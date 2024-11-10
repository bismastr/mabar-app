package gamingSession

type GamingSession struct {
	CreatedAt      string     `firestore:"created_at,omitempty" json:"created_at,omitempty"`
	CreatedBy      *CreatedBy `firestore:"created_by,omitempty" json:"created_by,omitempty"`
	SessionEnd     string     `firestore:"session_end,omitempty" json:"session_end,omitempty"`
	SessionStart   string     `firestore:"session_start,omitempty" json:"session_start,omitempty"`
	MembersSession []string   `firestore:"members_sessions,omitempty" json:"members_sessions,omitempty"`
	GameName       string     `firestore:"game_name,omitempty" json:"game_name,omitempty"`
	IsFinish       bool       `json:"is_finish,omitempty"`
}

type CreatedBy struct {
	Id       string `firestore:"id,omitempty" json:"id,omitempty"`
	Username string `firestore:"username,omitempty" json:"username,omitempty"`
}
