package device

import (
	"fops/application/fss/client"
	"fops/application/fss/eumTaskType"
	"fops/application/fss/runLog"
	"fops/application/fss/task"
	"fops/application/fss/taskgroup"
	"fops/domain/fss"
	"github.com/farseernet/farseer.go/core"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/core/eumLogLevel"
	"github.com/farseernet/farseer.go/utils/http"
)

func init() {
	_ = container.Register(func() fss.IFssDevice { return &fssDevice{} })
}

type fssDevice struct {
}

func (fssDevice) GetClientList(fssServer string) []client.Dto {
	url := fssServer + "/meta/GetClientList"
	var result = http.PostJson[core.ApiResponse[[]client.Dto]](url, nil, 2000)
	return result.Data
}

func (fssDevice) GetClientCount(fssServer string) int64 {
	url := fssServer + "/meta/GetClientCount"
	var result = http.PostJson[core.ApiResponseLong](url, nil, 2000)
	return result.Data
}

func (fssDevice) CopyTaskGroup(fssServer string, taskGroupId int) int {
	url := fssServer + "/meta/CopyTaskGroup"
	var result = http.PostJson[core.ApiResponseInt](url, map[string]int{"id": taskGroupId}, 2000)
	return result.Data
}

func (fssDevice) DeleteTaskGroup(fssServer string, taskGroupId int) core.ApiResponseString {
	url := fssServer + "/meta/DeleteTaskGroup"
	var result = http.PostJson[core.ApiResponseString](url, map[string]int{"id": taskGroupId}, 2000)
	return result
}

func (fssDevice) GetTaskGroupInfo(fssServer string, taskGroupId int) taskgroup.Dto {
	url := fssServer + "/meta/GetTaskGroupInfo"
	var result = http.PostJson[core.ApiResponse[taskgroup.Dto]](url, map[string]int{"id": taskGroupId}, 2000)
	return result.Data
}

func (fssDevice) SyncCacheToDb(fssServer string) {
	url := fssServer + "/meta/SyncCacheToDb"
	http.PostJson[core.ApiResponseString](url, nil, 2000)
}

func (fssDevice) GetTaskGroupList(fssServer string) []taskgroup.Dto {
	url := fssServer + "/meta/GetTaskGroupList"
	var result = http.PostJson[core.ApiResponse[[]taskgroup.Dto]](url, nil, 2000)
	return result.Data
}

func (fssDevice) GetTaskGroupCount(fssServer string) int64 {
	url := fssServer + "/meta/GetTaskGroupCount"
	var result = http.PostJson[core.ApiResponseLong](url, nil, 2000)
	return result.Data
}

func (fssDevice) GetTaskGroupUnRunCount(fssServer string) int {
	url := fssServer + "/meta/GetTaskGroupUnRunCount"
	var result = http.PostJson[core.ApiResponseInt](url, nil, 2000)
	return result.Data
}

func (fssDevice) AddTaskGroup(fssServer string, dto taskgroup.Dto) int {
	url := fssServer + "/meta/AddTaskGroup"
	var result = http.PostJson[core.ApiResponseInt](url, dto, 2000)
	return result.Data
}

func (fssDevice) SaveTaskGroup(fssServer string, dto taskgroup.Dto) {
	url := fssServer + "/meta/SaveTaskGroup"
	http.PostJson[core.ApiResponseInt](url, dto, 2000)
}

func (fssDevice) TodayTaskFailCount(fssServer string) int {
	url := fssServer + "/meta/TodayTaskFailCount"
	var result = http.PostJson[core.ApiResponseInt](url, nil, 2000)
	return result.Data
}

func (fssDevice) GetTaskUnFinishList(fssServer string, top int) []task.Dto {
	url := fssServer + "/meta/GetTaskUnFinishList"
	var result = http.PostJson[core.ApiResponse[[]task.Dto]](url, map[string]int{"top": top}, 2000)
	return result.Data
}

func (fssDevice) GetEnableTaskList(fssServer string, status eumTaskType.Enum, pageSize int, pageIndex int) []task.Dto {
	url := fssServer + "/meta/GetEnableTaskList"
	var result = http.PostJson[core.ApiResponse[[]task.Dto]](url, map[string]int{"Status": int(status), "PageSize": pageSize, "PageIndex": pageIndex}, 2000)
	return result.Data
}

func (fssDevice) GetTaskList(fssServer string, groupId int, pageSize int, pageIndex int) []task.Dto {
	url := fssServer + "/meta/GetTaskList"
	var result = http.PostJson[core.ApiResponse[[]task.Dto]](url, map[string]int{"GroupId": groupId, "PageSize": pageSize, "PageIndex": pageIndex}, 2000)
	return result.Data
}

func (fssDevice) GetTaskFinishList(fssServer string, pageSize int, pageIndex int) []task.Dto {
	url := fssServer + "/meta/GetTaskFinishList"
	var result = http.PostJson[core.ApiResponse[[]task.Dto]](url, map[string]int{"PageSize": pageSize, "PageIndex": pageIndex}, 2000)
	return result.Data
}

func (fssDevice) CancelTask(fssServer string, taskGroupId int) {
	url := fssServer + "/meta/CancelTask"
	http.PostJson[core.ApiResponseString](url, map[string]int{"Id": taskGroupId}, 2000)
}

func (fssDevice) GetRunLogList(fssServer string, jobName string, logLevel eumLogLevel.Enum, pageSize int, pageIndex int) core.PageList[runLog.Dto] {
	type dto struct {
		JobName   string
		LogLevel  eumLogLevel.Enum
		PageSize  int
		PageIndex int
	}
	url := fssServer + "/meta/GetRunLogList"
	var result = http.PostJson[core.ApiResponse[core.PageList[runLog.Dto]]](url, dto{
		JobName:   jobName,
		LogLevel:  logLevel,
		PageSize:  pageSize,
		PageIndex: pageIndex,
	}, 2000)
	return result.Data
}
