package repository

import (
	"fops/domain/metaData/dockerfileTpl"
	"fops/infrastructure/repository/agent/dockerfileTplAgent"
	"fs/core/container"
	"fs/mapper"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() dockerfileTpl.Repository { return &dockerfileTplRepository{} })
}

type dockerfileTplRepository struct {
}

// ToList Dockerfile模板列表
func (repository dockerfileTplRepository) ToList() []dockerfileTpl.DomainObject {
	lstPO := dockerfileTplAgent.ToList()
	return mapper.Array[dockerfileTpl.DomainObject](lstPO)
}

// ToInfo Dockerfile模板信息
func (repository dockerfileTplRepository) ToInfo(id int) dockerfileTpl.DomainObject {
	po := dockerfileTplAgent.ToInfo(id)
	return mapper.Single[dockerfileTpl.DomainObject](po)
}

// Count Dockerfile模板数量
func (repository dockerfileTplRepository) Count() int64 {
	return dockerfileTplAgent.Count()
}

// Add 添加Dockerfile模板
func (repository dockerfileTplRepository) Add(do dockerfileTpl.DomainObject) {
	po := mapper.Single[dockerfileTplAgent.PO](do)
	dockerfileTplAgent.Add(po)
}

// Update 修改Dockerfile模板
func (repository dockerfileTplRepository) Update(id int, do dockerfileTpl.DomainObject) {
	po := mapper.Single[dockerfileTplAgent.PO](do)
	dockerfileTplAgent.Update(id, po)
}

// Delete 删除Dockerfile模板
func (repository dockerfileTplRepository) Delete(id int) {
	dockerfileTplAgent.Delete(id)
}
