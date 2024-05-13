package apiHandler

import (
	"net/http"

	"github.com/Avon11/Chotu-Go/internal/handler/api"
	"github.com/gorilla/mux"
)

func Handler() {
	r := mux.NewRouter()

	// Add all Urls
	r.HandleFunc("/post-url",api.PostUrl) // Add method

	http.Handle("/",r)
	http.ListenAndServe(":3000", nil)
}