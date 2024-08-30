package gaming_session

var (
	membersSession []string
	mabarSession   bool
)

func CheckJoin(userId string) bool {
	for _, u := range membersSession {
		if u == userId {
			return true
		}
	}
	return false
}
