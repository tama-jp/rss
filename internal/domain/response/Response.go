package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tama-jp/rss/internal/utils/message"
)

type Response struct {
	C *gin.Context
}

type Payload struct {
	Code    string      `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

func (g *Response) Payload(httpCode int, statusCode string, data interface{}) {
	requestURI := g.C.Request.RequestURI

	if data != nil {
		fmt.Println("RequestURI: "+requestURI, "data", data)
	}

	g.C.JSON(httpCode, Payload{
		Code:    statusCode,
		Data:    data,
		Message: message.GetApiNameAndTime(g.C, ""),
	})
}

func (g *Response) ErrorPayload(httpCode int, statusCode string, err error) {
	requestURI := g.C.Request.RequestURI

	if err != nil {
		fmt.Println("RequestURI: "+requestURI, "error", err)
	}
	msg := err.Error()

	g.C.AbortWithStatusJSON(httpCode, Payload{
		Code:    statusCode,
		Message: message.GetApiNameAndTime(g.C, msg),
	})
}
