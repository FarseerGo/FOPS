package repository

import (
	"fops/domain/k8s/cluster"
	"fops/infrastructure/repository/context"
	"fops/infrastructure/repository/model"
	"fs/core/container"
	"fs/data"
	"fs/mapper"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() cluster.Repository { return &clusterRepository{data.Init[context.MysqlContext]().Cluster} })
}

type clusterRepository struct {
	data.TableSet[model.ClusterPO]
}

// ToList 集群列表
func (repository clusterRepository) ToList() []cluster.DomainObject {
	lst := repository.Select("Id", "Name", "RuntimeEnvType").Asc("Sort").ToList()
	return mapper.Array[cluster.DomainObject](lst)
}

// ToInfo 集群信息
func (repository clusterRepository) ToInfo(id int) cluster.DomainObject {
	po := repository.Where("Id = ?", id).ToEntity()
	return mapper.Single[cluster.DomainObject](po)
}

//// Count 集群数量
//func (repository clusterRepository) Count() int64 {
//	return repository.Count()
//}

// Add 添加集群
func (repository clusterRepository) Add(do cluster.DomainObject) int {
	po := mapper.Single[model.ClusterPO](do)
	repository.Insert(&po)
	return po.Id
}

// Update 修改集群
func (repository clusterRepository) Update(id int, do cluster.DomainObject) {
	po := mapper.Single[model.ClusterPO](do)
	repository.Where("Id = ?", id).Update(po)
}

// Delete 删除集群
func (repository clusterRepository) Delete(id int) {
	repository.Where("Id = ?", id).Delete()
}
