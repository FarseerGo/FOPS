package metaData

import (
	"fops/domain/metaData/projectGroup"
	"fs/core/container"
	"fs/mapper"
)

type ProjectGroupApp struct {
	repository projectGroup.Repository
}

func NewProjectGroupApp() *ProjectGroupApp {
	return &ProjectGroupApp{
		repository: container.Resolve[projectGroup.Repository](),
	}
}

// ToList 项目组列表
func (app *ProjectGroupApp) ToList() []ProjectGroupDto {
	lstDo := app.repository.ToList()
	return mapper.Array[ProjectGroupDto](lstDo)
}

// ToInfo 项目组信息
func (app *ProjectGroupApp) ToInfo(id int) ProjectGroupDto {
	do := app.repository.ToInfo(id)
	return mapper.Single[ProjectGroupDto](do)
}

// Count 项目组数量
func (app *ProjectGroupApp) Count() int {
	return app.repository.Count()
}

// Add 添加项目组
func (app *ProjectGroupApp) Add(dto ProjectGroupDto) {
	do := mapper.Single[projectGroup.DomainObject](dto)
	app.repository.Add(do)
}

// Update 修改项目组
func (app *ProjectGroupApp) Update(dto ProjectGroupDto) {
	do := mapper.Single[projectGroup.DomainObject](dto)
	app.repository.Update(dto.Id, do)
}

// Delete 删除项目组
func (app *ProjectGroupApp) Delete(id int) {
	app.repository.Delete(id)
}
