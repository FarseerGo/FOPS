package fss

import (
	"fops/application/fss/client"
	"fops/application/fss/eumTaskType"
	"fops/application/fss/runLog"
	"fops/application/fss/task"
	"fops/application/fss/taskgroup"
	"fops/domain/_/eumLogLevel"
	"fs/core"
)

type Repository interface {
	// GetClientList 取出全局客户端列表
	GetClientList() []client.Dto
	// GetClientCount 取出全局客户端数量
	GetClientCount() int64
	// CopyTaskGroup 复制任务组信息
	CopyTaskGroup(taskGroupId int) core.ApiResponseInt
	// DeleteTaskGroup 删除任务组
	DeleteTaskGroup(taskGroupId int) core.ApiResponseString
	// GetTaskGroupInfo 获取任务组信息
	GetTaskGroupInfo(taskGroupId int) taskgroup.Dto
	// SyncCacheToDb 同步缓存到数据库
	SyncCacheToDb()
	// GetTaskGroupList 获取全部任务组列表
	GetTaskGroupList() []taskgroup.Dto
	// GetTaskGroupCount 获取任务组数量
	GetTaskGroupCount() int64
	// GetTaskGroupUnRunCount 获取未执行的任务数量
	GetTaskGroupUnRunCount() int
	// AddTaskGroup 添加任务组信息
	AddTaskGroup(dto taskgroup.Dto) core.ApiResponseInt
	// SaveTaskGroup 保存TaskGroup
	SaveTaskGroup(dto taskgroup.Dto)
	// TodayTaskFailCount 今日执行失败数量
	TodayTaskFailCount() int
	// GetTaskUnFinishList 获取进行中的任务
	GetTaskUnFinishList(top int) []task.Dto
	// GetEnableTaskList 获取在用的任务
	GetEnableTaskList(status eumTaskType.Enum, pageSize int, pageIndex int) []task.Dto
	// GetTaskList 获取指定任务组的任务列表
	GetTaskList(groupId int, pageSize int, pageIndex int) []task.Dto
	// GetTaskFinishList 获取已完成的任务列表
	GetTaskFinishList(pageSize int, pageIndex int) []task.Dto
	// CancelTask 取消任务
	CancelTask(taskGroupId int)
	// GetRunLogList 获取日志
	GetRunLogList(jobName string, logLevel eumLogLevel.Enum, pageSize int, pageIndex int) core.PageList[runLog.Dto]
}
