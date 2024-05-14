package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Avon11/Chotu-Go/internal/service"
)

type PostUrlRequest struct {
	Url string `json:"url"`
}

type PostUrlResponse struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Model interface{} `json:"model"`
}
type PostUrlResponseModel struct {
	Url string `json:"url"`
}

func PostUrl(w http.ResponseWriter, r *http.Request) {
	var urlStruct PostUrlRequest
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&urlStruct)
	if err != nil {
		log.Fatalln("Error !!! ", err)
	}
	_, _ = service.CreateShortUrl(urlStruct.Url)
	respModel := &PostUrlResponseModel{
		Url: urlStruct.Url,
	}
	response := &PostUrlResponse{
		Code:  http.StatusOK,
		Msg:   "success",
		Model: respModel,
	}
	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		log.Fatalln("Error !!! ", err)
	}
}
