package projectGroup

import (
	"fops/domain/metaData/projectGroup"
	"fs/core/container"
	"fs/mapper"
)

type app struct {
	repository projectGroup.Repository
}

func NewApp() *app {
	return &app{
		repository: container.Resolve[projectGroup.Repository](),
	}
}

// ToList 项目组列表
func (app *app) ToList() []Dto {
	lstDo := app.repository.ToList()
	return mapper.Array[Dto](lstDo)
}

// ToInfo 项目组信息
func (app *app) ToInfo(id int) Dto {
	do := app.repository.ToInfo(id)
	return mapper.Single[Dto](do)
}

// Count 项目组数量
func (app *app) Count() int64 {
	return app.repository.Count()
}

// Add 添加项目组
func (app *app) Add(dto Dto) {
	do := mapper.Single[projectGroup.DomainObject](dto)
	app.repository.Add(do)
}

// Update 修改项目组
func (app *app) Update(dto Dto) {
	do := mapper.Single[projectGroup.DomainObject](dto)
	app.repository.Update(dto.Id, do)
}

// Delete 删除项目组
func (app *app) Delete(id int) {
	app.repository.Delete(id)
}
