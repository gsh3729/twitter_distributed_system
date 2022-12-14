package auth

import (
	"log"
	globals "backend/globals"
	// helpers "backend/helpers"
)

func SignUp(username string, password string) string {
	globals.UserPass[username] = password
	return username
}

func SignIn(username string, password string) bool {
	return CheckUserPass(username, password)
}

func CheckUserPass(username, password string) bool {
	userpass := globals.UserPass

	log.Println("checkUserPass", username, password, userpass)

	if val, ok := userpass[username]; ok {
		log.Println(val, ok)
		if val == password {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}