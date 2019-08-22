package controllers

//This file will handle the routing of http requests that comes into the webserver to the correct method detected
import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/tripleaceinnovations/api/models"
)

// Handles requests on message collection and request to manipulate messages
// Uses regular expression to match incoming http request
type messageController struct {
	messageIDPattern *regexp.Regexp
}

// Creating a function and making it a method (ServeHTTP) by specifying the type which the function binds
func (mc messageController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/v1/messages" {
		switch r.Method {
		case http.MethodGet:
			mc.getAll(w, r)
		case http.MethodPost:
			mc.post(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			log.Println("... HTTP method used is not allowed here ...")
		}
	} else {
		matches := mc.messageIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method {
		case http.MethodGet:
			mc.get(id, w)
		case http.MethodPut:
			mc.put(id, w, r)
		case http.MethodDelete:
			mc.delete(id, w)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			log.Println("... HTTP method used is not allowed ...")
		}

	}

}

//getAll method will retrieves all the messages from our model layer and returning it back out
func (mc *messageController) getAll(w http.ResponseWriter, r *http.Request) {
	log.Println("... getALL: calling GetMessages ...")
	encodeResponseAsJSON(models.GetMessages(), w)
}

//get method will get the id of the resource we gonna be workin with, accept the response writer from the serveHTTP method,
func (mc *messageController) get(id int, w http.ResponseWriter) {
	log.Println("...GET: calling GetMessageByID ...")
	m, err := models.GetMessageByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(m, w)
}

func (mc *messageController) post(w http.ResponseWriter, r *http.Request) {
	m, err := mc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse message object"))
		log.Println("... POST: message parsing error ...")
		return
	}
	log.Println("...calling AddMessage ...")
	m, err = models.AddMessage(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(m, w)
}

func (mc *messageController) put(id int, w http.ResponseWriter, r *http.Request) {
	m, err := mc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse message object"))
		log.Println("... PUT: message parsing error ...")
		return
	}
	if id != m.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted message must match ID in URL"))
		return
	}
	log.Println("...calling UpdateMessage ...")
	m, err = models.UpdateMessage(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(m, w)
}

func (mc *messageController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveMessageByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println("... DELETE: some error in deleting message ...")
		return
	}
	w.WriteHeader(http.StatusOK)
}

//This will take any request coming in e.g update request and convert it into a msg object we can work with
func (mc *messageController) parseRequest(r *http.Request) (models.Phrase, error) {
	dec := json.NewDecoder(r.Body)
	var u models.Phrase
	err := dec.Decode(&u)
	if err != nil {
		log.Println("parseRequest validation error: ", err)
		return models.Phrase{}, err
	}
	return u, nil
}

func newMessageController() *messageController {
	return &messageController{
		messageIDPattern: regexp.MustCompile(`^/api/v1/messages/(\d+)/?`),
	}
}
