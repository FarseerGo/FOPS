package repository

import (
	domain "fops/domain/_"
	"fops/domain/metaData/project"
	"fops/infrastructure/repository/context"
	"fops/infrastructure/repository/model"
	"fs/core/container"
	"fs/data"
	"fs/mapper"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() project.Repository { return &projectRepository{data.Init[context.MysqlContext]().Project} })
}

type projectRepository struct {
	data.TableSet[model.ProjectPO]
}

// ToList 项目列表
func (repository projectRepository) ToList() []project.DomainObject {
	lst := repository.ToList()
	return mapper.Array[project.DomainObject](lst)
}

// ToListByGroupId 项目列表
func (repository projectRepository) ToListByGroupId(groupId int) []project.DomainObject {
	lst := repository.Where("GroupId = ?", groupId).Select("Id", "Name", "DockerVer", "ClusterVer", "DockerHub", "GroupId", "GitId", "BuildType", "DockerfileTpl").ToList()
	return mapper.Array[project.DomainObject](lst)
}

// ToAppList 应用列表
func (repository projectRepository) ToAppList() []project.DomainObject {
	lst := repository.Where("AppId <> ''").ToList()
	return mapper.Array[project.DomainObject](lst)
}

// ToInfo 项目信息
func (repository projectRepository) ToInfo(id int) project.DomainObject {
	po := repository.Where("Id = ?", id).ToEntity()
	return mapper.Single[project.DomainObject](po)
}

//// Count 项目数量
//func (repository projectRepository) Count() int64 {
//	return repository.Count()
//}

// GroupCount 使用项目组的数量
func (repository projectRepository) GroupCount(groupId int) int64 {
	return repository.Where("GroupId = ?", groupId).Count()
}

// GitCount 使用Git的数量
func (repository projectRepository) GitCount(gitId int) int64 {
	return repository.Where("GitId = ?", gitId).Count()
}

// Add 添加项目
func (repository projectRepository) Add(do project.DomainObject) int {
	po := mapper.Single[model.ProjectPO](do)
	repository.Insert(&po)
	return po.Id
}

// Update 修改项目
func (repository projectRepository) Update(id int, do project.DomainObject, args ...interface{}) {
	po := mapper.Single[model.ProjectPO](do)
	repository.Where("Id = ?", id).Select(args).Update(po)
}

// UpdateYamlId 修改yaml脚本ID
func (repository projectRepository) UpdateYamlId(id int, deploymentId int, serviceId int, ingressId int, configId int) {
	repository.Where("Id = ?", id).Select("K8STplDeployment", "K8STplService", "K8STplIngress", "K8STplConfig").Update(model.ProjectPO{
		K8STplDeployment: deploymentId,
		K8STplService:    serviceId,
		K8STplIngress:    ingressId,
		K8STplConfig:     configId,
	})
}

// UpdateDockerVer 修改镜像版本
func (repository projectRepository) UpdateDockerVer(id int, dockerVer string) {
	repository.Where("Id = ?", id).UpdateValue("DockerVer", dockerVer)
}

// UpdateClusterVer 修改集群的镜像版本
func (repository projectRepository) UpdateClusterVer(id int, dicClusterVer map[int]*domain.ClusterVerVO) {
	repository.Where("Id = ?", id).UpdateValue("ClusterVer", dicClusterVer)
}

// Delete 删除项目
func (repository projectRepository) Delete(id int) {
	repository.Where("Id = ?", id).Delete()
}
