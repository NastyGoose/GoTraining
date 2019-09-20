package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

type Array_Request struct {
	Commands []string `json:"commands"`
}

type String_Request struct {
	Command string `json:"commands"`
}

// CommandsHandler is a function for executing commands passed in request
func CommandsHandler(w http.ResponseWriter, req *http.Request) {
	var Commands Array_Request
	var Command String_Request

	body, err1 := ioutil.ReadAll(req.Body)

	if err1 != nil {
		panic(err1)
	}

	err := json.Unmarshal(body, &Commands)
	if err != nil {
		err2 := json.Unmarshal(body, &Command)

		if err2 != nil {
			log.Println(err2)
			panic(err2)
		}

		executeCommand(Command.Command)

	}

	for i := 0; i < len(Commands.Commands); i++ {
		executeCommand(Commands.Commands[i])
	}

}

func executeCommand(command string) {
	if CheckCommand(command) == true {
		out, err := exec.Command("sh", "-c", command).Output()

		if err != nil {
			fmt.Printf("%s", err)
		}

		fmt.Printf("%s", out)

	} else {
		log.Println("Passed command is invalid and will not be executed")
	}
}
