package dockerHub

import (
	"fops/domain/metaData/dockerHub"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/mapper"
)

type app struct {
	repository dockerHub.Repository
}

func NewApp() *app {
	return &app{
		repository: container.Resolve[dockerHub.Repository](),
	}
}

// Add 添加
func (app *app) Add(dto Dto) {
	do := mapper.Single[dockerHub.DomainObject](dto)
	app.repository.Add(do)
}

// Update 修改
func (app *app) Update(dto Dto) {
	do := mapper.Single[dockerHub.DomainObject](dto)
	app.repository.Update(dto.Id, do)
}

// ToList DockerHub列表
func (app *app) ToList() []Dto {
	lst := app.repository.ToList()
	return mapper.Array[Dto](lst)
}

// ToInfo DockerHub信息
func (app *app) ToInfo(id int) Dto {
	do := app.repository.ToInfo(id)
	return mapper.Single[Dto](do)
}

// Count DockerHub数量
func (app *app) Count() int64 {
	return app.repository.Count()
}
