package entity

type UserAuth struct {
	Model
	UserID      uint   `gorm:"index;comment:ユーザID"`
	User        User   `gorm:"comment:ユーザ"`
	AccessToken string `gorm:"index:unique;comment:アクセストークン"`
}
