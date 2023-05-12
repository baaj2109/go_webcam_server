package model

import "github.com/baaj2109/webcam_server/global"

type Auth struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(email, passwd string) bool {
	// var user User
	users := []Auth{}
	global.SQLLiteDb.Select("id").Where("email == ? ?", email).Find(&users)

	if len(users) > 0 {
		return global.CheckPasswordHash(passwd, users[0].Password)
	}
	return false
}
