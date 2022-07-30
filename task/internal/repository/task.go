package respository

import (
	"com.go-pro-study/todolist/go-grpc-todolist-study/task/internal/service"
)

type Task struct {
	TaskID    uint `gorm:"primarykey"`
	UserID    uint `gorm:"index"`
	Status    int  `gorm:"default:0"`
	Title     string
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64
}

func (*Task) TaskCreate(req *service.TaskRequest) error {
	task := Task{
		UserID:    uint(req.UserID),
		Title:     req.Title,
		Content:   req.Content,
		StartTime: int64(req.StartTime),
		EndTime:   int64(req.EndTime),
	}
	//fmt.Println("task:", task)
	return DB.Create(&task).Error
}

func (t *Task) TaskDelete(req *service.TaskRequest) (err error) {
	return DB.Model(&Task{}).Where("task_id=?", req.TaskID).Delete(&Task{}).Error
}

func (t *Task) TaskShow(req *service.TaskRequest) (tasklist []Task, err error) {
	err = DB.Model(&Task{}).Where("user_id=?", req.UserID).Find(&tasklist).Error
	if err != nil {
		return nil, err
	}
	return tasklist, nil
}

func (*Task) TaskUpdate(req *service.TaskRequest) (err error) {
	t := Task{}
	err = DB.Where("task_id=?", req.TaskID).First(&t).Error
	if err != nil {
		return
	}
	t.UserID = uint(req.UserID)
	t.Title = req.Title
	t.Content = req.Content
	t.StartTime = int64(req.StartTime)
	t.EndTime = int64(req.EndTime)
	return DB.Save(&t).Error
}

func BuildTask(item Task) *service.TaskModel {
	return &service.TaskModel{
		TaskID:    uint32(item.TaskID),
		UserID:    uint32(item.UserID),
		Status:    uint32(item.Status),
		Title:     item.Title,
		Content:   item.Content,
		StartTime: uint32(item.StartTime),
		EndTime:   uint32(item.EndTime),
	}
}

func BuildTasks(item []Task) (tList []*service.TaskModel) {
	for _, v := range item {
		t := BuildTask(v)
		tList = append(tList, t)
	}
	return tList
}
