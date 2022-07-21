package metaData

import (
	"fops/domain/metaData/dockerfileTpl"
	"fs/core/container"
	"fs/mapper"
)

type DockerfileTplApp struct {
	repository dockerfileTpl.Repository
}

func NewDockerfileTplApp() *DockerfileTplApp {
	return &DockerfileTplApp{
		repository: container.Resolve[dockerfileTpl.Repository](),
	}
}

// Add 添加Dockerfile模板
func (app *DockerfileTplApp) Add(dto DockerfileTplDto) {
	do := mapper.Single[dockerfileTpl.DomainObject](dto)
	app.repository.Add(do)
}

// Update 修改Dockerfile模板
func (app *DockerfileTplApp) Update(dto DockerfileTplDto) {
	do := mapper.Single[dockerfileTpl.DomainObject](dto)
	app.repository.Update(dto.Id, do)
}

// ToList Dockerfile模板列表
func (app *DockerfileTplApp) ToList() []DockerfileTplDto {
	lstDo := app.repository.ToList()
	return mapper.Array[DockerfileTplDto](lstDo)
}

// ToInfo Dockerfile模板信息
func (app *DockerfileTplApp) ToInfo(id int) DockerfileTplDto {
	do := app.repository.ToInfo(id)
	return mapper.Single[DockerfileTplDto](do)
}

// Count Dockerfile模板数量
func (app *DockerfileTplApp) Count() int {
	return app.repository.Count()
}
