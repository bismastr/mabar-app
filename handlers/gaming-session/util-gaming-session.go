package gaming_session

import (
	"fmt"
)

var (
	membersSession []string
	mabarSession   bool
)

func GenerateMemberMention(members []string) string {
	result := ""
	for _, s := range members {
		result += fmt.Sprintf("<@%v>", s)
	}
	return result
}

func CheckJoin(userId string) bool {
	for _, u := range membersSession {
		if u == userId {
			return true
		}
	}
	return false
}
