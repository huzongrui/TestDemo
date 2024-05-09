package models

import (
	"time"
	"web_gin/dao"
)

type User struct {
	Id         int
	Name       string
	Password   string
	AddTime    int64
	UpdateTime int64
}

func (User) TableName() string {
	return "user"
}

func GetUserInfoByName(name string) (User, error) {
	var user User
	err := dao.Db.Where("name = ?", name).First(&user).Error
	return user, err
}
func AddUser(name string, password string) (int, error) {
	user := User{Name: name, Password: password, AddTime: time.Now().Unix(), UpdateTime: time.Now().Unix()}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}

// func GetUserTest(id int) (User, error) {
// 	var user User
// 	//err := dao.Db.Where("id = ?", id).First(&user).Error
// 	err := dao.Db.Where("id = ?", id).First(&user).Error
// 	return user, err
// }
// func GetUserListTest() ([]User, error) {
// 	var users []User
// 	err := dao.Db.Where("id < ?", 3).Find(&users).Error
// 	return users, err

// }
// func AddUserTest(name string) (int, error) {
// 	user := User{Name: name}
// 	err := dao.Db.Create(&user).Error
// 	return user.Id, err
// }
// func UpdateUserTest(id int, name string) {
// 	//user,id:=User{Name: name,Id: id}
// 	dao.Db.Model(&User{}).Where("id = ?", id).Update("name", name)

// }
// func DeleleUserTest(id int) error {
// 	err := dao.Db.Delete(&User{}, id).Error
// 	return err
// }
