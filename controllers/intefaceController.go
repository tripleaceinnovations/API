package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// This interface will handle all the routing in the emtire app
// When network request is received, it will go o the correct controller to be processed

func RegisterControllers() {
	mc := newMessageController()
	http.Handle("/api/v1/messages", *mc)
	http.Handle("/api/v1/messages/", *mc)
	log.Println("... interface handler ...")
}

//takes the user oject and convert it to json representation
func encodeResponseAsJSON(data interface{}, w io.Writer) {
	json.NewEncoder(w).Encode(data)
	log.Println("... JSON Encoding completed ...")
}
