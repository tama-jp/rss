package repository

import (
	"fmt"
	entity "github.com/tama-jp/rss/internal/domain/entities"
	db "github.com/tama-jp/rss/internal/frameworks/database"
	"github.com/tama-jp/rss/internal/frameworks/logger"
	"github.com/tama-jp/rss/internal/usecases/port"
)

type userRoleRepository struct {
	db  *db.DataBase
	log *logger.LogBase
}

func NewUserRoleRepository(db *db.DataBase, log *logger.LogBase) port.UserRolePort {
	return &userRoleRepository{db, log}
}

func (r *userRoleRepository) FindList() (*[]entity.UserRole, error) {
	r.log.PrintInfo("Start", "userRoleRepository:FindList", "")
	fmt.Println("Start", "userRoleRepository:FindList", "")

	var userRoles []entity.UserRole

	result := r.db.Connect().Order("id asc").Find(&userRoles)
	if result.Error != nil {
		return nil, result.Error
	}

	r.log.PrintInfo("End", "userRoleRepository:FindList", "")
	fmt.Println("End", "userRoleRepository:FindList", "")
	return &userRoles, nil
}

func (r *userRoleRepository) FindId(id uint) (*entity.UserRole, error) {
	r.log.PrintInfo("Start", "userRoleRepository:FindId", "")
	fmt.Println("Start", "userRoleRepository:FindId", "")

	var userRole entity.UserRole
	// データベースマッチ

	result := r.db.Connect().Where(&entity.UserRole{Model: entity.Model{ID: id}}).First(&userRole)

	if result.Error != nil {
		return nil, result.Error
	}

	r.log.PrintInfo("End", "userRoleRepository:FindId", "")
	fmt.Println("End", "userRoleRepository:FindId", "")

	return &userRole, nil
}
