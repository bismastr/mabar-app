package model

type GamingSession struct {
	CreatedAt      string   `firestore:"created_at,omitempty"`
	CreadetBy      string   `firestore:"created_by,omitempty"`
	SessionEnd     string   `firestore:"session_end,omitempty"`
	SessionStart   string   `firestore:"session_start,omitempty"`
	MembersSession []string `firestore:"members_sessions,omitempty"`
}
