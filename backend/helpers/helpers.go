package helpers

// import (
// 	globals "backend/globals"
// )

// func CheckUserPass(username, password string) bool {
// 	userpass := globals.UserPass

// 	if val, ok := userpass[username]; ok {
// 		if val == password {
// 			return true
// 		} else {
// 			return false
// 		}
// 	} else {
// 		return false
// 	}
// }

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func RemoveFromSlice(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
