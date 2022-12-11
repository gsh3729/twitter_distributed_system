package findPeople
import (
	// "log"
	// "strings"

	globals "proj/web/globals"
)

func GetPeopleForUser(username string) []string {
	people := []string{}
	for key := range globals.UserPass {
		if key != username { //check the comparison if its working
			people = append(people, key)
		}
	}
	return people
} 