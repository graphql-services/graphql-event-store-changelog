package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/graphql-services/go-saga/eventstore"

	"github.com/gorilla/mux"
	"github.com/graphql-services/graphql-event-store-changelog/src"
)

// EventsHandler ...
func EventsHandler(r *mux.Router, db *src.DB) error {

	r.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var e eventstore.Event
		err = json.Unmarshal(body, &e)
		fmt.Println("Event!!", string(body))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = src.ImportEvents([]eventstore.Event{e}, db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}).Methods("POST")

	e := r.PathPrefix("/events").Subrouter()
	e.HandleFunc("/latest", func(w http.ResponseWriter, r *http.Request) {
		event, found, err := src.GetLatestEvent(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !found {
			http.NotFound(w, r)
			return
		}

		if event == nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		data, err := json.Marshal(event)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("content-type", "application/json")
		w.Write(data)
	}).Methods("GET")

	e.HandleFunc("/import", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var e []eventstore.Event
		err = json.Unmarshal(body, &e)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = src.ImportEvents(e, db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}).Methods("POST")

	return nil
}
