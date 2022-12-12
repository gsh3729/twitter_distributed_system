package auth

import (
	globals "backend/globals"
	helpers "backend/helpers"
)

func SignUp(username string, password string) string {
	globals.UserPass[username] = password
	return username
}

func SignIn(username string, password string) bool {
	return helpers.CheckUserPass(username, password)
}
