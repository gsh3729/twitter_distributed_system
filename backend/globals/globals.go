package globals

import "time"

type Tweet struct {
	Time time.Time
	Text string
	User string
}

var Endpoints = []string{"localhost:2379", "localhost:3379", "localhost:4379"}
var Timeout = 5 * time.Second

type User struct {
	Username string
	Password string
}
