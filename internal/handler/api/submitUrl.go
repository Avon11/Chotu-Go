package api

import (
	"net/http"
)

type PostUrlJson struct {
	Url string	`json:"url"`
}

func PostUrl(w http.ResponseWriter, r *http.Request) {
	// Process the request and Call the service 
	// Get the reponse and write in w
}