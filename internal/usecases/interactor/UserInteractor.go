package interactor

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tama-jp/rss/internal/domain/request"
	"github.com/tama-jp/rss/internal/domain/response"
	db "github.com/tama-jp/rss/internal/frameworks/database"
	"github.com/tama-jp/rss/internal/usecases/port"
	"github.com/tama-jp/rss/internal/utils/message"
	"net/http"
	"strconv"
	"strings"
)

type UserInteractor struct {
	userRepo     port.UserPort
	jwtRepo      port.JwtRepository
	userAuthRepo port.UserAuthPort
	loggerRepo   port.LoggerPort
}

func NewUserInteractor(jwtRepo port.JwtRepository,
	userRepo port.UserPort,
	userAuthRepo port.UserAuthPort,
	loggerRepo port.LoggerPort,
) *UserInteractor {
	return &UserInteractor{
		jwtRepo:      jwtRepo,
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
		loggerRepo:   loggerRepo}
}

func (interactor *UserInteractor) PrintInfo(num string, group string, message string) {
	interactor.loggerRepo.PrintInfo(num, group, message)
}

func (interactor *UserInteractor) PrintError(num string, group string, message string) {
	interactor.loggerRepo.PrintError(num, group, message)
}

func (interactor *UserInteractor) PrintDebug(num string, group string, message string) {
	interactor.loggerRepo.PrintDebug(num, group, message)
}

func (interactor *UserInteractor) Search(c *gin.Context) (httpCode int, statusCode string, data interface{}, error error) {
	interactor.loggerRepo.PrintInfo("Start", "UserInteractor:Search", "")
	fmt.Println("Start", "UserInteractor:Search")

	//targetUser, resMyUserRoleDescription, err := getTargetUserInfo(c, interactor.jwtRepo, interactor.userRepo, interactor.loggerRepo)
	targetUser, resMyUserRoleDescription, err := getTargetUserInfo(c, interactor.jwtRepo, interactor.userRepo, interactor.userAuthRepo, interactor.loggerRepo)

	if err != nil {
		fmt.Println("1-1", "UserInteractor:Search", "err", err.Error())

		return http.StatusBadRequest, message.ERR000, nil, err
	}

	userIdStr := c.Param("user_id")

	fmt.Println("2", "UserInteractor:Search")

	userId, err := strconv.Atoi(userIdStr)

	user, err := interactor.userRepo.FindUserId(uint(userId))

	if err != nil {
		return 0, "", nil, err
	}

	if !(resMyUserRoleDescription.SuperUser == true || resMyUserRoleDescription.Ceo == true) {

		if targetUser.ID != user.ID {
			fmt.Println("2-1", "UserInteractor:Search", "err", err.Error())
			err := errors.New("ユーザーが違います")
			return http.StatusBadRequest, message.ERR999, nil, err
		}
	}

	userRoleDescription := getResUserRoleDescription(user.RoleBitCode, interactor.loggerRepo)

	resLoginSearch := response.UserInfoSearchResponse{
		UserId:          user.ID,
		UserName:        user.UserName,
		LastName:        user.LastName,
		FirstName:       user.FirstName,
		EmployeeNumber:  user.EmployeeNumber,
		RoleBitCode:     user.RoleBitCode,
		RoleDescription: *userRoleDescription,
	}
	interactor.loggerRepo.PrintInfo("End", "UserInteractor:Search", "")
	fmt.Println("End", "UserInteractor:Search")

	return http.StatusOK, message.STS000, resLoginSearch, nil
}

func (interactor *UserInteractor) PasswordChange(c *gin.Context) (httpCode int, statusCode string, data interface{}, error error) {
	interactor.loggerRepo.PrintInfo("Start", "UserInteractor:PasswordChange", "")
	fmt.Println("Start", "UserInteractor:PasswordChange")

	//targetUser, resMyUserRoleDescription, err := getTargetUserInfo(c, interactor.jwtRepo, interactor.userRepo, interactor.loggerRepo)
	targetUser, resMyUserRoleDescription, err := getTargetUserInfo(c, interactor.jwtRepo, interactor.userRepo, interactor.userAuthRepo, interactor.loggerRepo)

	if err != nil {
		return http.StatusBadRequest, message.ERR000, nil, err
	}

	if resMyUserRoleDescription.SuperUser != true {
		return http.StatusBadRequest, message.ERR008, nil, err
	}

	var req request.ReqUserPasswordChange //alice型の変数の定義
	err = c.BindJSON(&req)

	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	// バリデーションチェック
	if err := req.Validate(c); err != nil {
		fmt.Println("Validation Error")
		return http.StatusBadRequest, message.ERR999, nil, err
	}
	result := strings.Compare(req.OldPassword, req.NewPassword)

	if result == 0 {
		return http.StatusBadRequest, message.ERR009, nil, err
	}

	result = strings.Compare(targetUser.Password, db.GetSHA512String(req.OldPassword))

	if result != 0 {
		return http.StatusBadRequest, message.ERR009, nil, err
	}

	retUser, err := interactor.userRepo.UpdateUse(
		targetUser.ID,
		nil,
		nil,
		nil,
		nil,
		&req.NewPassword,
		0)

	if err != nil {
		return http.StatusBadRequest, message.ERR999, nil, err
	}

	resPostUserInfo := response.PostUserInfoResponse{
		UserId:          retUser.ID,
		UserName:        retUser.UserName,
		LastName:        retUser.LastName,
		FirstName:       retUser.FirstName,
		EmployeeNumber:  retUser.EmployeeNumber,
		RoleBitCode:     retUser.RoleBitCode,
		RoleDescription: *resMyUserRoleDescription,
	}

	interactor.loggerRepo.PrintInfo("End", "UserInteractor:PasswordChange", "")
	fmt.Println("End", "UserInteractor:PasswordChange")

	return http.StatusOK, message.STS000, resPostUserInfo, nil
}
