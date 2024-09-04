package repository

import (
	"fmt"
	entity "github.com/tama-jp/rss/internal/domain/entities"
	db "github.com/tama-jp/rss/internal/frameworks/database"
	"github.com/tama-jp/rss/internal/frameworks/logger"
	"github.com/tama-jp/rss/internal/usecases/port"
)

type userAuthRepository struct {
	db  *db.DataBase
	log *logger.LogBase
}

func NewUserAuthRepository(db *db.DataBase, log *logger.LogBase) port.UserAuthPort {
	return &userAuthRepository{db, log}
}

func (r *userAuthRepository) FindAccessToken(accessToken string) (*entity.UserAuth, error) {
	r.log.PrintInfo("Start", "attendanceRepository:FindAccessToken", "")
	fmt.Println("Start", "attendanceRepository:FindAccessToken", "")

	var userAuth entity.UserAuth

	fmt.Println("1", "attendanceRepository:FindAccessToken:accessToken", accessToken)

	// データベースマッチ
	result := r.db.Connect().Where(&entity.UserAuth{AccessToken: accessToken}).First(&userAuth)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, nil
	}

	r.log.PrintInfo("End", "attendanceRepository:FindAccessToken", "")
	fmt.Println("End", "attendanceRepository:FindAccessToken", "")

	return &userAuth, nil
}

func (r *userAuthRepository) DeleteInsertAccessToken(userID uint, deleteAccessToken string, insertAccessToken string) (*entity.UserAuth, error) {
	r.log.PrintInfo("Start", "attendanceRepository:DeleteInsertAccessToken", "")
	fmt.Println("Start", "attendanceRepository:DeleteInsertAccessToken", "")

	// トランザクション 対応
	var userAuth entity.UserAuth

	result := r.db.Connect().Where(&entity.UserAuth{AccessToken: deleteAccessToken}).Delete(&userAuth)

	userAuth = entity.UserAuth{UserID: userID, AccessToken: insertAccessToken}

	result = r.db.Connect().Create(&userAuth) // pass pointer of data to Create
	if result.Error != nil {
		return nil, result.Error
	}

	r.log.PrintInfo("End", "attendanceRepository:DeleteInsertAccessToken", "")
	fmt.Println("End", "attendanceRepository:DeleteInsertAccessToken", "")

	return &userAuth, nil
}

func (r *userAuthRepository) DeleteAccessToken(userID uint, accessToken string) (*entity.UserAuth, error) {
	r.log.PrintInfo("Start", "attendanceRepository:DeleteAccessToken", "")
	fmt.Println("Start", "attendanceRepository:DeleteAccessToken", "")

	// トランザクション 対応
	var userAuth entity.UserAuth

	tx := r.db.Connect().Begin()

	result := tx.Where(&entity.UserAuth{UserID: userID, AccessToken: accessToken}).First(&userAuth)

	if result.Error != nil {
		fmt.Printf("result[%s]\n", result.Error)

		tx.Rollback()
		return nil, result.Error
	}

	tx.Delete(&userAuth)
	tx.Commit()

	r.log.PrintInfo("End", "attendanceRepository:DeleteAccessToken", "")
	fmt.Println("End", "attendanceRepository:DeleteAccessToken", "")

	return &userAuth, nil
}

func (r *userAuthRepository) InsertAccessToken(userID uint, accessToken string) (*entity.UserAuth, error) {
	r.log.PrintInfo("Start", "attendanceRepository:InsertAccessToken", "")
	fmt.Println("Start", "attendanceRepository:InsertAccessToken", "")

	userAuth := entity.UserAuth{UserID: userID, AccessToken: accessToken}

	result := r.db.Connect().Create(&userAuth) // pass pointer of data to Create
	if result.Error != nil {
		return nil, result.Error
	}

	result = r.db.Connect().Where(&entity.UserAuth{AccessToken: accessToken}).First(&userAuth)

	if result.Error != nil {
		return nil, result.Error
	}

	r.log.PrintInfo("End", "attendanceRepository:InsertAccessToken", "")
	fmt.Println("End", "attendanceRepository:InsertAccessToken", "")

	return &userAuth, nil
}
