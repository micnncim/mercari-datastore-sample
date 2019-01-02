package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/micnncim/mercari-datastore-sample/datastore"
	"github.com/micnncim/mercari-datastore-sample/entity"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	u := new(entity.User)
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		fmt.Fprintf(w, "failed to decode json: %s", err.Error())
		return
	}
	defer r.Body.Close()

	ctx := context.Background()
	c, err := datastore.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if err := c.CreateUser(ctx, u); err != nil {
		fmt.Fprintf(w, "failed to create user: %s", err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	c, err := datastore.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	us, err := c.ListUsers(ctx)
	if err != nil {
		fmt.Fprintf(w, "failed to list users: %s", err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(us); err != nil {
		fmt.Fprintf(w, "failed to encode json: %s", err.Error())
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	u := new(entity.User)
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		fmt.Fprintf(w, "failed to decode json: %s", err.Error())
		return
	}
	defer r.Body.Close()

	ctx := context.Background()
	c, err := datastore.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	u, err = c.UpdateUser(ctx, u)
	if err != nil {
		fmt.Fprintf(w, "failed to create user: %s", err.Error())
		return
	}
	if err := json.NewEncoder(w).Encode(u); err != nil {
		fmt.Fprintf(w, "failed to encode json: %s", err.Error())
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	u := new(entity.User)
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		fmt.Fprintf(w, "failed to decode json: %s", err.Error())
		return
	}
	defer r.Body.Close()

	ctx := context.Background()
	c, err := datastore.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if err := c.DeleteUser(ctx, u); err != nil {
		fmt.Fprintf(w, "failed to create user: %s", err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
