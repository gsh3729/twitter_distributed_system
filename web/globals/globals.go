package globals

import "time"

var Secret = []byte("secret")

var UserPass = make(map[string]string)

var Following = make(map[string][]string)

var Followers = make(map[string][]string)

type Tweet struct {
	Time time.Time
	Text string
	User string
}

var Tweets = make(map[string][]Tweet)

const Userkey = "user"
