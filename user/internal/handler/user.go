package handler

import (
	"com.go-pro-study/todolist/go-grpc-todolist-study/user/internal/respository"
	"com.go-pro-study/todolist/go-grpc-todolist-study/user/internal/service"
	"com.go-pro-study/todolist/go-grpc-todolist-study/user/pkg/e"
	"context"
)

type UserService struct {
	*service.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

//用户登录
func (u *UserService) UserLogin(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var usr respository.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.Sucess
	err = usr.ShowUserInfo(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.UserDetail = respository.BuildUser(usr)
	return resp, nil
}

//UserRegister 用户注册
func (u *UserService) UserRegister(ctx context.Context, req *service.UserRequest) (resp *service.UserDetailResponse, err error) {
	var usr respository.User
	resp = new(service.UserDetailResponse)
	resp.Code = e.Sucess
	usr, err = usr.UserCreate(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.UserDetail = respository.BuildUser(usr)
	return resp, nil
}
