package cluster

import (
	"fops/domain/k8s/cluster"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/mapper"
)

type app struct {
	repository cluster.Repository
}

func NewApp() *app {
	return &app{
		repository: container.Resolve[cluster.Repository](),
	}
}

// ToList 集群列表
func (app *app) ToList() []Dto {
	lstDo := app.repository.ToList()
	return mapper.Array[Dto](lstDo)
}

// Add 添加集群
func (app *app) Add(dto Dto) {
	do := mapper.Single[cluster.DomainObject](dto)
	app.repository.Add(do)
}

// Update 修改集群
func (app *app) Update(dto Dto) {
	do := mapper.Single[cluster.DomainObject](dto)
	app.repository.Update(dto.Id, do)
}

// ToInfo 集群信息
func (app *app) ToInfo(id int) Dto {
	do := app.repository.ToInfo(id)
	return mapper.Single[Dto](do)
}

// Count 集群数量
func (app *app) Count() int64 {
	return app.repository.Count()
}
