package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type GetRedirectUrlResponse struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Model interface{} `json:"model"`
}

type GetRedirectUrl struct {
	ShortCode string
	Url       string
}

type GetUrlResponseModel struct {
	ShortCode string `json:"shortCode"`
	Url       string `json:"url"`
}

func (h *Handler) GetUrl(c *gin.Context) {
	shortCode := cast.ToString(c.Query("code"))
	if len(shortCode) <= 0 {
		c.JSON(200, GetRedirectUrlResponse{
			Code: 400,
			Msg:  "invalid code",
		})
		return
	}

	shortUrl, errResp := h.Service.RedirectUrl(c, shortCode)
	if errResp != nil {
		response := &GetRedirectUrlResponse{
			Code: errResp.ErrCode,
			Msg:  errResp.ErrMsg,
		}
		c.JSON(200, response)
		return
	}

	resp := &GetUrlResponseModel{
		ShortCode: shortUrl.ShortCode,
		Url:       shortUrl.Url,
	}

	response := &GetRedirectUrlResponse{
		Code:  shortUrl.Code,
		Msg:   "success",
		Model: resp,
	}
	c.JSON(200, response)
}
