package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tama-jp/rss/internal/domain/response"
	"github.com/tama-jp/rss/internal/usecases/interactor"
)

type AccessTokenController struct {
	Interactor *interactor.AccessTokenInteractor
}

func NewAccessTokenController(interactor *interactor.AccessTokenInteractor) *AccessTokenController {
	return &AccessTokenController{Interactor: interactor}
}

func (controller *AccessTokenController) GetAccessToken(c *gin.Context) {
	controller.Interactor.PrintInfo("Start", "AccessTokenController:GetAccessToken", c.Request.RequestURI)
	fmt.Println(c.Request.RequestURI, "Start", "AccessTokenController:GetAccessToken")

	res := response.Response{C: c}

	status, statusCode, data, err := controller.Interactor.GetAccessToken(c)

	if err != nil {
		fmt.Println("err", err)
		res.ErrorPayload(status, statusCode, err)
		return
	} else {
		fmt.Println("OK", data)
		res.Payload(status, statusCode, data)
	}

	out, _ := json.Marshal(data)
	controller.Interactor.PrintInfo("End", "AccessTokenController:GetAccessToken", string(out))
	fmt.Println(c.Request.RequestURI, "End", "AccessTokenController:GetAccessToken")
}

func (controller *AccessTokenController) PutAccessToken(c *gin.Context) {
	controller.Interactor.PrintInfo("Start", "AccessTokenController:PutAccessToken", c.Request.RequestURI)
	fmt.Println(c.Request.RequestURI, "Start", "AccessTokenController:PutAccessToken")

	res := response.Response{C: c}

	status, statusCode, data, err := controller.Interactor.PutAccessToken(c)

	if err != nil {
		fmt.Println("err", err)
		res.ErrorPayload(status, statusCode, err)

		return
	} else {

		fmt.Println("OK", data)
		res.Payload(status, statusCode, data)
	}
	out, _ := json.Marshal(data)
	controller.Interactor.PrintInfo("End", "AccessTokenController:PutAccessToken", string(out))
	fmt.Println(c.Request.RequestURI, "End", "AccessTokenController:PutAccessToken")
}

func (controller *AccessTokenController) DeleteAccessToken(c *gin.Context) {
	controller.Interactor.PrintInfo("Start", "AccessTokenController:DeleteAccessToken", c.Request.RequestURI)
	fmt.Println(c.Request.RequestURI, "Start", "AccessTokenController:DeleteAccessToken")

	res := response.Response{C: c}

	status, statusCode, data, err := controller.Interactor.DeleteAccessToken(c)

	if err != nil {
		fmt.Println("err", err)
		res.ErrorPayload(status, statusCode, err)

		return
	} else {

		fmt.Println("OK", data)
		res.Payload(status, statusCode, data)
	}

	out, _ := json.Marshal(data)
	controller.Interactor.PrintInfo("End", "AccessTokenController:DeleteAccessToken", string(out))
	fmt.Println(c.Request.RequestURI, "End", "AccessTokenController:DeleteAccessToken")
}
