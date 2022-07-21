package metaData

import (
	"fops/domain/metaData/project"
	"fs/core/container"
	"fs/mapper"
)

type ProjectApp struct {
	repository project.Repository
}

func NewProjectApp() *ProjectApp {
	return &ProjectApp{
		repository: container.Resolve[project.Repository](),
	}
}

// ToList 项目列表
func (app *ProjectApp) ToList() []ProjectDto {
	lst := app.repository.ToList()
	return mapper.Array[ProjectDto](lst)
}

// ToInfo 项目信息
func (app *ProjectApp) ToInfo(id int) ProjectDto {
	do := app.repository.ToInfo(id)
	return mapper.Single[ProjectDto](do)
}

// Add 添加项目
func (app *ProjectApp) Add(dto ProjectDto) int {
	do := mapper.Single[project.DomainObject](dto)
	do.CheckForSave()
	return app.repository.Add(do)
}

// Update 修改项目
func (app *ProjectApp) Update(dto ProjectDto) {
	do := mapper.Single[project.DomainObject](dto)
	do.CheckForSave()
	app.repository.Update(do.Id, do)
}

// GitCount 使用Git的数量
func (app *ProjectApp) GitCount(id int) int {
	return app.repository.GitCount(id)
}

// Count 项目数量
func (app *ProjectApp) Count() int {
	return app.repository.Count()
}

// ToAppList 应用列表
func (app *ProjectApp) ToAppList() []ProjectDto {
	lst := app.repository.ToAppList()
	return mapper.Array[ProjectDto](lst)
}

// GroupCount 使用项目组的数量
func (app *ProjectApp) GroupCount(id int) int {
	return app.repository.GroupCount(id)
}

// Delete 删除项目
func (app *ProjectApp) Delete(id int) {
	app.repository.Delete(id)
}
