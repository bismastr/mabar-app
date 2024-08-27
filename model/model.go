package model

type CreatedBy struct {
	Id       string `firestore:"id,omitempty"`
	Username string `firestore:"username,omitempty"`
}

type GamingSession struct {
	CreatedAt      string     `firestore:"created_at,omitempty"`
	CreatedBy      *CreatedBy `firestore:"created_by,omitempty"`
	SessionEnd     string     `firestore:"session_end,omitempty"`
	SessionStart   string     `firestore:"session_start,omitempty"`
	MembersSession []string   `firestore:"members_sessions,omitempty"`
	IsFinish       bool
}
