package interactor

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tama-jp/rss/internal/domain/request"
	"github.com/tama-jp/rss/internal/domain/response"
	"github.com/tama-jp/rss/internal/usecases/port"
	"github.com/tama-jp/rss/internal/utils/message"
	"net/http"
)

type AccessTokenInteractor struct {
	userRepo     port.UserPort
	jwtRepo      port.JwtRepository
	userAuthRepo port.UserAuthPort
	loggerRepo   port.LoggerPort
}

func NewAccessTokenInteractor(jwtRepo port.JwtRepository,
	userRepo port.UserPort,
	userAuthRepo port.UserAuthPort,
	loggerRepo port.LoggerPort,
) *AccessTokenInteractor {
	return &AccessTokenInteractor{
		jwtRepo:      jwtRepo,
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
		loggerRepo:   loggerRepo,
	}
}

func (interactor *AccessTokenInteractor) PrintInfo(num string, group string, message string) {
	interactor.loggerRepo.PrintInfo(num, group, message)
}

func (interactor *AccessTokenInteractor) PrintError(num string, group string, message string) {
	interactor.loggerRepo.PrintError(num, group, message)
}

func (interactor *AccessTokenInteractor) PrintDebug(num string, group string, message string) {
	interactor.loggerRepo.PrintDebug(num, group, message)
}

func (interactor *AccessTokenInteractor) GetAccessToken(c *gin.Context) (httpCode int, statusCode string, data interface{}, error error) {
	interactor.loggerRepo.PrintInfo("Start", "AccessTokenInteractor:GetAccessToken", "")
	fmt.Println("Start", "AccessTokenInteractor:GetAccessToken")

	userName := c.Request.Header.Get(message.X_USER_NAME)

	interactor.loggerRepo.PrintInfo("userName", "AccessTokenInteractor:GetAccessToken", "userName["+userName+"]")
	fmt.Println("userName", "AccessTokenInteractor:GetAccessToken", "userName["+userName+"]")

	password := c.Request.Header.Get(message.X_PASSWORD)

	interactor.loggerRepo.PrintInfo("userName", "AccessTokenInteractor:GetAccessToken", "password["+password+"]")
	fmt.Println("userName", "AccessTokenInteractor:GetAccessToken", "userName["+password+"]")

	req := request.LoginSearchRequest{
		UserName: userName,
		Password: password,
	}

	// バリデーションチェック
	if err := req.Validate(c); err != nil {
		fmt.Println("Validation Error")
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	// Users の取得
	targetUser, err := interactor.userRepo.LoginUserNamePassword(userName, password)
	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	if targetUser.ID == 0 {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	// アクセストークン取得
	auth, accessToken, err := interactor.jwtRepo.GenerateTokenProc(req.UserName)

	_, err = interactor.userAuthRepo.InsertAccessToken(targetUser.ID, accessToken)

	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	userRoleDescription := getResUserRoleDescription(targetUser.RoleBitCode, interactor.loggerRepo)
	resLoginSearch := response.GetAccessTokenResponse{
		UserID:             targetUser.ID,
		EmployeeNumber:     targetUser.EmployeeNumber,
		UserName:           targetUser.UserName,
		LastName:           targetUser.LastName,
		FirstName:          targetUser.FirstName,
		AccessToken:        accessToken,
		AccessTokenExpires: auth.Exp,
		RoleBitCode:        targetUser.RoleBitCode,
		RoleDescription:    *userRoleDescription,
	}
	interactor.loggerRepo.PrintInfo("End", "AccessTokenInteractor:GetAccessToken", "")
	fmt.Println("End", "AccessTokenInteractor:GetAccessToken")

	return http.StatusOK, message.STS000, resLoginSearch, nil
}

func (interactor *AccessTokenInteractor) PutAccessToken(c *gin.Context) (httpCode int, statusCode string, data interface{}, error error) {
	interactor.loggerRepo.PrintInfo("Start", "AccessTokenInteractor:PutAccessToken", "")
	fmt.Println("Start", "AccessTokenInteractor:PutAccessToken")

	oldAccessToken, err := interactor.jwtRepo.GetToken(c)
	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	auth, err := interactor.jwtRepo.ParseProc(oldAccessToken)

	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	// アクセストークン取得
	auth, newAccessToken, err := interactor.jwtRepo.GenerateTokenProc(auth.UserName)

	// Users の取得
	targetUser, err := interactor.userRepo.FindUserName(auth.UserName)

	if targetUser.ID == 0 {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	_, err = interactor.userAuthRepo.DeleteInsertAccessToken(targetUser.ID, oldAccessToken, newAccessToken)

	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	userRoleDescription := getResUserRoleDescription(targetUser.RoleBitCode, interactor.loggerRepo)

	resLoginSearch := response.GetAccessTokenResponse{
		UserID:             targetUser.ID,
		EmployeeNumber:     targetUser.EmployeeNumber,
		UserName:           targetUser.UserName,
		LastName:           targetUser.LastName,
		FirstName:          targetUser.FirstName,
		AccessToken:        newAccessToken,
		AccessTokenExpires: auth.Exp,
		RoleBitCode:        targetUser.RoleBitCode,
		RoleDescription:    *userRoleDescription,
	}
	interactor.loggerRepo.PrintInfo("End", "AccessTokenInteractor:PutAccessToken", "")
	fmt.Println("End", "AccessTokenInteractor:PutAccessToken")

	return http.StatusOK, message.STS000, resLoginSearch, nil
}

func (interactor *AccessTokenInteractor) DeleteAccessToken(c *gin.Context) (httpCode int, statusCode string, data interface{}, error error) {
	interactor.loggerRepo.PrintInfo("Start", "AccessTokenInteractor:DeleteAccessToken", "")
	fmt.Println("Start", "AccessTokenInteractor:DeleteAccessToken")

	accessToken, err := interactor.jwtRepo.GetToken(c)
	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	auth, err := interactor.jwtRepo.ParseProc(accessToken)

	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	// Users の取得
	targetUser, err := interactor.userRepo.FindUserName(auth.UserName)

	_, err = interactor.userAuthRepo.DeleteAccessToken(targetUser.ID, accessToken)

	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	if targetUser.ID == 0 {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	resLoginSearch := response.DeleteAccessTokenResponse{
		UserName:    targetUser.UserName,
		AccessToken: accessToken,
	}
	interactor.loggerRepo.PrintInfo("End", "AccessTokenInteractor:DeleteAccessToken", "")
	fmt.Println("End", "AccessTokenInteractor:DeleteAccessToken")

	return http.StatusOK, message.STS000, resLoginSearch, nil
}
