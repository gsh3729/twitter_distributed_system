package globals

import "time"

var UserPass = make(map[string]string)

var Following = make(map[string][]string)

var Followers = make(map[string][]string)

type Tweet struct {
	Time time.Time
	Text string
	User string
}

var Tweets = make(map[string][]Tweet)

var Endpoints = []string{"localhost:2379", "localhost:3379", "localhost:4379"}
var Timeout = 5 * time.Second

type User struct {
	Username string
	Password string
}
