package repository

import (
	domain "fops/domain/_"
	"fops/domain/metaData/project"
	"fops/infrastructure/repository/agent/projectAgent"
	"fs/core/container"
	"fs/mapper"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() project.Repository { return &projectRepository{} })
}

type projectRepository struct {
}

// ToList 项目列表
func (repository projectRepository) ToList() []project.DomainObject {
	lst := projectAgent.ToList()
	return mapper.Array[project.DomainObject](lst)
}

// ToListByGroupId 项目列表
func (repository projectRepository) ToListByGroupId(groupId int) []project.DomainObject {
	lst := projectAgent.ToListByGroupId(groupId)
	return mapper.Array[project.DomainObject](lst)
}

// ToAppList 应用列表
func (repository projectRepository) ToAppList() []project.DomainObject {
	lst := projectAgent.ToAppList()
	return mapper.Array[project.DomainObject](lst)
}

// ToInfo 项目信息
func (repository projectRepository) ToInfo(id int) project.DomainObject {
	po := projectAgent.ToInfo(id)
	return mapper.Single[project.DomainObject](po)
}

// Count 项目数量
func (repository projectRepository) Count() int64 {
	return projectAgent.Count()
}

// GroupCount 使用项目组的数量
func (repository projectRepository) GroupCount(groupId int) int64 {
	return projectAgent.GroupCount(groupId)
}

// GitCount 使用Git的数量
func (repository projectRepository) GitCount(gitId int) int64 {
	return projectAgent.GitCount(gitId)
}

// Add 添加项目
func (repository projectRepository) Add(do project.DomainObject) int {
	po := mapper.Single[projectAgent.PO](do)
	return projectAgent.Add(po)
}

// Update 修改项目
func (repository projectRepository) Update(id int, do project.DomainObject) {
	po := mapper.Single[projectAgent.PO](do)
	projectAgent.Update(id, po)
}

// UpdateDockerVer 修改镜像版本
func (repository projectRepository) UpdateDockerVer(id int, dockerVer string) {
	projectAgent.UpdateDockerVer(id, dockerVer)
}

// UpdateClusterVer 修改集群的镜像版本
func (repository projectRepository) UpdateClusterVer(id int, dicClusterVer map[int]*domain.ClusterVerVO) {
	projectAgent.UpdateClusterVer(id, dicClusterVer)
}

// Delete 删除项目
func (repository projectRepository) Delete(id int) {
	projectAgent.Delete(id)
}
