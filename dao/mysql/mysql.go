package mysql

import (
	"errors"
	"fmt"
	"github.com/xws117/bluebell/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn string = "root:123456@tcp(127.0.0.1:3306)/123?charset=utf8mb4&parseTime=True&loc=Local"
var db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func Login(user *model.User) (err error) {
	//oripassword := user.Password
	var u model.User
	db.AutoMigrate(u)
	fmt.Println(user.UserName, user.Password)
	e := db.Where("user_name = ? AND password = ?", user.UserName, user.Password).First(&u).Error
	if e != nil {
		return errors.New("username or password error")
	} else {
		//TODO
	}
	return nil
}
