package k8s

import (
	"fops/domain/k8s/cluster"
	"fs/core/container"
	"fs/mapper"
)

type ClusterApp struct {
	repository cluster.Repository
}

func NewClusterApp() *ClusterApp {
	return &ClusterApp{
		repository: container.Resolve[cluster.Repository](),
	}
}

// ToList 集群列表
func (app *ClusterApp) ToList() []ClusterDto {
	lstDo := app.repository.ToList()
	return mapper.Array[ClusterDto](lstDo)
}

// Add 添加集群
func (app *ClusterApp) Add(dto ClusterDto) {
	do := mapper.Single[cluster.DomainObject](dto)
	app.repository.Add(do)
}

// Update 修改集群
func (app *ClusterApp) Update(dto ClusterDto) {
	do := mapper.Single[cluster.DomainObject](dto)
	app.repository.Update(dto.Id, do)
}

// ToInfo 集群信息
func (app *ClusterApp) ToInfo(id int) ClusterDto {
	do := app.repository.ToInfo(id)
	return mapper.Single[ClusterDto](do)
}

// Count 集群数量
func (app *ClusterApp) Count() int {
	return app.repository.Count()
}
