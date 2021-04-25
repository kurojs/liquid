package store

import "gorm.io/gorm"

// User reflects users data from DB
type User struct {
	*gorm.Model
	Username string `gorm:"index,type=VARCHAR(256)"`
	Password string
}
