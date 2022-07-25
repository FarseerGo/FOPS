package repository

import (
	"fops/domain/k8s/cluster"
	"fops/infrastructure/repository/agent/clusterAgent"
	"fs/core/container"
	"fs/mapper"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() cluster.Repository { return &clusterRepository{} })
}

type clusterRepository struct {
}

// ToList 集群列表
func (repository clusterRepository) ToList() []cluster.DomainObject {
	lst := clusterAgent.ToList()
	return mapper.Array[cluster.DomainObject](lst)
}

// ToInfo 集群信息
func (repository clusterRepository) ToInfo(id int) cluster.DomainObject {
	po := clusterAgent.ToInfo(id)
	return mapper.Single[cluster.DomainObject](po)
}

// Count 集群数量
func (repository clusterRepository) Count() int64 {
	return clusterAgent.Count()
}

// Add 添加集群
func (repository clusterRepository) Add(do cluster.DomainObject) int {
	po := mapper.Single[clusterAgent.PO](do)
	return clusterAgent.Add(po)
}

// Update 修改集群
func (repository clusterRepository) Update(id int, do cluster.DomainObject) {
	po := mapper.Single[clusterAgent.PO](do)
	clusterAgent.Update(id, po)
}

// Delete 删除集群
func (repository clusterRepository) Delete(id int) {
	clusterAgent.Delete(id)
}
