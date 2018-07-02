package models

import (
	"github.com/lfuture/easygin/pkg/database/mysql"
)

type User struct {
	mysql.Model
	FirstName string `gorm:"Column:first_name;varchar(20)"`
	LastName string `gorm:"Column:last_name;varchar(20)"`
}


func (User) TableName() string {
	return "user"
}

func AddUser(FirstName string, LastName string) *User {
	user := &User{
		FirstName:FirstName,
		LastName:LastName,
	}
	mysql.DB.NewRecord(user)
	mysql.DB.Create(user)
	mysql.DB.Save(user)
	return user
}