package project

import "fops/infrastructure/repository/context"

type agent struct {
}

func NewAgent() agent { return agent{} }

// ToList 项目列表
func (agent) ToList() []PO {
	return context.NewContext().Project.ToList()
}

// ToAppList 应用列表
func (agent) ToAppList() []PO {
	return context.NewContext().Project.Where("AppId <> ''").ToList()
}

// ToList 项目列表
func (agent) ToList(groupId int) []PO {
	return context.NewContext().Project.Where("GroupId = ?", groupId).Select("Id", "Name", "DockerVer", "ClusterVer", "DockerHub", "GroupId", "GitId", "BuildType", "DockerfileTpl").ToList()
}

// ToInfo 项目信息
func (agent) ToInfo(id int) PO {
	return context.NewContext().Project.Where("Id = ?", id).ToEntity()
}

// Count 项目数量
func (agent) Count() int64 {
	return context.NewContext().Project.Count()
}

// GroupCount 使用项目组的数量
func (agent) GroupCount(groupId int) int64 {
	return context.NewContext().Project.Where("GroupId = ?", groupId).Count()
}

// GitCount 使用Git的数量
func (agent) GitCount(gitId int) int64 {
	return context.NewContext().Project.Where("GitId = ?", gitId).Count()
}

// Add 添加项目
func (agent) Add(po PO) int {
	context.NewContext().Project.Insert(&po)
	return po.Id
}

// Update 修改项目
func (agent) Update(id int, po PO) {
	context.NewContext().Project.Where("Id = ?", id).Update(po)
}

// Delete 删除项目
func (agent) Delete(id int) {
	context.NewContext().Project.Where("Id = ?", id).Delete()
}
