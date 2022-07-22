package git

import (
	"fops/domain/building/build/vo"
	"fops/domain/building/device"
	"fops/domain/metaData/git"
	"fs/core/container"
	"fs/mapper"
	"time"
)

type app struct {
	repository git.Repository
	device     device.IGitDevice
}

func NewApp() *app {
	return &app{
		repository: container.Resolve[git.Repository](),
		device:     container.Resolve[device.IGitDevice](),
	}
}

// ToList Git列表
func (app *app) ToList() []Dto {
	lst := app.repository.ToList()
	return mapper.Array[Dto](lst)
}

// Add 添加GIT
func (app *app) Add(dto Dto) {
	do := mapper.Single[git.DomainObject](dto)
	app.repository.Add(do)
}

// UpdateForTime 修改GIT的拉取时间
func (app *app) UpdateForTime(gitId int, time time.Time) {
	app.repository.UpdateForTime(gitId, time)
}

// Update 修改GIT
func (app *app) Update(dto Dto) {
	do := mapper.Single[git.DomainObject](dto)
	app.repository.Update(dto.Id, do)
}

// ToInfo Git信息
func (app *app) ToInfo(id int) Dto {
	do := app.repository.ToInfo(id)
	return mapper.Single[Dto](do)
}

// Delete 删除GIT
func (app *app) Delete(id int) {
	app.repository.Delete(id)
}

// Count Git数量
func (app *app) Count() int {
	return app.repository.Count()
}

// Clear 清除仓库
func (app *app) Clear(gitId int, progress chan string) bool {
	if progress == nil {
		progress = make(chan string, 10)
	}

	var do = app.repository.ToInfo(gitId)
	return app.device.Clear(do.Hub, progress)
}

// CloneOrPull 根据判断是否存在Git目录，来决定返回Clone or pull
func (app *app) CloneOrPull(dto Dto, lstLog []string) bool {
	progress := make(chan string, 10)
	for _, log := range lstLog {
		progress <- log
	}
	vo := mapper.Single[vo.GitVO](dto)
	return app.device.CloneOrPull(vo, progress)
}
