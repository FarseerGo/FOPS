package fss

import (
	"fops/application/fss"
	"fops/application/fss/eumTaskType"
	"fops/domain/_/eumLogLevel"
	"fs/core"
)

type Repository interface {
	// GetClientList 取出全局客户端列表
	GetClientList() []fss.ClientDto
	// GetClientCount 取出全局客户端数量
	GetClientCount() int64
	// CopyTaskGroup 复制任务组信息
	CopyTaskGroup(taskGroupId int) core.ApiResponseInt
	// DeleteTaskGroup 删除任务组
	DeleteTaskGroup(taskGroupId int) core.ApiResponseString
	// GetTaskGroupInfo 获取任务组信息
	GetTaskGroupInfo(taskGroupId int) fss.TaskGroupDTO
	// SyncCacheToDb 同步缓存到数据库
	SyncCacheToDb()
	// GetTaskGroupList 获取全部任务组列表
	GetTaskGroupList() []fss.TaskGroupDTO
	// GetTaskGroupCount 获取任务组数量
	GetTaskGroupCount() int64
	// GetTaskGroupUnRunCount 获取未执行的任务数量
	GetTaskGroupUnRunCount() int
	// AddTaskGroup 添加任务组信息
	AddTaskGroup(dto fss.TaskGroupDTO) core.ApiResponseInt
	// SaveTaskGroup 保存TaskGroup
	SaveTaskGroup(dto fss.TaskGroupDTO)
	// TodayTaskFailCount 今日执行失败数量
	TodayTaskFailCount() int
	// GetTaskUnFinishList 获取进行中的任务
	GetTaskUnFinishList(top int) []fss.TaskDTO
	// GetEnableTaskList 获取在用的任务
	GetEnableTaskList(status eumTaskType.Enum, pageSize int, pageIndex int) []fss.TaskDTO
	// GetTaskList 获取指定任务组的任务列表
	GetTaskList(groupId int, pageSize int, pageIndex int) []fss.TaskDTO
	// GetTaskFinishList 获取已完成的任务列表
	GetTaskFinishList(pageSize int, pageIndex int) []fss.TaskDTO
	// CancelTask 取消任务
	CancelTask(taskGroupId int)
	// GetRunLogList 获取日志
	GetRunLogList(jobName string, logLevel eumLogLevel.Enum, pageSize int, pageIndex int) core.PageList[fss.RunLogDTO]
}
