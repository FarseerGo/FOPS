package project

import (
	"fops/domain/metaData/project"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/mapper"
)

type app struct {
	repository project.Repository
}

func NewApp() *app {
	return &app{
		repository: container.Resolve[project.Repository](),
	}
}

// ToList 项目列表
func (app *app) ToList() []Dto {
	lst := app.repository.ToList()
	return mapper.Array[Dto](lst)
}

// ToInfo 项目信息
func (app *app) ToInfo(id int) Dto {
	do := app.repository.ToInfo(id)
	return mapper.Single[Dto](do)
}

// Add 添加项目
func (app *app) Add(dto Dto) int {
	do := mapper.Single[project.DomainObject](dto)
	do.CheckForSave()
	return app.repository.Add(do)
}

// Update 修改项目
func (app *app) Update(dto Dto) {
	do := mapper.Single[project.DomainObject](dto)
	do.CheckForSave()
	app.repository.Update(do.Id, do)
}

// GitCount 使用Git的数量
func (app *app) GitCount(id int) int64 {
	return app.repository.GitCount(id)
}

// Count 项目数量
func (app *app) Count() int64 {
	return app.repository.Count()
}

// ToAppList 应用列表
func (app *app) ToAppList() []Dto {
	lst := app.repository.ToAppList()
	return mapper.Array[Dto](lst)
}

// GroupCount 使用项目组的数量
func (app *app) GroupCount(id int) int64 {
	return app.repository.GroupCount(id)
}

// Delete 删除项目
func (app *app) Delete(id int) {
	app.repository.Delete(id)
}
