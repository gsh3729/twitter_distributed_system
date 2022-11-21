package globals

var Secret = []byte("secret")

var UserPass = make(map[string]string)

var Following = make(map[string][]string)

var Followers = make(map[string][]string)

type Tweet struct {
	time int
}

var Tweets = make(map[string][]string)

const Userkey = "user"
