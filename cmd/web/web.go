// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"strings"
//     "html/template"
// 	"github.com/gin-contrib/sessions"
// 	"github.com/gin-contrib/sessions/cookie"
// 	"github.com/gin-gonic/gin"
// )

// var Secret = []byte("secret")

// const Userkey = "user"


// func sayhelloName(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()       // parse arguments, you have to call this by yourself
// 	fmt.Println(r.Form) // print form information in server side
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}
// 	fmt.Fprintf(w, "Hello astaxie!") // send data to client side
// }

// func login(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()       // parse arguments, you have to call this by yourself
// 	fmt.Println(r.Form) // print form information in server side
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}
// 	fmt.Fprintf(w, "Login!") // send data to client side
// }

// func signup(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()       // parse arguments, you have to call this by yourself
// 	fmt.Println(r.Form) // print form information in server side
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}
// 	fmt.Fprintf(w, "Signup") // send data to client side
// }

// func home(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()       // parse arguments, you have to call this by yourself
// 	fmt.Println(r.Form) // print form information in server side
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}
// 	fmt.Fprintf(w, "Home") // send data to client side
// }
// type ProfileDetails struct {
//     ProfileName string
// }

// func profile(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()       // parse arguments, you have to call this by yourself
// 	fmt.Println(r.Form) // print form information in server side
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}
//     userId := r.URL.Path[len("/profile/"):]
//     tmplt, _ := template.ParseFiles("../../templates/profile.html")
//     event := ProfileDetails{
//         ProfileName: "userId",
//     }
//     err := tmplt.Execute(w, event)
//     if err != nil {
//         return
//     }
// 	fmt.Fprintf(w, "Profile %s!", userId) // send data to client side
// }

// func composeTweet(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()       // parse arguments, you have to call this by yourself
// 	fmt.Println(r.Form) // print form information in server side
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}
// 	fmt.Fprintf(w, "Compose tweet") // send data to client side
// }

// func connectPeople(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()       // parse arguments, you have to call this by yourself
// 	fmt.Println(r.Form) // print form information in server side
// 	fmt.Println("path", r.URL.Path)
// 	fmt.Println("scheme", r.URL.Scheme)
// 	fmt.Println(r.Form["url_long"])
// 	for k, v := range r.Form {
// 		fmt.Println("key:", k)
// 		fmt.Println("val:", strings.Join(v, ""))
// 	}
// 	fmt.Fprintf(w, "Connect people") // send data to client side
// }

// func main() {
// 	http.HandleFunc("/", sayhelloName) // set router
// 	http.HandleFunc("/login/", login)
// 	http.HandleFunc("/signup/", signup)
// 	http.HandleFunc("/home/", home)
// 	http.HandleFunc("/profile/", profile)
// 	http.HandleFunc("/compose/tweet/", composeTweet)
// 	http.HandleFunc("/connect/people/", connectPeople)
// 	err := http.ListenAndServe(":9090", nil) // set listen port
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }

package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	//"html/template"
	//"strings"

	// globals "web/globals"
	globals "proj/web/globals"
	// middleware "web/auth"
	// routes "web/routes"
)

func main() {
	router := gin.Default()

	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

	public := router.Group("/")
	router.PublicRoutes(public)

	private := router.Group("/")
	private.Use(middleware.AuthRequired)
	routes.PrivateRoutes(private)

	router.Run("localhost:8080")
}