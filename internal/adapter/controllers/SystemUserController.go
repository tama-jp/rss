package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tama-jp/rss/internal/domain/response"
	"github.com/tama-jp/rss/internal/usecases/interactor"
)

type SystemUserController struct {
	Interactor *interactor.SystemUserInteractor
}

func NewSystemUserController(interactor *interactor.SystemUserInteractor) *SystemUserController {
	return &SystemUserController{Interactor: interactor}
}

func (controller *SystemUserController) GetUserRoleList(c *gin.Context) {
	controller.Interactor.PrintInfo("Start", "SystemUserController:GetUserRoleList", c.Request.RequestURI)
	fmt.Println(c.Request.RequestURI, "Start", "SystemUserController:GetUserRoleList")

	res := response.Response{C: c}

	status, statusCode, data, err := controller.Interactor.GetUserRoleList(c)
	if err != nil {
		fmt.Println("err", err)
		res.ErrorPayload(status, statusCode, err)

		return
	} else {

		fmt.Println("OK", data)
		res.Payload(status, statusCode, data)
	}

	out, _ := json.Marshal(data)
	controller.Interactor.PrintInfo("End", "SystemUserController:GetUserRoleList", string(out))
	fmt.Println(c.Request.RequestURI, "End", "SystemUserController:GetUserRoleList")
}

func (controller *SystemUserController) GetUserList(c *gin.Context) {
	controller.Interactor.PrintInfo("Start", "SystemUserController:GetUserList", c.Request.RequestURI)
	fmt.Println(c.Request.RequestURI, "Start", "SystemUserController:GetUserList")

	res := response.Response{C: c}

	status, statusCode, data, err := controller.Interactor.GetUserList(c)
	if err != nil {
		fmt.Println("err", err)
		res.ErrorPayload(status, statusCode, err)

		return
	} else {

		fmt.Println("OK", data)
		res.Payload(status, statusCode, data)
	}

	out, _ := json.Marshal(data)
	controller.Interactor.PrintInfo("End", "SystemUserController:GetUserList", string(out))
	fmt.Println(c.Request.RequestURI, "End", "SystemUserController:GetUserList")

}

func (controller *SystemUserController) SearchUserInfo(c *gin.Context) {
	controller.Interactor.PrintInfo("Start", "SystemUserController:SearchUserInfo", c.Request.RequestURI)
	fmt.Println(c.Request.RequestURI, "Start", "SystemUserController:SearchUserInfo")

	res := response.Response{C: c}

	status, statusCode, data, err := controller.Interactor.SearchUserInfo(c)
	if err != nil {
		fmt.Println("err", err)
		res.ErrorPayload(status, statusCode, err)

		return
	} else {

		fmt.Println("OK", data)
		res.Payload(status, statusCode, data)
	}

	out, _ := json.Marshal(data)
	controller.Interactor.PrintInfo("End", "SystemUserController:SearchUserInfo", string(out))
	fmt.Println(c.Request.RequestURI, "End", "SystemUserController:SearchUserInfo")

}

func (controller *SystemUserController) PutUserInfo(c *gin.Context) {
	controller.Interactor.PrintInfo("Start", "SystemUserController:PutUserInfo", c.Request.RequestURI)
	fmt.Println(c.Request.RequestURI, "Start", "SystemUserController:PutUserInfo")

	res := response.Response{C: c}

	status, statusCode, data, err := controller.Interactor.PutUserInfo(c)
	if err != nil {
		fmt.Println("err", err)
		res.ErrorPayload(status, statusCode, err)

		return
	} else {

		fmt.Println("OK", data)
		res.Payload(status, statusCode, data)
	}

	out, _ := json.Marshal(data)
	controller.Interactor.PrintInfo("End", "SystemUserController:PutUserInfo", string(out))
	fmt.Println(c.Request.RequestURI, "End", "SystemUserController:PutUserInfo")
}

func (controller *SystemUserController) DeleteUserInfo(c *gin.Context) {
	controller.Interactor.PrintInfo("Start", "SystemUserController:DeleteUserInfo", c.Request.RequestURI)
	fmt.Println(c.Request.RequestURI, "Start", "SystemUserController:DeleteUserInfo")

	res := response.Response{C: c}

	status, statusCode, data, err := controller.Interactor.DeleteUserInfo(c)

	if err != nil {
		fmt.Println("err", err)
		res.ErrorPayload(status, statusCode, err)

		return
	} else {

		fmt.Println("OK", data)
		res.Payload(status, statusCode, data)
	}

	out, _ := json.Marshal(data)
	controller.Interactor.PrintInfo("End", "SystemUserController:DeleteUserInfo", string(out))
	fmt.Println(c.Request.RequestURI, "End", "SystemUserController:DeleteUserInfo")

}

func (controller *SystemUserController) PostUserInfo(c *gin.Context) {
	controller.Interactor.PrintInfo("Start", "SystemUserController:PostUserInfo", c.Request.RequestURI)
	fmt.Println(c.Request.RequestURI, "Start", "SystemUserController:PostUserInfo")

	res := response.Response{C: c}

	status, statusCode, data, err := controller.Interactor.PostUserInfo(c)

	if err != nil {
		fmt.Println("err", err)
		res.ErrorPayload(status, statusCode, err)

		return
	} else {

		fmt.Println("OK", data)
		res.Payload(status, statusCode, data)
	}

	out, _ := json.Marshal(data)
	controller.Interactor.PrintInfo("End", "SystemUserController:PostUserInfo", string(out))
	fmt.Println(c.Request.RequestURI, "End", "SystemUserController:PostUserInfo")

}
