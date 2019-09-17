package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

type req_struct struct {
	Commands []string
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/commands", commandsHandler).Methods(http.MethodPost)

	log.Println("Server is up!")

	http.ListenAndServe(":8080", r)

}

func commandsHandler(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t req_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	log.Println(t.Commands)

	for i := 0; i < len(t.Commands); i++ {
		out, err := exec.Command("sh", "-c", t.Commands[i]).Output()

		if err != nil {
			fmt.Printf("%s", err)
		}

		fmt.Printf("%s", out)
	}
}
