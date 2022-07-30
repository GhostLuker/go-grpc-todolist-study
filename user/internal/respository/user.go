package respository

import (
	"com.go-pro-study/todolist/go-grpc-todolist-study/user/internal/service"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserId         uint   `gorm:"primarykey"`
	UserName       string `gorm:"unique"`
	NickName       string
	PasswordDigest string
}

const (
	PasswordCost = 12 //密码加密难度
)

//检查用户是否存在
func (user *User) CheckUserExist(req *service.UserRequest) bool {
	if err := DB.Where("user_name=?", req.UserName).First(&user).Error; err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

//获取用户信息
func (user *User) ShowUserInfo(req *service.UserRequest) (err error) {
	if exist := user.CheckUserExist(req); exist {
		return nil
	}
	return errors.New("Username not exist")
}

//
func (user *User) UserCreate(req *service.UserRequest) (usr User, err error) {
	var count int64
	DB.Where("user_name=?", req.UserName).Count(&count)
	if count != 0 {
		return User{}, errors.New("UserName exists")
	}
	usr = User{
		UserName: req.UserName,
		NickName: req.NickName,
	}
	//密码加密
	usr.SetPassword(req.Password)
	err = DB.Create(&usr).Error
	return usr, err
}

//密码加密
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

//校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

//序列化User
func BuildUser(item User) *service.UserModel {
	return &service.UserModel{
		UserID:   uint32(item.UserId),
		UserName: item.UserName,
		NickName: item.NickName,
	}
}
