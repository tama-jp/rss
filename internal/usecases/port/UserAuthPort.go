package port

import entity "github.com/tama-jp/rss/internal/domain/entities"

type UserAuthPort interface {
	FindAccessToken(accessToken string) (*entity.UserAuth, error)
	DeleteAccessToken(userID uint, accessToken string) (*entity.UserAuth, error)
	InsertAccessToken(userID uint, accessToken string) (*entity.UserAuth, error)
	DeleteInsertAccessToken(userID uint, deleteAccessToken string, insertAccessToken string) (*entity.UserAuth, error)
}
