package task

import (
	"fops/application/fss/eumTaskType"
	"fops/domain/fss"
	"github.com/farseernet/farseer.go/core/container"
)

type app struct {
	repository fss.IFssDevice
}

func NewApp() *app {
	return &app{
		repository: container.Resolve[fss.IFssDevice](),
	}
}

// TodayFailCount 今日执行失败数量
func (app *app) TodayFailCount() int {
	return app.repository.TodayTaskFailCount()
}

// GetUnFinishList 获取进行中的任务
func (app *app) GetUnFinishList(top int) []Dto {
	return app.repository.GetTaskUnFinishList(top)
}

// Cancel 取消任务
func (app *app) Cancel(taskGroupId int) {
	app.repository.CancelTask(taskGroupId)
}

// GetEnableList 获取在用的任务
func (app *app) GetEnableList(status eumTaskType.Enum, pageSize int, pageIndex int) []Dto {
	return app.repository.GetEnableTaskList(status, pageSize, pageIndex)
}

// GetFinishList 获取已完成的任务列表
func (app *app) GetFinishList(pageSize int, pageIndex int) []Dto {
	return app.repository.GetTaskFinishList(pageSize, pageIndex)
}

// GetList 获取指定任务组的任务列表
func (app *app) GetList(groupId int, pageSize int, pageIndex int) []Dto {
	return app.repository.GetTaskList(groupId, pageSize, pageIndex)
}
