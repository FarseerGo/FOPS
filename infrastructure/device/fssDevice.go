package device

import (
	"fops/application/fss/client"
	"fops/application/fss/eumTaskType"
	"fops/application/fss/runLog"
	"fops/application/fss/task"
	"fops/application/fss/taskgroup"
	"fops/domain/_/eumLogLevel"
	"fops/domain/fss"
)

func init() {
	_ = container.Register(func() fss.IFssDevice { return &fssDevice{} })
}

type fssDevice struct {
}

func (fssDevice) GetClientList() []client.Dto {
	//TODO implement me
	panic("implement me")
}

func (fssDevice) GetClientCount() int64 {
	//TODO implement me
	panic("implement me")
}

func (fssDevice) CopyTaskGroup(taskGroupId int) core.ApiResponseInt {
	//TODO implement me
	panic("implement me")
}

func (fssDevice) DeleteTaskGroup(taskGroupId int) core.ApiResponseString {
	//TODO implement me
	panic("implement me")
}

func (fssDevice) GetTaskGroupInfo(taskGroupId int) taskgroup.Dto {
	//TODO implement me
	panic("implement me")
}

func (fssDevice) SyncCacheToDb() {
	//TODO implement me
	panic("implement me")
}

func (fssDevice) GetTaskGroupList() []taskgroup.Dto {
	//TODO implement me
	panic("implement me")
}

func (fssDevice) GetTaskGroupCount() int64 {
	//TODO implement me
	panic("implement me")
}

func (fssDevice) GetTaskGroupUnRunCount() int {
	//TODO implement me
	panic("implement me")
}

func (fssDevice) AddTaskGroup(dto taskgroup.Dto) core.ApiResponseInt {
	//TODO implement me
	panic("implement me")
}

func (fssDevice) SaveTaskGroup(dto taskgroup.Dto) {
	//TODO implement me
	panic("implement me")
}

func (fssDevice) TodayTaskFailCount() int {
	//TODO implement me
	panic("implement me")
}

func (fssDevice) GetTaskUnFinishList(top int) []task.Dto {
	//TODO implement me
	panic("implement me")
}

func (fssDevice) GetEnableTaskList(status eumTaskType.Enum, pageSize int, pageIndex int) []task.Dto {
	//TODO implement me
	panic("implement me")
}

func (fssDevice) GetTaskList(groupId int, pageSize int, pageIndex int) []task.Dto {
	//TODO implement me
	panic("implement me")
}

func (fssDevice) GetTaskFinishList(pageSize int, pageIndex int) []task.Dto {
	//TODO implement me
	panic("implement me")
}

func (fssDevice) CancelTask(taskGroupId int) {
	//TODO implement me
	panic("implement me")
}

func (fssDevice) GetRunLogList(jobName string, logLevel eumLogLevel.Enum, pageSize int, pageIndex int) core.PageList[runLog.Dto] {
	//TODO implement me
	panic("implement me")
}
