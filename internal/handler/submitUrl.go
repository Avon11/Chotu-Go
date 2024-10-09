package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

func (h *Handler) PostUrl(c *gin.Context) {
	var urlStruct PostUrlRequest

	if err := c.ShouldBindJSON(&urlStruct); err != nil {
		log.Println("Error while decoding request: ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "invalid request",
		})
		return
	}

	code := http.StatusOK
	msg := "success"
	respModel := &PostUrlResponseModel{}

	shortUrl, errResp := h.Service.CreateShortUrl(c, urlStruct.Url)
	if errResp != nil {
		code = http.StatusInternalServerError
		msg = "internal server error"
	} else {
		respModel = &PostUrlResponseModel{
			Url: shortUrl.Url,
		}
	}

	response := &PostUrlResponse{
		Code:  code,
		Msg:   msg,
		Model: respModel,
	}

	c.JSON(code, response)
}
