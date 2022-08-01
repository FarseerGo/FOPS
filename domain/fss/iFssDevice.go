package fss

import (
	"fops/application/fss/client"
	"fops/application/fss/eumTaskType"
	"fops/application/fss/runLog"
	"fops/application/fss/task"
	"fops/application/fss/taskgroup"
	"github.com/farseernet/farseer.go/core"
	"github.com/farseernet/farseer.go/core/eumLogLevel"
)

type IFssDevice interface {
	// GetClientList 取出全局客户端列表
	GetClientList(fssServer string) []client.Dto
	// GetClientCount 取出全局客户端数量
	GetClientCount(fssServer string) int64
	// CopyTaskGroup 复制任务组信息
	CopyTaskGroup(fssServer string, taskGroupId int) int
	// DeleteTaskGroup 删除任务组
	DeleteTaskGroup(fssServer string, taskGroupId int) core.ApiResponseString
	// GetTaskGroupInfo 获取任务组信息
	GetTaskGroupInfo(fssServer string, taskGroupId int) taskgroup.Dto
	// SyncCacheToDb 同步缓存到数据库
	SyncCacheToDb(fssServer string)
	// GetTaskGroupList 获取全部任务组列表
	GetTaskGroupList(fssServer string) []taskgroup.Dto
	// GetTaskGroupCount 获取任务组数量
	GetTaskGroupCount(fssServer string) int64
	// GetTaskGroupUnRunCount 获取未执行的任务数量
	GetTaskGroupUnRunCount(fssServer string) int
	// AddTaskGroup 添加任务组信息
	AddTaskGroup(fssServer string, dto taskgroup.Dto) int
	// SaveTaskGroup 保存TaskGroup
	SaveTaskGroup(fssServer string, dto taskgroup.Dto)
	// TodayTaskFailCount 今日执行失败数量
	TodayTaskFailCount(fssServer string) int
	// GetTaskUnFinishList 获取进行中的任务
	GetTaskUnFinishList(fssServer string, top int) []task.Dto
	// GetEnableTaskList 获取在用的任务
	GetEnableTaskList(fssServer string, status eumTaskType.Enum, pageSize int, pageIndex int) []task.Dto
	// GetTaskList 获取指定任务组的任务列表
	GetTaskList(fssServer string, groupId int, pageSize int, pageIndex int) []task.Dto
	// GetTaskFinishList 获取已完成的任务列表
	GetTaskFinishList(fssServer string, pageSize int, pageIndex int) []task.Dto
	// CancelTask 取消任务
	CancelTask(fssServer string, taskGroupId int)
	// GetRunLogList 获取日志
	GetRunLogList(fssServer string, jobName string, logLevel eumLogLevel.Enum, pageSize int, pageIndex int) core.PageList[runLog.Dto]
}
