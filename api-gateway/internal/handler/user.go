package handler

import (
	"com.go-pro-study/todolist/go-grpc-todolist-study/api-gateway/internal/service"
	"com.go-pro-study/todolist/go-grpc-todolist-study/api-gateway/pkg/e"
	"com.go-pro-study/todolist/go-grpc-todolist-study/api-gateway/pkg/res"
	"com.go-pro-study/todolist/go-grpc-todolist-study/api-gateway/pkg/util"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

//用户注册
func UserRegister(ginContext *gin.Context) {
	var usrReq service.UserRequest
	PanicIfUserError(ginContext.Bind(&usrReq))

	//gin.Key中去除服务实现
	userService := ginContext.Keys["user"].(service.UserServiceClient)
	userResp, err := userService.UserRegister(context.Background(), &usrReq)
	PanicIfUserError(err)
	r := res.Response{
		Data:   userResp,
		Status: uint(userResp.Code),
		Msg:    e.GetMsg(uint(userResp.Code)),
	}
	ginContext.JSON(http.StatusOK, r)
}

//用户登录
func UserLogin(ginContext *gin.Context) {
	var usrReq service.UserRequest
	PanicIfUserError(ginContext.Bind(&usrReq))

	//gin.Key中获取服务实现
	userService := ginContext.Keys["user"].(service.UserServiceClient)
	userResp, err := userService.UserLogin(context.Background(), &usrReq)
	PanicIfUserError(err)
	token, err := util.GenerateToken(uint(userResp.UserDetail.UserID))
	r := res.Response{
		Data: res.TokenData{
			User:  userResp.UserDetail,
			Token: token,
		},
		Status: uint(userResp.Code),
		Msg:    e.GetMsg(uint(userResp.Code)),
	}
	ginContext.JSON(http.StatusOK, r)
}
