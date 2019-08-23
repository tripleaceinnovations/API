package main

import (
	"net/http"

	"log"

	"github.com/tripleaceinnovations/api/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
	log.Println("... listening on port 3000 ...")
}
