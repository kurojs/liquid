package store

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mySQLStore *MySQLStore
	once       sync.Once
)

type MySQLStore struct {
	db *gorm.DB
}

func GetMySQLStore(username, pwd, dbname, host string, port int) *MySQLStore {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			username, pwd, host, port, dbname,
		)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("init DB connection failed %s\n", err)
		}

		if strings.EqualFold(os.Getenv("DB_DEBUG"), "true") {
			db = db.Debug()
		}

		if err := db.AutoMigrate(&User{}); err != nil {
			log.Fatalf("migrate db failed %s\n", err)
		}

		// These users is for testing only
		db.Create(&User{
			Username: "kuro",
			Password: "kuro",
		}).Create(&User{
			Username: "liquid",
			Password: "liquid",
		}).Create(&User{
			Username: "skywalker",
			Password: "skywalker",
		})

		mySQLStore = &MySQLStore{
			db: db,
		}
	})

	return mySQLStore
}

// ValidateUser valid user by getting user existed in db and check if its hashed password matched
func (s *MySQLStore) ValidateUser(ctx context.Context, user *User) (bool, error) {
	gotUser := &User{}
	if err := s.db.WithContext(ctx).First(gotUser, "username = ? AND password = ?", user.Username, user.Password).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	/*  TODO: use bcrypt for password
	if err := s.db.WithContext(ctx).First(user, "username = ?", user.Username).Error; err != nil {
		return false, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(gotUser.Password), []byte(user.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil

		}
		return false, err
	}
	*/

	return true, nil
}
