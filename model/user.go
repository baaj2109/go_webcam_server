package model

import (
	"time"

	"github.com/baaj2109/webcam_server/global"
)

type User struct {
	Id int `gorm:"primaryKey" json:"id"`
	// Account  string
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Addtime  time.Time `gorm:"default:current_timestamp"`
	LastIP   string
	Status   int
}

func CheckUser(email, passwd string) bool {
	// var user User
	users := []User{}
	global.SQLLiteDb.Select("id").Where("email == ? AND Password == ?", email, passwd).Find(&users)

	return len(users) > 0
}
