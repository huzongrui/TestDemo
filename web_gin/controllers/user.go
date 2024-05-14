package controllers

import (
	"strconv"
	"web_gin/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

// func (u UserController) GetUserInfo(c *gin.Context) {
// 	idstr := c.Param("id")
// 	//name := c.Param("name")
// 	fmt.Println(idstr)
// 	id, _ := strconv.Atoi(idstr)
// 	user, _ := models.GetUserTest(id)
// 	returnSuccess(c, 0, user.Id, user, 1)
// }
// func (u UserController) GETList(c *gin.Context) {
// 	// logger.Write("日志信息", "user")
// 	// defer func() {
// 	// 	if err := recover(); err != nil {
// 	// 		fmt.Println("捕获异常：", err)
// 	// 	}
// 	// }()
// 	num1 := 1
// 	num2 := 0
// 	num3 := num1 / num2
// 	returnError(c, 4004, num3)
// }
// func (u UserController) AddUserTest(c *gin.Context) {
// 	name := c.DefaultPostForm("name", "")
// 	id, err := models.AddUserTest(name)
// 	if err != nil {
// 		returnError(c, 4002, "保存错误")
// 		return
// 	}
// 	returnSuccess(c, 0, "保存成功", id, 1)
// }
// func (u UserController) UpdateUserTest(c *gin.Context) {
// 	name := c.DefaultPostForm("name", "")
// 	idstr := c.DefaultPostForm("id", "")
// 	id, _ := strconv.Atoi(idstr)
// 	models.UpdateUserTest(id, name)
// 	returnSuccess(c, 0, "更新成功", true, 1)

// }
// func (u UserController) DeleleUserTest(c *gin.Context) {
// 	idstr := c.DefaultPostForm("id", "")
// 	id, _ := strconv.Atoi(idstr)
// 	err := models.DeleleUserTest(id)
// 	if err != nil {
// 		returnError(c, 4002, "删除错误")
// 		return
// 	}
// 	returnSuccess(c, 0, "删除成功", true, 1)

// }
//
//	func (u UserController) GetUserListTest(c *gin.Context) {
//		users, err := models.GetUserListTest()
//		if err != nil {
//			returnError(c, 4004, "没有相关数据")
//			return
//		}
//		returnSuccess(c, 0, "获取成功", users, 1)
//	}
func (u UserController) Register(c *gin.Context) {
	name := c.DefaultPostForm("name", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirmPassword", "")

	if name == "" || password == "" || confirmPassword == "" {
		returnError(c, 4001, "请输入正确信息")
		return
	}
	if password != confirmPassword {
		returnError(c, 4001, "密码和确认密码不匹配")
		return
	}
	user, _ := models.GetUserInfoByName(name)
	if user.Id != 0 {
		returnError(c, 4001, "用户名已经存在")
		return
	}
	_, err := models.AddUser(name, EncryMd5(password))
	if err != nil {
		returnError(c, 4001, "注册失败")
		return
	}
	returnSuccess(c, 0, "注册成功", "", 1)
}

type UserApi struct {
	Id   int
	Name string
}

func (u UserController) Login(c *gin.Context) {
	name := c.DefaultPostForm("name", "")
	password := c.DefaultPostForm("password", "")
	if name == "" || password == "" {
		returnError(c, 4001, "请输入正确信息")
		return
	}
	user, _ := models.GetUserInfoByName(name)
	if user.Id == 0 {
		returnError(c, 4004, "用户名或者密码错误")
		return
	}
	if user.Password != EncryMd5(password) {
		returnError(c, 4004, "用户名或者密码错误")
		return
	}
	sessions := sessions.Default(c)
	sessions.Set("login:"+strconv.Itoa(user.Id), user.Id)
	sessions.Save()
	data := UserApi{Id: user.Id, Name: user.Name}
	returnSuccess(c, 0, "登录成功", data, 1)

}
