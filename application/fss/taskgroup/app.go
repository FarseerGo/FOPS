package taskgroup

import (
	"fops/domain/fss"
	"github.com/farseernet/farseer.go/core"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/linq"
	"github.com/farseernet/farseer.go/utils/parse"
	"github.com/robfig/cron/v3"
)

type app struct {
	repository fss.IFssDevice
	fssServer  string
}

func NewApp() *app {
	return &app{
		repository: container.Resolve[fss.IFssDevice](),
	}
}

// SyncCacheToDb 同步缓存到数据库
func (app *app) SyncCacheToDb() {
	app.repository.SyncCacheToDb(app.fssServer)
}

// ToPageList 获取任务组列表
func (app *app) ToPageList(jobName string, enable int, pageSize int, pageIndex int) core.PageList[Dto] {
	lst := app.repository.GetTaskGroupList(app.fssServer)

	// 筛选JobName
	if jobName != "" {
		lst = linq.From(lst).Where(func(item Dto) bool { return item.JobName == jobName }).ToArray()
	}

	// 筛选状态
	if enable > -1 {
		isEnable := parse.Convert(enable, false)
		lst = linq.From(lst).Where(func(item Dto) bool { return item.IsEnable == isEnable }).ToArray()
	}

	lst = linq.From(lst).OrderBy(func(item Dto) any { return item.JobName })

	return linq.From(lst).ToPageList(pageSize, pageIndex)
}

// ToList 获取任务组列表
func (app *app) ToList() []Dto {
	lst := app.repository.GetTaskGroupList(app.fssServer)
	return linq.From[Dto, string](lst).OrderBy(func(item Dto) any { return item.JobName })
}

// ToInfo 获取任务组信息
func (app *app) ToInfo(taskGroupId int) Dto {
	var taskGroup = app.repository.GetTaskGroupInfo(app.fssServer, taskGroupId)

	if taskGroup.Cron == "" {
		taskGroup.Cron = parse.Convert(taskGroup.IntervalMs, "")
	}

	return taskGroup
}

// Add 添加任务组
func (app *app) Add(taskGroup Dto) {
	// 是否为数字
	if parse.IsInt(taskGroup.Cron) {
		if _, err := cron.ParseStandard(taskGroup.Cron); err != nil {
			panic("Cron格式错误")
		}
	}

	var result = app.repository.AddTaskGroup(app.fssServer, taskGroup)

	if result < 1 {
		panic("添加任务组失败")
	}
}

// Update 保存TaskGroup
func (app *app) Update(taskGroup Dto) {
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
	app.repository.SaveTaskGroup(app.fssServer, taskGroup)
}

// Copy 复制任务组信息
func (app *app) Copy(taskGroupId int) int {
	return app.repository.CopyTaskGroup(app.fssServer, taskGroupId)
}

// Save 保存TaskGroup
func (app *app) Save(taskGroup Dto) {
	app.repository.SaveTaskGroup(app.fssServer, taskGroup)
}

// Delete 删除任务组
func (app *app) Delete(taskGroupId int) {
	app.repository.DeleteTaskGroup(app.fssServer, taskGroupId)
}

// GetCount 获取任务组数量
func (app *app) GetCount() int64 {
	return app.repository.GetTaskGroupCount(app.fssServer)
}

// GetUnRunCount 获取未执行的任务数量
func (app *app) GetUnRunCount() int {
	return app.repository.GetTaskGroupUnRunCount(app.fssServer)
}
