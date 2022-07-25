package projectAgent

import (
	domain "fops/domain/_"
	"fops/infrastructure/repository/context"
)

// ToList 项目列表
func ToList() []PO {
	return context.NewContext().Project.ToList()
}

// ToAppList 应用列表
func ToAppList() []PO {
	return context.NewContext().Project.Where("AppId <> ''").ToList()
}

// ToListByGroupId 项目列表
func ToListByGroupId(groupId int) []PO {
	return context.NewContext().Project.Where("GroupId = ?", groupId).Select("Id", "Name", "DockerVer", "ClusterVer", "DockerHub", "GroupId", "GitId", "BuildType", "DockerfileTpl").ToList()
}

// ToInfo 项目信息
func ToInfo(id int) PO {
	return context.NewContext().Project.Where("Id = ?", id).ToEntity()
}

// Count 项目数量
func Count() int64 {
	return context.NewContext().Project.Count()
}

// GroupCount 使用项目组的数量
func GroupCount(groupId int) int64 {
	return context.NewContext().Project.Where("GroupId = ?", groupId).Count()
}

// GitCount 使用Git的数量
func GitCount(gitId int) int64 {
	return context.NewContext().Project.Where("GitId = ?", gitId).Count()
}

// Add 添加项目
func Add(po PO) int {
	context.NewContext().Project.Insert(&po)
	return po.Id
}

// Update 修改项目
func Update(id int, po PO) {
	context.NewContext().Project.Where("Id = ?", id).Update(po)
}

// UpdateDockerVer 修改项目
func UpdateDockerVer(id int, dockerVer string) {
	context.NewContext().Project.Where("Id = ?", id).UpdateValue("DockerVer", dockerVer)
}

// UpdateClusterVer 修改项目
func UpdateClusterVer(id int, dicClusterVer map[int]*domain.ClusterVerVO) {
	context.NewContext().Project.Where("Id = ?", id).UpdateValue("ClusterVer", dicClusterVer)
}

// Delete 删除项目
func Delete(id int) {
	context.NewContext().Project.Where("Id = ?", id).Delete()
}
