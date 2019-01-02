package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/micnncim/mercari-datastore-sample/handler"
)

func main() {
	r := chi.NewRouter()

	r.Post("/CreateUser", handler.CreateUser)
	r.Post("/ListUsers", handler.ListUsers)
	r.Post("/UpdateUser", handler.UpdateUser)
	r.Post("/DeleteUser", handler.DeleteUser)

	fmt.Println("starting server :8080...")
	http.ListenAndServe(":8080", r)
}
