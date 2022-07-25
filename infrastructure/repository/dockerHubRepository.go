package repository

import (
	"fops/domain/metaData/dockerHub"
	"fops/infrastructure/repository/agent/dockerHubAgent"
	"fs/core/container"
	"fs/mapper"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() dockerHub.Repository { return &dockerHubRepository{} })
}

type dockerHubRepository struct {
}

// ToList DockerHub列表
func (repository dockerHubRepository) ToList() []dockerHub.DomainObject {
	lstPO := dockerHubAgent.ToList()
	return mapper.Array[dockerHub.DomainObject](lstPO)
}

// ToInfo DockerHub信息
func (repository dockerHubRepository) ToInfo(id int) dockerHub.DomainObject {
	po := dockerHubAgent.ToInfo(id)
	return mapper.Single[dockerHub.DomainObject](po)
}

// Count DockerHub数量
func (repository dockerHubRepository) Count() int64 {
	return dockerHubAgent.Count()
}

// Add 添加仓库
func (repository dockerHubRepository) Add(do dockerHub.DomainObject) {
	po := mapper.Single[dockerHubAgent.PO](do)
	dockerHubAgent.Add(po)
}

// Update 修改仓库
func (repository dockerHubRepository) Update(id int, do dockerHub.DomainObject) {
	po := mapper.Single[dockerHubAgent.PO](do)
	dockerHubAgent.Update(id, po)
}

// Delete 删除仓库
func (repository dockerHubRepository) Delete(id int) {
	dockerHubAgent.Delete(id)
}
