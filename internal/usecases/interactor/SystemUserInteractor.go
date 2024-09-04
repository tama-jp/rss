package interactor

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tama-jp/rss/internal/domain/request"
	"github.com/tama-jp/rss/internal/domain/response"
	"github.com/tama-jp/rss/internal/usecases/port"
	"github.com/tama-jp/rss/internal/utils/message"
	"net/http"
	"strconv"
)

type SystemUserInteractor struct {
	userRepo     port.UserPort
	userRoleRepo port.UserRolePort
	jwtRepo      port.JwtRepository
	userAuthRepo port.UserAuthPort
	loggerRepo   port.LoggerPort
}

func NewSystemUserInteractor(
	jwtRepo port.JwtRepository,
	userRepo port.UserPort,
	userRoleRepo port.UserRolePort,
	userAuthRepo port.UserAuthPort,
	loggerRepo port.LoggerPort,
) *SystemUserInteractor {
	return &SystemUserInteractor{
		jwtRepo:      jwtRepo,
		userRepo:     userRepo,
		userRoleRepo: userRoleRepo,
		userAuthRepo: userAuthRepo,
		loggerRepo:   loggerRepo,
	}
}

func (interactor *SystemUserInteractor) PrintInfo(num string, group string, message string) {
	interactor.loggerRepo.PrintInfo(num, group, message)
}

func (interactor *SystemUserInteractor) PrintError(num string, group string, message string) {
	interactor.loggerRepo.PrintError(num, group, message)
}

func (interactor *SystemUserInteractor) PrintDebug(num string, group string, message string) {
	interactor.loggerRepo.PrintDebug(num, group, message)
}

func (interactor *SystemUserInteractor) GetUserRoleList(c *gin.Context) (httpCode int, statusCode string, data interface{}, error error) {
	interactor.loggerRepo.PrintInfo("Start", "SystemUserInteractor:GetUserRoleList", "")
	fmt.Println("Start", "SystemUserInteractor:GetUserRoleList")

	userRoles, err := interactor.userRoleRepo.FindList()

	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	var resUserRoleList []response.UserRoleListResponse

	for _, userRole := range *userRoles {
		resUserRoleList = append(resUserRoleList, response.UserRoleListResponse{
			Id:       userRole.ID,
			Name:     userRole.Name,
			RoleName: userRole.RoleName,
		})
	}

	interactor.loggerRepo.PrintInfo("End", "SystemUserInteractor:GetUserRoleList", "")
	fmt.Println("End", "SystemUserInteractor:GetUserRoleList")

	return http.StatusOK, message.STS000, resUserRoleList, nil
}

func (interactor *SystemUserInteractor) GetUserList(c *gin.Context) (httpCode int, statusCode string, data interface{}, error error) {
	interactor.loggerRepo.PrintInfo("Start", "SystemUserInteractor:GetUserList", "")
	fmt.Println("Start", "SystemUserInteractor:GetUserList")

	//_, resMyUserRoleDescription, err := getTargetUserInfo(c, interactor.jwtRepo, interactor.userRepo, interactor.loggerRepo)
	_, resMyUserRoleDescription, err := getTargetUserInfo(c, interactor.jwtRepo, interactor.userRepo, interactor.userAuthRepo, interactor.loggerRepo)

	if !(resMyUserRoleDescription.SuperUser == true || resMyUserRoleDescription.Ceo == true) {
		return http.StatusBadRequest, message.ERR000, nil, err
	}

	users, err := interactor.userRepo.GetUserList()

	var resUserList []response.UserListResponse

	for _, user := range *users {
		userRoleDescription := getResUserRoleDescription(user.RoleBitCode, interactor.loggerRepo)
		resUserList = append(resUserList, response.UserListResponse{
			UserId:          user.ID,
			UserName:        user.UserName,
			LastName:        user.LastName,
			FirstName:       user.FirstName,
			EmployeeNumber:  user.EmployeeNumber,
			RoleBitCode:     user.RoleBitCode,
			RoleDescription: *userRoleDescription,
		})
	}

	interactor.loggerRepo.PrintInfo("End", "SystemUserInteractor:GetUserList", "")
	fmt.Println("End", "SystemUserInteractor:GetUserList")

	return http.StatusOK, message.STS000, resUserList, nil
}

func (interactor *SystemUserInteractor) SearchUserInfo(c *gin.Context) (httpCode int, statusCode string, data interface{}, error error) {
	interactor.loggerRepo.PrintInfo("Start", "SystemUserInteractor:SearchUserInfo", "")
	fmt.Println("Start", "SystemUserInteractor:SearchUserInfo")

	//_, resMyUserRoleDescription, err := getTargetUserInfo(c, interactor.jwtRepo, interactor.userRepo, interactor.loggerRepo)
	_, resMyUserRoleDescription, err := getTargetUserInfo(c, interactor.jwtRepo, interactor.userRepo, interactor.userAuthRepo, interactor.loggerRepo)

	if !(resMyUserRoleDescription.SuperUser == true || resMyUserRoleDescription.Ceo == true) {
		return http.StatusBadRequest, message.ERR000, nil, err
	}

	userIdStr := c.Param("user_id")

	userId, _ := strconv.Atoi(userIdStr)

	user, resUserRoleDescription, err := getUser(uint(userId), interactor.userRepo, interactor.loggerRepo)
	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	interactor.loggerRepo.PrintInfo(" user.ID", "SystemUserInteractor:SearchUserInfo", string(user.ID))
	fmt.Println(" user.ID", "SystemUserInteractor:SearchUserInfo", string(user.ID))

	resUserSearch := response.UserSearchResponse{
		UserId:          user.ID,
		UserName:        user.UserName,
		LastName:        user.LastName,
		FirstName:       user.FirstName,
		EmployeeNumber:  user.EmployeeNumber,
		RoleBitCode:     user.RoleBitCode,
		RoleDescription: *resUserRoleDescription,
	}
	interactor.loggerRepo.PrintInfo("End", "SystemUserInteractor:SearchUserInfo", "")
	fmt.Println("End", "SystemUserInteractor:SearchUserInfo")

	return http.StatusOK, message.STS000, resUserSearch, nil
}

func (interactor *SystemUserInteractor) PutUserInfo(c *gin.Context) (httpCode int, statusCode string, data interface{}, error error) {
	interactor.loggerRepo.PrintInfo("Start", "SystemUserInteractor:PutUserInfo", "")
	fmt.Println("Start", "SystemUserInteractor:PutUserInfo")

	//_, resMyUserRoleDescription, err := getTargetUserInfo(c, interactor.jwtRepo, interactor.userRepo, interactor.loggerRepo)
	_, resMyUserRoleDescription, err := getTargetUserInfo(c, interactor.jwtRepo, interactor.userRepo, interactor.userAuthRepo, interactor.loggerRepo)

	if !(resMyUserRoleDescription.SuperUser == true || resMyUserRoleDescription.Ceo == true) {
		return http.StatusBadRequest, message.ERR000, nil, err
	}

	fmt.Println("SystemUserInteractor PutUserInfo", 1)

	userIdStr := c.Param("user_id")
	var req request.UserPutRequest //alice型の変数の定義
	err = c.BindJSON(&req)

	fmt.Println("SystemUserInteractor PutUserInfo", 2)

	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	// TODO: ここできていないNull許容の仕方理解できていない
	// バリデーションチェック
	//if err := req.Validate(c); err != nil {
	//	fmt.Println("Validation Error")
	//	return http.StatusBadRequest, message.ERR999, nil, err
	//}

	fmt.Println("SystemUserInteractor PutUserInfo", 3)
	userId, err := strconv.Atoi(userIdStr)
	fmt.Println("SystemUserInteractor PutUserInfo", 4)

	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	fmt.Println("SystemUserInteractor PutUserInfo", 5)

	_, err = interactor.userRepo.FindUserId(uint(userId))

	fmt.Println("SystemUserInteractor PutUserInfo", 6)

	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	fmt.Println("SystemUserInteractor PutUserInfo", 7)

	fmt.Println("SystemUserInteractor PutUserInfo", req)
	// TODO:ユーザーロールチェックいるかも

	retUser, err := interactor.userRepo.UpdateUse(
		uint(userId),
		req.UserName,
		req.LastName,
		req.FirstName,
		req.EmployeeNumber,
		req.Password,
		req.RoleBitCode)

	fmt.Println("SystemUserInteractor PutUserInfo", 8)

	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	fmt.Println("SystemUserInteractor PutUserInfo", 9)

	_, resUserRoleDescription, err := getUser(retUser.ID, interactor.userRepo, interactor.loggerRepo)

	fmt.Println("SystemUserInteractor PutUserInfo", 10)

	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	fmt.Println("SystemUserInteractor PutUserInfo", 11)

	resPostUserInfo := response.PostUserInfoResponse{
		UserId:          retUser.ID,
		UserName:        retUser.UserName,
		LastName:        retUser.LastName,
		FirstName:       retUser.FirstName,
		EmployeeNumber:  retUser.EmployeeNumber,
		RoleBitCode:     retUser.RoleBitCode,
		RoleDescription: *resUserRoleDescription,
	}

	interactor.loggerRepo.PrintInfo("End", "SystemUserInteractor:PutUserInfo", "")
	fmt.Println("End", "SystemUserInteractor:PutUserInfo")

	return http.StatusOK, message.STS000, resPostUserInfo, nil
}

func (interactor *SystemUserInteractor) DeleteUserInfo(c *gin.Context) (httpCode int, statusCode string, data interface{}, error error) {
	interactor.loggerRepo.PrintInfo("Start", "SystemUserInteractor:DeleteUserInfo", "")
	fmt.Println("Start", "SystemUserInteractor:DeleteUserInfo")

	//_, resMyUserRoleDescription, err := getTargetUserInfo(c, interactor.jwtRepo, interactor.userRepo, interactor.loggerRepo)
	_, resMyUserRoleDescription, err := getTargetUserInfo(c, interactor.jwtRepo, interactor.userRepo, interactor.userAuthRepo, interactor.loggerRepo)

	if !(resMyUserRoleDescription.SuperUser == true || resMyUserRoleDescription.Ceo == true) {
		return http.StatusBadRequest, message.ERR000, nil, err
	}

	userIdStr := c.Param("user_id")

	userId, err := strconv.Atoi(userIdStr)

	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	if userId == 1 {
		messageStr := message.GetMsg(message.ERR016)
		return http.StatusBadRequest, message.ERR016, nil, errors.New(messageStr)

	}

	user, err := interactor.userRepo.FindUserId(uint(userId))
	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	retUser, err := interactor.userRepo.DeleteUser(user.UserName)

	resDeleteUserInfo := response.DeleteUserInfoResponse{
		UserName: retUser.UserName,
	}
	interactor.loggerRepo.PrintInfo("End", "SystemUserInteractor:DeleteUserInfo", "")
	fmt.Println("End", "SystemUserInteractor:DeleteUserInfo")

	return http.StatusOK, message.STS000, resDeleteUserInfo, nil
}

func (interactor *SystemUserInteractor) PostUserInfo(c *gin.Context) (httpCode int, statusCode string, data interface{}, error error) {
	interactor.loggerRepo.PrintInfo("Start", "SystemUserInteractor:PostUserInfo", "")
	fmt.Println("Start", "SystemUserInteractor:PostUserInfo")

	//_, resMyUserRoleDescription, err := getTargetUserInfo(c, interactor.jwtRepo, interactor.userRepo, interactor.loggerRepo)
	_, resMyUserRoleDescription, err := getTargetUserInfo(c, interactor.jwtRepo, interactor.userRepo, interactor.userAuthRepo, interactor.loggerRepo)

	if !(resMyUserRoleDescription.SuperUser == true || resMyUserRoleDescription.Ceo == true) {
		return http.StatusBadRequest, message.ERR000, nil, err
	}

	var req request.UserPostRequest
	err = c.BindJSON(&req)

	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	// バリデーションチェック
	// TODO: ここできていないRoleCodeの部分がおかしい
	if err := req.Validate(c); err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	// TODO:ユーザーロールチェックいるかも

	user, err := interactor.userRepo.InsertUser(
		req.UserName,
		req.LastName,
		req.FirstName,
		req.EmployeeNumber,
		req.Password,
		req.RoleBitCode)

	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	user, resUserRoleDescription, err := getUser(user.ID, interactor.userRepo, interactor.loggerRepo)

	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	resPostUserInfo := response.PostUserInfoResponse{
		UserId:          user.ID,
		UserName:        user.UserName,
		LastName:        user.LastName,
		FirstName:       user.FirstName,
		EmployeeNumber:  user.EmployeeNumber,
		RoleBitCode:     user.RoleBitCode,
		RoleDescription: *resUserRoleDescription,
	}

	interactor.loggerRepo.PrintInfo("End", "SystemUserInteractor:PostUserInfo", "")
	fmt.Println("End", "SystemUserInteractor:PostUserInfo")

	return http.StatusOK, message.STS000, resPostUserInfo, nil
}
