package yaml

import (
	"fops/domain/k8s/yamlTpl"
	"fs/core/container"
	"fs/mapper"
)

type app struct {
	repository yamlTpl.Repository
}

func NewApp() *app {
	return &app{
		repository: container.Resolve[yamlTpl.Repository](),
	}
}

// ToList Yaml模板列表
func (app *app) ToList() []Dto {
	lst := app.repository.ToList()
	return mapper.Array[Dto](lst)
}

// Add 添加Yaml模板
func (app *app) Add(dto Dto) {
	do := mapper.Single[yamlTpl.DomainObject](dto)
	app.repository.Add(do)
}

// Update 修改Yaml模板
func (app *app) Update(dto Dto) {
	do := mapper.Single[yamlTpl.DomainObject](dto)
	app.repository.Update(dto.Id, do)
}

// ToInfo Yaml模板信息
func (app *app) ToInfo(id int) Dto {
	do := app.repository.ToInfo(id)
	return mapper.Single[Dto](do)
}

// Count Yaml模板数量
func (app *app) Count() int {
	return app.repository.Count()
}
