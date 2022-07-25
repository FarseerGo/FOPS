package buildAgent

import (
	"fops/domain/_/eumBuildStatus"
	"fops/infrastructure/repository/context"
	"fs"
	"fs/core"
)

// GetBuildNumber 获取构建的编号
func GetBuildNumber(projectId int) int {
	return context.NewContext().Build.Where("ProjectId = ?", projectId).Order("Id desc").GetInt("BuildNumber")
}

// Add 添加构建任务
func Add(po PO) int {
	context.NewContext().Build.Insert(&po)
	return po.Id
}

// Update 修改任务
func Update(id int, po PO) int64 {
	return context.NewContext().Build.Where("Id = ?", id).Update(po)
}

// GetBuildId 获取构建任务的主键
func GetBuildId(projectId int, buildNumber int) int {
	return context.NewContext().Build.Where("BuildNumber = ? and ProjectId = ?", buildNumber, projectId).GetInt("Id")
}

// Count 当前构建的队列数量
func Count() int64 {
	return context.NewContext().Build.Where("Status <> ?", eumBuildStatus.Finish).Count()
}

// ToBuildingList 获取构建队列前30
func ToBuildingList(pageSize int, pageIndex int) core.PageList[PO] {
	return context.NewContext().Build.Select("Id", "Status", "BuildNumber", "IsSuccess", "ProjectId", "ProjectName", "CreateAt", "FinishAt", "ClusterId").Order("Id desc").ToPageList(pageSize, pageIndex)
}

// ToInfo 查看构建信息
func ToInfo(id int) PO {
	return context.NewContext().Build.Where("Id = ?", id).ToEntity()
}

// GetStatus 获取构建状态
func GetStatus(id int) eumBuildStatus.Enum {
	return eumBuildStatus.Enum(context.NewContext().Build.Where("Id = ?", id).GetInt("Status"))
}

// ToUnBuildInfo 获取未构建的任务
func ToUnBuildInfo() PO {
	return context.NewContext().Build.Where("Status = ? and BuildServerId = ?", eumBuildStatus.None, fs.AppId).Asc("Id").ToEntity()
}
