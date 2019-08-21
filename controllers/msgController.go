package controllers

//This file will handle the routing of http requests that comes into the webserver to the correct method detected
import (
	"encoding/json"
	"net/http" //using to pull in the response writer and request object
	"regexp"
	"strconv"

	"github.com/tripleaceinnovations/api/models"
)

type messageController struct { //custome type that I can add a method to
	//Will have some routing responsiilities. It will handles two types of requests
	//1.) resource requests on the users collection
	//2.) request to manipulate individual users.
	// So in order to discern which type of request it will be working with, we need to use regular expression to match the incoming http request
	messageIDPattern *regexp.Regexp //defines userIDPattern as regular exp from the *regexp package. remember to add "regexp" in the import block
}

//Below, I am creating a function and making it a method by specifying the type which I wanna bind the function to
//(uc userController) - binding uc, a local var name to a userController type
//ServeHTTP - method name
//(w http.ResponseWriter, r *http.Request) - has a signature that require response writer obj from the http package and also a requst obj. So this works with the info coming from  the http request directly
// this fn will recieve http request in, decide which method to pass the request to for it to be processed
//First we need to understand if we are dealing with entire users collection or a single user object
func (mc messageController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hello from the User Controller")) //writing back to the http respose obj with a slice of byte and doing a string conversion to display a string
	if r.URL.Path == "/api/v1/messages" { //check if the URL.Path of the request object is /api/v1/users; if it is, then we are dealing with entire users collection
		//two things cal happen when dealing with the entire user collection: 1. return all of the users from the collectn back to the requester;
		// 2. I am going to use a POST on the user collection to show that someone is adding a user into the collection so we use a switch
		switch r.Method { //checks the method of the incoming request
		case http.MethodGet:
			mc.getAll(w, r)
		case http.MethodPost:
			mc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := mc.messageIDPattern.FindStringSubmatch(r.URL.Path) //compares incoming info (r.URL.Path) to the regular expression uc.userIDPattern
		//if it finds a match, it will populate the matches var with a slice containimg all of the matches.
		//In the regular exp, we have '\d+' that will define a sub-group for api/v1/users/ that is going to be populated if we have a match to the regular expression
		if len(matches) == 0 { // checking if the length matches the slice.. i.e do we have a ID for our regular expression? bcos if someone request /users/harey, our id is expected to be nos and not string
			w.WriteHeader(http.StatusNotFound)
		}
		id, err := strconv.Atoi(matches[1]) //converts the string response (regular exp are made of strings) to a numerial data type for go to work with it
		//even though we have match and we have a number, based on the reg exp that we define, we still have to convert to nuemerical data type for Go to work with it
		//We are passing the index 1 of the matches colection and that will be the subgroup match containing the id value
		if err != nil { //we are checking if the convesrion above fails and we return StatusNotFound if that is the case.
			w.WriteHeader(http.StatusNotFound)
		}
		switch r.Method { //request coming to switch will support 3 verbs - Get, Put & Delete to allows us to retrieve, update and remove a user from the collection.
		case http.MethodGet:
			mc.get(id, w)
		case http.MethodPut:
			mc.put(id, w, r)
		case http.MethodDelete:
			mc.delete(id, w)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}

	}

}

//getAll method will retrieves all the users from our model layer and returning it back out
func (mc *messageController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetMessages(), w)
}

//get method will get the id of the resource we gonna be workin with, accept the response writer from the serveHTTP method,
func (mc *messageController) get(id int, w http.ResponseWriter) {
	m, err := models.GetMessageByID(id) //call into the model layer, retrieve the user by user id
	if err != nil {                     //if doesnt find the user id, it returns an error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(m, w) //if does find a user id, it will take the user obj and return it out to the requester
}

func (mc *messageController) post(w http.ResponseWriter, r *http.Request) {
	m, err := mc.parseRequest(r)
	if err != nil { //if doesnt find the user id, it returns an error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	m, err = models.AddMessage(m)
	if err != nil { //if doesnt find the user id, it returns an error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(m, w)
}

func (mc *messageController) put(id int, w http.ResponseWriter, r *http.Request) {
	m, err := mc.parseRequest(r)
	if err != nil { //if doesnt find the user id, it returns an error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	if id != m.ID { //if doesnt find the user id, it returns an error
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}
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
	if err != nil { //if doesnt find the user id, it returns an error
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

//This will take any request coming in e.g update request and convert it into a user object we can work with
func (mc *messageController) parseRequest(r *http.Request) (models.Phrase, error) {
	dec := json.NewDecoder(r.Body)
	var u models.Phrase
	err := dec.Decode(&u)
	if err != nil {
		return models.Phrase{}, err
	}
	return u, nil
}

//create a constrcution func. Convention - newTypeOfObjToBeConstrcuted
//we return a pointer. We can return the value itself but since the constructor fn is setting up an obj for somebody else to use, we are going to return pointers out so we avod unncessary copy operation

func newMessageController() *messageController {
	return &messageController{ //with struct, we can do this. this is a local var and Go will promote it to whatever level it need to be so we can return the address of the local var and not lose it.
		messageIDPattern: regexp.MustCompile(`^/api/v1/messages/(\d+)/?`),
		//looking for api/v1/users/ followed by a no. So when we have that, we rae going to be dealing with a specific resource that we can manipulate

	}
}
