package fss

import (
	"fops/domain/fss"
	"fs/core"
	"fs/core/container"
	"fs/linq"
	"fs/utils/parse"
	"github.com/robfig/cron/v3"
)

type taskGroupApp struct {
	repository fss.Repository
}

func NewTaskGroupApp() *taskGroupApp {
	return &taskGroupApp{
		repository: container.Resolve[fss.Repository](),
	}
}

// SyncCacheToDb 同步缓存到数据库
func (app *taskGroupApp) SyncCacheToDb() {
	app.repository.SyncCacheToDb()
}

// ToPageList 获取任务组列表
func (app *taskGroupApp) ToPageList(jobName string, enable int, pageSize int, pageIndex int) core.PageList[TaskGroupDTO] {
	lst := app.repository.GetTaskGroupList()

	// 筛选JobName
	if jobName != "" {
		lst = linq.From(lst).Where(func(item TaskGroupDTO) bool { return item.JobName == jobName }).ToArray()
	}

	// 筛选状态
	if enable > -1 {
		isEnable := parse.Convert(enable, false)
		lst = linq.From(lst).Where(func(item TaskGroupDTO) bool { return item.IsEnable == isEnable }).ToArray()
	}

	lst = linq.FromOrder[TaskGroupDTO, string](lst).OrderBy(func(item TaskGroupDTO) string { return item.JobName })

	return linq.From(lst).ToPageList(pageSize, pageIndex)
}

// ToList 获取任务组列表
func (app *taskGroupApp) ToList() []TaskGroupDTO {
	lst := app.repository.GetTaskGroupList()
	return linq.FromOrder[TaskGroupDTO, string](lst).OrderBy(func(item TaskGroupDTO) string { return item.JobName })
}

// ToInfo 获取任务组信息
func (app *taskGroupApp) ToInfo(taskGroupId int) TaskGroupDTO {
	var taskGroup = app.repository.GetTaskGroupInfo(taskGroupId)

	if taskGroup.Cron == "" {
		taskGroup.Cron = parse.Convert(taskGroup.IntervalMs, "")
	}

	return taskGroup
}

// Add 添加任务组
func (app *taskGroupApp) Add(taskGroup TaskGroupDTO) {
	// 是否为数字
	if parse.IsInt(taskGroup.Cron) {
		if _, err := cron.ParseStandard(taskGroup.Cron); err != nil {
			panic("Cron格式错误")
		}
	}

	var result = app.repository.AddTaskGroup(taskGroup)
	if !result.Status {
		panic(result.StatusMessage)
	}
}

// 保存TaskGroup

func (app *taskGroupApp) Update(taskGroup TaskGroupDTO) {
	if taskGroup.Id < 1 {
		panic("任务组不存在")
	}
	if parse.IsInt(taskGroup.Cron) {
		if _, err := cron.ParseStandard(taskGroup.Cron); err != nil {
			panic("Cron格式错误")
		}
	}

	// 是否为数字
	taskGroup.IntervalMs = 0
	app.repository.SaveTaskGroup(taskGroup)
}

// Copy 复制任务组信息
func (app *taskGroupApp) Copy(taskGroupId int) int {
	return app.repository.CopyTaskGroup(taskGroupId).Data
}

// Save 保存TaskGroup
func (app *taskGroupApp) Save(taskGroup TaskGroupDTO) {
	app.repository.SaveTaskGroup(taskGroup)
}

// Delete 删除任务组
func (app *taskGroupApp) Delete(taskGroupId int) {
	app.repository.DeleteTaskGroup(taskGroupId)
}

// GetCount 获取任务组数量
func (app *taskGroupApp) GetCount() int64 {
	return app.repository.GetTaskGroupCount()
}

// GetUnRunCount 获取未执行的任务数量
func (app *taskGroupApp) GetUnRunCount() int {
	return app.repository.GetTaskGroupUnRunCount()
}
