package handler

import (
	"com.go-pro-study/todolist/go-grpc-todolist-study/task/internal/repository"
	"com.go-pro-study/todolist/go-grpc-todolist-study/task/internal/service"
	"com.go-pro-study/todolist/go-grpc-todolist-study/task/pkg/e"
	"context"
)

type TaskService struct {
	service.UnimplementedTaskServiceServer
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (t *TaskService) TaskCreate(ctx context.Context, req *service.TaskRequest) (resp *service.CommonResponse, err error) {
	var task respository.Task
	resp = new(service.CommonResponse)
	resp.Code = e.Sucess
	err = task.TaskCreate(req)
	if err != nil {
		resp.Code = e.Error
		resp.Msg = e.GetMsg(e.Error)
		resp.Data = err.Error()
		return resp, err
	}
	resp.Msg = e.GetMsg(uint(resp.Code))
	return resp, nil
}

func (t *TaskService) TaskShow(ctx context.Context, req *service.TaskRequest) (resp *service.TasksDetailResponse, err error) {
	var task respository.Task
	resp = new(service.TasksDetailResponse)
	resp.Code = e.Sucess
	tasklist, err := task.TaskShow(req)
	if err != nil {
		resp.Code = e.Error
		return resp, err
	}
	resp.TaskDetail = respository.BuildTasks(tasklist)
	return resp, nil
}

func (t *TaskService) TaskUpdate(ctx context.Context, req *service.TaskRequest) (resp *service.CommonResponse, err error) {
	var task respository.Task
	resp = new(service.CommonResponse)
	resp.Code = e.Sucess
	err = task.TaskUpdate(req)
	if err != nil {
		resp.Code = e.Error
		resp.Msg = e.GetMsg(e.Error)
		resp.Data = err.Error()
		return resp, err
	}
	resp.Msg = e.GetMsg(uint(resp.Code))
	return resp, nil
}

func (t *TaskService) TaskDelete(ctx context.Context, req *service.TaskRequest) (resp *service.CommonResponse, err error) {
	var task respository.Task
	resp = new(service.CommonResponse)
	resp.Code = e.Sucess
	err = task.TaskDelete(req)
	if err != nil {
		resp.Code = e.Error
		resp.Msg = e.GetMsg(e.Error)
		resp.Data = err.Error()
		return resp, err
	}
	resp.Msg = e.GetMsg(uint(resp.Code))
	return resp, nil
}
