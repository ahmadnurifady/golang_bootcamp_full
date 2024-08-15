package testhttp

import (
	"net/http"
)

func HandlerSimpleTest(w http.ResponseWriter, r *http.Request) {

	greeting := "testing"

	w.Write([]byte(greeting))
}

func HandlerGreeting(w http.ResponseWriter, r *http.Request) {

	greeting := "Hello Nigga"

	w.Write([]byte(greeting))
}

func SimpleHttp() {
	http.HandleFunc("/simple", HandlerSimpleTest)
	http.HandleFunc("/greeting", HandlerGreeting)
	// http.ListenAndServe(":8080", nil)

	server := new(http.Server)
	server.Addr = ":8080"

	server.ListenAndServe()

}
