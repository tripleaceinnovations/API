package main

import (
	"net/http"

	"github.com/tripleaceinnovations/learngo/controllers"
	//"fmt"
	//"github.com/tripleaceinnovations/learngo/models"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
	//ip/port and serve multiplexer which will handle the request coming in ie.e front controller and decide which back controller(our user controller) handles the request via the handle methods we have in RegisterController()
	// nil means we say the default is ok for us. Basically, this is our http server
}
