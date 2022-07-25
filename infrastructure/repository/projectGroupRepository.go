package repository

import (
	"fops/domain/metaData/projectGroup"
	projectGroupAgent "fops/infrastructure/repository/agent/projectGroup"
	"fs/core/container"
	"fs/mapper"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() projectGroup.Repository {
		return &projectGroupRepository{
			agent: projectGroupAgent.NewAgent(),
		}
	})
}

type projectGroupRepository struct {
	agent projectGroupAgent.Agent
}

// ToList 项目组列表
func (repository *projectGroupRepository) ToList() []projectGroup.DomainObject {
	lstPO := repository.agent.ToList()
	return mapper.Array[projectGroup.DomainObject](lstPO)
}

// ToInfo 项目组信息
func (repository *projectGroupRepository) ToInfo(id int) *projectGroup.DomainObject {
	po := repository.agent.ToInfo(id)
	do := mapper.Single[projectGroup.DomainObject](po)
	return &do
}

// Count 项目组数量
func (repository *projectGroupRepository) Count() int64 {
	return repository.agent.Count()
}

// Add 添加项目组
func (repository *projectGroupRepository) Add(do projectGroup.DomainObject) int {
	po := mapper.Single[projectGroupAgent.PO](do)
	return repository.agent.Add(po)
}

// Update 修改项目组
func (repository *projectGroupRepository) Update(id int, do projectGroup.DomainObject) {
	po := mapper.Single[projectGroupAgent.PO](do)
	repository.agent.Update(id, po)
}

// Delete 删除项目组
func (repository *projectGroupRepository) Delete(id int) {
	repository.agent.Delete(id)
}
