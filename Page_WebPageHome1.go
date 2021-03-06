package main

import (
	"fmt"
	"net/http"
)

func homeFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Home</h1>")
	fmt.Fprint(w, "<B>Kalyan </B>")
}
func aboutusFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>AboutUs</h1>")
}
func loginFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Login</h1>")
}
func logoutFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>logout</h1>")
}
func saveFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>File Has been saved </h1>")
}
func delFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>File Has been Deleted </h1>")
}
func dataFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Get Data from Db </h1>")
}
func fbFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>FB link</h1>")
}
func tFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>T link</h1>")
}



func main() {
	http.HandleFunc("/home", homeFunc)
	http.HandleFunc("/aboutus", aboutusFunc)
	http.HandleFunc("/login", loginFunc)
	http.HandleFunc("/logout", logoutFunc)
	http.HandleFunc("/save", saveFunc)
	http.HandleFunc("/del", delFunc)
	http.HandleFunc("/data", dataFunc)
	http.HandleFunc("/fb", fbFunc)
	http.HandleFunc("/t", tFunc)
	http.ListenAndServe(":3000", nil)
}
