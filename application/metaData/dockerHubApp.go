package metaData

import (
	"fops/domain/metaData/dockerHub"
	"fs/core/container"
	"fs/mapper"
)

type DockerHubApp struct {
	repository dockerHub.Repository
}

func NewDockerHubApp() *DockerHubApp {
	return &DockerHubApp{
		repository: container.Resolve[dockerHub.Repository](),
	}
}

// Add 添加
func (app *DockerHubApp) Add(dto DockerHubDto) {
	do := mapper.Single[dockerHub.DomainObject](dto)
	app.repository.Add(do)
}

// Update 修改
func (app *DockerHubApp) Update(dto DockerHubDto) {
	do := mapper.Single[dockerHub.DomainObject](dto)
	app.repository.Update(dto.Id, do)
}

// ToList DockerHub列表
func (app *DockerHubApp) ToList() []DockerHubDto {
	lst := app.repository.ToList()
	return mapper.Array[DockerHubDto](lst)
}

// ToInfo DockerHub信息
func (app *DockerHubApp) ToInfo(id int) DockerHubDto {
	do := app.repository.ToInfo(id)
	return mapper.Single[DockerHubDto](do)
}

// Count DockerHub数量
func (app *DockerHubApp) Count() int {
	return app.repository.Count()
}
