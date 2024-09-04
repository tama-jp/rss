package entity

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `gorm:"primary_key;comment:id"`
	CreatedAt time.Time      `gorm:"index;comment:作成日"`
	UpdatedAt time.Time      `gorm:"index;comment:更新日"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:削除日"`
	CreatedBy string         `gorm:"index;comment:作成者"`
	UpdatedBy string         `gorm:"index;comment:更新者"`
}
