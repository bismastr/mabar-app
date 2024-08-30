package utils

import "fmt"

func GenerateMemberMention(members []string) string {
	result := ""
	for _, s := range members {
		result += fmt.Sprintf("<@%v>", s)
	}
	return result
}
