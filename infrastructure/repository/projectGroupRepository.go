package repository

import (
	"fops/domain/metaData/projectGroup"
	"fops/infrastructure/repository/agent/projectGroupAgent"
	"fs/core/container"
	"fs/mapper"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() projectGroup.Repository {
		return &projectGroupRepository{}
	})
}

type projectGroupRepository struct {
}

// ToList 项目组列表
func (repository *projectGroupRepository) ToList() []projectGroup.DomainObject {
	lstPO := projectGroupAgent.ToList()
	return mapper.Array[projectGroup.DomainObject](lstPO)
}

// ToInfo 项目组信息
func (repository *projectGroupRepository) ToInfo(id int) *projectGroup.DomainObject {
	po := projectGroupAgent.ToInfo(id)
	do := mapper.Single[projectGroup.DomainObject](po)
	return &do
}

// Count 项目组数量
func (repository *projectGroupRepository) Count() int64 {
	return projectGroupAgent.Count()
}

// Add 添加项目组
func (repository *projectGroupRepository) Add(do projectGroup.DomainObject) int {
	po := mapper.Single[projectGroupAgent.PO](do)
	return projectGroupAgent.Add(po)
}

// Update 修改项目组
func (repository *projectGroupRepository) Update(id int, do projectGroup.DomainObject) {
	po := mapper.Single[projectGroupAgent.PO](do)
	projectGroupAgent.Update(id, po)
}

// Delete 删除项目组
func (repository *projectGroupRepository) Delete(id int) {
	projectGroupAgent.Delete(id)
}
