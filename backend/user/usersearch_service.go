package user

import (
	globals "backend/globals"
)

func GetPeopleForUser(username string) []string {
	people := []string{}
	for key := range globals.UserPass {
		if key != username {
			people = append(people, key)
		}
	}
	return people
}
