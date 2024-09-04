package port

import entity "github.com/tama-jp/rss/internal/domain/entities"

type UserPort interface {
	GetUserList() (*[]entity.User, error)
	FindUserId(userId uint) (*entity.User, error)
	LoginUserNamePassword(userName string, password string) (*entity.User, error)
	FindUserName(userName string) (*entity.User, error)
	InsertUser(userName string, lastName string, firstName string, employeeNumber string, password string, roleBitCode uint64) (*entity.User, error)
	UpdateUse(userId uint, userName *string, lastName *string, firstName *string, employeeNumber *string, password *string, roleBitCode uint64) (*entity.User, error)
	DeleteUser(userName string) (*entity.User, error)
}
