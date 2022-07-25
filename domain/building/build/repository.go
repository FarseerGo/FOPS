package build

import (
	"fops/domain/_/eumBuildStatus"
)

type Repository interface {
	// GetBuildNumber 获取构建的编号
	GetBuildNumber(projectId int) int
	// GetBuildId 获取构建任务的主键
	GetBuildId(projectId int, buildNumber int) int
	// Add 添加构建任务
	Add(do DomainObject) int
	// Cancel 主动取消任务
	Cancel(id int)
	// Count 当前构建的队列数量
	Count() int64
	// ToBuildingList 获取构建队列前30
	ToBuildingList(pageSize int, pageIndex int) []DomainObject
	// ToInfo 查看构建信息
	ToInfo(id int) DomainObject
	// GetUnBuildInfo / 获取未构建的任务
	GetUnBuildInfo() DomainObject
	// SetBuilding 设置任务为构建中
	SetBuilding(buildId int) int64
	// Success 任务完成
	Success(id int)
	// GetStatus 获取构建状态
	GetStatus(id int) eumBuildStatus.Enum
}
