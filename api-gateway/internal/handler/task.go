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

func ListTask(ginCtx *gin.Context) {
	var tReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&tReq))
	claim, _ := util.ParseToken(ginCtx.GetHeader("Authorization"))
	tReq.UserID = uint32(claim.UserId)
	taskService := ginCtx.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.TaskShow(context.Background(), &tReq)

	PanicIfTaskError(err)
	r := res.Response{
		Status: uint(taskResp.Code),
		Data:   taskResp,
		Msg:    e.GetMsg(uint(taskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}

func CreateTask(ginCtx *gin.Context) {
	var tReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&tReq))
	claim, _ := util.ParseToken(ginCtx.GetHeader("Authorization"))
	tReq.UserID = uint32(claim.UserId)
	taskService := ginCtx.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.TaskCreate(context.Background(), &tReq)

	PanicIfTaskError(err)
	r := res.Response{
		Status: uint(taskResp.Code),
		Data:   taskResp,
		Msg:    e.GetMsg(uint(taskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}

func UpdateTask(ginCtx *gin.Context) {
	var tReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&tReq))
	claim, _ := util.ParseToken(ginCtx.GetHeader("Authorization"))
	tReq.UserID = uint32(claim.UserId)
	taskService := ginCtx.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.TaskUpdate(context.Background(), &tReq)

	PanicIfTaskError(err)
	r := res.Response{
		Status: uint(taskResp.Code),
		Data:   taskResp,
		Msg:    e.GetMsg(uint(taskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}

func DeleteTask(ginCtx *gin.Context) {
	var tReq service.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&tReq))
	claim, _ := util.ParseToken(ginCtx.GetHeader("Authorization"))
	tReq.UserID = uint32(claim.UserId)
	taskService := ginCtx.Keys["task"].(service.TaskServiceClient)
	taskResp, err := taskService.TaskDelete(context.Background(), &tReq)

	PanicIfTaskError(err)
	r := res.Response{
		Status: uint(taskResp.Code),
		Data:   taskResp,
		Msg:    e.GetMsg(uint(taskResp.Code)),
	}
	ginCtx.JSON(http.StatusOK, r)
}
