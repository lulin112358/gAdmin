package admin

import (
	"awesomeProject/model"
	"awesomeProject/model/admin"
	"awesomeProject/serializer"
	admin2 "awesomeProject/serializer/admin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
)

// 用户注册字段
type UserRegister struct {
	UserName        string `form:"user_name" json:"user_name" binding:"required"`
	Status          int    `form:"status" json:"status" binding:"required"`
	Password        string `form:"password" json:"password" binding:"required"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required"`
	Avatar          string `form:"avatar" json:"avatar" binding:"required"`
}

// 用户登陆字段
type UserLogin struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// 用户修改字段
type UserEdit struct {
	UserName        string `form:"user_name" json:"user_name" binding:"required"`
	Status          int    `form:"status" json:"status" binding:"required"`
	Password        string `form:"password" json:"password"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm"`
	Avatar          string `form:"avatar" json:"avatar" binding:"required"`
	ID              int    `form:"id" json:"id" binding:"required"`
}

// 用户列表字段
type UserList struct {
	Key  string `form:"key" json:"key"`
	Page int    `form:"page" json:"page"`
}

// 用户删除字段
type UserDel struct {
	ID int `form:"id" json:"id" binding:"required"`
}

// 用户注册
func (register *UserRegister) Register() serializer.Response {
	u := admin.User{
		UserName: register.UserName,
		Avatar:   register.Avatar,
		Status:   register.Status,
	}

	// 验证
	if register.Password != register.PasswordConfirm {
		return serializer.Response{
			Code: 0,
			Msg:  "两次输入密码不一致",
		}
	}

	count := 0
	model.DB.Model(&admin.User{}).Where("user_name = ?", register.UserName).Count(&count)
	if count > 0 {
		return serializer.Response{
			Code: 0,
			Msg:  "该用户名已经被占用",
		}
	}

	// 设置密码
	bytes, err := bcrypt.GenerateFromPassword([]byte(register.Password), 12)
	if err != nil {
		return serializer.Response{
			Code: 0,
			Msg:  "密码加密失败",
		}
	}
	u.Password = string(bytes)

	// 创建用户
	if err := model.DB.Create(&u).Error; err == nil {
		return serializer.Response{
			Code: 1,
			Msg:  "创建成功",
		}
	} else {
		return serializer.Response{
			Code:  0,
			Msg:   "创建失败",
			Error: err.Error(),
		}
	}
}

// 用户登陆
func (login *UserLogin) Login(c *gin.Context) serializer.Response {
	var u admin.User
	if err := model.DB.Where("user_name = ?", login.UserName).First(&u).Error; err != nil {
		return serializer.Response{
			Code: 0,
			Msg:  "账号或密码错误",
		}
	}

	// 检查密码
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(login.Password)) != nil {
		return serializer.Response{
			Code: 0,
			Msg:  "密码错误",
		}
	}

	// 设置session
	session := sessions.Default(c)
	session.Clear()
	session.Set("user_id", u.ID)
	session.Save()

	return serializer.Response{
		Code: 1,
		Msg:  "登陆成功",
	}
}

// 获取用户列表
func (list *UserList) GetUserList() serializer.Response {
	perPage, _ := strconv.Atoi(os.Getenv("PER_PAGE"))
	if list.Page == 0 {
		list.Page = 1
	}

	var users []admin.User
	count := 0
	model.DB.Model(&admin.User{}).Count(&count)
	if err := model.DB.Where("user_name like ?", "%"+list.Key+"%").Offset((list.Page - 1) * perPage).Limit(perPage).Find(&users).Error; err != nil {
		return serializer.Response{
			Code: 0,
			Msg:  err.Error(),
		}
	} else {
		return serializer.Response{
			Code: 1,
			Data: serializer.BuildList(admin2.UserSerializer(users), count),
			Msg:  "success",
		}
	}
}

// 获取用户信息
func GetUserInfo(id string) serializer.Response {
	var u admin.User
	if err := model.DB.First(&u, id).Error; err != nil {
		return serializer.Response{
			Code: 0,
			Msg:  "数据库查询错误",
		}
	} else {
		return serializer.Response{
			Code: 1,
			Data: u,
			Msg:  "success",
		}
	}
}

// 修改用户信息
func (edit *UserEdit) UserEdit() serializer.Response {
	// 验证
	if edit.Password != edit.PasswordConfirm {
		return serializer.Response{
			Code: 0,
			Msg:  "两次密码不一致",
		}
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(edit.Password), 12)
	if err != nil {
		return serializer.Response{
			Code: 0,
			Msg:  "密码加密错误",
		}
	}

	// 修改的数据
	update := map[string]interface{}{
		"user_name": edit.UserName,
		"avatar":    edit.Avatar,
		"password":  string(bytes),
		"status":    edit.Status,
	}

	// 修改资料
	if err := model.DB.Model(&admin.User{}).Where("id = ?", edit.ID).Updates(update).Error; err == nil {
		return serializer.Response{
			Code: 1,
			Msg:  "修改成功",
		}
	} else {
		return serializer.Response{
			Code:  0,
			Error: err.Error(),
		}
	}
}

//用户删除
func (del *UserDel) UserDel() serializer.Response {
	var u admin.User
	if err := model.DB.Where("id = ?", del.ID).Delete(&u).Error; err == nil {
		return serializer.Response{
			Code: 1,
			Msg:  "删除成功",
		}
	} else {
		return serializer.Response{
			Code: 0,
			Msg:  "删除失败",
		}
	}
}
