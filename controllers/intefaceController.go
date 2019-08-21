package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	mc := newMessageController() //creates a new user controller via the constructor controller created earlier
	//pc := newPalindromeController()
	http.Handle("/api/v1/messages", *mc)  //handles any request from api/v1/users route and uses the controler obj
	http.Handle("/api/v1/messages/", *mc) //handles any request from api/v1/users/ route and uses the controler obj
	//http.Handle("/api/v1/palindrome", *pc)  //handles any request from api/v1/palindrome route and uses the controler obj
	//http.Handle("/api/v1/palindrome/", *pc) //handles any request from api/v1/palindrome/ route and uses the controler obj
}

//takes the user oject and convert it to json representation
//it reaches into the json package imported, create an encoder to encode go obj into json representation and call the encode method on the encoder passing in whatever data we received
func encodeResponseAsJSON(data interface{}, w io.Writer) {
	//enc := json.NewEncoder(w)
	//enc.Encode(data)
	json.NewEncoder(w).Encode(data)
}

/* This interface will handle all the routing in the emtire app
when network request is received, it will go o the correct controller to be processed
*/
