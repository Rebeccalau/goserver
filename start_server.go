package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	myconfig "start_server/config"
	"start_server/state"
)

type StateRequest struct {
	RequestType string
}

type StateHandler struct {
	Config       myconfig.Configuration
	StateMachine state.StateMachine
}

func (stateArgs *StateHandler) change(w http.ResponseWriter, request *http.Request) {
	var parsedRequest StateRequest

	switch request.Method {
	case "GET":
		err := json.NewDecoder(request.Body).Decode(&parsedRequest)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		response := stateArgs.StateMachine.ProcessStateChange(parsedRequest.RequestType)

		w.Write([]byte(response))

	default:
		fmt.Println("Unsupported Request Method", request.Method)
	}
}

func (stateArgs *StateHandler) current(w http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case "GET":

		response := stateArgs.StateMachine.CurrentState.Name
		w.Write([]byte(response))

	default:
		fmt.Println("Unsupported Request Method", request.Method)
	}
}

func headers(_ http.ResponseWriter, request *http.Request) {
	for name, headers := range request.Header {
		for _, h := range headers {
			fmt.Println("\n", name, h)
		}
	}
}

func main() {
	arguments := os.Args[1]
	StartHttpServer(arguments)
}

func StartHttpServer(arguments string) {
	config := myconfig.NewConfiguration(arguments)
	stateMachine := state.NewStateMachine(config)
	stateRoute := &StateHandler{Config: config, StateMachine: stateMachine}

	http.HandleFunc("/state", stateRoute.change)
	http.HandleFunc("/current", stateRoute.current)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8080", nil)
}
