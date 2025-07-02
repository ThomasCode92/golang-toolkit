package main

import (
	"log"
	"net/http"

	"github.com/ThomasCode92/toolkit"
)

type RequestPayload struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

type ResponsePayload struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message"`
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/receive-post", receivePost)
	mux.HandleFunc("/remote-server", remoteServer)

	log.Println("Starting service...")

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func receivePost(w http.ResponseWriter, r *http.Request) {
	var reqPayload RequestPayload
	var t toolkit.Tools

	err := t.ReadJSON(w, r, &reqPayload)
	if err != nil {
		t.ErrorJSON(w, err)
		return
	}

	responsePayload := ResponsePayload{
		Message: "hit the handler okay, and sending response",
	}

	err = t.WriteJSON(w, http.StatusAccepted, responsePayload)
	if err != nil {
		log.Println(err)
	}
}

func remoteServer(w http.ResponseWriter, r *http.Request) {
}
