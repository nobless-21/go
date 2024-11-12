package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You see user list\n")
}

func ListVK(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "You see user list for subdomain %s\n", vars["subdomain"])
}

func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "you try to see user %s\n", vars["id"])
}

/*
curl -v -X PUT -H "Content-Type: application/json" -d '{"login":"rvasily"}' http://localhost:8080/users
*/

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you try to create new user\n")
}

/*
curl -v -X POST -H "Content-Type: application/json"  -H "X-Auth: test" -d '{"name":"Vasily Romanov"}' http://localhost:8080/users/rvasily
*/

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "you try to update %s\n", vars["login"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", List)

	r.HandleFunc("/users", List).
		Host("localhost")

	r.HandleFunc("/users", ListVK).
		Host("{subdomain}.vk.ru")

	r.HandleFunc("/users", Update).
		Methods("PUT")

		/*
		   See anything wrong?
		*/
	r.HandleFunc("/users/{login:[0-9a-z]+}", Update)
	r.HandleFunc("/users/{id:[0-9]+}", Get)

	r.HandleFunc("/users/{login}", Create).
		Methods("POST").
		Headers("X-Auth", "test")

	fmt.Println("starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
