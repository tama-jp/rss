package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tama-jp/rss/internal/domain/response"
	"github.com/tama-jp/rss/internal/usecases/interactor"

	"encoding/json"
)

type UserController struct {
	Interactor *interactor.UserInteractor
}

func NewUserController(interactor *interactor.UserInteractor) *UserController {
	return &UserController{Interactor: interactor}
}

func (controller *UserController) Search(c *gin.Context) {
	controller.Interactor.PrintInfo("Start", "UserController:Search", c.Request.RequestURI)
	fmt.Println(c.Request.RequestURI, "Start", "UserController:Search")

	res := response.Response{C: c}

	status, statusCode, data, err := controller.Interactor.Search(c)

	if err != nil {
		fmt.Println("error")
		res.ErrorPayload(status, statusCode, err)

		return
	} else {
		fmt.Println("OK", data)

		res.Payload(status, statusCode, data)
	}

	out, _ := json.Marshal(data)
	controller.Interactor.PrintInfo("End", "UserController:Search", string(out))
	fmt.Println(c.Request.RequestURI, "End", "UserController:Search")

}

func (controller *UserController) PasswordChange(c *gin.Context) {
	controller.Interactor.PrintInfo("Start", "UserController:PasswordChange", c.Request.RequestURI)
	fmt.Println(c.Request.RequestURI, "Start", "UserController:PasswordChange")

	res := response.Response{C: c}

	status, statusCode, data, err := controller.Interactor.PasswordChange(c)

	if err != nil {
		fmt.Println("error")
		res.ErrorPayload(status, statusCode, err)

		return
	} else {
		fmt.Println("OK", data)

		res.Payload(status, statusCode, data)

	}

	out, _ := json.Marshal(data)
	controller.Interactor.PrintInfo("End", "UserController:PasswordChange", string(out))
	fmt.Println(c.Request.RequestURI, "End", "UserController:PasswordChange")

}
