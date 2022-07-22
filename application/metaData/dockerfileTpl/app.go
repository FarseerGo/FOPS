package dockerfileTpl

import (
	"fops/domain/metaData/dockerfileTpl"
	"fs/core/container"
	"fs/mapper"
)

type app struct {
	repository dockerfileTpl.Repository
}

func NewApp() *app {
	return &app{
		repository: container.Resolve[dockerfileTpl.Repository](),
	}
}

// Add 添加Dockerfile模板
func (app *app) Add(dto Dto) {
	do := mapper.Single[dockerfileTpl.DomainObject](dto)
	app.repository.Add(do)
}

// Update 修改Dockerfile模板
func (app *app) Update(dto Dto) {
	do := mapper.Single[dockerfileTpl.DomainObject](dto)
	app.repository.Update(dto.Id, do)
}

// ToList Dockerfile模板列表
func (app *app) ToList() []Dto {
	lstDo := app.repository.ToList()
	return mapper.Array[Dto](lstDo)
}

// ToInfo Dockerfile模板信息
func (app *app) ToInfo(id int) Dto {
	do := app.repository.ToInfo(id)
	return mapper.Single[Dto](do)
}

// Count Dockerfile模板数量
func (app *app) Count() int {
	return app.repository.Count()
}
