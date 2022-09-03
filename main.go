package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"catchMeStore/app/controller"
)

func main() {
	// start the router to serve requests
	r := httprouter.New()
	routes(r)

	err := http.ListenAndServe("localhost:4444", r)
	if err != nil {
		log.Fatal(err)
	}
}

func routes(r *httprouter.Router) {
	r.ServeFiles("/public/*filepatch", http.Dir("public"))
	r.GET("/", controller.StartPage)
	r.GET("/users", controller.GetUsers)
}
