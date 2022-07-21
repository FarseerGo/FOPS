package fss

import (
	"fops/application/fss/eumTaskType"
	"fops/domain/fss"
	"fs/core/container"
)

type taskApp struct {
	repository fss.Repository
}

func NewTaskApp() *taskApp {
	return &taskApp{
		repository: container.Resolve[fss.Repository](),
	}
}

// TodayFailCount 今日执行失败数量
func (app *taskApp) TodayFailCount() int {
	return app.repository.TodayTaskFailCount()
}

// GetUnFinishList 获取进行中的任务
func (app *taskApp) GetUnFinishList(top int) []TaskDTO {
	return app.repository.GetTaskUnFinishList(top)
}

// Cancel 取消任务
func (app *taskApp) Cancel(taskGroupId int) {
	app.repository.CancelTask(taskGroupId)
}

// GetEnableList 获取在用的任务
func (app *taskApp) GetEnableList(status eumTaskType.Enum, pageSize int, pageIndex int) []TaskDTO {
	return app.repository.GetEnableTaskList(status, pageSize, pageIndex)
}

// GetFinishList 获取已完成的任务列表
func (app *taskApp) GetFinishList(pageSize int, pageIndex int) []TaskDTO {
	return app.repository.GetTaskFinishList(pageSize, pageIndex)
}

// GetList 获取指定任务组的任务列表
func (app *taskApp) GetList(groupId int, pageSize int, pageIndex int) []TaskDTO {
	return app.repository.GetTaskList(groupId, pageSize, pageIndex)
}
