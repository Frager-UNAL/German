package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router *mux.Router
)

func initRouter() {
	println("Initializing router")
	router = mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to FRAGER Admin backend! (made in go)")
	})

	admin_initialize()
	router_initialize()
}

func getRouter() *mux.Router {
	if router == nil {
		initRouter()
	}

	return router
}

func verifyRequest(r *http.Request, values []string) (bool, map[string]interface{}) {
	r.ParseForm()

	exitMap := make(map[string]interface{})
	correct := true

	for _, tag := range values {
		if len(r.Form.Get(tag)) == 0 {
			correct = false
			exitMap[tag] = "Missing value"
		}
	}

	return correct, exitMap
}

func StartServe(port string) error {
	err := http.ListenAndServe(port, getRouter()) //Launching app
	return err
}
