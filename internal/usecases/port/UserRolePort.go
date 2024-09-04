package port

import entity "github.com/tama-jp/rss/internal/domain/entities"

type UserRolePort interface {
	FindId(id uint) (*entity.UserRole, error)
	FindList() (*[]entity.UserRole, error)
}
