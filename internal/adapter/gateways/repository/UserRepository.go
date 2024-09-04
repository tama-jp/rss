package repository

import (
	"errors"
	"fmt"
	"github.com/tama-jp/rss/internal/domain/entities"
	db "github.com/tama-jp/rss/internal/frameworks/database"
	"github.com/tama-jp/rss/internal/frameworks/logger"
	"github.com/tama-jp/rss/internal/usecases/port"
	"gorm.io/gorm"
)

type userRepository struct {
	db  *db.DataBase
	log *logger.LogBase
}

func NewUserRepository(db *db.DataBase, log *logger.LogBase) port.UserPort {
	return &userRepository{db, log}
}

func (r *userRepository) GetUserList() (*[]entity.User, error) {
	r.log.PrintInfo("Start", "userRepository:GetUserList", "")
	fmt.Println("Start", "userRepository:GetUserList", "")

	var users []entity.User

	result := r.db.Connect().Model(&entity.User{}).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	r.log.PrintInfo("End", "userRepository:GetUserList", "")
	fmt.Println("End", "userRepository:GetUserList", "")

	return &users, nil
}

func (r *userRepository) FindUserId(userId uint) (*entity.User, error) {
	r.log.PrintInfo("Start", "userRepository:FindUserId", "")
	fmt.Println("Start", "userRepository:FindUserId", "")

	var user entity.User
	// データベースマッチ

	fmt.Println("1", "FindUserId:FindUserId")
	result := r.db.Connect().Where(&entity.User{Model: entity.Model{ID: userId}}).Find(&user)
	//result := r.db.Connect().Where(&entity.User{Model: entity.Model{ID: userId}}).First(&user)
	fmt.Println("2", "FindUserId:FindUserId")

	if result.Error != nil {
		return nil, result.Error
	}

	r.log.PrintInfo("End", "userRepository:FindUserId", "")
	fmt.Println("End", "userRepository:FindUserId", "")

	return &user, nil
}

func (r *userRepository) LoginUserNamePassword(userName string, password string) (*entity.User, error) {
	r.log.PrintInfo("Start", "userRepository:LoginUserNamePassword", "")
	fmt.Println("Start", "userRepository:LoginUserNamePassword", "")

	var user entity.User
	// データベースマッチ

	result := r.db.Connect().Where(&entity.User{UserName: userName, Password: db.GetSHA512String(password)}).Model(&entity.User{}).Find(&user)
	//result := r.db.Connect().Where(&entity.User{UserName: userName}).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	r.log.PrintInfo("End", "userRepository:LoginUserNamePassword", "")
	fmt.Println("End", "userRepository:LoginUserNamePassword", "")

	return &user, nil
}

func (r *userRepository) FindUserName(userName string) (*entity.User, error) {
	r.log.PrintInfo("Start", "userRepository:FindUserName", "")
	fmt.Println("Start", "userRepository:FindUserName", "")

	var user entity.User
	// データベースマッチ

	result := r.db.Connect().Where(&entity.User{UserName: userName}).Model(&entity.User{}).Find(&user)
	//result := r.db.Connect().Where(&entity.User{UserName: userName}).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	r.log.PrintInfo("End", "userRepository:FindUserName", "")
	fmt.Println("End", "userRepository:FindUserName", "")

	return &user, nil
}

func (r *userRepository) InsertUser(userName string, lastName string, firstName string, employeeNumber string, password string, roleBitCode uint64) (*entity.User, error) {
	r.log.PrintInfo("End", "userRepository:InsertUser", "")
	fmt.Println("End", "userRepository:InsertUser", "")

	//result := r.db.Connect().Where(&entity.UserRole{Model: entity.Model{ID: UserRoleID}})
	//
	//if result.Error != nil {
	//	return nil, result.Error
	//}

	result := r.db.Connect().Where("user_name = ?", userName).First(&entity.User{})

	var user entity.User
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		user = entity.User{
			UserName:       userName,
			LastName:       lastName,
			FirstName:      firstName,
			EmployeeNumber: employeeNumber,
			Password:       db.GetSHA512String(password),
			RoleBitCode:    roleBitCode,
		}

		result = r.db.Connect().Create(&user) // pass pointer of data to Create
		if result.Error != nil {
			return nil, result.Error
		}

		result = r.db.Connect().Where(&entity.User{UserName: userName})

		if result.Error != nil {
			return nil, result.Error
		}

	} else {
		return nil, errors.New("already have a user")
	}

	r.log.PrintInfo("End", "userRepository:InsertUser", "")
	fmt.Println("End", "userRepository:InsertUser", "")

	return &user, nil

}

func (r *userRepository) UpdateUse(
	userId uint,
	userName *string,
	lastName *string,
	firstName *string,
	employeeNumber *string,
	password *string,
	roleBitCode uint64) (*entity.User, error) {
	r.log.PrintInfo("Start", "userRepository:UpdateUse", "")
	fmt.Println("Start", "userRepository:UpdateUse", "")

	user := entity.User{}

	if err := r.db.Connect().Where(&entity.User{Model: entity.Model{ID: userId}}).First(&user).Error; err == nil {

		if userName != nil {
			user.UserName = *userName
		}

		if lastName != nil {
			user.LastName = *lastName
		}

		if firstName != nil {
			user.FirstName = *firstName
		}

		if employeeNumber != nil {
			user.EmployeeNumber = *employeeNumber
		}

		if password != nil {
			passwordEnc := db.GetSHA512String(*password)
			user.Password = passwordEnc
		}

		if roleBitCode != 0 {
			user.RoleBitCode = roleBitCode
		}

		r.db.Connect().Save(&user)
		result := r.db.Connect().Where(&entity.User{Model: entity.Model{ID: userId}}).Model(&entity.User{}).Find(&user)

		if result.Error != nil {
			return nil, result.Error
		}

	} else {
		return nil, err
	}

	r.log.PrintInfo("End", "userRepository:UpdateUse", "")
	fmt.Println("End", "userRepository:UpdateUse", "")

	return &user, nil
}

func (r *userRepository) DeleteUser(userName string) (*entity.User, error) {
	r.log.PrintInfo("Start", "userRepository:DeleteUser", "")
	fmt.Println("Start", "userRepository:DeleteUser", "")

	user := entity.User{}

	if err := r.db.Connect().Where(&entity.User{UserName: userName}).First(&user).Error; err == nil {
		user.RoleBitCode = entity.UserRoleNoAuthority
		r.db.Connect().Save(&user)
		result := r.db.Connect().Where(&entity.User{UserName: userName})

		if result.Error != nil {
			return nil, result.Error
		}

	} else {
		return nil, err
	}

	r.log.PrintInfo("End", "userRepository:DeleteUser", "")
	fmt.Println("End", "userRepository:DeleteUser", "")

	return &user, nil

}
