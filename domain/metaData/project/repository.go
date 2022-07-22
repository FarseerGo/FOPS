package project

import domain "fops/domain/_"

type Repository interface {
	// ToList 项目列表
	ToList() []DomainObject
	// ToListByGroupId 项目列表
	ToListByGroupId(groupId int) []DomainObject
	// ToAppList 应用列表
	ToAppList() []DomainObject
	// ToInfo 项目信息
	ToInfo(id int) DomainObject
	// Count 项目数量
	Count() int
	// GroupCount 使用项目组的数量
	GroupCount(groupId int) int
	// GitCount 使用Git的数量
	GitCount(gitId int) int
	// Add 添加项目
	Add(project DomainObject) int
	// Update 修改项目
	Update(id int, project DomainObject)
	// UpdateDockerVer 修改镜像版本
	UpdateDockerVer(id int, dockerVer string)
	// UpdateClusterVer 修改集群的镜像版本
	UpdateClusterVer(id int, dicClusterVer map[int]*domain.ClusterVerVO)
	// Delete 删除项目
	Delete(id int)
}
